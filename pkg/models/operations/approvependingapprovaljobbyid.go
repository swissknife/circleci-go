package operations

type ApprovePendingApprovalJobByIDPathParams struct {
	ApprovalRequestID string `pathParam:"style=simple,explode=false,name=approval_request_id"`
	ID                string `pathParam:"style=simple,explode=false,name=id"`
}

// ApprovePendingApprovalJobByIDMessageResponse
// message response
type ApprovePendingApprovalJobByIDMessageResponse struct {
	Message string `json:"message"`
}

type ApprovePendingApprovalJobByIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ApprovePendingApprovalJobByIDRequest struct {
	PathParams ApprovePendingApprovalJobByIDPathParams
}

type ApprovePendingApprovalJobByIDResponse struct {
	ContentType                                               string
	MessageResponse                                           *ApprovePendingApprovalJobByIDMessageResponse
	StatusCode                                                int64
	ApprovePendingApprovalJobByIDDefaultApplicationJSONObject *ApprovePendingApprovalJobByIDDefaultApplicationJSON
}
