package operations

import (
	"time"
)

type GetWebhooksScopeTypeEnum string

const (
	GetWebhooksScopeTypeEnumProject GetWebhooksScopeTypeEnum = "project"
)

type GetWebhooksQueryParams struct {
	ScopeID   string                   `queryParam:"style=form,explode=true,name=scope-id"`
	ScopeType GetWebhooksScopeTypeEnum `queryParam:"style=form,explode=true,name=scope-type"`
}

type GetWebhooks200ApplicationJSONWebhookEventsEnum string

const (
	GetWebhooks200ApplicationJSONWebhookEventsEnumWorkflowCompleted GetWebhooks200ApplicationJSONWebhookEventsEnum = "workflow-completed"
	GetWebhooks200ApplicationJSONWebhookEventsEnumJobCompleted      GetWebhooks200ApplicationJSONWebhookEventsEnum = "job-completed"
)

// GetWebhooks200ApplicationJSONWebhookScope
// The scope in which the relevant events that will trigger webhooks
type GetWebhooks200ApplicationJSONWebhookScope struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type GetWebhooks200ApplicationJSONWebhook struct {
	CreatedAt     time.Time                                        `json:"created-at"`
	Events        []GetWebhooks200ApplicationJSONWebhookEventsEnum `json:"events"`
	ID            string                                           `json:"id"`
	Name          string                                           `json:"name"`
	Scope         GetWebhooks200ApplicationJSONWebhookScope        `json:"scope"`
	SigningSecret string                                           `json:"signing-secret"`
	UpdatedAt     time.Time                                        `json:"updated-at"`
	URL           string                                           `json:"url"`
	VerifyTLS     bool                                             `json:"verify-tls"`
}

// GetWebhooks200ApplicationJSON
// A list of webhooks
type GetWebhooks200ApplicationJSON struct {
	Items         []GetWebhooks200ApplicationJSONWebhook `json:"items"`
	NextPageToken string                                 `json:"next_page_token"`
}

type GetWebhooksDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetWebhooksRequest struct {
	QueryParams GetWebhooksQueryParams
}

type GetWebhooksResponse struct {
	ContentType                             string
	StatusCode                              int64
	GetWebhooks200ApplicationJSONObject     *GetWebhooks200ApplicationJSON
	GetWebhooksDefaultApplicationJSONObject *GetWebhooksDefaultApplicationJSON
}
