package operations

import (
	"time"
)

type GetWorkflowSummaryPathParams struct {
	ProjectSlug  string `pathParam:"style=simple,explode=false,name=project-slug"`
	WorkflowName string `pathParam:"style=simple,explode=false,name=workflow-name"`
}

type GetWorkflowSummaryQueryParams struct {
	AllBranches *bool                  `queryParam:"style=form,explode=true,name=all-branches"`
	Branches    map[string]interface{} `queryParam:"style=form,explode=true,name=branches"`
}

// GetWorkflowSummary200ApplicationJSONMetricsDurationMetrics
// Metrics relating to the duration of runs for a workflow.
type GetWorkflowSummary200ApplicationJSONMetricsDurationMetrics struct {
	Max               int64   `json:"max"`
	Mean              int64   `json:"mean"`
	Median            int64   `json:"median"`
	Min               int64   `json:"min"`
	P95               int64   `json:"p95"`
	StandardDeviation float32 `json:"standard_deviation"`
}

// GetWorkflowSummary200ApplicationJSONMetrics
// Metrics aggregated across a workflow for a given time window.
type GetWorkflowSummary200ApplicationJSONMetrics struct {
	DurationMetrics  GetWorkflowSummary200ApplicationJSONMetricsDurationMetrics `json:"duration_metrics"`
	FailedRuns       int64                                                      `json:"failed_runs"`
	Mttr             int64                                                      `json:"mttr"`
	SuccessRate      float32                                                    `json:"success_rate"`
	SuccessfulRuns   int64                                                      `json:"successful_runs"`
	Throughput       float32                                                    `json:"throughput"`
	TotalCreditsUsed int64                                                      `json:"total_credits_used"`
	TotalRuns        int64                                                      `json:"total_runs"`
	WindowEnd        time.Time                                                  `json:"window_end"`
	WindowStart      time.Time                                                  `json:"window_start"`
}

// GetWorkflowSummary200ApplicationJSONTrends
// Trends for aggregated metrics across a workflow for a given time window.
type GetWorkflowSummary200ApplicationJSONTrends struct {
	FailedRuns         float32 `json:"failed_runs"`
	MedianDurationSecs float32 `json:"median_duration_secs"`
	Mttr               float32 `json:"mttr"`
	P95DurationSecs    float32 `json:"p95_duration_secs"`
	SuccessRate        float32 `json:"success_rate"`
	Throughput         float32 `json:"throughput"`
	TotalCreditsUsed   float32 `json:"total_credits_used"`
	TotalRuns          float32 `json:"total_runs"`
}

// GetWorkflowSummary200ApplicationJSON
// Workflow level aggregated metrics and trends response
type GetWorkflowSummary200ApplicationJSON struct {
	Metrics       GetWorkflowSummary200ApplicationJSONMetrics `json:"metrics"`
	Trends        GetWorkflowSummary200ApplicationJSONTrends  `json:"trends"`
	WorkflowNames []string                                    `json:"workflow_names"`
}

type GetWorkflowSummaryDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetWorkflowSummaryRequest struct {
	PathParams  GetWorkflowSummaryPathParams
	QueryParams GetWorkflowSummaryQueryParams
}

type GetWorkflowSummaryResponse struct {
	ContentType                                    string
	StatusCode                                     int64
	GetWorkflowSummary200ApplicationJSONObject     *GetWorkflowSummary200ApplicationJSON
	GetWorkflowSummaryDefaultApplicationJSONObject *GetWorkflowSummaryDefaultApplicationJSON
}
