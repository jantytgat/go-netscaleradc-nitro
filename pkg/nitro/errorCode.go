package nitro

type errorCode float64

const (
	NSGO_CLIENT_ERROR_CODE                       errorCode = 10000000
	NSGO_CLIENT_ERROR_MESSAGE                    string    = "client error"
	NSGO_CLIENT_CREATE_ERROR_CODE                errorCode = 10000001
	NSGO_CLIENT_CREATE_ERROR_MESSAGE             string    = "error creating client"
	NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_CODE     errorCode = 10000002
	NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE  string    = "error creating http request"
	NSGO_CLIENT_EXECUTEREQUEST_ERROR_CODE        errorCode = 10000003
	NSGO_CLIENT_EXECUTEREQUEST_ERROR_MESSAGE     string    = "error executing nitro request"
	NSGO_CLIENT_LOGIN_ERROR_CODE                 errorCode = 10000004
	NSGO_CLIENT_LOGIN_ERROR_MESSAGE              string    = "login error"
	NSGO_CLIENT_LOGOUT_ERROR_CODE                errorCode = 10000005
	NSGO_CLIENT_LOGOUT_ERROR_MESSAGE             string    = "logout error"
	NSGO_CLIENT_CONNECTIONSETTINGS_ERROR_CODE    errorCode = 10000006
	NSGO_CLIENT_CONNECTIONSETTINGS_ERROR_MESSAGE string    = "error in connection settings"
	NSGO_CLIENT_SAVECONFIG_ERROR_CODE            errorCode = 10000007
	NSGO_CLIENT_SAVECONFIG_ERROR_MESSAGE         string    = "error saving config"
	NSGO_RESOURCE_ERROR_CODE                     errorCode = 40000000
	NSGO_RESOURCE_ERROR_MESSAGE                  string    = "resource error"
	NSGO_RESOURCE_VALIDATION_ERROR_CODE          errorCode = 20000001
	NSGO_RESOURCE_VALIDATION_ERROR_MESSAGE       string    = "validation error"
	NSGO_RESOURCE_INVALIDTYPE_ERROR_CODE         errorCode = 20000002
	NSGO_RESOURCE_INVALIDTYPE_ERROR_MESSAGE      string    = "invalid resource type"
	NSGO_RESOURCE_INVALIDFIELD_ERROR_CODE        errorCode = 20000003
	NSGO_RESOURCE_INVALIDFIELD_ERROR_MESSAGE     string    = "invalid field"
	NSGO_RESOURCE_SERIALIZATION_ERROR_CODE       errorCode = 20000004
	NSGO_RESOURCE_SERIALIZATION_ERROR_MESSAGE    string    = "serialization error"
	NSGO_RESOURCE_DESERIALIZATION_ERROR_CODE     errorCode = 20000005
	NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE  string    = "deserialization error"
	NSGO_RESOURCE_EXTRACTION_ERROR_CODE          errorCode = 20000006
	NSGO_RESOURCE_EXTRACTION_ERROR_MESSAGE       string    = "data extraction error"
	NSGO_RESOURCE_TAG_ERROR_CODE                 errorCode = 20000007
	NSGO_RESOURCE_TAG_ERROR_MESSAGE              string    = "tag error"
	NSGO_ERROR_CODE                              errorCode = 30000000
	NSGO_ERROR_MESSAGE                           string    = "controller Error"
	NSGO_OPERATION_NOTIMPLEMENTED_ERROR_CODE     errorCode = 30000001
	NSGO_OPERATION_NOTIMPLEMENTED_ERROR_MESSAGE  string    = "operation not implemented"
	NSGO_API_ERROR_CODE                          errorCode = 40000000
	NSGO_API_ERROR_MESSAGE                       string    = "nitro api specific error"
)
