package operations

type GetFlakyTestsPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type GetFlakyTests200ApplicationJSONFlakyTests struct {
	Classname         string                 `json:"classname"`
	File              string                 `json:"file"`
	JobName           string                 `json:"job-name"`
	JobNumber         map[string]interface{} `json:"job-number"`
	PipelineNumber    map[string]interface{} `json:"pipeline-number"`
	Source            string                 `json:"source"`
	TestName          string                 `json:"test-name"`
	TimeWasted        map[string]interface{} `json:"time-wasted,omitempty"`
	TimesFlaked       int64                  `json:"times-flaked"`
	WorkflowCreatedAt interface{}            `json:"workflow-created-at"`
	WorkflowID        interface{}            `json:"workflow-id"`
	WorkflowName      string                 `json:"workflow-name"`
}

// GetFlakyTests200ApplicationJSON
// Flaky tests response
type GetFlakyTests200ApplicationJSON struct {
	FlakyTests      []GetFlakyTests200ApplicationJSONFlakyTests `json:"flaky-tests"`
	TotalFlakyTests float64                                     `json:"total-flaky-tests"`
}

type GetFlakyTestsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetFlakyTestsRequest struct {
	PathParams GetFlakyTestsPathParams
}

type GetFlakyTestsResponse struct {
	ContentType                               string
	StatusCode                                int64
	GetFlakyTests200ApplicationJSONObject     *GetFlakyTests200ApplicationJSON
	GetFlakyTestsDefaultApplicationJSONObject *GetFlakyTestsDefaultApplicationJSON
}
