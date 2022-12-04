package operations

type ListEnvVarsPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type ListEnvVarsEnvironmentVariableListResponseEnvironmentVariablePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ListEnvVarsEnvironmentVariableListResponse struct {
	Items         []ListEnvVarsEnvironmentVariableListResponseEnvironmentVariablePair `json:"items"`
	NextPageToken string                                                              `json:"next_page_token"`
}

type ListEnvVarsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListEnvVarsRequest struct {
	PathParams ListEnvVarsPathParams
}

type ListEnvVarsResponse struct {
	ContentType                             string
	EnvironmentVariableListResponse         *ListEnvVarsEnvironmentVariableListResponse
	StatusCode                              int64
	ListEnvVarsDefaultApplicationJSONObject *ListEnvVarsDefaultApplicationJSON
}
