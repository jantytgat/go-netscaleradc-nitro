/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package nitro

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

type Request[T ResourceReader] struct {
	Method       string
	ResourceName string
	Action       Action
	Arguments    map[string]string
	Filter       map[string]string
	Attributes   []string
	Data         []T
}

func (r *Request[T]) GetResourceTypeName() string {
	t := *new(T)

	return t.GetTypeName()
}

func (r *Request[T]) getResourceType() ResourceType {
	t := *new(T)
	path := reflect.TypeOf(t).PkgPath()
	path = strings.TrimPrefix(path, ResourcePackagePath)

	switch path {
	case "config":
		return ResourceTypeConfig
	case "stat":
		return ResourceTypeStat
	default:
		return ResourceTypeUnkown
	}
}

func (r *Request[T]) ValidateResourceType() error {
	if r.getResourceType() == ResourceTypeUnkown {
		return ResourceInvalidTypeError
	}
	return nil
}

func (r *Request[T]) ValidateArguments() error {
	for k := range r.Arguments {
		if !contains(r.getValidFieldNames(), k) {
			return ResourceInvalidFieldError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_INVALIDFIELD_ERROR_MESSAGE+" %s in arguments", k))
		}
	}
	return nil
}

func (r *Request[T]) ValidateFilter() error {
	for k := range r.Filter {
		if !contains(r.getValidFieldNames(), k) {
			return ResourceInvalidFieldError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_INVALIDFIELD_ERROR_MESSAGE+" %s in filter", k))
		}
	}
	return nil
}

func (r *Request[T]) ValidateAttributes() error {
	for _, k := range r.Attributes {
		if !contains(r.getValidFieldNames(), k) {
			return ResourceInvalidFieldError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_INVALIDFIELD_ERROR_MESSAGE+" %s in attributes", k))
		}
	}
	return nil
}

// TODO REWORK ValidateData() error handling
// TODO REWORK ValidateData() logic
func (r *Request[T]) ValidateData(reader io.Reader) error {
	var (
		err              error
		resourceTypeName string
	)

	resourceTypeName = r.GetResourceTypeName()

	// Read Serialized JSON data
	var b []byte
	b, err = io.ReadAll(reader)

	if reader == nil || len(b) == 0 {
		return ResourceValidationError.WithError(err)
	}

	// Convert JSON to map[string]interface
	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return ResourceValidationError.WithError(err)
	}

	// Standardize data to []interface{} for resources that don't post an array in the body
	var data []interface{}
	switch resourceTypeName {
	case "login":
		data = append(data, m[resourceTypeName])
	default:
		data = m[resourceTypeName].([]interface{})
	}

	// Get Tags for type T
	var tags map[string]Tag
	tags, err = GetNitroTags[T]()
	if err != nil {
		return ResourceValidationError.WithError(err)
	}

	// Iterate over data in request
	for _, ds := range data {
		d := ds.(map[string]interface{})

		for k := range d {
			if tags[k].Permission == "readonly" {
				return ResourceValidationError.WithMessage(fmt.Sprintf("Validation error for read-only data field %s", k))
			}
		}
	}

	return nil
}

func (r *Request[T]) GetUrlPathAndQuery() (string, error) {
	var (
		err    error
		output strings.Builder
	)

	// Run validation functions before building the Url path and query
	err = r.validateUrl([]func() error{
		r.ValidateResourceType,
		r.ValidateArguments,
		r.ValidateFilter,
		r.ValidateAttributes})
	if err != nil {
		return "", ResourceValidationError.WithError(err)
	}

	output.WriteString(r.getResourceType().UrlPath())
	output.WriteString(r.getResourceName())
	output.WriteString(r.getUrlQuery())

	return output.String(), nil

}

func (r *Request[T]) getUrlPath() string {
	return r.getResourceType().UrlPath()
}

func (r *Request[T]) getResourceName() string {
	switch len(r.ResourceName) {
	case 0:
		return ""
	default:
		return "/" + r.ResourceName
	}
}

func (r *Request[T]) getUrlQuery() string {
	var output strings.Builder

	output.WriteString(getUrlQueryAction(r.Action))
	output.WriteString(getUrlQueryMapString(output.Len(), ArgumentsQueryType, r.Arguments, ""))
	output.WriteString(getUrlQueryMapString(output.Len(), FilterQueryType, r.Filter, "/"))
	output.WriteString(getUrlQueryListString(output.Len(), AttributesQueryType, r.Attributes))

	return output.String()
}

func (r *Request[T]) getValidFieldNames() []string {
	var fieldNames []string

	t := *new(T)
	rv := reflect.ValueOf(t)

	for i := 0; i < rv.NumField(); i++ {
		jsonTag := rv.Type().Field(i).Tag.Get("json")
		jsonTagElem := strings.Split(jsonTag, ",")
		fieldNames = append(fieldNames, strings.ToLower(jsonTagElem[0]))
	}

	return fieldNames
}

func (r *Request[T]) validateUrl(f []func() error) error {
	for _, v := range f {
		err := v()
		if err != nil {
			return err
		}
	}

	return nil
}

func contains(keys []string, key string) bool {
	for _, v := range keys {
		if v == key {
			return true
		}
	}

	return false
}

func (r *Request[T]) serializeBody() (io.Reader, error) {
	var (
		err error
		t   reflect.Type
		v   reflect.Value
	)

	resource := *new(T)
	tag := `json:"` + resource.GetTypeName() + `"`

	switch resource.GetTypeName() {
	// Login is a special case in Nitro API
	case "login":
		t = reflect.StructOf([]reflect.StructField{
			{
				Name: "Data",
				Type: reflect.TypeOf(resource),
				Tag:  reflect.StructTag(tag),
			},
		})

		// Create a new struct of type t
		v = reflect.New(t).Elem()
		v.Field(0).Set(reflect.ValueOf(r.Data[0]))
	default:
		t = reflect.StructOf([]reflect.StructField{
			{
				Name: "Data",
				Type: reflect.TypeOf(r.Data),
				Tag:  reflect.StructTag(tag),
			},
		})
		v = reflect.New(t).Elem()
		v.Field(0).Set(reflect.ValueOf(r.Data))
	}

	// Encode v to JSON
	w := bytes.Buffer{}
	err = json.NewEncoder(&w).Encode(v.Addr().Interface())
	if err != nil {
		return nil, ResourceSerializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_SERIALIZATION_ERROR_MESSAGE+" for %s", resource.GetTypeName())).WithError(err)
	}
	return bytes.NewReader(w.Bytes()), nil
}

func getUrlQueryAction(a Action) string {
	switch a {
	case ActionNone:
		return ""
	default:
		return "?action=" + a.string
	}
}

func getUrlQueryListString(urlQueryLength int, queryType UrlQueryType, queryList []string) string {
	if len(queryList) == 0 {
		return ""
	}

	var output strings.Builder
	output.WriteString(getUrlQueryStringSeparator(urlQueryLength))
	output.WriteString(queryType.Prefix())
	output.WriteString(strings.Join(queryList, ","))

	return output.String()
}

func getUrlQueryMapString(urlQueryLength int, queryType UrlQueryType, queryMap map[string]string, wrapCharacter string) string {
	if len(queryMap) == 0 {
		return ""
	}

	var output strings.Builder
	output.WriteString(getUrlQueryStringSeparator(urlQueryLength))
	output.WriteString(queryType.Prefix())
	output.WriteString(getUrlQueryMapEntriesAsString(queryMap, wrapCharacter))

	return output.String()
}

func getUrlQueryMapEntriesAsString(queryMap map[string]string, wrapCharacter string) string {
	var output strings.Builder

	keys := getUrlQueryMapSortedKeys(queryMap)
	lastIndex := len(keys) - 1

	for index, key := range keys {
		value := wrapCharacter + queryMap[key] + wrapCharacter
		output.WriteString(key)
		output.WriteString(":")
		output.WriteString(url.PathEscape(value))
		output.WriteString(getUrlQueryMapStringSeparator(index, lastIndex))
	}

	return output.String()
}

func getUrlQueryMapSortedKeys(queryMap map[string]string) []string {
	keys := make([]string, 0, len(queryMap))
	for k := range queryMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func getUrlQueryMapStringSeparator(index int, lastIndex int) string {
	if index < lastIndex {
		return ","
	} else {
		return ""
	}
}

func getUrlQueryStringSeparator(length int) string {
	if length == 0 {
		return "?"
	}
	return "&"
}
