package operations

type DeleteEnvVarPathParams struct {
	Name        string `pathParam:"style=simple,explode=false,name=name"`
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

// DeleteEnvVarMessageResponse
// message response
type DeleteEnvVarMessageResponse struct {
	Message string `json:"message"`
}

type DeleteEnvVarDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteEnvVarRequest struct {
	PathParams DeleteEnvVarPathParams
}

type DeleteEnvVarResponse struct {
	ContentType                              string
	MessageResponse                          *DeleteEnvVarMessageResponse
	StatusCode                               int64
	DeleteEnvVarDefaultApplicationJSONObject *DeleteEnvVarDefaultApplicationJSON
}
