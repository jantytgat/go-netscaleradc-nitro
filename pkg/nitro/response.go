package nitro

import (
	"fmt"
	"reflect"
)

type Response[T ResourceReader] struct {
	ErrorCode float64 `json:"errorcode"`
	Message   string  `json:"message"`
	Severity  string  `json:"severity"`
	Data      []T
}

func (n *Response[T]) GetResourceTypeName() string {
	t := *new(T)
	return t.GetTypeName()
}

func (n *Response[T]) ExtractData(data interface{}) error {
	var (
		err error
		m   []interface{}
	)

	// Nitro responses either contain an array of the requested resources, or just the resource
	// Convert data to []interface[] for uniform operation
	rt := reflect.TypeOf(data)
	switch rt.Kind() {
	case reflect.Map:
		m = append(m, data)
	default:
		m = data.([]interface{})
	}

	for _, i := range m {
		var d T
		if err = mapToStruct(&d, i.(map[string]interface{})); err != nil {
			return ResourceExtractionError.WithMessage(fmt.Sprintf("failed to extract data for %s", d.GetTypeName()+": "+err.Error())).WithError(err)
		}
		n.Data = append(n.Data, d)
	}
	return nil
}
