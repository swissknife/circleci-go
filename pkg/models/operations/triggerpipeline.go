package operations

import (
	"time"
)

type TriggerPipelinePathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

// TriggerPipelineTriggerPipelineParameters
// The information you can supply when triggering a pipeline.
type TriggerPipelineTriggerPipelineParameters struct {
	Branch     *string                `json:"branch,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Tag        *string                `json:"tag,omitempty"`
}

type TriggerPipelinePipelineCreationStateEnum string

const (
	TriggerPipelinePipelineCreationStateEnumCreated      TriggerPipelinePipelineCreationStateEnum = "created"
	TriggerPipelinePipelineCreationStateEnumErrored      TriggerPipelinePipelineCreationStateEnum = "errored"
	TriggerPipelinePipelineCreationStateEnumSetupPending TriggerPipelinePipelineCreationStateEnum = "setup-pending"
	TriggerPipelinePipelineCreationStateEnumSetup        TriggerPipelinePipelineCreationStateEnum = "setup"
	TriggerPipelinePipelineCreationStateEnumPending      TriggerPipelinePipelineCreationStateEnum = "pending"
)

// TriggerPipelinePipelineCreation
// A pipeline creation response.
type TriggerPipelinePipelineCreation struct {
	CreatedAt time.Time                                `json:"created_at"`
	ID        string                                   `json:"id"`
	Number    int64                                    `json:"number"`
	State     TriggerPipelinePipelineCreationStateEnum `json:"state"`
}

type TriggerPipelineDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type TriggerPipelineRequest struct {
	PathParams TriggerPipelinePathParams
	Request    *TriggerPipelineTriggerPipelineParameters `request:"mediaType=application/json"`
}

type TriggerPipelineResponse struct {
	ContentType                                 string
	PipelineCreation                            *TriggerPipelinePipelineCreation
	StatusCode                                  int64
	TriggerPipelineDefaultApplicationJSONObject *TriggerPipelineDefaultApplicationJSON
}
