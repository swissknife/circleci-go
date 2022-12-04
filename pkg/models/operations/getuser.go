package operations

type GetUserPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetUserUser struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

type GetUserDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetUserRequest struct {
	PathParams GetUserPathParams
}

type GetUserResponse struct {
	ContentType                         string
	StatusCode                          int64
	User                                *GetUserUser
	GetUserDefaultApplicationJSONObject *GetUserDefaultApplicationJSON
}
