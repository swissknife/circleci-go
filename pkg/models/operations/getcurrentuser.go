package operations

type GetCurrentUserUser struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

type GetCurrentUserDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetCurrentUserResponse struct {
	ContentType                                string
	StatusCode                                 int64
	User                                       *GetCurrentUserUser
	GetCurrentUserDefaultApplicationJSONObject *GetCurrentUserDefaultApplicationJSON
}
