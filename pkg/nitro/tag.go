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
			return nil, ResourceValidationError.WithMessage(fmt.Sprintf("Validation error for %s: field %s not found", t.Name(), t.Field(i).Name))
		}

		var jsonElements []string
		jsonElements, err = splitTag(jsonTag, ",")
		if err != nil {
			return nil, err
		}

		fieldName := jsonElements[0]

		nitroTag, ok := t.Field(i).Tag.Lookup("nitro")
		if !ok {
			// TODO cleanup
			err = fmt.Errorf("could not find nitro tag for %v", t.Field(i).Name)
			continue
		}

		var parsedTag Tag
		parsedTag, err = parseNitroTag(nitroTag)
		if err != nil {
			return m, err
		}

		m[fieldName] = parsedTag
	}

	return m, err
}

func parseNitroTag(tag string) (Tag, error) {
	// Tag Format: name=value,name=value,...
	var (
		err      error
		t        Tag
		m        = make(map[string]interface{})
		elements []string
	)

	elements, err = splitTag(tag, ",")
	if err != nil {
		return t, ResourceTagError.WithMessage(fmt.Sprintf("Cannot split tag elements for %s", tag))
	}

	for _, element := range elements {
		var e []string
		e, err = splitTag(element, "=")
		if err != nil {
			return t, ResourceTagError.WithMessage(fmt.Sprintf("Cannot parse tag elements in %s", tag))
		}
		// Add element name to map with value
		m[e[0]] = e[1]
	}

	err = mapToStruct[Tag](&t, m)
	if err != nil {
		return t, ResourceTagError.WithMessage(fmt.Sprintf("Cannot convert map to tag %s", tag)).WithError(err)
	}
	return t, nil
}

func splitTag(tag string, separator string) ([]string, error) {
	var (
		err    error
		output []string
	)

	output = strings.Split(tag, separator)
	if len(output) == 0 {
		return nil, ResourceTagError.WithMessage(fmt.Sprintf("Cannot split tag %s", tag))
	}

	return output, err
}
