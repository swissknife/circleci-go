package operations

type GetProjectWorkflowsPageDataPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type GetProjectWorkflowsPageDataReportingWindowEnum string

const (
	GetProjectWorkflowsPageDataReportingWindowEnumLast7Days   GetProjectWorkflowsPageDataReportingWindowEnum = "last-7-days"
	GetProjectWorkflowsPageDataReportingWindowEnumLast90Days  GetProjectWorkflowsPageDataReportingWindowEnum = "last-90-days"
	GetProjectWorkflowsPageDataReportingWindowEnumLast24Hours GetProjectWorkflowsPageDataReportingWindowEnum = "last-24-hours"
	GetProjectWorkflowsPageDataReportingWindowEnumLast30Days  GetProjectWorkflowsPageDataReportingWindowEnum = "last-30-days"
	GetProjectWorkflowsPageDataReportingWindowEnumLast60Days  GetProjectWorkflowsPageDataReportingWindowEnum = "last-60-days"
)

type GetProjectWorkflowsPageDataQueryParams struct {
	Branches        map[string]interface{}                          `queryParam:"style=form,explode=true,name=branches"`
	ReportingWindow *GetProjectWorkflowsPageDataReportingWindowEnum `queryParam:"style=form,explode=true,name=reporting-window"`
	WorkflowNames   map[string]interface{}                          `queryParam:"style=form,explode=true,name=workflow-names"`
}

// GetProjectWorkflowsPageData200ApplicationJSONProjectDataMetrics
// Metrics aggregated across all workflows and branches for a project.
type GetProjectWorkflowsPageData200ApplicationJSONProjectDataMetrics struct {
	SuccessRate       float32 `json:"success_rate"`
	Throughput        float32 `json:"throughput"`
	TotalCreditsUsed  int64   `json:"total_credits_used"`
	TotalDurationSecs int64   `json:"total_duration_secs"`
	TotalRuns         int64   `json:"total_runs"`
}

// GetProjectWorkflowsPageData200ApplicationJSONProjectDataTrends
// Metric trends aggregated across all workflows and branches for a project.
type GetProjectWorkflowsPageData200ApplicationJSONProjectDataTrends struct {
	SuccessRate       float32 `json:"success_rate"`
	Throughput        float32 `json:"throughput"`
	TotalCreditsUsed  float32 `json:"total_credits_used"`
	TotalDurationSecs float32 `json:"total_duration_secs"`
	TotalRuns         float32 `json:"total_runs"`
}

// GetProjectWorkflowsPageData200ApplicationJSONProjectData
// Metrics and trends data aggregated for a given project.
type GetProjectWorkflowsPageData200ApplicationJSONProjectData struct {
	Metrics GetProjectWorkflowsPageData200ApplicationJSONProjectDataMetrics `json:"metrics"`
	Trends  GetProjectWorkflowsPageData200ApplicationJSONProjectDataTrends  `json:"trends"`
}

// GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowBranchDataMetrics
// Metrics aggregated across a workflow or branchfor a project.
type GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowBranchDataMetrics struct {
	P95DurationSecs  float32 `json:"p95_duration_secs"`
	SuccessRate      float32 `json:"success_rate"`
	TotalCreditsUsed int64   `json:"total_credits_used"`
	TotalRuns        int64   `json:"total_runs"`
}

// GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowBranchDataTrends
// Trends aggregated across a workflow or branch for a project.
type GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowBranchDataTrends struct {
	P95DurationSecs  float32 `json:"p95_duration_secs"`
	SuccessRate      float32 `json:"success_rate"`
	TotalCreditsUsed float32 `json:"total_credits_used"`
	TotalRuns        float32 `json:"total_runs"`
}

type GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowBranchData struct {
	Branch       string                                                                        `json:"branch"`
	Metrics      GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowBranchDataMetrics `json:"metrics"`
	Trends       GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowBranchDataTrends  `json:"trends"`
	WorkflowName string                                                                        `json:"workflow_name"`
}

// GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowDataMetrics
// Metrics aggregated across a workflow or branchfor a project.
type GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowDataMetrics struct {
	P95DurationSecs  float32 `json:"p95_duration_secs"`
	SuccessRate      float32 `json:"success_rate"`
	TotalCreditsUsed int64   `json:"total_credits_used"`
	TotalRuns        int64   `json:"total_runs"`
}

// GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowDataTrends
// Trends aggregated across a workflow or branch for a project.
type GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowDataTrends struct {
	P95DurationSecs  float32 `json:"p95_duration_secs"`
	SuccessRate      float32 `json:"success_rate"`
	TotalCreditsUsed float32 `json:"total_credits_used"`
	TotalRuns        float32 `json:"total_runs"`
}

type GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowData struct {
	Metrics      GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowDataMetrics `json:"metrics"`
	Trends       GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowDataTrends  `json:"trends"`
	WorkflowName string                                                                  `json:"workflow_name"`
}

type GetProjectWorkflowsPageData200ApplicationJSON struct {
	AllBranches               []string                                                                 `json:"all_branches,omitempty"`
	AllWorkflows              []string                                                                 `json:"all_workflows,omitempty"`
	OrgID                     *interface{}                                                             `json:"org_id,omitempty"`
	ProjectData               *GetProjectWorkflowsPageData200ApplicationJSONProjectData                `json:"project_data,omitempty"`
	ProjectID                 *interface{}                                                             `json:"project_id,omitempty"`
	ProjectWorkflowBranchData []GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowBranchData `json:"project_workflow_branch_data,omitempty"`
	ProjectWorkflowData       []GetProjectWorkflowsPageData200ApplicationJSONProjectWorkflowData       `json:"project_workflow_data,omitempty"`
}

type GetProjectWorkflowsPageDataDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetProjectWorkflowsPageDataRequest struct {
	PathParams  GetProjectWorkflowsPageDataPathParams
	QueryParams GetProjectWorkflowsPageDataQueryParams
}

type GetProjectWorkflowsPageDataResponse struct {
	ContentType                                             string
	StatusCode                                              int64
	GetProjectWorkflowsPageData200ApplicationJSONObject     *GetProjectWorkflowsPageData200ApplicationJSON
	GetProjectWorkflowsPageDataDefaultApplicationJSONObject *GetProjectWorkflowsPageDataDefaultApplicationJSON
}
