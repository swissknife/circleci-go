package operations

import (
	"time"
)

type ListSchedulesForProjectPathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type ListSchedulesForProjectQueryParams struct {
	PageToken *string `queryParam:"style=form,explode=true,name=page-token"`
}

// ListSchedulesForProject200ApplicationJSONScheduleUser
// The attribution actor who will run the scheduled pipeline.
type ListSchedulesForProject200ApplicationJSONScheduleUser struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

type ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnum string

const (
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnumTue ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnum = "TUE"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnumSat ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnum = "SAT"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnumSun ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnum = "SUN"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnumMon ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnum = "MON"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnumThu ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnum = "THU"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnumWed ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnum = "WED"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnumFri ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnum = "FRI"
)

type ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum string

const (
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumMar ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "MAR"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumNov ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "NOV"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumDec ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "DEC"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumJun ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "JUN"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumMay ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "MAY"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumOct ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "OCT"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumFeb ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "FEB"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumApr ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "APR"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumSep ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "SEP"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumAug ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "AUG"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumJan ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "JAN"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnumJul ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum = "JUL"
)

type ListSchedulesForProject200ApplicationJSONScheduleTimetable1 struct {
	DaysOfMonth []int64                                                                     `json:"days-of-month,omitempty"`
	DaysOfWeek  []ListSchedulesForProject200ApplicationJSONScheduleTimetable1DaysOfWeekEnum `json:"days-of-week"`
	HoursOfDay  []int64                                                                     `json:"hours-of-day"`
	Months      []ListSchedulesForProject200ApplicationJSONScheduleTimetable1MonthsEnum     `json:"months,omitempty"`
	PerHour     int64                                                                       `json:"per-hour"`
}

type ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnum string

const (
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnumTue ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnum = "TUE"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnumSat ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnum = "SAT"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnumSun ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnum = "SUN"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnumMon ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnum = "MON"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnumThu ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnum = "THU"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnumWed ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnum = "WED"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnumFri ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnum = "FRI"
)

type ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum string

const (
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumMar ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "MAR"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumNov ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "NOV"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumDec ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "DEC"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumJun ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "JUN"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumMay ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "MAY"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumOct ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "OCT"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumFeb ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "FEB"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumApr ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "APR"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumSep ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "SEP"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumAug ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "AUG"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumJan ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "JAN"
	ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnumJul ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum = "JUL"
)

type ListSchedulesForProject200ApplicationJSONScheduleTimetable2 struct {
	DaysOfMonth []int64                                                                     `json:"days-of-month"`
	DaysOfWeek  []ListSchedulesForProject200ApplicationJSONScheduleTimetable2DaysOfWeekEnum `json:"days-of-week,omitempty"`
	HoursOfDay  []int64                                                                     `json:"hours-of-day"`
	Months      []ListSchedulesForProject200ApplicationJSONScheduleTimetable2MonthsEnum     `json:"months,omitempty"`
	PerHour     int64                                                                       `json:"per-hour"`
}

// ListSchedulesForProject200ApplicationJSONSchedule
// A schedule response
type ListSchedulesForProject200ApplicationJSONSchedule struct {
	Actor       ListSchedulesForProject200ApplicationJSONScheduleUser `json:"actor"`
	CreatedAt   time.Time                                             `json:"created-at"`
	Description string                                                `json:"description"`
	ID          string                                                `json:"id"`
	Name        string                                                `json:"name"`
	Parameters  map[string]interface{}                                `json:"parameters"`
	ProjectSlug string                                                `json:"project-slug"`
	Timetable   interface{}                                           `json:"timetable"`
	UpdatedAt   time.Time                                             `json:"updated-at"`
}

// ListSchedulesForProject200ApplicationJSON
// A sequence of schedules
type ListSchedulesForProject200ApplicationJSON struct {
	Items         []ListSchedulesForProject200ApplicationJSONSchedule `json:"items"`
	NextPageToken string                                              `json:"next_page_token"`
}

type ListSchedulesForProjectDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListSchedulesForProjectRequest struct {
	PathParams  ListSchedulesForProjectPathParams
	QueryParams ListSchedulesForProjectQueryParams
}

type ListSchedulesForProjectResponse struct {
	ContentType                                         string
	StatusCode                                          int64
	ListSchedulesForProject200ApplicationJSONObject     *ListSchedulesForProject200ApplicationJSON
	ListSchedulesForProjectDefaultApplicationJSONObject *ListSchedulesForProjectDefaultApplicationJSON
}
