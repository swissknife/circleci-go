package operations

import (
	"time"
)

type ListEnvironmentVariablesFromContextPathParams struct {
	ContextID string `pathParam:"style=simple,explode=false,name=context-id"`
}

type ListEnvironmentVariablesFromContextQueryParams struct {
	PageToken *string `queryParam:"style=form,explode=true,name=page-token"`
}

type ListEnvironmentVariablesFromContext200ApplicationJSONItems struct {
	ContextID string    `json:"context_id"`
	CreatedAt time.Time `json:"created_at"`
	Variable  string    `json:"variable"`
}

type ListEnvironmentVariablesFromContext200ApplicationJSON struct {
	Items         []ListEnvironmentVariablesFromContext200ApplicationJSONItems `json:"items"`
	NextPageToken string                                                       `json:"next_page_token"`
}

type ListEnvironmentVariablesFromContextDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListEnvironmentVariablesFromContextRequest struct {
	PathParams  ListEnvironmentVariablesFromContextPathParams
	QueryParams ListEnvironmentVariablesFromContextQueryParams
}

type ListEnvironmentVariablesFromContextResponse struct {
	ContentType                                                     string
	StatusCode                                                      int64
	ListEnvironmentVariablesFromContext200ApplicationJSONObject     *ListEnvironmentVariablesFromContext200ApplicationJSON
	ListEnvironmentVariablesFromContextDefaultApplicationJSONObject *ListEnvironmentVariablesFromContextDefaultApplicationJSON
}
