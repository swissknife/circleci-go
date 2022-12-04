package operations

import (
	"time"
)

type CreateWebhookRequestBodyEventsEnum string

const (
	CreateWebhookRequestBodyEventsEnumWorkflowCompleted CreateWebhookRequestBodyEventsEnum = "workflow-completed"
	CreateWebhookRequestBodyEventsEnumJobCompleted      CreateWebhookRequestBodyEventsEnum = "job-completed"
)

type CreateWebhookRequestBodyScopeTypeEnum string

const (
	CreateWebhookRequestBodyScopeTypeEnumProject CreateWebhookRequestBodyScopeTypeEnum = "project"
)

// CreateWebhookRequestBodyScope
// The scope in which the relevant events that will trigger webhooks
type CreateWebhookRequestBodyScope struct {
	ID   string                                `json:"id"`
	Type CreateWebhookRequestBodyScopeTypeEnum `json:"type"`
}

// CreateWebhookRequestBody
// The parameters for a create webhook request
type CreateWebhookRequestBody struct {
	Events        []CreateWebhookRequestBodyEventsEnum `json:"events"`
	Name          string                               `json:"name"`
	Scope         CreateWebhookRequestBodyScope        `json:"scope"`
	SigningSecret string                               `json:"signing-secret"`
	URL           string                               `json:"url"`
	VerifyTLS     bool                                 `json:"verify-tls"`
}

type CreateWebhookWebhookEventsEnum string

const (
	CreateWebhookWebhookEventsEnumWorkflowCompleted CreateWebhookWebhookEventsEnum = "workflow-completed"
	CreateWebhookWebhookEventsEnumJobCompleted      CreateWebhookWebhookEventsEnum = "job-completed"
)

// CreateWebhookWebhookScope
// The scope in which the relevant events that will trigger webhooks
type CreateWebhookWebhookScope struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type CreateWebhookWebhook struct {
	CreatedAt     time.Time                        `json:"created-at"`
	Events        []CreateWebhookWebhookEventsEnum `json:"events"`
	ID            string                           `json:"id"`
	Name          string                           `json:"name"`
	Scope         CreateWebhookWebhookScope        `json:"scope"`
	SigningSecret string                           `json:"signing-secret"`
	UpdatedAt     time.Time                        `json:"updated-at"`
	URL           string                           `json:"url"`
	VerifyTLS     bool                             `json:"verify-tls"`
}

type CreateWebhookDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateWebhookRequest struct {
	Request *CreateWebhookRequestBody `request:"mediaType=application/json"`
}

type CreateWebhookResponse struct {
	ContentType                               string
	StatusCode                                int64
	Webhook                                   *CreateWebhookWebhook
	CreateWebhookDefaultApplicationJSONObject *CreateWebhookDefaultApplicationJSON
}
