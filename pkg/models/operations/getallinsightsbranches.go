package operations

type GetAllInsightsBranchesPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type GetAllInsightsBranchesQueryParams struct {
	WorkflowName *string `queryParam:"style=form,explode=true,name=workflow-name"`
}

type GetAllInsightsBranchesDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetAllInsightsBranchesRequest struct {
	PathParams  GetAllInsightsBranchesPathParams
	QueryParams GetAllInsightsBranchesQueryParams
}

type GetAllInsightsBranchesResponse struct {
	ContentType                                        string
	StatusCode                                         int64
	GetAllInsightsBranches200ApplicationJSONAny        *interface{}
	GetAllInsightsBranchesDefaultApplicationJSONObject *GetAllInsightsBranchesDefaultApplicationJSON
}
