package operations

type DeleteEnvironmentVariableFromContextPathParams struct {
	ContextID  string `pathParam:"style=simple,explode=false,name=context-id"`
	EnvVarName string `pathParam:"style=simple,explode=false,name=env-var-name"`
}

// DeleteEnvironmentVariableFromContextMessageResponse
// message response
type DeleteEnvironmentVariableFromContextMessageResponse struct {
	Message string `json:"message"`
}

type DeleteEnvironmentVariableFromContextDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteEnvironmentVariableFromContextRequest struct {
	PathParams DeleteEnvironmentVariableFromContextPathParams
}

type DeleteEnvironmentVariableFromContextResponse struct {
	ContentType                                                      string
	MessageResponse                                                  *DeleteEnvironmentVariableFromContextMessageResponse
	StatusCode                                                       int64
	DeleteEnvironmentVariableFromContextDefaultApplicationJSONObject *DeleteEnvironmentVariableFromContextDefaultApplicationJSON
}
