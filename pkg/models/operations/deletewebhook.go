package operations

type DeleteWebhookPathParams struct {
	WebhookID string `pathParam:"style=simple,explode=false,name=webhook-id"`
}

// DeleteWebhookMessageResponse
// message response
type DeleteWebhookMessageResponse struct {
	Message string `json:"message"`
}

type DeleteWebhookDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteWebhookRequest struct {
	PathParams DeleteWebhookPathParams
}

type DeleteWebhookResponse struct {
	ContentType                               string
	MessageResponse                           *DeleteWebhookMessageResponse
	StatusCode                                int64
	DeleteWebhookDefaultApplicationJSONObject *DeleteWebhookDefaultApplicationJSON
}
