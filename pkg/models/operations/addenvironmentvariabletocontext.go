package operations

import (
	"time"
)

type AddEnvironmentVariableToContextPathParams struct {
	ContextID  string `pathParam:"style=simple,explode=false,name=context-id"`
	EnvVarName string `pathParam:"style=simple,explode=false,name=env-var-name"`
}

type AddEnvironmentVariableToContextRequestBody struct {
	Value string `json:"value"`
}

type AddEnvironmentVariableToContext200ApplicationJSON1 struct {
	ContextID string    `json:"context_id"`
	CreatedAt time.Time `json:"created_at"`
	Variable  string    `json:"variable"`
}

// AddEnvironmentVariableToContext200ApplicationJSONMessageResponse
// message response
type AddEnvironmentVariableToContext200ApplicationJSONMessageResponse struct {
	Message string `json:"message"`
}

type AddEnvironmentVariableToContextDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type AddEnvironmentVariableToContextRequest struct {
	PathParams AddEnvironmentVariableToContextPathParams
	Request    *AddEnvironmentVariableToContextRequestBody `request:"mediaType=application/json"`
}

type AddEnvironmentVariableToContextResponse struct {
	ContentType                                                 string
	StatusCode                                                  int64
	AddEnvironmentVariableToContext200ApplicationJSONAnyOf      *interface{}
	AddEnvironmentVariableToContextDefaultApplicationJSONObject *AddEnvironmentVariableToContextDefaultApplicationJSON
}
