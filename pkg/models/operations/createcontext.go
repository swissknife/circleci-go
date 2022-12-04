package operations

import (
	"time"
)

type CreateContextRequestBodyOwner1TypeEnum string

const (
	CreateContextRequestBodyOwner1TypeEnumAccount      CreateContextRequestBodyOwner1TypeEnum = "account"
	CreateContextRequestBodyOwner1TypeEnumOrganization CreateContextRequestBodyOwner1TypeEnum = "organization"
)

type CreateContextRequestBodyOwner1 struct {
	ID   string                                  `json:"id"`
	Type *CreateContextRequestBodyOwner1TypeEnum `json:"type,omitempty"`
}

type CreateContextRequestBodyOwner2TypeEnum string

const (
	CreateContextRequestBodyOwner2TypeEnumOrganization CreateContextRequestBodyOwner2TypeEnum = "organization"
)

type CreateContextRequestBodyOwner2 struct {
	Slug string                                  `json:"slug"`
	Type *CreateContextRequestBodyOwner2TypeEnum `json:"type,omitempty"`
}

type CreateContextRequestBody struct {
	Name  string      `json:"name"`
	Owner interface{} `json:"owner"`
}

type CreateContextContext struct {
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
}

type CreateContextDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateContextRequest struct {
	Request *CreateContextRequestBody `request:"mediaType=application/json"`
}

type CreateContextResponse struct {
	ContentType                               string
	Context                                   *CreateContextContext
	StatusCode                                int64
	CreateContextDefaultApplicationJSONObject *CreateContextDefaultApplicationJSON
}
