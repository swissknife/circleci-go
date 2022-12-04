package operations

import (
	"time"
)

type ListContextsOwnerTypeEnum string

const (
	ListContextsOwnerTypeEnumAccount      ListContextsOwnerTypeEnum = "account"
	ListContextsOwnerTypeEnumOrganization ListContextsOwnerTypeEnum = "organization"
)

type ListContextsQueryParams struct {
	OwnerID   *string                    `queryParam:"style=form,explode=true,name=owner-id"`
	OwnerSlug *string                    `queryParam:"style=form,explode=true,name=owner-slug"`
	OwnerType *ListContextsOwnerTypeEnum `queryParam:"style=form,explode=true,name=owner-type"`
	PageToken *string                    `queryParam:"style=form,explode=true,name=page-token"`
}

type ListContexts200ApplicationJSONContext struct {
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
}

type ListContexts200ApplicationJSON struct {
	Items         []ListContexts200ApplicationJSONContext `json:"items"`
	NextPageToken string                                  `json:"next_page_token"`
}

type ListContextsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListContextsRequest struct {
	QueryParams ListContextsQueryParams
}

type ListContextsResponse struct {
	ContentType                              string
	StatusCode                               int64
	ListContexts200ApplicationJSONObject     *ListContexts200ApplicationJSON
	ListContextsDefaultApplicationJSONObject *ListContextsDefaultApplicationJSON
}
