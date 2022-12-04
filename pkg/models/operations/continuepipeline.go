package operations

type ContinuePipelineRequestBody struct {
	Configuration   string                 `json:"configuration"`
	ContinuationKey string                 `json:"continuation-key"`
	Parameters      map[string]interface{} `json:"parameters,omitempty"`
}

// ContinuePipelineMessageResponse
// message response
type ContinuePipelineMessageResponse struct {
	Message string `json:"message"`
}

type ContinuePipelineDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ContinuePipelineRequest struct {
	Request *ContinuePipelineRequestBody `request:"mediaType=application/json"`
}

type ContinuePipelineResponse struct {
	ContentType                                  string
	MessageResponse                              *ContinuePipelineMessageResponse
	StatusCode                                   int64
	ContinuePipelineDefaultApplicationJSONObject *ContinuePipelineDefaultApplicationJSON
}
