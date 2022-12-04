package operations

type CreateEnvVarPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type CreateEnvVarEnvironmentVariablePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CreateEnvVarDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateEnvVarRequest struct {
	PathParams CreateEnvVarPathParams
	Request    *CreateEnvVarEnvironmentVariablePair `request:"mediaType=application/json"`
}

type CreateEnvVarResponse struct {
	ContentType                              string
	EnvironmentVariablePair                  *CreateEnvVarEnvironmentVariablePair
	StatusCode                               int64
	CreateEnvVarDefaultApplicationJSONObject *CreateEnvVarDefaultApplicationJSON
}
