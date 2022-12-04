package operations

type DeleteCheckoutKeyPathParams struct {
	Fingerprint string `pathParam:"style=simple,explode=false,name=fingerprint"`
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

// DeleteCheckoutKeyMessageResponse
// message response
type DeleteCheckoutKeyMessageResponse struct {
	Message string `json:"message"`
}

type DeleteCheckoutKeyDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteCheckoutKeyRequest struct {
	PathParams DeleteCheckoutKeyPathParams
}

type DeleteCheckoutKeyResponse struct {
	ContentType                                   string
	MessageResponse                               *DeleteCheckoutKeyMessageResponse
	StatusCode                                    int64
	DeleteCheckoutKeyDefaultApplicationJSONObject *DeleteCheckoutKeyDefaultApplicationJSON
}
