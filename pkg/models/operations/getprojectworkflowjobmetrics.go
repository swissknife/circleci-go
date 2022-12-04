package operations

import (
	"time"
)

type GetProjectWorkflowJobMetricsPathParams struct {
	ProjectSlug  string `pathParam:"style=simple,explode=false,name=project-slug"`
	WorkflowName string `pathParam:"style=simple,explode=false,name=workflow-name"`
}

type GetProjectWorkflowJobMetricsReportingWindowEnum string

const (
	GetProjectWorkflowJobMetricsReportingWindowEnumLast7Days   GetProjectWorkflowJobMetricsReportingWindowEnum = "last-7-days"
	GetProjectWorkflowJobMetricsReportingWindowEnumLast90Days  GetProjectWorkflowJobMetricsReportingWindowEnum = "last-90-days"
	GetProjectWorkflowJobMetricsReportingWindowEnumLast24Hours GetProjectWorkflowJobMetricsReportingWindowEnum = "last-24-hours"
	GetProjectWorkflowJobMetricsReportingWindowEnumLast30Days  GetProjectWorkflowJobMetricsReportingWindowEnum = "last-30-days"
	GetProjectWorkflowJobMetricsReportingWindowEnumLast60Days  GetProjectWorkflowJobMetricsReportingWindowEnum = "last-60-days"
)

type GetProjectWorkflowJobMetricsQueryParams struct {
	AllBranches     *bool                                            `queryParam:"style=form,explode=true,name=all-branches"`
	Branch          *string                                          `queryParam:"style=form,explode=true,name=branch"`
	PageToken       *string                                          `queryParam:"style=form,explode=true,name=page-token"`
	ReportingWindow *GetProjectWorkflowJobMetricsReportingWindowEnum `queryParam:"style=form,explode=true,name=reporting-window"`
}

// GetProjectWorkflowJobMetrics200ApplicationJSONItemsMetricsDurationMetrics
// Metrics relating to the duration of runs for a workflow job.
type GetProjectWorkflowJobMetrics200ApplicationJSONItemsMetricsDurationMetrics struct {
	Max               int64   `json:"max"`
	Mean              int64   `json:"mean"`
	Median            int64   `json:"median"`
	Min               int64   `json:"min"`
	P95               int64   `json:"p95"`
	StandardDeviation float32 `json:"standard_deviation"`
}

// GetProjectWorkflowJobMetrics200ApplicationJSONItemsMetrics
// Metrics relating to a workflow job's runs.
type GetProjectWorkflowJobMetrics200ApplicationJSONItemsMetrics struct {
	DurationMetrics  GetProjectWorkflowJobMetrics200ApplicationJSONItemsMetricsDurationMetrics `json:"duration_metrics"`
	FailedRuns       int64                                                                     `json:"failed_runs"`
	SuccessRate      float32                                                                   `json:"success_rate"`
	SuccessfulRuns   int64                                                                     `json:"successful_runs"`
	Throughput       float32                                                                   `json:"throughput"`
	TotalCreditsUsed int64                                                                     `json:"total_credits_used"`
	TotalRuns        int64                                                                     `json:"total_runs"`
}

type GetProjectWorkflowJobMetrics200ApplicationJSONItems struct {
	Metrics     GetProjectWorkflowJobMetrics200ApplicationJSONItemsMetrics `json:"metrics"`
	Name        string                                                     `json:"name"`
	WindowEnd   time.Time                                                  `json:"window_end"`
	WindowStart time.Time                                                  `json:"window_start"`
}

// GetProjectWorkflowJobMetrics200ApplicationJSON
// Paginated workflow job summary metrics.
type GetProjectWorkflowJobMetrics200ApplicationJSON struct {
	Items         []GetProjectWorkflowJobMetrics200ApplicationJSONItems `json:"items"`
	NextPageToken string                                                `json:"next_page_token"`
}

type GetProjectWorkflowJobMetricsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetProjectWorkflowJobMetricsRequest struct {
	PathParams  GetProjectWorkflowJobMetricsPathParams
	QueryParams GetProjectWorkflowJobMetricsQueryParams
}

type GetProjectWorkflowJobMetricsResponse struct {
	ContentType                                              string
	StatusCode                                               int64
	GetProjectWorkflowJobMetrics200ApplicationJSONObject     *GetProjectWorkflowJobMetrics200ApplicationJSON
	GetProjectWorkflowJobMetricsDefaultApplicationJSONObject *GetProjectWorkflowJobMetricsDefaultApplicationJSON
}
