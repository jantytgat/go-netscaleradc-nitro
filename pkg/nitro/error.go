package nitro

import "errors"

var (
	ClientCreateError             = Error{code: NSGO_CLIENT_CREATE_ERROR_CODE, message: NSGO_CLIENT_CREATE_ERROR_MESSAGE}
	ClientCreateHttpRequestError  = Error{code: NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_CODE, message: NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE}
	ClientExecuteRequestError     = Error{code: NSGO_CLIENT_EXECUTEREQUEST_ERROR_CODE, message: NSGO_CLIENT_EXECUTEREQUEST_ERROR_MESSAGE}
	ClientLoginError              = Error{code: NSGO_CLIENT_LOGIN_ERROR_CODE, message: NSGO_CLIENT_LOGIN_ERROR_MESSAGE}
	ClientLogoutError             = Error{code: NSGO_CLIENT_LOGOUT_ERROR_CODE, message: NSGO_CLIENT_LOGOUT_ERROR_MESSAGE}
	ClientSaveConfigError         = Error{code: NSGO_CLIENT_SAVECONFIG_ERROR_CODE, message: NSGO_CLIENT_SAVECONFIG_ERROR_MESSAGE}
	ClientConnectionSettingsError = Error{code: NSGO_CLIENT_CONNECTIONSETTINGS_ERROR_CODE, message: NSGO_CLIENT_CONNECTIONSETTINGS_ERROR_MESSAGE}

	ResourceValidationError      = Error{code: NSGO_RESOURCE_VALIDATION_ERROR_CODE, message: NSGO_RESOURCE_VALIDATION_ERROR_MESSAGE}
	ResourceInvalidTypeError     = Error{code: NSGO_RESOURCE_INVALIDTYPE_ERROR_CODE, message: NSGO_RESOURCE_INVALIDTYPE_ERROR_MESSAGE}
	ResourceInvalidFieldError    = Error{code: NSGO_RESOURCE_INVALIDFIELD_ERROR_CODE, message: NSGO_RESOURCE_INVALIDFIELD_ERROR_MESSAGE}
	ResourceSerializationError   = Error{code: NSGO_RESOURCE_SERIALIZATION_ERROR_CODE, message: NSGO_RESOURCE_SERIALIZATION_ERROR_MESSAGE}
	ResourceDeserializationError = Error{code: NSGO_RESOURCE_DESERIALIZATION_ERROR_CODE, message: NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE}
	ResourceExtractionError      = Error{code: NSGO_RESOURCE_EXTRACTION_ERROR_CODE, message: NSGO_RESOURCE_EXTRACTION_ERROR_MESSAGE}
	ResourceTagError             = Error{code: NSGO_RESOURCE_TAG_ERROR_CODE, message: NSGO_RESOURCE_TAG_ERROR_MESSAGE}

	OperationNotImplementedError = Error{code: NSGO_OPERATION_NOTIMPLEMENTED_ERROR_CODE, message: NSGO_OPERATION_NOTIMPLEMENTED_ERROR_MESSAGE}

	ApiError = Error{code: NSGO_API_ERROR_CODE, message: NSGO_API_ERROR_MESSAGE}
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
	if e.code >= NSGO_CLIENT_ERROR_CODE {
		return e.message
	}

	if e.err != nil {
		return NSGO_API_ERROR_MESSAGE + ": " + e.message + " - " + e.err.Error()
	}
	return NSGO_API_ERROR_MESSAGE + ": " + e.message
}

func (e Error) Unwrap() error {
	return e.err
}

func (e Error) Is(err error) bool {
	var t Error
	ok := errors.As(err, &t)
	if !ok {
		return false
	}
	return e.code == t.code
}
