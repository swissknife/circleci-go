package operations

import (
	"time"
)

type GetJobTimeseriesPathParams struct {
	ProjectSlug  string `pathParam:"style=simple,explode=false,name=project-slug"`
	WorkflowName string `pathParam:"style=simple,explode=false,name=workflow-name"`
}

type GetJobTimeseriesGranularityEnum string

const (
	GetJobTimeseriesGranularityEnumDaily  GetJobTimeseriesGranularityEnum = "daily"
	GetJobTimeseriesGranularityEnumHourly GetJobTimeseriesGranularityEnum = "hourly"
)

type GetJobTimeseriesQueryParams struct {
	Branch      *string                          `queryParam:"style=form,explode=true,name=branch"`
	EndDate     *time.Time                       `queryParam:"style=form,explode=true,name=end-date"`
	Granularity *GetJobTimeseriesGranularityEnum `queryParam:"style=form,explode=true,name=granularity"`
	StartDate   *time.Time                       `queryParam:"style=form,explode=true,name=start-date"`
}

// GetJobTimeseries200ApplicationJSONItemsMetricsDurationMetrics
// Metrics relating to the duration of runs for a workflow.
type GetJobTimeseries200ApplicationJSONItemsMetricsDurationMetrics struct {
	Max    int64 `json:"max"`
	Median int64 `json:"median"`
	Min    int64 `json:"min"`
	P95    int64 `json:"p95"`
	Total  int64 `json:"total"`
}

// GetJobTimeseries200ApplicationJSONItemsMetrics
// Metrics relating to a workflow's runs.
type GetJobTimeseries200ApplicationJSONItemsMetrics struct {
	DurationMetrics   GetJobTimeseries200ApplicationJSONItemsMetricsDurationMetrics `json:"duration_metrics"`
	FailedRuns        int64                                                         `json:"failed_runs"`
	MedianCreditsUsed int64                                                         `json:"median_credits_used"`
	SuccessfulRuns    int64                                                         `json:"successful_runs"`
	Throughput        float32                                                       `json:"throughput"`
	TotalCreditsUsed  int64                                                         `json:"total_credits_used"`
	TotalRuns         int64                                                         `json:"total_runs"`
}

type GetJobTimeseries200ApplicationJSONItems struct {
	MaxEndedAt   time.Time                                      `json:"max_ended_at"`
	Metrics      GetJobTimeseries200ApplicationJSONItemsMetrics `json:"metrics"`
	MinStartedAt time.Time                                      `json:"min_started_at"`
	Name         string                                         `json:"name"`
	Timestamp    time.Time                                      `json:"timestamp"`
}

// GetJobTimeseries200ApplicationJSON
// Project level timeseries metrics response
type GetJobTimeseries200ApplicationJSON struct {
	Items         []GetJobTimeseries200ApplicationJSONItems `json:"items"`
	NextPageToken string                                    `json:"next_page_token"`
}

type GetJobTimeseriesDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetJobTimeseriesRequest struct {
	PathParams  GetJobTimeseriesPathParams
	QueryParams GetJobTimeseriesQueryParams
}

type GetJobTimeseriesResponse struct {
	ContentType                                  string
	StatusCode                                   int64
	GetJobTimeseries200ApplicationJSONObject     *GetJobTimeseries200ApplicationJSON
	GetJobTimeseriesDefaultApplicationJSONObject *GetJobTimeseriesDefaultApplicationJSON
}
