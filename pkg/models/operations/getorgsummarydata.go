package operations

type GetOrgSummaryDataPathParams struct {
	OrgSlug string `pathParam:"style=simple,explode=false,name=org-slug"`
}

type GetOrgSummaryDataReportingWindowEnum string

const (
	GetOrgSummaryDataReportingWindowEnumLast7Days   GetOrgSummaryDataReportingWindowEnum = "last-7-days"
	GetOrgSummaryDataReportingWindowEnumLast90Days  GetOrgSummaryDataReportingWindowEnum = "last-90-days"
	GetOrgSummaryDataReportingWindowEnumLast24Hours GetOrgSummaryDataReportingWindowEnum = "last-24-hours"
	GetOrgSummaryDataReportingWindowEnumLast30Days  GetOrgSummaryDataReportingWindowEnum = "last-30-days"
	GetOrgSummaryDataReportingWindowEnumLast60Days  GetOrgSummaryDataReportingWindowEnum = "last-60-days"
)

type GetOrgSummaryDataQueryParams struct {
	ProjectNames    map[string]interface{}                `queryParam:"style=form,explode=true,name=project-names"`
	ReportingWindow *GetOrgSummaryDataReportingWindowEnum `queryParam:"style=form,explode=true,name=reporting-window"`
}

// GetOrgSummaryData200ApplicationJSONOrgDataMetrics
// Metrics for a single org metrics.
type GetOrgSummaryData200ApplicationJSONOrgDataMetrics struct {
	SuccessRate       float32 `json:"success_rate"`
	Throughput        float32 `json:"throughput"`
	TotalCreditsUsed  int64   `json:"total_credits_used"`
	TotalDurationSecs int64   `json:"total_duration_secs"`
	TotalRuns         int64   `json:"total_runs"`
}

// GetOrgSummaryData200ApplicationJSONOrgDataTrends
// Trends for a single org.
type GetOrgSummaryData200ApplicationJSONOrgDataTrends struct {
	SuccessRate       float32 `json:"success_rate"`
	Throughput        float32 `json:"throughput"`
	TotalCreditsUsed  float32 `json:"total_credits_used"`
	TotalDurationSecs float32 `json:"total_duration_secs"`
	TotalRuns         float32 `json:"total_runs"`
}

// GetOrgSummaryData200ApplicationJSONOrgData
// Aggregated metrics for an org, with trends.
type GetOrgSummaryData200ApplicationJSONOrgData struct {
	Metrics GetOrgSummaryData200ApplicationJSONOrgDataMetrics `json:"metrics"`
	Trends  GetOrgSummaryData200ApplicationJSONOrgDataTrends  `json:"trends"`
}

// GetOrgSummaryData200ApplicationJSONOrgProjectDataMetrics
// Metrics for a single project, across all branches.
type GetOrgSummaryData200ApplicationJSONOrgProjectDataMetrics struct {
	SuccessRate       float32 `json:"success_rate"`
	TotalCreditsUsed  int64   `json:"total_credits_used"`
	TotalDurationSecs int64   `json:"total_duration_secs"`
	TotalRuns         int64   `json:"total_runs"`
}

// GetOrgSummaryData200ApplicationJSONOrgProjectDataTrends
// Trends for a single project, across all branches.
type GetOrgSummaryData200ApplicationJSONOrgProjectDataTrends struct {
	SuccessRate       float32 `json:"success_rate"`
	TotalCreditsUsed  float32 `json:"total_credits_used"`
	TotalDurationSecs float32 `json:"total_duration_secs"`
	TotalRuns         float32 `json:"total_runs"`
}

type GetOrgSummaryData200ApplicationJSONOrgProjectData struct {
	Metrics     GetOrgSummaryData200ApplicationJSONOrgProjectDataMetrics `json:"metrics"`
	ProjectName string                                                   `json:"project_name"`
	Trends      GetOrgSummaryData200ApplicationJSONOrgProjectDataTrends  `json:"trends"`
}

// GetOrgSummaryData200ApplicationJSON
// Summary metrics with trends for the entire org, and for each project.
type GetOrgSummaryData200ApplicationJSON struct {
	AllProjects    []string                                            `json:"all_projects"`
	OrgData        GetOrgSummaryData200ApplicationJSONOrgData          `json:"org_data"`
	OrgProjectData []GetOrgSummaryData200ApplicationJSONOrgProjectData `json:"org_project_data"`
}

type GetOrgSummaryDataDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetOrgSummaryDataRequest struct {
	PathParams  GetOrgSummaryDataPathParams
	QueryParams GetOrgSummaryDataQueryParams
}

type GetOrgSummaryDataResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	GetOrgSummaryData200ApplicationJSONObject     *GetOrgSummaryData200ApplicationJSON
	GetOrgSummaryDataDefaultApplicationJSONObject *GetOrgSummaryDataDefaultApplicationJSON
}
