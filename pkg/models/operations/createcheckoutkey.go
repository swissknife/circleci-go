package operations

import (
	"time"
)

type CreateCheckoutKeyPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum string

const (
	CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnumUserKey   CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum = "user-key"
	CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnumDeployKey CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum = "deploy-key"
)

type CreateCheckoutKeyCheckoutKeyInput struct {
	Type CreateCheckoutKeyCheckoutKeyInputCheckoutKeyInputTypeEnum `json:"type"`
}

type CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum string

const (
	CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnumDeployKey     CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum = "deploy-key"
	CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnumGithubUserKey CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum = "github-user-key"
)

type CreateCheckoutKeyCheckoutKey struct {
	CreatedAt   time.Time                                       `json:"created-at"`
	Fingerprint string                                          `json:"fingerprint"`
	Preferred   bool                                            `json:"preferred"`
	PublicKey   string                                          `json:"public-key"`
	Type        CreateCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum `json:"type"`
}

type CreateCheckoutKeyDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateCheckoutKeyRequest struct {
	PathParams CreateCheckoutKeyPathParams
	Request    *CreateCheckoutKeyCheckoutKeyInput `request:"mediaType=application/json"`
}

type CreateCheckoutKeyResponse struct {
	CheckoutKey                                   *CreateCheckoutKeyCheckoutKey
	ContentType                                   string
	StatusCode                                    int64
	CreateCheckoutKeyDefaultApplicationJSONObject *CreateCheckoutKeyDefaultApplicationJSON
}
