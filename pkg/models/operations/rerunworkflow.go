package operations

type RerunWorkflowPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

// RerunWorkflowRerunWorkflowParameters
// The information you can supply when rerunning a workflow.
type RerunWorkflowRerunWorkflowParameters struct {
	EnableSSH  *bool    `json:"enable_ssh,omitempty"`
	FromFailed *bool    `json:"from_failed,omitempty"`
	Jobs       []string `json:"jobs,omitempty"`
	SparseTree *bool    `json:"sparse_tree,omitempty"`
}

// RerunWorkflow202ApplicationJSON
// A response to rerunning a workflow
type RerunWorkflow202ApplicationJSON struct {
	WorkflowID string `json:"workflow_id"`
}

type RerunWorkflowDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type RerunWorkflowRequest struct {
	PathParams RerunWorkflowPathParams
	Request    *RerunWorkflowRerunWorkflowParameters `request:"mediaType=application/json"`
}

type RerunWorkflowResponse struct {
	ContentType                               string
	StatusCode                                int64
	RerunWorkflow202ApplicationJSONObject     *RerunWorkflow202ApplicationJSON
	RerunWorkflowDefaultApplicationJSONObject *RerunWorkflowDefaultApplicationJSON
}
