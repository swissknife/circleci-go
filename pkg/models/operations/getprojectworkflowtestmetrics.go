package operations

type GetProjectWorkflowTestMetricsPathParams struct {
	ProjectSlug  string `pathParam:"style=simple,explode=false,name=project-slug"`
	WorkflowName string `pathParam:"style=simple,explode=false,name=workflow-name"`
}

type GetProjectWorkflowTestMetricsQueryParams struct {
	AllBranches *bool   `queryParam:"style=form,explode=true,name=all-branches"`
	Branch      *string `queryParam:"style=form,explode=true,name=branch"`
}

type GetProjectWorkflowTestMetrics200ApplicationJSONMostFailedTests struct {
	Classname   string  `json:"classname"`
	FailedRuns  int64   `json:"failed_runs"`
	File        string  `json:"file"`
	Flaky       bool    `json:"flaky"`
	JobName     string  `json:"job_name"`
	P95Duration float64 `json:"p95_duration"`
	Source      string  `json:"source"`
	TestName    string  `json:"test_name"`
	TotalRuns   int64   `json:"total_runs"`
}

type GetProjectWorkflowTestMetrics200ApplicationJSONSlowestTests struct {
	Classname   string  `json:"classname"`
	FailedRuns  int64   `json:"failed_runs"`
	File        string  `json:"file"`
	Flaky       bool    `json:"flaky"`
	JobName     string  `json:"job_name"`
	P95Duration float64 `json:"p95_duration"`
	Source      string  `json:"source"`
	TestName    string  `json:"test_name"`
	TotalRuns   int64   `json:"total_runs"`
}

// GetProjectWorkflowTestMetrics200ApplicationJSONTestRunsTestCounts
// Test counts for a given pipeline number
type GetProjectWorkflowTestMetrics200ApplicationJSONTestRunsTestCounts struct {
	Error   int64 `json:"error"`
	Failure int64 `json:"failure"`
	Skipped int64 `json:"skipped"`
	Success int64 `json:"success"`
	Total   int64 `json:"total"`
}

type GetProjectWorkflowTestMetrics200ApplicationJSONTestRuns struct {
	PipelineNumber int64                                                             `json:"pipeline_number"`
	SuccessRate    float32                                                           `json:"success_rate"`
	TestCounts     GetProjectWorkflowTestMetrics200ApplicationJSONTestRunsTestCounts `json:"test_counts"`
	WorkflowID     interface{}                                                       `json:"workflow_id"`
}

// GetProjectWorkflowTestMetrics200ApplicationJSON
// Project level test metrics response
type GetProjectWorkflowTestMetrics200ApplicationJSON struct {
	AverageTestCount     int64                                                            `json:"average_test_count"`
	MostFailedTests      []GetProjectWorkflowTestMetrics200ApplicationJSONMostFailedTests `json:"most_failed_tests"`
	MostFailedTestsExtra int64                                                            `json:"most_failed_tests_extra"`
	SlowestTests         []GetProjectWorkflowTestMetrics200ApplicationJSONSlowestTests    `json:"slowest_tests"`
	SlowestTestsExtra    int64                                                            `json:"slowest_tests_extra"`
	TestRuns             []GetProjectWorkflowTestMetrics200ApplicationJSONTestRuns        `json:"test_runs"`
	TotalTestRuns        int64                                                            `json:"total_test_runs"`
}

type GetProjectWorkflowTestMetricsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetProjectWorkflowTestMetricsRequest struct {
	PathParams  GetProjectWorkflowTestMetricsPathParams
	QueryParams GetProjectWorkflowTestMetricsQueryParams
}

type GetProjectWorkflowTestMetricsResponse struct {
	ContentType                                               string
	StatusCode                                                int64
	GetProjectWorkflowTestMetrics200ApplicationJSONObject     *GetProjectWorkflowTestMetrics200ApplicationJSON
	GetProjectWorkflowTestMetricsDefaultApplicationJSONObject *GetProjectWorkflowTestMetricsDefaultApplicationJSON
}
