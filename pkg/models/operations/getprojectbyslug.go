package operations

type GetProjectBySlugPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type GetProjectBySlugProjectVcsInfoProviderEnum string

const (
	GetProjectBySlugProjectVcsInfoProviderEnumBitbucket GetProjectBySlugProjectVcsInfoProviderEnum = "Bitbucket"
	GetProjectBySlugProjectVcsInfoProviderEnumCircleCi  GetProjectBySlugProjectVcsInfoProviderEnum = "CircleCI"
	GetProjectBySlugProjectVcsInfoProviderEnumGitHub    GetProjectBySlugProjectVcsInfoProviderEnum = "GitHub"
)

// GetProjectBySlugProjectVcsInfo
// Information about the VCS that hosts the project source code.
type GetProjectBySlugProjectVcsInfo struct {
	DefaultBranch string                                     `json:"default_branch"`
	Provider      GetProjectBySlugProjectVcsInfoProviderEnum `json:"provider"`
	VcsURL        string                                     `json:"vcs_url"`
}

// GetProjectBySlugProject
// NOTE: The definition of Project is subject to change.
type GetProjectBySlugProject struct {
	ID               string                         `json:"id"`
	Name             string                         `json:"name"`
	OrganizationID   string                         `json:"organization_id"`
	OrganizationName string                         `json:"organization_name"`
	OrganizationSlug string                         `json:"organization_slug"`
	Slug             string                         `json:"slug"`
	VcsInfo          GetProjectBySlugProjectVcsInfo `json:"vcs_info"`
}

type GetProjectBySlugDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetProjectBySlugRequest struct {
	PathParams GetProjectBySlugPathParams
}

type GetProjectBySlugResponse struct {
	ContentType                                  string
	Project                                      *GetProjectBySlugProject
	StatusCode                                   int64
	GetProjectBySlugDefaultApplicationJSONObject *GetProjectBySlugDefaultApplicationJSON
}
