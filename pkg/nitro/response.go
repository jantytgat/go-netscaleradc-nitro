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
		err = mapToStruct(&d, i.(map[string]interface{}))
		if err != nil {
			return ResourceExtractionError.WithMessage(fmt.Sprintf("Failed to extract data for %s", rt.Name())).WithError(err)
		}
		n.Data = append(n.Data, d)
	}
	return nil
}
