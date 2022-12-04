package operations

import (
	"time"
)

type GetContextPathParams struct {
	ContextID string `pathParam:"style=simple,explode=false,name=context-id"`
}

type GetContextContext struct {
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
}

type GetContextDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetContextRequest struct {
	PathParams GetContextPathParams
}

type GetContextResponse struct {
	ContentType                            string
	Context                                *GetContextContext
	StatusCode                             int64
	GetContextDefaultApplicationJSONObject *GetContextDefaultApplicationJSON
}
