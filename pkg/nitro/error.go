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
	"errors"
	"fmt"
)

var (
	SpecificNitroError        = errors.New("nitro error")
	SpecificNitroErrorMessage = "%s: %w"

	DataValidationError        = errors.New("data validation error")
	DataValidationErrorMessage = "failed validating data for %s with message %s --> %w"

	SerializeBodyError        = errors.New("body serialization error")
	SerializeBodyErrorMessage = "failed serializing body for %s with message %s --> %w"

	InvalidFieldError        = errors.New("field name not found error")
	InvalidFieldErrorMessage = "cannot find field '%s' --> %w"

	CreateHttpRequestError        = errors.New("create http request error")
	CreateHttpRequestErrorMessage = "createhttprequest[%s]: %s --> %w"

	ExecuteNitroRequestError        = errors.New("execute nitro request error")
	ExecuteNitroRequestErrorMessage = "executenitrorequest: %s --> %w"

	DeserializeResponseError        = errors.New("unmarshal body error")
	DeserializeResponseErrorMessage = "deserializeresponse: %s --> %w"

	ExtractDataError        = errors.New("extract data error")
	ExtractDataErrorMessage = "failed to extract data for %s with message %s --> %w"

	NotImplementedError        = errors.New("not implemented error")
	NotImplementedErrorMessage = "%s --> %w"
)

func FormatSpecificNitroError(msg string) error {
	return fmt.Errorf(SpecificNitroErrorMessage, msg, SpecificNitroError)
}

func FormatDataValidationError(name string, err error) error {
	return fmt.Errorf(DataValidationErrorMessage, name, err.Error(), DataValidationError)
}

func FormatSerializeBodyError(name string, err error) error {
	return fmt.Errorf(SerializeBodyErrorMessage, name, err.Error(), SerializeBodyError)
}

func FormatInvalidFieldError(field string) error {
	return fmt.Errorf(InvalidFieldErrorMessage, field, InvalidFieldError)
}

func FormatCreateHttpRequestError(name string, err error) error {
	return fmt.Errorf(CreateHttpRequestErrorMessage, name, err.Error(), CreateHttpRequestError)
}

func FormatExecuteNitroRequestError(err error) error {
	return fmt.Errorf(ExecuteNitroRequestErrorMessage, err.Error(), ExecuteNitroRequestError)
}

func FormatDeserializeResponseError(err error) error {
	return fmt.Errorf(DeserializeResponseErrorMessage, err.Error(), DeserializeResponseError)
}

func FormatExtractDataError(t string, err error) error {
	return fmt.Errorf(ExtractDataErrorMessage, t, err.Error(), ExtractDataError)
}

func FormatNotImplementedError(msg string) error {
	return fmt.Errorf(NotImplementedErrorMessage, msg, NotImplementedError)
}
