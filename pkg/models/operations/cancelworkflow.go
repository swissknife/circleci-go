package operations

type CancelWorkflowPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

// CancelWorkflowMessageResponse
// message response
type CancelWorkflowMessageResponse struct {
	Message string `json:"message"`
}

type CancelWorkflowDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CancelWorkflowRequest struct {
	PathParams CancelWorkflowPathParams
}

type CancelWorkflowResponse struct {
	ContentType                                string
	MessageResponse                            *CancelWorkflowMessageResponse
	StatusCode                                 int64
	CancelWorkflowDefaultApplicationJSONObject *CancelWorkflowDefaultApplicationJSON
}
