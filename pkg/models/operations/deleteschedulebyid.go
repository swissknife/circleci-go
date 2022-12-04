package operations

type DeleteScheduleByIDPathParams struct {
	ScheduleID string `pathParam:"style=simple,explode=false,name=schedule-id"`
}

// DeleteScheduleByIDMessageResponse
// message response
type DeleteScheduleByIDMessageResponse struct {
	Message string `json:"message"`
}

type DeleteScheduleByIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteScheduleByIDRequest struct {
	PathParams DeleteScheduleByIDPathParams
}

type DeleteScheduleByIDResponse struct {
	ContentType                                    string
	MessageResponse                                *DeleteScheduleByIDMessageResponse
	StatusCode                                     int64
	DeleteScheduleByIDDefaultApplicationJSONObject *DeleteScheduleByIDDefaultApplicationJSON
}
