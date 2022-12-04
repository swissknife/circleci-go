package operations

import (
	"time"
)

type GetWorkflowByIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetWorkflowByIDWorkflowStatusEnum string

const (
	GetWorkflowByIDWorkflowStatusEnumSuccess      GetWorkflowByIDWorkflowStatusEnum = "success"
	GetWorkflowByIDWorkflowStatusEnumRunning      GetWorkflowByIDWorkflowStatusEnum = "running"
	GetWorkflowByIDWorkflowStatusEnumNotRun       GetWorkflowByIDWorkflowStatusEnum = "not_run"
	GetWorkflowByIDWorkflowStatusEnumFailed       GetWorkflowByIDWorkflowStatusEnum = "failed"
	GetWorkflowByIDWorkflowStatusEnumError        GetWorkflowByIDWorkflowStatusEnum = "error"
	GetWorkflowByIDWorkflowStatusEnumFailing      GetWorkflowByIDWorkflowStatusEnum = "failing"
	GetWorkflowByIDWorkflowStatusEnumOnHold       GetWorkflowByIDWorkflowStatusEnum = "on_hold"
	GetWorkflowByIDWorkflowStatusEnumCanceled     GetWorkflowByIDWorkflowStatusEnum = "canceled"
	GetWorkflowByIDWorkflowStatusEnumUnauthorized GetWorkflowByIDWorkflowStatusEnum = "unauthorized"
)

type GetWorkflowByIDWorkflowTagEnum string

const (
	GetWorkflowByIDWorkflowTagEnumSetup GetWorkflowByIDWorkflowTagEnum = "setup"
)

// GetWorkflowByIDWorkflow
// A workflow
type GetWorkflowByIDWorkflow struct {
	CanceledBy     *string                           `json:"canceled_by,omitempty"`
	CreatedAt      time.Time                         `json:"created_at"`
	ErroredBy      *string                           `json:"errored_by,omitempty"`
	ID             string                            `json:"id"`
	Name           string                            `json:"name"`
	PipelineID     string                            `json:"pipeline_id"`
	PipelineNumber int64                             `json:"pipeline_number"`
	ProjectSlug    string                            `json:"project_slug"`
	StartedBy      string                            `json:"started_by"`
	Status         GetWorkflowByIDWorkflowStatusEnum `json:"status"`
	StoppedAt      time.Time                         `json:"stopped_at"`
	Tag            *GetWorkflowByIDWorkflowTagEnum   `json:"tag,omitempty"`
}

type GetWorkflowByIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetWorkflowByIDRequest struct {
	PathParams GetWorkflowByIDPathParams
}

type GetWorkflowByIDResponse struct {
	ContentType                                 string
	StatusCode                                  int64
	Workflow                                    *GetWorkflowByIDWorkflow
	GetWorkflowByIDDefaultApplicationJSONObject *GetWorkflowByIDDefaultApplicationJSON
}
