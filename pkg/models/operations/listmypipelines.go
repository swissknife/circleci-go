package operations

import (
	"time"
)

type ListMyPipelinesPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type ListMyPipelinesQueryParams struct {
	PageToken *string `queryParam:"style=form,explode=true,name=page-token"`
}

type ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnum string

const (
	ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnumConfig      ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnum = "config"
	ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnumConfigFetch ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnum = "config-fetch"
	ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnumTimeout     ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnum = "timeout"
	ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnumPermission  ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnum = "permission"
	ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnumOther       ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnum = "other"
	ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnumPlan        ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnum = "plan"
)

// ListMyPipelinesPipelineListResponsePipelineErrors
// An error with a type and message.
type ListMyPipelinesPipelineListResponsePipelineErrors struct {
	Message string                                                    `json:"message"`
	Type    ListMyPipelinesPipelineListResponsePipelineErrorsTypeEnum `json:"type"`
}

type ListMyPipelinesPipelineListResponsePipelineStateEnum string

const (
	ListMyPipelinesPipelineListResponsePipelineStateEnumCreated      ListMyPipelinesPipelineListResponsePipelineStateEnum = "created"
	ListMyPipelinesPipelineListResponsePipelineStateEnumErrored      ListMyPipelinesPipelineListResponsePipelineStateEnum = "errored"
	ListMyPipelinesPipelineListResponsePipelineStateEnumSetupPending ListMyPipelinesPipelineListResponsePipelineStateEnum = "setup-pending"
	ListMyPipelinesPipelineListResponsePipelineStateEnumSetup        ListMyPipelinesPipelineListResponsePipelineStateEnum = "setup"
	ListMyPipelinesPipelineListResponsePipelineStateEnumPending      ListMyPipelinesPipelineListResponsePipelineStateEnum = "pending"
)

// ListMyPipelinesPipelineListResponsePipelineTriggerActor
// The user who triggered the Pipeline.
type ListMyPipelinesPipelineListResponsePipelineTriggerActor struct {
	AvatarURL string `json:"avatar_url"`
	Login     string `json:"login"`
}

type ListMyPipelinesPipelineListResponsePipelineTriggerTypeEnum string

const (
	ListMyPipelinesPipelineListResponsePipelineTriggerTypeEnumScheduledPipeline ListMyPipelinesPipelineListResponsePipelineTriggerTypeEnum = "scheduled_pipeline"
	ListMyPipelinesPipelineListResponsePipelineTriggerTypeEnumExplicit          ListMyPipelinesPipelineListResponsePipelineTriggerTypeEnum = "explicit"
	ListMyPipelinesPipelineListResponsePipelineTriggerTypeEnumAPI               ListMyPipelinesPipelineListResponsePipelineTriggerTypeEnum = "api"
	ListMyPipelinesPipelineListResponsePipelineTriggerTypeEnumWebhook           ListMyPipelinesPipelineListResponsePipelineTriggerTypeEnum = "webhook"
)

// ListMyPipelinesPipelineListResponsePipelineTrigger
// A summary of the trigger.
type ListMyPipelinesPipelineListResponsePipelineTrigger struct {
	Actor      ListMyPipelinesPipelineListResponsePipelineTriggerActor    `json:"actor"`
	ReceivedAt time.Time                                                  `json:"received_at"`
	Type       ListMyPipelinesPipelineListResponsePipelineTriggerTypeEnum `json:"type"`
}

// ListMyPipelinesPipelineListResponsePipelineVcsCommit
// The latest commit in the pipeline.
type ListMyPipelinesPipelineListResponsePipelineVcsCommit struct {
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

// ListMyPipelinesPipelineListResponsePipelineVcs
// VCS information for the pipeline.
type ListMyPipelinesPipelineListResponsePipelineVcs struct {
	Branch              *string                                               `json:"branch,omitempty"`
	Commit              *ListMyPipelinesPipelineListResponsePipelineVcsCommit `json:"commit,omitempty"`
	OriginRepositoryURL string                                                `json:"origin_repository_url"`
	ProviderName        string                                                `json:"provider_name"`
	ReviewID            *string                                               `json:"review_id,omitempty"`
	ReviewURL           *string                                               `json:"review_url,omitempty"`
	Revision            string                                                `json:"revision"`
	Tag                 *string                                               `json:"tag,omitempty"`
	TargetRepositoryURL string                                                `json:"target_repository_url"`
}

// ListMyPipelinesPipelineListResponsePipeline
// A pipeline response.
type ListMyPipelinesPipelineListResponsePipeline struct {
	CreatedAt         time.Time                                            `json:"created_at"`
	Errors            []ListMyPipelinesPipelineListResponsePipelineErrors  `json:"errors"`
	ID                string                                               `json:"id"`
	Number            int64                                                `json:"number"`
	ProjectSlug       string                                               `json:"project_slug"`
	State             ListMyPipelinesPipelineListResponsePipelineStateEnum `json:"state"`
	Trigger           ListMyPipelinesPipelineListResponsePipelineTrigger   `json:"trigger"`
	TriggerParameters map[string]interface{}                               `json:"trigger_parameters,omitempty"`
	UpdatedAt         *time.Time                                           `json:"updated_at,omitempty"`
	Vcs               *ListMyPipelinesPipelineListResponsePipelineVcs      `json:"vcs,omitempty"`
}

// ListMyPipelinesPipelineListResponse
// List of pipelines
type ListMyPipelinesPipelineListResponse struct {
	Items         []ListMyPipelinesPipelineListResponsePipeline `json:"items"`
	NextPageToken string                                        `json:"next_page_token"`
}

type ListMyPipelinesDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListMyPipelinesRequest struct {
	PathParams  ListMyPipelinesPathParams
	QueryParams ListMyPipelinesQueryParams
}

type ListMyPipelinesResponse struct {
	ContentType                                 string
	PipelineListResponse                        *ListMyPipelinesPipelineListResponse
	StatusCode                                  int64
	ListMyPipelinesDefaultApplicationJSONObject *ListMyPipelinesDefaultApplicationJSON
}
