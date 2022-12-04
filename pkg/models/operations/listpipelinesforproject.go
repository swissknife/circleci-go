package operations

import (
	"time"
)

type ListPipelinesForProjectPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type ListPipelinesForProjectQueryParams struct {
	Branch    *string `queryParam:"style=form,explode=true,name=branch"`
	PageToken *string `queryParam:"style=form,explode=true,name=page-token"`
}

type ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnum string

const (
	ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnumConfig      ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnum = "config"
	ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnumConfigFetch ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnum = "config-fetch"
	ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnumTimeout     ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnum = "timeout"
	ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnumPermission  ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnum = "permission"
	ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnumOther       ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnum = "other"
	ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnumPlan        ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnum = "plan"
)

// ListPipelinesForProjectPipelineListResponsePipelineErrors
// An error with a type and message.
type ListPipelinesForProjectPipelineListResponsePipelineErrors struct {
	Message string                                                            `json:"message"`
	Type    ListPipelinesForProjectPipelineListResponsePipelineErrorsTypeEnum `json:"type"`
}

type ListPipelinesForProjectPipelineListResponsePipelineStateEnum string

const (
	ListPipelinesForProjectPipelineListResponsePipelineStateEnumCreated      ListPipelinesForProjectPipelineListResponsePipelineStateEnum = "created"
	ListPipelinesForProjectPipelineListResponsePipelineStateEnumErrored      ListPipelinesForProjectPipelineListResponsePipelineStateEnum = "errored"
	ListPipelinesForProjectPipelineListResponsePipelineStateEnumSetupPending ListPipelinesForProjectPipelineListResponsePipelineStateEnum = "setup-pending"
	ListPipelinesForProjectPipelineListResponsePipelineStateEnumSetup        ListPipelinesForProjectPipelineListResponsePipelineStateEnum = "setup"
	ListPipelinesForProjectPipelineListResponsePipelineStateEnumPending      ListPipelinesForProjectPipelineListResponsePipelineStateEnum = "pending"
)

// ListPipelinesForProjectPipelineListResponsePipelineTriggerActor
// The user who triggered the Pipeline.
type ListPipelinesForProjectPipelineListResponsePipelineTriggerActor struct {
	AvatarURL string `json:"avatar_url"`
	Login     string `json:"login"`
}

type ListPipelinesForProjectPipelineListResponsePipelineTriggerTypeEnum string

const (
	ListPipelinesForProjectPipelineListResponsePipelineTriggerTypeEnumScheduledPipeline ListPipelinesForProjectPipelineListResponsePipelineTriggerTypeEnum = "scheduled_pipeline"
	ListPipelinesForProjectPipelineListResponsePipelineTriggerTypeEnumExplicit          ListPipelinesForProjectPipelineListResponsePipelineTriggerTypeEnum = "explicit"
	ListPipelinesForProjectPipelineListResponsePipelineTriggerTypeEnumAPI               ListPipelinesForProjectPipelineListResponsePipelineTriggerTypeEnum = "api"
	ListPipelinesForProjectPipelineListResponsePipelineTriggerTypeEnumWebhook           ListPipelinesForProjectPipelineListResponsePipelineTriggerTypeEnum = "webhook"
)

// ListPipelinesForProjectPipelineListResponsePipelineTrigger
// A summary of the trigger.
type ListPipelinesForProjectPipelineListResponsePipelineTrigger struct {
	Actor      ListPipelinesForProjectPipelineListResponsePipelineTriggerActor    `json:"actor"`
	ReceivedAt time.Time                                                          `json:"received_at"`
	Type       ListPipelinesForProjectPipelineListResponsePipelineTriggerTypeEnum `json:"type"`
}

// ListPipelinesForProjectPipelineListResponsePipelineVcsCommit
// The latest commit in the pipeline.
type ListPipelinesForProjectPipelineListResponsePipelineVcsCommit struct {
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

// ListPipelinesForProjectPipelineListResponsePipelineVcs
// VCS information for the pipeline.
type ListPipelinesForProjectPipelineListResponsePipelineVcs struct {
	Branch              *string                                                       `json:"branch,omitempty"`
	Commit              *ListPipelinesForProjectPipelineListResponsePipelineVcsCommit `json:"commit,omitempty"`
	OriginRepositoryURL string                                                        `json:"origin_repository_url"`
	ProviderName        string                                                        `json:"provider_name"`
	ReviewID            *string                                                       `json:"review_id,omitempty"`
	ReviewURL           *string                                                       `json:"review_url,omitempty"`
	Revision            string                                                        `json:"revision"`
	Tag                 *string                                                       `json:"tag,omitempty"`
	TargetRepositoryURL string                                                        `json:"target_repository_url"`
}

// ListPipelinesForProjectPipelineListResponsePipeline
// A pipeline response.
type ListPipelinesForProjectPipelineListResponsePipeline struct {
	CreatedAt         time.Time                                                    `json:"created_at"`
	Errors            []ListPipelinesForProjectPipelineListResponsePipelineErrors  `json:"errors"`
	ID                string                                                       `json:"id"`
	Number            int64                                                        `json:"number"`
	ProjectSlug       string                                                       `json:"project_slug"`
	State             ListPipelinesForProjectPipelineListResponsePipelineStateEnum `json:"state"`
	Trigger           ListPipelinesForProjectPipelineListResponsePipelineTrigger   `json:"trigger"`
	TriggerParameters map[string]interface{}                                       `json:"trigger_parameters,omitempty"`
	UpdatedAt         *time.Time                                                   `json:"updated_at,omitempty"`
	Vcs               *ListPipelinesForProjectPipelineListResponsePipelineVcs      `json:"vcs,omitempty"`
}

// ListPipelinesForProjectPipelineListResponse
// List of pipelines
type ListPipelinesForProjectPipelineListResponse struct {
	Items         []ListPipelinesForProjectPipelineListResponsePipeline `json:"items"`
	NextPageToken string                                                `json:"next_page_token"`
}

type ListPipelinesForProjectDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListPipelinesForProjectRequest struct {
	PathParams  ListPipelinesForProjectPathParams
	QueryParams ListPipelinesForProjectQueryParams
}

type ListPipelinesForProjectResponse struct {
	ContentType                                         string
	PipelineListResponse                                *ListPipelinesForProjectPipelineListResponse
	StatusCode                                          int64
	ListPipelinesForProjectDefaultApplicationJSONObject *ListPipelinesForProjectDefaultApplicationJSON
}
