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

type errorCode float64

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
	NSGO_CLIENT_SAVECONFIG_ERROR_CODE                      errorCode = 10000007
	NSGO_CLIENT_SAVECONFIG_ERROR_MESSAGE                   string    = "Error saving config"
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
	NSGO_RESOURCE_EXTRACTION_ERROR_CODE                    errorCode = 20000006
	NSGO_RESOURCE_EXTRACTION_ERROR_MESSAGE                 string    = "Data extraction error"
	NSGO_RESOURCE_TAG_ERROR_CODE                           errorCode = 20000007
	NSGO_RESOURCE_TAG_ERROR_MESSAGE                        string    = "Tag error"
	NSGO_CONTROLLER_ERROR_CODE                             errorCode = 30000000
	NSGO_CONTROLLER_ERROR_MESSAGE                          string    = "Controller Error"
	NSGO_CONTROLLER_OPERATION_NOTIMPLEMENTED_ERROR_CODE    errorCode = 30000001
	NSGO_CONTROLLER_OPERATION_NOTIMPLEMENTED_ERROR_MESSAGE string    = "Operation not implemented"
	NSGO_API_ERROR_CODE                                    errorCode = 40000000
	NSGO_API_ERROR_MESSAGE                                 string    = "Nitro API specific error"
)
