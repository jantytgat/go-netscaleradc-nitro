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

const (
	NSGO_CLIENT_ERROR_CODE                                 errorCode = 10000000
	NSGO_CLIENT_ERROR_MESSAGE                              string    = "Client error"
	NSGO_CLIENT_CREATE_ERROR_CODE                          errorCode = 10000001
	NSGO_CLIENT_CREATE_ERROR_MESSAGE                       string    = "Error creating client"
	NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_CODE               errorCode = 10000002
	NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE            string    = "Error creating http request"
	NSGO_CLIENT_EXECUTEREQUEST_ERROR_CODE                  errorCode = 10000003
	NSGO_CLIENT_EXECUTEREQUEST_ERROR_MESSAGE               string    = "Error executing nitro request"
	NSGO_CLIENT_LOGIN_ERROR_CODE                           errorCode = 10000004
	NSGO_CLIENT_LOGIN_ERROR_MESSAGE                        string    = "Login error"
	NSGO_CLIENT_LOGOUT_ERROR_CODE                          errorCode = 10000005
	NSGO_CLIENT_LOGOUT_ERROR_MESSAGE                       string    = "Logout error"
	NSGO_CLIENT_CONNECTIONSETTINGS_ERROR_CODE              errorCode = 10000006
	NSGO_CLIENT_CONNECTIONSETTINGS_ERROR_MESSAGE           string    = "Error in connection settings"
	NSGO_RESOURCE_ERROR_CODE                               errorCode = 40000000
	NSGO_RESOURCE_ERROR_MESSAGE                            string    = "Resource error"
	NSGO_RESOURCE_VALIDATION_ERROR_CODE                    errorCode = 20000001
	NSGO_RESOURCE_VALIDATION_ERROR_MESSAGE                 string    = "Validation error"
	NSGO_RESOURCE_INVALIDTYPE_ERROR_CODE                   errorCode = 20000002
	NSGO_RESOURCE_INVALIDTYPE_ERROR_MESSAGE                string    = "Invalid Resource Type"
	NSGO_RESOURCE_INVALIDFIELD_ERROR_CODE                  errorCode = 20000003
	NSGO_RESOURCE_INVALIDFIELD_ERROR_MESSAGE               string    = "Invalid field"
	NSGO_RESOURCE_SERIALIZATION_ERROR_CODE                 errorCode = 20000004
	NSGO_RESOURCE_SERIALIZATION_ERROR_MESSAGE              string    = "Serialization error"
	NSGO_RESOURCE_DESERIALIZATION_ERROR_CODE               errorCode = 20000005
	NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE            string    = "Deserialization error"
	NSGO_RESOURCE_TAG_ERROR_CODE                           errorCode = 20000006
	NSGO_RESOURCE_TAG_ERROR_MESSAGE                        string    = "Tag error"
	NSGO_CONTROLLER_ERROR_CODE                             errorCode = 30000000
	NSGO_CONTROLLER_ERROR_MESSAGE                          string    = "Controller Error"
	NSGO_CONTROLLER_OPERATION_NOTIMPLEMENTED_ERROR_CODE    errorCode = 30000001
	NSGO_CONTROLLER_OPERATION_NOTIMPLEMENTED_ERROR_MESSAGE string    = "Operation not implemented"
	NSGO_API_ERROR_CODE                                    errorCode = 40000000
	NSGO_API_ERROR_MESSAGE                                 string    = "Nitro API specific error"

	NSERR_SSL_NOCERT_ERROR_CODE    errorCode = 1540
	NSERR_SSL_NOCERT_ERROR_MESSAGE string    = "Certificate does not exist"
)

var (
	ClientCreateError             = Error{code: NSGO_CLIENT_CREATE_ERROR_CODE, message: NSGO_CLIENT_CREATE_ERROR_MESSAGE}
	ClientCreateHttpRequestError  = Error{code: NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_CODE, message: NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE}
	ClientExecuteRequestError     = Error{code: NSGO_CLIENT_EXECUTEREQUEST_ERROR_CODE, message: NSGO_CLIENT_EXECUTEREQUEST_ERROR_MESSAGE}
	ClientLoginError              = Error{code: NSGO_CLIENT_LOGIN_ERROR_CODE, message: NSGO_CLIENT_LOGIN_ERROR_MESSAGE}
	ClientLogoutError             = Error{code: NSGO_CLIENT_LOGOUT_ERROR_CODE, message: NSGO_CLIENT_LOGOUT_ERROR_MESSAGE}
	ClientConnectionSettingsError = Error{code: NSGO_CLIENT_CONNECTIONSETTINGS_ERROR_CODE, message: NSGO_CLIENT_CONNECTIONSETTINGS_ERROR_MESSAGE}

	ResourceValidationError      = Error{code: NSGO_RESOURCE_VALIDATION_ERROR_CODE, message: NSGO_RESOURCE_VALIDATION_ERROR_MESSAGE}
	ResourceInvalidTypeError     = Error{code: NSGO_RESOURCE_INVALIDTYPE_ERROR_CODE, message: NSGO_RESOURCE_INVALIDTYPE_ERROR_MESSAGE}
	ResourceInvalidFieldError    = Error{code: NSGO_RESOURCE_INVALIDFIELD_ERROR_CODE, message: NSGO_RESOURCE_INVALIDFIELD_ERROR_MESSAGE}
	ResourceSerializationError   = Error{code: NSGO_RESOURCE_SERIALIZATION_ERROR_CODE, message: NSGO_RESOURCE_SERIALIZATION_ERROR_MESSAGE}
	ResourceDeserializationError = Error{code: NSGO_RESOURCE_DESERIALIZATION_ERROR_CODE, message: NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE}
	ResourceTagError             = Error{code: NSGO_RESOURCE_TAG_ERROR_CODE, message: NSGO_RESOURCE_TAG_ERROR_MESSAGE}

	ControllerOperationNotImplementedError = Error{code: NSGO_CONTROLLER_OPERATION_NOTIMPLEMENTED_ERROR_CODE, message: NSGO_CONTROLLER_OPERATION_NOTIMPLEMENTED_ERROR_MESSAGE}

	ApiError = Error{code: NSGO_API_ERROR_CODE, message: NSGO_API_ERROR_MESSAGE}

	NSERR_SSL_NOCERT = Error{code: NSERR_SSL_NOCERT_ERROR_CODE, message: NSERR_SSL_NOCERT_ERROR_MESSAGE}
)

type Error struct {
	code    errorCode
	message string
	err     error
}

func (e Error) WithMessage(msg string) Error {
	e.message = msg
	return e
}

func (e Error) WithCode(code float64) Error {
	e.code = errorCode(code)
	return e
}

func (e Error) WithError(err error) Error {
	e.err = err
	return e
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Unwrap() error {
	return e.err
}

func (e Error) Is(err error) bool {
	t, ok := err.(Error)
	if !ok {
		return false
	}
	return e.code == t.code
}
