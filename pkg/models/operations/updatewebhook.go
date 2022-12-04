package operations

import (
	"time"
)

type UpdateWebhookPathParams struct {
	WebhookID string `pathParam:"style=simple,explode=false,name=webhook-id"`
}

type UpdateWebhookRequestBodyEventsEnum string

const (
	UpdateWebhookRequestBodyEventsEnumWorkflowCompleted UpdateWebhookRequestBodyEventsEnum = "workflow-completed"
	UpdateWebhookRequestBodyEventsEnumJobCompleted      UpdateWebhookRequestBodyEventsEnum = "job-completed"
)

// UpdateWebhookRequestBody
// The parameters for an update webhook request
type UpdateWebhookRequestBody struct {
	Events        []UpdateWebhookRequestBodyEventsEnum `json:"events,omitempty"`
	Name          *string                              `json:"name,omitempty"`
	SigningSecret *string                              `json:"signing-secret,omitempty"`
	URL           *string                              `json:"url,omitempty"`
	VerifyTLS     *bool                                `json:"verify-tls,omitempty"`
}

type UpdateWebhookWebhookEventsEnum string

const (
	UpdateWebhookWebhookEventsEnumWorkflowCompleted UpdateWebhookWebhookEventsEnum = "workflow-completed"
	UpdateWebhookWebhookEventsEnumJobCompleted      UpdateWebhookWebhookEventsEnum = "job-completed"
)

// UpdateWebhookWebhookScope
// The scope in which the relevant events that will trigger webhooks
type UpdateWebhookWebhookScope struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type UpdateWebhookWebhook struct {
	CreatedAt     time.Time                        `json:"created-at"`
	Events        []UpdateWebhookWebhookEventsEnum `json:"events"`
	ID            string                           `json:"id"`
	Name          string                           `json:"name"`
	Scope         UpdateWebhookWebhookScope        `json:"scope"`
	SigningSecret string                           `json:"signing-secret"`
	UpdatedAt     time.Time                        `json:"updated-at"`
	URL           string                           `json:"url"`
	VerifyTLS     bool                             `json:"verify-tls"`
}

type UpdateWebhookDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateWebhookRequest struct {
	PathParams UpdateWebhookPathParams
	Request    *UpdateWebhookRequestBody `request:"mediaType=application/json"`
}

type UpdateWebhookResponse struct {
	ContentType                               string
	StatusCode                                int64
	Webhook                                   *UpdateWebhookWebhook
	UpdateWebhookDefaultApplicationJSONObject *UpdateWebhookDefaultApplicationJSON
}
