package operations

type GetJobArtifactsPathParams struct {
	JobNumber   interface{} `pathParam:"style=simple,explode=false,name=job-number"`
	ProjectSlug string      `pathParam:"style=simple,explode=false,name=project-slug"`
}

// GetJobArtifactsArtifactListResponseArtifact
// An artifact
type GetJobArtifactsArtifactListResponseArtifact struct {
	NodeIndex int64  `json:"node_index"`
	Path      string `json:"path"`
	URL       string `json:"url"`
}

type GetJobArtifactsArtifactListResponse struct {
	Items         []GetJobArtifactsArtifactListResponseArtifact `json:"items"`
	NextPageToken string                                        `json:"next_page_token"`
}

type GetJobArtifactsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetJobArtifactsRequest struct {
	PathParams GetJobArtifactsPathParams
}

type GetJobArtifactsResponse struct {
	ArtifactListResponse                        *GetJobArtifactsArtifactListResponse
	ContentType                                 string
	StatusCode                                  int64
	GetJobArtifactsDefaultApplicationJSONObject *GetJobArtifactsDefaultApplicationJSON
}
