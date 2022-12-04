package operations

import (
	"time"
)

type GetJobDetailsPathParams struct {
	JobNumber   interface{} `pathParam:"style=simple,explode=false,name=job-number"`
	ProjectSlug string      `pathParam:"style=simple,explode=false,name=project-slug"`
}

// GetJobDetailsJobDetailsContexts
// Information about the context.
type GetJobDetailsJobDetailsContexts struct {
	Name string `json:"name"`
}

// GetJobDetailsJobDetailsExecutor
// Information about executor used for a job.
type GetJobDetailsJobDetailsExecutor struct {
	ResourceClass string  `json:"resource_class"`
	Type          *string `json:"type,omitempty"`
}

// GetJobDetailsJobDetailsLatestWorkflow
// Info about the latest workflow the job was a part of.
type GetJobDetailsJobDetailsLatestWorkflow struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetJobDetailsJobDetailsMessages
// Message from CircleCI execution platform.
type GetJobDetailsJobDetailsMessages struct {
	Message string  `json:"message"`
	Reason  *string `json:"reason,omitempty"`
	Type    string  `json:"type"`
}

// GetJobDetailsJobDetailsOrganization
// Information about an organization.
type GetJobDetailsJobDetailsOrganization struct {
	Name string `json:"name"`
}

// GetJobDetailsJobDetailsParallelRuns
// Info about a status of the parallel run.
type GetJobDetailsJobDetailsParallelRuns struct {
	Index  int64  `json:"index"`
	Status string `json:"status"`
}

// GetJobDetailsJobDetailsPipeline
// Info about a pipeline the job is a part of.
type GetJobDetailsJobDetailsPipeline struct {
	ID string `json:"id"`
}

// GetJobDetailsJobDetailsProject
// Information about a project.
type GetJobDetailsJobDetailsProject struct {
	ExternalURL string `json:"external_url"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
}

type GetJobDetailsJobDetailsStatusEnum string

const (
	GetJobDetailsJobDetailsStatusEnumSuccess            GetJobDetailsJobDetailsStatusEnum = "success"
	GetJobDetailsJobDetailsStatusEnumRunning            GetJobDetailsJobDetailsStatusEnum = "running"
	GetJobDetailsJobDetailsStatusEnumNotRun             GetJobDetailsJobDetailsStatusEnum = "not_run"
	GetJobDetailsJobDetailsStatusEnumFailed             GetJobDetailsJobDetailsStatusEnum = "failed"
	GetJobDetailsJobDetailsStatusEnumRetried            GetJobDetailsJobDetailsStatusEnum = "retried"
	GetJobDetailsJobDetailsStatusEnumQueued             GetJobDetailsJobDetailsStatusEnum = "queued"
	GetJobDetailsJobDetailsStatusEnumNotRunning         GetJobDetailsJobDetailsStatusEnum = "not_running"
	GetJobDetailsJobDetailsStatusEnumInfrastructureFail GetJobDetailsJobDetailsStatusEnum = "infrastructure_fail"
	GetJobDetailsJobDetailsStatusEnumTimedout           GetJobDetailsJobDetailsStatusEnum = "timedout"
	GetJobDetailsJobDetailsStatusEnumOnHold             GetJobDetailsJobDetailsStatusEnum = "on_hold"
	GetJobDetailsJobDetailsStatusEnumTerminatedUnknown  GetJobDetailsJobDetailsStatusEnum = "terminated-unknown"
	GetJobDetailsJobDetailsStatusEnumBlocked            GetJobDetailsJobDetailsStatusEnum = "blocked"
	GetJobDetailsJobDetailsStatusEnumCanceled           GetJobDetailsJobDetailsStatusEnum = "canceled"
	GetJobDetailsJobDetailsStatusEnumUnauthorized       GetJobDetailsJobDetailsStatusEnum = "unauthorized"
)

// GetJobDetailsJobDetails
// Job Details
type GetJobDetailsJobDetails struct {
	Contexts       []GetJobDetailsJobDetailsContexts     `json:"contexts"`
	CreatedAt      time.Time                             `json:"created_at"`
	Duration       int64                                 `json:"duration"`
	Executor       GetJobDetailsJobDetailsExecutor       `json:"executor"`
	LatestWorkflow GetJobDetailsJobDetailsLatestWorkflow `json:"latest_workflow"`
	Messages       []GetJobDetailsJobDetailsMessages     `json:"messages"`
	Name           string                                `json:"name"`
	Number         int64                                 `json:"number"`
	Organization   GetJobDetailsJobDetailsOrganization   `json:"organization"`
	ParallelRuns   []GetJobDetailsJobDetailsParallelRuns `json:"parallel_runs"`
	Parallelism    int64                                 `json:"parallelism"`
	Pipeline       GetJobDetailsJobDetailsPipeline       `json:"pipeline"`
	Project        GetJobDetailsJobDetailsProject        `json:"project"`
	QueuedAt       time.Time                             `json:"queued_at"`
	StartedAt      time.Time                             `json:"started_at"`
	Status         GetJobDetailsJobDetailsStatusEnum     `json:"status"`
	StoppedAt      *time.Time                            `json:"stopped_at,omitempty"`
	WebURL         string                                `json:"web_url"`
}

type GetJobDetailsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetJobDetailsRequest struct {
	PathParams GetJobDetailsPathParams
}

type GetJobDetailsResponse struct {
	ContentType                               string
	JobDetails                                *GetJobDetailsJobDetails
	StatusCode                                int64
	GetJobDetailsDefaultApplicationJSONObject *GetJobDetailsDefaultApplicationJSON
}
