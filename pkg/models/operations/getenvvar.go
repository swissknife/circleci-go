package operations

type GetEnvVarPathParams struct {
	Name        string `pathParam:"style=simple,explode=false,name=name"`
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type GetEnvVarEnvironmentVariablePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetEnvVarDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetEnvVarRequest struct {
	PathParams GetEnvVarPathParams
}

type GetEnvVarResponse struct {
	ContentType                           string
	EnvironmentVariablePair               *GetEnvVarEnvironmentVariablePair
	StatusCode                            int64
	GetEnvVarDefaultApplicationJSONObject *GetEnvVarDefaultApplicationJSON
}
