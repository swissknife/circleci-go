package operations

type GetPipelineConfigByIDPathParams struct {
	PipelineID string `pathParam:"style=simple,explode=false,name=pipeline-id"`
}

// GetPipelineConfigByIDPipelineConfig
// The configuration strings for the pipeline.
type GetPipelineConfigByIDPipelineConfig struct {
	Compiled            string  `json:"compiled"`
	CompiledSetupConfig *string `json:"compiled-setup-config,omitempty"`
	SetupConfig         *string `json:"setup-config,omitempty"`
	Source              string  `json:"source"`
}

type GetPipelineConfigByIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetPipelineConfigByIDRequest struct {
	PathParams GetPipelineConfigByIDPathParams
}

type GetPipelineConfigByIDResponse struct {
	ContentType                                       string
	PipelineConfig                                    *GetPipelineConfigByIDPipelineConfig
	StatusCode                                        int64
	GetPipelineConfigByIDDefaultApplicationJSONObject *GetPipelineConfigByIDDefaultApplicationJSON
}
