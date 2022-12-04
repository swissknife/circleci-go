package operations

import (
	"time"
)

type GetWebhookByIDPathParams struct {
	WebhookID string `pathParam:"style=simple,explode=false,name=webhook-id"`
}

type GetWebhookByIDWebhookEventsEnum string

const (
	GetWebhookByIDWebhookEventsEnumWorkflowCompleted GetWebhookByIDWebhookEventsEnum = "workflow-completed"
	GetWebhookByIDWebhookEventsEnumJobCompleted      GetWebhookByIDWebhookEventsEnum = "job-completed"
)

// GetWebhookByIDWebhookScope
// The scope in which the relevant events that will trigger webhooks
type GetWebhookByIDWebhookScope struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type GetWebhookByIDWebhook struct {
	CreatedAt     time.Time                         `json:"created-at"`
	Events        []GetWebhookByIDWebhookEventsEnum `json:"events"`
	ID            string                            `json:"id"`
	Name          string                            `json:"name"`
	Scope         GetWebhookByIDWebhookScope        `json:"scope"`
	SigningSecret string                            `json:"signing-secret"`
	UpdatedAt     time.Time                         `json:"updated-at"`
	URL           string                            `json:"url"`
	VerifyTLS     bool                              `json:"verify-tls"`
}

type GetWebhookByIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetWebhookByIDRequest struct {
	PathParams GetWebhookByIDPathParams
}

type GetWebhookByIDResponse struct {
	ContentType                                string
	StatusCode                                 int64
	Webhook                                    *GetWebhookByIDWebhook
	GetWebhookByIDDefaultApplicationJSONObject *GetWebhookByIDDefaultApplicationJSON
}
