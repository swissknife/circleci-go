package operations

import (
	"time"
)

type GetProjectWorkflowMetricsPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type GetProjectWorkflowMetricsReportingWindowEnum string

const (
	GetProjectWorkflowMetricsReportingWindowEnumLast7Days   GetProjectWorkflowMetricsReportingWindowEnum = "last-7-days"
	GetProjectWorkflowMetricsReportingWindowEnumLast90Days  GetProjectWorkflowMetricsReportingWindowEnum = "last-90-days"
	GetProjectWorkflowMetricsReportingWindowEnumLast24Hours GetProjectWorkflowMetricsReportingWindowEnum = "last-24-hours"
	GetProjectWorkflowMetricsReportingWindowEnumLast30Days  GetProjectWorkflowMetricsReportingWindowEnum = "last-30-days"
	GetProjectWorkflowMetricsReportingWindowEnumLast60Days  GetProjectWorkflowMetricsReportingWindowEnum = "last-60-days"
)

type GetProjectWorkflowMetricsQueryParams struct {
	AllBranches     *bool                                         `queryParam:"style=form,explode=true,name=all-branches"`
	Branch          *string                                       `queryParam:"style=form,explode=true,name=branch"`
	PageToken       *string                                       `queryParam:"style=form,explode=true,name=page-token"`
	ReportingWindow *GetProjectWorkflowMetricsReportingWindowEnum `queryParam:"style=form,explode=true,name=reporting-window"`
}

// GetProjectWorkflowMetrics200ApplicationJSONItemsMetricsDurationMetrics
// Metrics relating to the duration of runs for a workflow.
type GetProjectWorkflowMetrics200ApplicationJSONItemsMetricsDurationMetrics struct {
	Max               int64   `json:"max"`
	Mean              int64   `json:"mean"`
	Median            int64   `json:"median"`
	Min               int64   `json:"min"`
	P95               int64   `json:"p95"`
	StandardDeviation float32 `json:"standard_deviation"`
}

// GetProjectWorkflowMetrics200ApplicationJSONItemsMetrics
// Metrics relating to a workflow's runs.
type GetProjectWorkflowMetrics200ApplicationJSONItemsMetrics struct {
	DurationMetrics  GetProjectWorkflowMetrics200ApplicationJSONItemsMetricsDurationMetrics `json:"duration_metrics"`
	FailedRuns       int64                                                                  `json:"failed_runs"`
	Mttr             int64                                                                  `json:"mttr"`
	SuccessRate      float32                                                                `json:"success_rate"`
	SuccessfulRuns   int64                                                                  `json:"successful_runs"`
	Throughput       float32                                                                `json:"throughput"`
	TotalCreditsUsed int64                                                                  `json:"total_credits_used"`
	TotalRecoveries  int64                                                                  `json:"total_recoveries"`
	TotalRuns        int64                                                                  `json:"total_runs"`
}

type GetProjectWorkflowMetrics200ApplicationJSONItems struct {
	Metrics     GetProjectWorkflowMetrics200ApplicationJSONItemsMetrics `json:"metrics"`
	Name        string                                                  `json:"name"`
	WindowEnd   time.Time                                               `json:"window_end"`
	WindowStart time.Time                                               `json:"window_start"`
}

// GetProjectWorkflowMetrics200ApplicationJSON
// Paginated workflow summary metrics.
type GetProjectWorkflowMetrics200ApplicationJSON struct {
	Items         []GetProjectWorkflowMetrics200ApplicationJSONItems `json:"items"`
	NextPageToken string                                             `json:"next_page_token"`
}

type GetProjectWorkflowMetricsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetProjectWorkflowMetricsRequest struct {
	PathParams  GetProjectWorkflowMetricsPathParams
	QueryParams GetProjectWorkflowMetricsQueryParams
}

type GetProjectWorkflowMetricsResponse struct {
	ContentType                                           string
	StatusCode                                            int64
	GetProjectWorkflowMetrics200ApplicationJSONObject     *GetProjectWorkflowMetrics200ApplicationJSON
	GetProjectWorkflowMetricsDefaultApplicationJSONObject *GetProjectWorkflowMetricsDefaultApplicationJSON
}
