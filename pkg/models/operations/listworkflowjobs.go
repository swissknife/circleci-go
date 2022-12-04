package operations

import (
	"time"
)

type ListWorkflowJobsPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type ListWorkflowJobsWorkflowJobListResponseJobStatusEnum string

const (
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumSuccess            ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "success"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumRunning            ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "running"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumNotRun             ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "not_run"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumFailed             ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "failed"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumRetried            ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "retried"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumQueued             ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "queued"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumNotRunning         ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "not_running"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumInfrastructureFail ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "infrastructure_fail"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumTimedout           ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "timedout"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumOnHold             ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "on_hold"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumTerminatedUnknown  ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "terminated-unknown"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumBlocked            ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "blocked"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumCanceled           ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "canceled"
	ListWorkflowJobsWorkflowJobListResponseJobStatusEnumUnauthorized       ListWorkflowJobsWorkflowJobListResponseJobStatusEnum = "unauthorized"
)

type ListWorkflowJobsWorkflowJobListResponseJobTypeEnum string

const (
	ListWorkflowJobsWorkflowJobListResponseJobTypeEnumBuild    ListWorkflowJobsWorkflowJobListResponseJobTypeEnum = "build"
	ListWorkflowJobsWorkflowJobListResponseJobTypeEnumApproval ListWorkflowJobsWorkflowJobListResponseJobTypeEnum = "approval"
)

// ListWorkflowJobsWorkflowJobListResponseJob
// Job
type ListWorkflowJobsWorkflowJobListResponseJob struct {
	ApprovalRequestID *string                                              `json:"approval_request_id,omitempty"`
	ApprovedBy        *string                                              `json:"approved_by,omitempty"`
	CanceledBy        *string                                              `json:"canceled_by,omitempty"`
	Dependencies      []string                                             `json:"dependencies"`
	ID                string                                               `json:"id"`
	JobNumber         *int64                                               `json:"job_number,omitempty"`
	Name              string                                               `json:"name"`
	ProjectSlug       string                                               `json:"project_slug"`
	StartedAt         time.Time                                            `json:"started_at"`
	Status            ListWorkflowJobsWorkflowJobListResponseJobStatusEnum `json:"status"`
	StoppedAt         *time.Time                                           `json:"stopped_at,omitempty"`
	Type              ListWorkflowJobsWorkflowJobListResponseJobTypeEnum   `json:"type"`
}

type ListWorkflowJobsWorkflowJobListResponse struct {
	Items         []ListWorkflowJobsWorkflowJobListResponseJob `json:"items"`
	NextPageToken string                                       `json:"next_page_token"`
}

type ListWorkflowJobsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListWorkflowJobsRequest struct {
	PathParams ListWorkflowJobsPathParams
}

type ListWorkflowJobsResponse struct {
	ContentType                                  string
	StatusCode                                   int64
	WorkflowJobListResponse                      *ListWorkflowJobsWorkflowJobListResponse
	ListWorkflowJobsDefaultApplicationJSONObject *ListWorkflowJobsDefaultApplicationJSON
}
