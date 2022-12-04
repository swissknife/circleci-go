package operations

import (
	"time"
)

type GetPipelineByIDPathParams struct {
	PipelineID string `pathParam:"style=simple,explode=false,name=pipeline-id"`
}

type GetPipelineByIDPipelineErrorsTypeEnum string

const (
	GetPipelineByIDPipelineErrorsTypeEnumConfig      GetPipelineByIDPipelineErrorsTypeEnum = "config"
	GetPipelineByIDPipelineErrorsTypeEnumConfigFetch GetPipelineByIDPipelineErrorsTypeEnum = "config-fetch"
	GetPipelineByIDPipelineErrorsTypeEnumTimeout     GetPipelineByIDPipelineErrorsTypeEnum = "timeout"
	GetPipelineByIDPipelineErrorsTypeEnumPermission  GetPipelineByIDPipelineErrorsTypeEnum = "permission"
	GetPipelineByIDPipelineErrorsTypeEnumOther       GetPipelineByIDPipelineErrorsTypeEnum = "other"
	GetPipelineByIDPipelineErrorsTypeEnumPlan        GetPipelineByIDPipelineErrorsTypeEnum = "plan"
)

// GetPipelineByIDPipelineErrors
// An error with a type and message.
type GetPipelineByIDPipelineErrors struct {
	Message string                                `json:"message"`
	Type    GetPipelineByIDPipelineErrorsTypeEnum `json:"type"`
}

type GetPipelineByIDPipelineStateEnum string

const (
	GetPipelineByIDPipelineStateEnumCreated      GetPipelineByIDPipelineStateEnum = "created"
	GetPipelineByIDPipelineStateEnumErrored      GetPipelineByIDPipelineStateEnum = "errored"
	GetPipelineByIDPipelineStateEnumSetupPending GetPipelineByIDPipelineStateEnum = "setup-pending"
	GetPipelineByIDPipelineStateEnumSetup        GetPipelineByIDPipelineStateEnum = "setup"
	GetPipelineByIDPipelineStateEnumPending      GetPipelineByIDPipelineStateEnum = "pending"
)

// GetPipelineByIDPipelineTriggerActor
// The user who triggered the Pipeline.
type GetPipelineByIDPipelineTriggerActor struct {
	AvatarURL string `json:"avatar_url"`
	Login     string `json:"login"`
}

type GetPipelineByIDPipelineTriggerTypeEnum string

const (
	GetPipelineByIDPipelineTriggerTypeEnumScheduledPipeline GetPipelineByIDPipelineTriggerTypeEnum = "scheduled_pipeline"
	GetPipelineByIDPipelineTriggerTypeEnumExplicit          GetPipelineByIDPipelineTriggerTypeEnum = "explicit"
	GetPipelineByIDPipelineTriggerTypeEnumAPI               GetPipelineByIDPipelineTriggerTypeEnum = "api"
	GetPipelineByIDPipelineTriggerTypeEnumWebhook           GetPipelineByIDPipelineTriggerTypeEnum = "webhook"
)

// GetPipelineByIDPipelineTrigger
// A summary of the trigger.
type GetPipelineByIDPipelineTrigger struct {
	Actor      GetPipelineByIDPipelineTriggerActor    `json:"actor"`
	ReceivedAt time.Time                              `json:"received_at"`
	Type       GetPipelineByIDPipelineTriggerTypeEnum `json:"type"`
}

// GetPipelineByIDPipelineVcsCommit
// The latest commit in the pipeline.
type GetPipelineByIDPipelineVcsCommit struct {
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

// GetPipelineByIDPipelineVcs
// VCS information for the pipeline.
type GetPipelineByIDPipelineVcs struct {
	Branch              *string                           `json:"branch,omitempty"`
	Commit              *GetPipelineByIDPipelineVcsCommit `json:"commit,omitempty"`
	OriginRepositoryURL string                            `json:"origin_repository_url"`
	ProviderName        string                            `json:"provider_name"`
	ReviewID            *string                           `json:"review_id,omitempty"`
	ReviewURL           *string                           `json:"review_url,omitempty"`
	Revision            string                            `json:"revision"`
	Tag                 *string                           `json:"tag,omitempty"`
	TargetRepositoryURL string                            `json:"target_repository_url"`
}

// GetPipelineByIDPipeline
// A pipeline response.
type GetPipelineByIDPipeline struct {
	CreatedAt         time.Time                        `json:"created_at"`
	Errors            []GetPipelineByIDPipelineErrors  `json:"errors"`
	ID                string                           `json:"id"`
	Number            int64                            `json:"number"`
	ProjectSlug       string                           `json:"project_slug"`
	State             GetPipelineByIDPipelineStateEnum `json:"state"`
	Trigger           GetPipelineByIDPipelineTrigger   `json:"trigger"`
	TriggerParameters map[string]interface{}           `json:"trigger_parameters,omitempty"`
	UpdatedAt         *time.Time                       `json:"updated_at,omitempty"`
	Vcs               *GetPipelineByIDPipelineVcs      `json:"vcs,omitempty"`
}

type GetPipelineByIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetPipelineByIDRequest struct {
	PathParams GetPipelineByIDPathParams
}

type GetPipelineByIDResponse struct {
	ContentType                                 string
	Pipeline                                    *GetPipelineByIDPipeline
	StatusCode                                  int64
	GetPipelineByIDDefaultApplicationJSONObject *GetPipelineByIDDefaultApplicationJSON
}
