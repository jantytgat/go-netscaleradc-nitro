package nitro

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func addResource[T ResourceReader](ctx context.Context, client *Client, r T) error {
	req := Request[T]{
		Method: http.MethodPost,
		Data: []T{
			r,
		},
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return err
	}

	if res.ErrorCode != 0 {
		return ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}
	return nil
}

func bindResource[T ResourceReader](ctx context.Context, client *Client, r T) error {
	req := Request[T]{
		Method: http.MethodPut,
		Data:   []T{r},
	}

	if _, err := executeNitroRequest[T](ctx, client, &req); err != nil {
		return err
	}
	return nil
}

func clearResource[T ResourceReader](ctx context.Context, client *Client, r T) error {
	req := Request[T]{
		Method: http.MethodPost,
		Action: ActionClear,
		Data:   []T{r},
	}

	if _, err := executeNitroRequest[T](ctx, client, &req); err != nil {
		return err
	}
	return nil
}

func countResource[T ResourceReader](ctx context.Context, client *Client) (T, error) {
	req := Request[T]{
		Method: http.MethodGet,
		Action: ActionCount,
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return *new(T), err
	}

	return res.Data[0], nil
}

func createHttpRequest[T ResourceReader](baseUrl string, req *Request[T]) (*http.Request, error) {
	var (
		err  error
		body io.Reader
		buf  = &bytes.Buffer{} // Used to validate data before creating the request
	)

	// Do not serialize body when there are no items to serialize
	if len(req.Data) > 0 {
		if body, err = req.serializeBody(); err != nil {
			return nil, ClientCreateHttpRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE+" for %s: %s", req.GetResourceTypeName(), err.Error())).WithError(err)
		}

		tee := io.TeeReader(body, buf)

		if err = req.ValidateData(tee); err != nil {
			return nil, ClientCreateHttpRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE+" for %s: %s", req.GetResourceTypeName(), err.Error())).WithError(err)
		}
	}

	// Build URL Path and Query
	var query string
	if query, err = req.GetUrlPathAndQuery(); err != nil {
		return nil, ClientCreateHttpRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE+" for %s: %s", req.GetResourceTypeName(), err.Error())).WithError(err)
	}
	url := baseUrl + query

	var r *http.Request
	if r, err = http.NewRequest(req.Method, url, bytes.NewReader(buf.Bytes())); err != nil {
		return nil, ClientCreateHttpRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE+" for %s: %s", req.GetResourceTypeName(), err.Error())).WithError(err)
	}

	return r, nil
}

func createResource[T ResourceReader](ctx context.Context, client *Client, r T) error {
	req := Request[T]{
		Method: http.MethodPost,
		Action: ActionCreate,
		Data: []T{
			r,
		},
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return err
	}

	if res.ErrorCode != 0 {
		return ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}
	return nil
}

func deleteResource[T ResourceReader](ctx context.Context, client *Client, name string, arguments map[string]string) error {
	req := Request[T]{
		Method:       http.MethodDelete,
		ResourceName: name,
		Arguments:    arguments,
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return err
	}

	if res.ErrorCode != 0 {
		return ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}
	return nil
}

func deserializeResponse[T ResourceReader](res *http.Response) (*Response[T], error) {
	var (
		err error
		r   Response[T]
	)

	defer func() {
		e := res.Body.Close()
		if e != nil {
			err = e
		}
	}()

	var rawBody []byte
	rawBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE + ": read body: " + err.Error())).WithError(err)
	}

	// Unmarshal the raw body json data into a map
	var bodyMap map[string]interface{}
	err = json.Unmarshal(rawBody, &bodyMap)
	if err != nil {
		return &r, ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE + ": unmarshal json: " + err.Error())).WithError(err)
	}

	var found bool
	if _, found = bodyMap["errorcode"]; !found {
		return &r, ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE + ": error code not found in body"))
	}
	r.ErrorCode = bodyMap["errorcode"].(float64)

	if _, found = bodyMap["message"]; !found {
		return &r, ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE + ": message not found in body"))

	}
	r.Message = bodyMap["message"].(string)

	if _, found = bodyMap["severity"]; !found {
		return &r, ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE + ": severity not found in body"))

	}
	r.Severity = bodyMap["severity"].(string)

	// Check if the Nitro API call return data for the resource itself
	// The response body should have a field with the name of the resource used for the call.
	t := *new(T)
	if _, ok := bodyMap[t.GetTypeName()]; !ok {
		// No resource data was returned, the HTTP response code details the result of the operation.
		// If something happened at the Nitro API level, the result details are in the errorCode and message
		switch res.StatusCode {
		case 200:
			return &r, nil
		case 201:
			return &r, nil
		default:
			return &r, ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE + ": " + r.Message)).WithCode(r.ErrorCode)
		}
	}

	// Extract the data for the resource and convert to a struct for the resource type
	if err = r.ExtractData(bodyMap[t.GetTypeName()]); err != nil {
		return &r, ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE + ": failed to extract data: " + err.Error())).WithError(err)
	}

	return &r, nil
}

func disableResource[T ResourceReader](ctx context.Context, client *Client, r T) error {
	req := Request[T]{
		Method: http.MethodPost,
		Action: ActionDisable,
		Data: []T{
			r,
		},
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return err
	}

	if res.ErrorCode != 0 {
		return ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}
	return nil
}

func enableResource[T ResourceReader](ctx context.Context, client *Client, r T) error {
	req := Request[T]{
		Method: http.MethodPost,
		Action: ActionEnable,
		Data: []T{
			r,
		},
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return err
	}

	if res.ErrorCode != 0 {
		return ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}
	return nil
}

func executeNitroRequest[T ResourceReader](ctx context.Context, c *Client, r *Request[T]) (*Response[T], error) {
	var (
		err error
		req *http.Request
		res *http.Response
	)

	if req, err = createHttpRequest[T](c.BaseUrl(), r); err != nil {
		return nil, ClientExecuteRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_EXECUTEREQUEST_ERROR_MESSAGE + ": " + err.Error())).WithError(err)
	}

	// Add context to http request
	req = req.WithContext(ctx)

	// Configure Nitro headers on http request
	c.setHeadersOnRequest(r.getResourceName(), req)

	if res, err = c.client.Do(req); err != nil {
		return nil, ClientExecuteRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_EXECUTEREQUEST_ERROR_MESSAGE + ": " + err.Error())).WithError(err)
	}

	var nitroRes *Response[T]
	if res.ContentLength == 0 {
		nitroRes = &Response[T]{
			ErrorCode: 0,
			Message:   "Done",
			Severity:  "NONE",
		}
		return nitroRes, nil
	}

	if nitroRes, err = deserializeResponse[T](res); err != nil {
		return nil, ClientExecuteRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_EXECUTEREQUEST_ERROR_MESSAGE + ": " + err.Error())).WithError(err)
	}

	return nitroRes, nil
}

func getResource[T ResourceReader](ctx context.Context, client *Client, name string, attributes []string) (T, error) {
	req := Request[T]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return *new(T), err
	}

	if res.ErrorCode != 0 {
		return *new(T), ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}

	return res.Data[0], nil
}

func linkResource[T ResourceReader](ctx context.Context, client *Client, r T) error {
	req := Request[T]{
		Method: http.MethodPost,
		Action: ActionLink,
		Data:   []T{r},
	}

	if _, err := executeNitroRequest[T](ctx, client, &req); err != nil {
		return err
	}
	return nil
}

func listResource[T ResourceReader](ctx context.Context, client *Client, attributes []string, filter map[string]string) ([]T, error) {
	req := Request[T]{
		Method:     http.MethodGet,
		Attributes: attributes,
		Filter:     filter,
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return nil, err
	}

	if res.ErrorCode != 0 {
		return nil, ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}

	return res.Data, nil
}

func listResourceWithName[T ResourceReader](ctx context.Context, client *Client, name string, attributes []string, filter map[string]string) ([]T, error) {
	req := Request[T]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
		Filter:       filter,
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return nil, err
	}

	if res.ErrorCode != 0 {
		return nil, ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}

	return res.Data, nil
}

func mapToStruct[T any](t *T, v map[string]interface{}) error {
	var err error
	var jsonData []byte

	rt := reflect.TypeOf(t)
	// Convert map to json
	jsonData, err = json.Marshal(v)
	if err != nil {
		return ResourceSerializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_SERIALIZATION_ERROR_MESSAGE + ": marshal json to: " + rt.String() + " failed: " + err.Error())).WithError(err)
	}

	// Convert json to struct of type T
	err = json.Unmarshal(jsonData, t)
	if err != nil {
		return ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_SERIALIZATION_ERROR_MESSAGE + ": marshal json to struct of " + rt.String() + " failed: " + err.Error())).WithError(err)
	}

	return nil
}

func putResource[T ResourceReader](ctx context.Context, client *Client, r T) error {
	req := Request[T]{
		Method: http.MethodPut,
		Data:   []T{r},
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return err
	}

	if res.ErrorCode != 0 {
		return ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}
	return nil
}

func renameResource[T ResourceReader](ctx context.Context, client *Client, r T) error {
	req := Request[T]{
		Method: http.MethodPost,
		Action: ActionRename,
		Data: []T{
			r,
		},
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return err
	}

	if res.ErrorCode != 0 {
		return ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}

	return nil
}

func stats[T ResourceReader](ctx context.Context, client *Client, name string, attributes []string) (T, error) {
	req := Request[T]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}

	var res *Response[T]
	var err error
	if res, err = executeNitroRequest(ctx, client, &req); err != nil {
		return *new(T), err
	}

	if res.ErrorCode != 0 {
		return *new(T), ApiError.WithCode(res.ErrorCode).WithMessage(res.Message)
	}

	return res.Data[0], nil
}

func unbindResource[T ResourceReader](ctx context.Context, client *Client, name string, arguments map[string]string) error {
	req := Request[T]{
		Method:       http.MethodDelete,
		ResourceName: name,
		Arguments:    arguments,
	}

	if _, err := executeNitroRequest[T](ctx, client, &req); err != nil {
		return err
	}
	return nil
}

func unlinkResource[T ResourceReader](ctx context.Context, client *Client, r T) error {
	req := Request[T]{
		Method: http.MethodPost,
		Action: ActionUnlink,
		Data:   []T{r},
	}

	if _, err := executeNitroRequest[T](ctx, client, &req); err != nil {
		return err
	}
	return nil
}
