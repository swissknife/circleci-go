package operations

type GetCollaborationsCollaboration struct {
	AvatarURL string `json:"avatar_url"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	VcsType   string `json:"vcs-type"`
}

type GetCollaborationsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetCollaborationsResponse struct {
	Collaborations                                []GetCollaborationsCollaboration
	ContentType                                   string
	StatusCode                                    int64
	GetCollaborationsDefaultApplicationJSONObject *GetCollaborationsDefaultApplicationJSON
}
