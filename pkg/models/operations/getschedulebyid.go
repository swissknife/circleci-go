package operations

import (
	"time"
)

type GetScheduleByIDPathParams struct {
	ScheduleID string `pathParam:"style=simple,explode=false,name=schedule-id"`
}

// GetScheduleByIDScheduleUser
// The attribution actor who will run the scheduled pipeline.
type GetScheduleByIDScheduleUser struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

type GetScheduleByIDScheduleTimetable1DaysOfWeekEnum string

const (
	GetScheduleByIDScheduleTimetable1DaysOfWeekEnumTue GetScheduleByIDScheduleTimetable1DaysOfWeekEnum = "TUE"
	GetScheduleByIDScheduleTimetable1DaysOfWeekEnumSat GetScheduleByIDScheduleTimetable1DaysOfWeekEnum = "SAT"
	GetScheduleByIDScheduleTimetable1DaysOfWeekEnumSun GetScheduleByIDScheduleTimetable1DaysOfWeekEnum = "SUN"
	GetScheduleByIDScheduleTimetable1DaysOfWeekEnumMon GetScheduleByIDScheduleTimetable1DaysOfWeekEnum = "MON"
	GetScheduleByIDScheduleTimetable1DaysOfWeekEnumThu GetScheduleByIDScheduleTimetable1DaysOfWeekEnum = "THU"
	GetScheduleByIDScheduleTimetable1DaysOfWeekEnumWed GetScheduleByIDScheduleTimetable1DaysOfWeekEnum = "WED"
	GetScheduleByIDScheduleTimetable1DaysOfWeekEnumFri GetScheduleByIDScheduleTimetable1DaysOfWeekEnum = "FRI"
)

type GetScheduleByIDScheduleTimetable1MonthsEnum string

const (
	GetScheduleByIDScheduleTimetable1MonthsEnumMar GetScheduleByIDScheduleTimetable1MonthsEnum = "MAR"
	GetScheduleByIDScheduleTimetable1MonthsEnumNov GetScheduleByIDScheduleTimetable1MonthsEnum = "NOV"
	GetScheduleByIDScheduleTimetable1MonthsEnumDec GetScheduleByIDScheduleTimetable1MonthsEnum = "DEC"
	GetScheduleByIDScheduleTimetable1MonthsEnumJun GetScheduleByIDScheduleTimetable1MonthsEnum = "JUN"
	GetScheduleByIDScheduleTimetable1MonthsEnumMay GetScheduleByIDScheduleTimetable1MonthsEnum = "MAY"
	GetScheduleByIDScheduleTimetable1MonthsEnumOct GetScheduleByIDScheduleTimetable1MonthsEnum = "OCT"
	GetScheduleByIDScheduleTimetable1MonthsEnumFeb GetScheduleByIDScheduleTimetable1MonthsEnum = "FEB"
	GetScheduleByIDScheduleTimetable1MonthsEnumApr GetScheduleByIDScheduleTimetable1MonthsEnum = "APR"
	GetScheduleByIDScheduleTimetable1MonthsEnumSep GetScheduleByIDScheduleTimetable1MonthsEnum = "SEP"
	GetScheduleByIDScheduleTimetable1MonthsEnumAug GetScheduleByIDScheduleTimetable1MonthsEnum = "AUG"
	GetScheduleByIDScheduleTimetable1MonthsEnumJan GetScheduleByIDScheduleTimetable1MonthsEnum = "JAN"
	GetScheduleByIDScheduleTimetable1MonthsEnumJul GetScheduleByIDScheduleTimetable1MonthsEnum = "JUL"
)

type GetScheduleByIDScheduleTimetable1 struct {
	DaysOfMonth []int64                                           `json:"days-of-month,omitempty"`
	DaysOfWeek  []GetScheduleByIDScheduleTimetable1DaysOfWeekEnum `json:"days-of-week"`
	HoursOfDay  []int64                                           `json:"hours-of-day"`
	Months      []GetScheduleByIDScheduleTimetable1MonthsEnum     `json:"months,omitempty"`
	PerHour     int64                                             `json:"per-hour"`
}

type GetScheduleByIDScheduleTimetable2DaysOfWeekEnum string

const (
	GetScheduleByIDScheduleTimetable2DaysOfWeekEnumTue GetScheduleByIDScheduleTimetable2DaysOfWeekEnum = "TUE"
	GetScheduleByIDScheduleTimetable2DaysOfWeekEnumSat GetScheduleByIDScheduleTimetable2DaysOfWeekEnum = "SAT"
	GetScheduleByIDScheduleTimetable2DaysOfWeekEnumSun GetScheduleByIDScheduleTimetable2DaysOfWeekEnum = "SUN"
	GetScheduleByIDScheduleTimetable2DaysOfWeekEnumMon GetScheduleByIDScheduleTimetable2DaysOfWeekEnum = "MON"
	GetScheduleByIDScheduleTimetable2DaysOfWeekEnumThu GetScheduleByIDScheduleTimetable2DaysOfWeekEnum = "THU"
	GetScheduleByIDScheduleTimetable2DaysOfWeekEnumWed GetScheduleByIDScheduleTimetable2DaysOfWeekEnum = "WED"
	GetScheduleByIDScheduleTimetable2DaysOfWeekEnumFri GetScheduleByIDScheduleTimetable2DaysOfWeekEnum = "FRI"
)

type GetScheduleByIDScheduleTimetable2MonthsEnum string

const (
	GetScheduleByIDScheduleTimetable2MonthsEnumMar GetScheduleByIDScheduleTimetable2MonthsEnum = "MAR"
	GetScheduleByIDScheduleTimetable2MonthsEnumNov GetScheduleByIDScheduleTimetable2MonthsEnum = "NOV"
	GetScheduleByIDScheduleTimetable2MonthsEnumDec GetScheduleByIDScheduleTimetable2MonthsEnum = "DEC"
	GetScheduleByIDScheduleTimetable2MonthsEnumJun GetScheduleByIDScheduleTimetable2MonthsEnum = "JUN"
	GetScheduleByIDScheduleTimetable2MonthsEnumMay GetScheduleByIDScheduleTimetable2MonthsEnum = "MAY"
	GetScheduleByIDScheduleTimetable2MonthsEnumOct GetScheduleByIDScheduleTimetable2MonthsEnum = "OCT"
	GetScheduleByIDScheduleTimetable2MonthsEnumFeb GetScheduleByIDScheduleTimetable2MonthsEnum = "FEB"
	GetScheduleByIDScheduleTimetable2MonthsEnumApr GetScheduleByIDScheduleTimetable2MonthsEnum = "APR"
	GetScheduleByIDScheduleTimetable2MonthsEnumSep GetScheduleByIDScheduleTimetable2MonthsEnum = "SEP"
	GetScheduleByIDScheduleTimetable2MonthsEnumAug GetScheduleByIDScheduleTimetable2MonthsEnum = "AUG"
	GetScheduleByIDScheduleTimetable2MonthsEnumJan GetScheduleByIDScheduleTimetable2MonthsEnum = "JAN"
	GetScheduleByIDScheduleTimetable2MonthsEnumJul GetScheduleByIDScheduleTimetable2MonthsEnum = "JUL"
)

type GetScheduleByIDScheduleTimetable2 struct {
	DaysOfMonth []int64                                           `json:"days-of-month"`
	DaysOfWeek  []GetScheduleByIDScheduleTimetable2DaysOfWeekEnum `json:"days-of-week,omitempty"`
	HoursOfDay  []int64                                           `json:"hours-of-day"`
	Months      []GetScheduleByIDScheduleTimetable2MonthsEnum     `json:"months,omitempty"`
	PerHour     int64                                             `json:"per-hour"`
}

// GetScheduleByIDSchedule
// A schedule response
type GetScheduleByIDSchedule struct {
	Actor       GetScheduleByIDScheduleUser `json:"actor"`
	CreatedAt   time.Time                   `json:"created-at"`
	Description string                      `json:"description"`
	ID          string                      `json:"id"`
	Name        string                      `json:"name"`
	Parameters  map[string]interface{}      `json:"parameters"`
	ProjectSlug string                      `json:"project-slug"`
	Timetable   interface{}                 `json:"timetable"`
	UpdatedAt   time.Time                   `json:"updated-at"`
}

type GetScheduleByIDDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetScheduleByIDRequest struct {
	PathParams GetScheduleByIDPathParams
}

type GetScheduleByIDResponse struct {
	ContentType                                 string
	Schedule                                    *GetScheduleByIDSchedule
	StatusCode                                  int64
	GetScheduleByIDDefaultApplicationJSONObject *GetScheduleByIDDefaultApplicationJSON
}
