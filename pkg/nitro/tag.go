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
	"strings"
)

type Tag struct {
	Permission string `json:"permission"`
}

func GetNitroTags[T ResourceReader]() (map[string]Tag, error) {
	var err error

	r := *new(T)
	t := reflect.TypeOf(r)

	var m = make(map[string]Tag, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		jsonTag, ok := t.Field(i).Tag.Lookup("json")
		if !ok {
			err = fmt.Errorf("could not find json tag for %v", t.Field(i).Name)
			return nil, err
		}

		var jsonElements []string
		jsonElements, err = splitTag(jsonTag, ",")
		if err != nil {
			return nil, err
		}

		fieldName := jsonElements[0]

		nitroTag, ok := t.Field(i).Tag.Lookup("nitro")
		if !ok {
			err = fmt.Errorf("could not find nitro tag for %v", t.Field(i).Name)
			continue
		}

		var parsedTag Tag
		parsedTag, err = parseNitroTag(nitroTag)
		// fmt.Printf("TAG: %v\n", parsedTag)
		if err != nil {
			return m, err
		}

		m[fieldName] = parsedTag
		if err != nil {
			return m, err
		}
	}

	return m, err
}

func parseNitroTag(tag string) (Tag, error) {
	// Tag Format: name=value,name=value,...
	var err error
	var t Tag
	var m = make(map[string]interface{})

	var elements []string
	elements, err = splitTag(tag, ",")
	if err != nil {
		return t, err
	}
	// fmt.Printf("Elements: %v\n", elements)

	for _, element := range elements {
		var e []string
		e, err = splitTag(element, "=")
		if err != nil {
			return t, err
		}
		// Add element name to map with value
		// fmt.Printf("Adding name %v with value %v\n", e[0], e[1])
		m[e[0]] = e[1]
	}

	err = mapToStruct[Tag](&t, m)
	if err != nil {
		// fmt.Printf("Error setting tag resource: %v\n", err)
		return t, err
	}
	// fmt.Printf("Tag returned: %v\n", t)
	return t, err
}

func splitTag(tag string, separator string) ([]string, error) {
	var err error

	r := strings.Split(tag, separator)
	if len(r) == 0 {
		err = fmt.Errorf("could not split tag %v", tag)
		return nil, err
	}

	return r, err
}
