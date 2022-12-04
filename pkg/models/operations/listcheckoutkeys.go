package operations

import (
	"time"
)

type ListCheckoutKeysPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum string

const (
	ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnumDeployKey     ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum = "deploy-key"
	ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnumGithubUserKey ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum = "github-user-key"
)

type ListCheckoutKeysCheckoutKeyListResponseCheckoutKey struct {
	CreatedAt   time.Time                                                             `json:"created-at"`
	Fingerprint string                                                                `json:"fingerprint"`
	Preferred   bool                                                                  `json:"preferred"`
	PublicKey   string                                                                `json:"public-key"`
	Type        ListCheckoutKeysCheckoutKeyListResponseCheckoutKeyCheckoutKeyTypeEnum `json:"type"`
}

type ListCheckoutKeysCheckoutKeyListResponse struct {
	Items         []ListCheckoutKeysCheckoutKeyListResponseCheckoutKey `json:"items"`
	NextPageToken string                                               `json:"next_page_token"`
}

type ListCheckoutKeysDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListCheckoutKeysRequest struct {
	PathParams ListCheckoutKeysPathParams
}

type ListCheckoutKeysResponse struct {
	CheckoutKeyListResponse                      *ListCheckoutKeysCheckoutKeyListResponse
	ContentType                                  string
	StatusCode                                   int64
	ListCheckoutKeysDefaultApplicationJSONObject *ListCheckoutKeysDefaultApplicationJSON
}
