package operations

import (
	"time"
)

type GetPipelineByNumberPathParams struct {
	PipelineNumber interface{} `pathParam:"style=simple,explode=false,name=pipeline-number"`
	ProjectSlug    string      `pathParam:"style=simple,explode=false,name=project-slug"`
}

type GetPipelineByNumberPipelineErrorsTypeEnum string

const (
	GetPipelineByNumberPipelineErrorsTypeEnumConfig      GetPipelineByNumberPipelineErrorsTypeEnum = "config"
	GetPipelineByNumberPipelineErrorsTypeEnumConfigFetch GetPipelineByNumberPipelineErrorsTypeEnum = "config-fetch"
	GetPipelineByNumberPipelineErrorsTypeEnumTimeout     GetPipelineByNumberPipelineErrorsTypeEnum = "timeout"
	GetPipelineByNumberPipelineErrorsTypeEnumPermission  GetPipelineByNumberPipelineErrorsTypeEnum = "permission"
	GetPipelineByNumberPipelineErrorsTypeEnumOther       GetPipelineByNumberPipelineErrorsTypeEnum = "other"
	GetPipelineByNumberPipelineErrorsTypeEnumPlan        GetPipelineByNumberPipelineErrorsTypeEnum = "plan"
)

// GetPipelineByNumberPipelineErrors
// An error with a type and message.
type GetPipelineByNumberPipelineErrors struct {
	Message string                                    `json:"message"`
	Type    GetPipelineByNumberPipelineErrorsTypeEnum `json:"type"`
}

type GetPipelineByNumberPipelineStateEnum string

const (
	GetPipelineByNumberPipelineStateEnumCreated      GetPipelineByNumberPipelineStateEnum = "created"
	GetPipelineByNumberPipelineStateEnumErrored      GetPipelineByNumberPipelineStateEnum = "errored"
	GetPipelineByNumberPipelineStateEnumSetupPending GetPipelineByNumberPipelineStateEnum = "setup-pending"
	GetPipelineByNumberPipelineStateEnumSetup        GetPipelineByNumberPipelineStateEnum = "setup"
	GetPipelineByNumberPipelineStateEnumPending      GetPipelineByNumberPipelineStateEnum = "pending"
)

// GetPipelineByNumberPipelineTriggerActor
// The user who triggered the Pipeline.
type GetPipelineByNumberPipelineTriggerActor struct {
	AvatarURL string `json:"avatar_url"`
	Login     string `json:"login"`
}

type GetPipelineByNumberPipelineTriggerTypeEnum string

const (
	GetPipelineByNumberPipelineTriggerTypeEnumScheduledPipeline GetPipelineByNumberPipelineTriggerTypeEnum = "scheduled_pipeline"
	GetPipelineByNumberPipelineTriggerTypeEnumExplicit          GetPipelineByNumberPipelineTriggerTypeEnum = "explicit"
	GetPipelineByNumberPipelineTriggerTypeEnumAPI               GetPipelineByNumberPipelineTriggerTypeEnum = "api"
	GetPipelineByNumberPipelineTriggerTypeEnumWebhook           GetPipelineByNumberPipelineTriggerTypeEnum = "webhook"
)

// GetPipelineByNumberPipelineTrigger
// A summary of the trigger.
type GetPipelineByNumberPipelineTrigger struct {
	Actor      GetPipelineByNumberPipelineTriggerActor    `json:"actor"`
	ReceivedAt time.Time                                  `json:"received_at"`
	Type       GetPipelineByNumberPipelineTriggerTypeEnum `json:"type"`
}

// GetPipelineByNumberPipelineVcsCommit
// The latest commit in the pipeline.
type GetPipelineByNumberPipelineVcsCommit struct {
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

// GetPipelineByNumberPipelineVcs
// VCS information for the pipeline.
type GetPipelineByNumberPipelineVcs struct {
	Branch              *string                               `json:"branch,omitempty"`
	Commit              *GetPipelineByNumberPipelineVcsCommit `json:"commit,omitempty"`
	OriginRepositoryURL string                                `json:"origin_repository_url"`
	ProviderName        string                                `json:"provider_name"`
	ReviewID            *string                               `json:"review_id,omitempty"`
	ReviewURL           *string                               `json:"review_url,omitempty"`
	Revision            string                                `json:"revision"`
	Tag                 *string                               `json:"tag,omitempty"`
	TargetRepositoryURL string                                `json:"target_repository_url"`
}

// GetPipelineByNumberPipeline
// A pipeline response.
type GetPipelineByNumberPipeline struct {
	CreatedAt         time.Time                            `json:"created_at"`
	Errors            []GetPipelineByNumberPipelineErrors  `json:"errors"`
	ID                string                               `json:"id"`
	Number            int64                                `json:"number"`
	ProjectSlug       string                               `json:"project_slug"`
	State             GetPipelineByNumberPipelineStateEnum `json:"state"`
	Trigger           GetPipelineByNumberPipelineTrigger   `json:"trigger"`
	TriggerParameters map[string]interface{}               `json:"trigger_parameters,omitempty"`
	UpdatedAt         *time.Time                           `json:"updated_at,omitempty"`
	Vcs               *GetPipelineByNumberPipelineVcs      `json:"vcs,omitempty"`
}

type GetPipelineByNumberDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetPipelineByNumberRequest struct {
	PathParams GetPipelineByNumberPathParams
}

type GetPipelineByNumberResponse struct {
	ContentType                                     string
	Pipeline                                        *GetPipelineByNumberPipeline
	StatusCode                                      int64
	GetPipelineByNumberDefaultApplicationJSONObject *GetPipelineByNumberDefaultApplicationJSON
}
