package operations

import (
	"time"
)

type GetProjectWorkflowRunsPathParams struct {
	ProjectSlug  string `pathParam:"style=simple,explode=false,name=project-slug"`
	WorkflowName string `pathParam:"style=simple,explode=false,name=workflow-name"`
}

type GetProjectWorkflowRunsQueryParams struct {
	AllBranches *bool      `queryParam:"style=form,explode=true,name=all-branches"`
	Branch      *string    `queryParam:"style=form,explode=true,name=branch"`
	EndDate     *time.Time `queryParam:"style=form,explode=true,name=end-date"`
	PageToken   *string    `queryParam:"style=form,explode=true,name=page-token"`
	StartDate   *time.Time `queryParam:"style=form,explode=true,name=start-date"`
}

type GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnum string

const (
	GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnumSuccess      GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnum = "success"
	GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnumFailed       GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnum = "failed"
	GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnumError        GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnum = "error"
	GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnumCanceled     GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnum = "canceled"
	GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnumUnauthorized GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnum = "unauthorized"
)

type GetProjectWorkflowRuns200ApplicationJSONItems struct {
	Branch      string                                                  `json:"branch"`
	CreatedAt   time.Time                                               `json:"created_at"`
	CreditsUsed int64                                                   `json:"credits_used"`
	Duration    int64                                                   `json:"duration"`
	ID          string                                                  `json:"id"`
	Status      GetProjectWorkflowRuns200ApplicationJSONItemsStatusEnum `json:"status"`
	StoppedAt   time.Time                                               `json:"stopped_at"`
}

// GetProjectWorkflowRuns200ApplicationJSON
// Paginated recent workflow runs.
type GetProjectWorkflowRuns200ApplicationJSON struct {
	Items         []GetProjectWorkflowRuns200ApplicationJSONItems `json:"items"`
	NextPageToken string                                          `json:"next_page_token"`
}

type GetProjectWorkflowRunsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetProjectWorkflowRunsRequest struct {
	PathParams  GetProjectWorkflowRunsPathParams
	QueryParams GetProjectWorkflowRunsQueryParams
}

type GetProjectWorkflowRunsResponse struct {
	ContentType                                        string
	StatusCode                                         int64
	GetProjectWorkflowRuns200ApplicationJSONObject     *GetProjectWorkflowRuns200ApplicationJSON
	GetProjectWorkflowRunsDefaultApplicationJSONObject *GetProjectWorkflowRunsDefaultApplicationJSON
}
