package operations

import (
	"time"
)

type ListWorkflowsByPipelineIDPathParams struct {
	PipelineID string `pathParam:"style=simple,explode=false,name=pipeline-id"`
}

type ListWorkflowsByPipelineIDQueryParams struct {
	PageToken *string `queryParam:"style=form,explode=true,name=page-token"`
}

type ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum string

const (
	ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnumSuccess      ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum = "success"
	ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnumRunning      ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum = "running"
	ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnumNotRun       ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum = "not_run"
	ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnumFailed       ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum = "failed"
	ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnumError        ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum = "error"
	ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnumFailing      ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum = "failing"
	ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnumOnHold       ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum = "on_hold"
	ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnumCanceled     ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum = "canceled"
	ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnumUnauthorized ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum = "unauthorized"
)

type ListWorkflowsByPipelineIDWorkflowListResponseWorkflowTagEnum string

const (
	ListWorkflowsByPipelineIDWorkflowListResponseWorkflowTagEnumSetup ListWorkflowsByPipelineIDWorkflowListResponseWorkflowTagEnum = "setup"
)

// ListWorkflowsByPipelineIDWorkflowListResponseWorkflow
// A workflow
type ListWorkflowsByPipelineIDWorkflowListResponseWorkflow struct {
	CanceledBy     *string                                                         `json:"canceled_by,omitempty"`
	CreatedAt      time.Time                                                       `json:"created_at"`
	ErroredBy      *string                                                         `json:"errored_by,omitempty"`
	ID             string                                                          `json:"id"`
	Name           string                                                          `json:"name"`
	PipelineID     string                                                          `json:"pipeline_id"`
	PipelineNumber int64                                                           `json:"pipeline_number"`
	ProjectSlug    string                                                          `json:"project_slug"`
	StartedBy      string                                                          `json:"started_by"`
	Status         ListWorkflowsByPipelineIDWorkflowListResponseWorkflowStatusEnum `json:"status"`
	StoppedAt      time.Time                                                       `json:"stopped_at"`
	Tag            *ListWorkflowsByPipelineIDWorkflowListResponseWorkflowTagEnum   `json:"tag,omitempty"`
}

// ListWorkflowsByPipelineIDWorkflowListResponse
// A list of workflows and associated pagination token.
type ListWorkflowsByPipelineIDWorkflowListResponse struct {
	Items         []ListWorkflowsByPipelineIDWorkflowListResponseWorkflow `json:"items"`
	NextPageToken string                                                  `json:"next_page_token"`
}

type ListWorkflowsByPipelineIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListWorkflowsByPipelineIDRequest struct {
	PathParams  ListWorkflowsByPipelineIDPathParams
	QueryParams ListWorkflowsByPipelineIDQueryParams
}

type ListWorkflowsByPipelineIDResponse struct {
	ContentType                                           string
	StatusCode                                            int64
	WorkflowListResponse                                  *ListWorkflowsByPipelineIDWorkflowListResponse
	ListWorkflowsByPipelineIDDefaultApplicationJSONObject *ListWorkflowsByPipelineIDDefaultApplicationJSON
}
