package operations

type DeleteContextPathParams struct {
	ContextID string `pathParam:"style=simple,explode=false,name=context-id"`
}

// DeleteContextMessageResponse
// message response
type DeleteContextMessageResponse struct {
	Message string `json:"message"`
}

type DeleteContextDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteContextRequest struct {
	PathParams DeleteContextPathParams
}

type DeleteContextResponse struct {
	ContentType                               string
	MessageResponse                           *DeleteContextMessageResponse
	StatusCode                                int64
	DeleteContextDefaultApplicationJSONObject *DeleteContextDefaultApplicationJSON
}
