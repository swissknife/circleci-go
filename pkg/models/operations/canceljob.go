package operations

type CancelJobPathParams struct {
	JobNumber   interface{} `pathParam:"style=simple,explode=false,name=job-number"`
	ProjectSlug string      `pathParam:"style=simple,explode=false,name=project-slug"`
}

// CancelJobMessageResponse
// message response
type CancelJobMessageResponse struct {
	Message string `json:"message"`
}

type CancelJobDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CancelJobRequest struct {
	PathParams CancelJobPathParams
}

type CancelJobResponse struct {
	ContentType                           string
	MessageResponse                       *CancelJobMessageResponse
	StatusCode                            int64
	CancelJobDefaultApplicationJSONObject *CancelJobDefaultApplicationJSON
}
