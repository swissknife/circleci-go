package operations

import (
	"time"
)

type GetCheckoutKeyPathParams struct {
	Fingerprint string `pathParam:"style=simple,explode=false,name=fingerprint"`
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type GetCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum string

const (
	GetCheckoutKeyCheckoutKeyCheckoutKeyTypeEnumDeployKey     GetCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum = "deploy-key"
	GetCheckoutKeyCheckoutKeyCheckoutKeyTypeEnumGithubUserKey GetCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum = "github-user-key"
)

type GetCheckoutKeyCheckoutKey struct {
	CreatedAt   time.Time                                    `json:"created-at"`
	Fingerprint string                                       `json:"fingerprint"`
	Preferred   bool                                         `json:"preferred"`
	PublicKey   string                                       `json:"public-key"`
	Type        GetCheckoutKeyCheckoutKeyCheckoutKeyTypeEnum `json:"type"`
}

type GetCheckoutKeyDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetCheckoutKeyRequest struct {
	PathParams GetCheckoutKeyPathParams
}

type GetCheckoutKeyResponse struct {
	CheckoutKey                                *GetCheckoutKeyCheckoutKey
	ContentType                                string
	StatusCode                                 int64
	GetCheckoutKeyDefaultApplicationJSONObject *GetCheckoutKeyDefaultApplicationJSON
}
