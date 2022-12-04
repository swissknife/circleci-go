package operations

import (
	"time"
)

type ListPipelinesQueryParams struct {
	Mine      *bool   `queryParam:"style=form,explode=true,name=mine"`
	OrgSlug   *string `queryParam:"style=form,explode=true,name=org-slug"`
	PageToken *string `queryParam:"style=form,explode=true,name=page-token"`
}

type ListPipelinesPipelineListResponsePipelineErrorsTypeEnum string

const (
	ListPipelinesPipelineListResponsePipelineErrorsTypeEnumConfig      ListPipelinesPipelineListResponsePipelineErrorsTypeEnum = "config"
	ListPipelinesPipelineListResponsePipelineErrorsTypeEnumConfigFetch ListPipelinesPipelineListResponsePipelineErrorsTypeEnum = "config-fetch"
	ListPipelinesPipelineListResponsePipelineErrorsTypeEnumTimeout     ListPipelinesPipelineListResponsePipelineErrorsTypeEnum = "timeout"
	ListPipelinesPipelineListResponsePipelineErrorsTypeEnumPermission  ListPipelinesPipelineListResponsePipelineErrorsTypeEnum = "permission"
	ListPipelinesPipelineListResponsePipelineErrorsTypeEnumOther       ListPipelinesPipelineListResponsePipelineErrorsTypeEnum = "other"
	ListPipelinesPipelineListResponsePipelineErrorsTypeEnumPlan        ListPipelinesPipelineListResponsePipelineErrorsTypeEnum = "plan"
)

// ListPipelinesPipelineListResponsePipelineErrors
// An error with a type and message.
type ListPipelinesPipelineListResponsePipelineErrors struct {
	Message string                                                  `json:"message"`
	Type    ListPipelinesPipelineListResponsePipelineErrorsTypeEnum `json:"type"`
}

type ListPipelinesPipelineListResponsePipelineStateEnum string

const (
	ListPipelinesPipelineListResponsePipelineStateEnumCreated      ListPipelinesPipelineListResponsePipelineStateEnum = "created"
	ListPipelinesPipelineListResponsePipelineStateEnumErrored      ListPipelinesPipelineListResponsePipelineStateEnum = "errored"
	ListPipelinesPipelineListResponsePipelineStateEnumSetupPending ListPipelinesPipelineListResponsePipelineStateEnum = "setup-pending"
	ListPipelinesPipelineListResponsePipelineStateEnumSetup        ListPipelinesPipelineListResponsePipelineStateEnum = "setup"
	ListPipelinesPipelineListResponsePipelineStateEnumPending      ListPipelinesPipelineListResponsePipelineStateEnum = "pending"
)

// ListPipelinesPipelineListResponsePipelineTriggerActor
// The user who triggered the Pipeline.
type ListPipelinesPipelineListResponsePipelineTriggerActor struct {
	AvatarURL string `json:"avatar_url"`
	Login     string `json:"login"`
}

type ListPipelinesPipelineListResponsePipelineTriggerTypeEnum string

const (
	ListPipelinesPipelineListResponsePipelineTriggerTypeEnumScheduledPipeline ListPipelinesPipelineListResponsePipelineTriggerTypeEnum = "scheduled_pipeline"
	ListPipelinesPipelineListResponsePipelineTriggerTypeEnumExplicit          ListPipelinesPipelineListResponsePipelineTriggerTypeEnum = "explicit"
	ListPipelinesPipelineListResponsePipelineTriggerTypeEnumAPI               ListPipelinesPipelineListResponsePipelineTriggerTypeEnum = "api"
	ListPipelinesPipelineListResponsePipelineTriggerTypeEnumWebhook           ListPipelinesPipelineListResponsePipelineTriggerTypeEnum = "webhook"
)

// ListPipelinesPipelineListResponsePipelineTrigger
// A summary of the trigger.
type ListPipelinesPipelineListResponsePipelineTrigger struct {
	Actor      ListPipelinesPipelineListResponsePipelineTriggerActor    `json:"actor"`
	ReceivedAt time.Time                                                `json:"received_at"`
	Type       ListPipelinesPipelineListResponsePipelineTriggerTypeEnum `json:"type"`
}

// ListPipelinesPipelineListResponsePipelineVcsCommit
// The latest commit in the pipeline.
type ListPipelinesPipelineListResponsePipelineVcsCommit struct {
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

// ListPipelinesPipelineListResponsePipelineVcs
// VCS information for the pipeline.
type ListPipelinesPipelineListResponsePipelineVcs struct {
	Branch              *string                                             `json:"branch,omitempty"`
	Commit              *ListPipelinesPipelineListResponsePipelineVcsCommit `json:"commit,omitempty"`
	OriginRepositoryURL string                                              `json:"origin_repository_url"`
	ProviderName        string                                              `json:"provider_name"`
	ReviewID            *string                                             `json:"review_id,omitempty"`
	ReviewURL           *string                                             `json:"review_url,omitempty"`
	Revision            string                                              `json:"revision"`
	Tag                 *string                                             `json:"tag,omitempty"`
	TargetRepositoryURL string                                              `json:"target_repository_url"`
}

// ListPipelinesPipelineListResponsePipeline
// A pipeline response.
type ListPipelinesPipelineListResponsePipeline struct {
	CreatedAt         time.Time                                          `json:"created_at"`
	Errors            []ListPipelinesPipelineListResponsePipelineErrors  `json:"errors"`
	ID                string                                             `json:"id"`
	Number            int64                                              `json:"number"`
	ProjectSlug       string                                             `json:"project_slug"`
	State             ListPipelinesPipelineListResponsePipelineStateEnum `json:"state"`
	Trigger           ListPipelinesPipelineListResponsePipelineTrigger   `json:"trigger"`
	TriggerParameters map[string]interface{}                             `json:"trigger_parameters,omitempty"`
	UpdatedAt         *time.Time                                         `json:"updated_at,omitempty"`
	Vcs               *ListPipelinesPipelineListResponsePipelineVcs      `json:"vcs,omitempty"`
}

// ListPipelinesPipelineListResponse
// List of pipelines
type ListPipelinesPipelineListResponse struct {
	Items         []ListPipelinesPipelineListResponsePipeline `json:"items"`
	NextPageToken string                                      `json:"next_page_token"`
}

type ListPipelinesDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListPipelinesRequest struct {
	QueryParams ListPipelinesQueryParams
}

type ListPipelinesResponse struct {
	ContentType                               string
	PipelineListResponse                      *ListPipelinesPipelineListResponse
	StatusCode                                int64
	ListPipelinesDefaultApplicationJSONObject *ListPipelinesDefaultApplicationJSON
}
