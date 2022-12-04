package operations

import (
	"time"
)

type UpdateSchedulePathParams struct {
	ScheduleID string `pathParam:"style=simple,explode=false,name=schedule-id"`
}

type UpdateScheduleUpdateScheduleParametersAttributionActorEnum string

const (
	UpdateScheduleUpdateScheduleParametersAttributionActorEnumCurrent UpdateScheduleUpdateScheduleParametersAttributionActorEnum = "current"
	UpdateScheduleUpdateScheduleParametersAttributionActorEnumSystem  UpdateScheduleUpdateScheduleParametersAttributionActorEnum = "system"
)

type UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnum string

const (
	UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumTue UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnum = "TUE"
	UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumSat UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnum = "SAT"
	UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumSun UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnum = "SUN"
	UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumMon UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnum = "MON"
	UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumThu UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnum = "THU"
	UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumWed UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnum = "WED"
	UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnumFri UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnum = "FRI"
)

type UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum string

const (
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumMar UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "MAR"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumNov UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "NOV"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumDec UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "DEC"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumJun UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "JUN"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumMay UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "MAY"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumOct UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "OCT"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumFeb UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "FEB"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumApr UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "APR"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumSep UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "SEP"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumAug UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "AUG"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumJan UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "JAN"
	UpdateScheduleUpdateScheduleParametersTimetableMonthsEnumJul UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum = "JUL"
)

// UpdateScheduleUpdateScheduleParametersTimetable
// Timetable that specifies when a schedule triggers.
type UpdateScheduleUpdateScheduleParametersTimetable struct {
	DaysOfMonth []int64                                                         `json:"days-of-month,omitempty"`
	DaysOfWeek  []UpdateScheduleUpdateScheduleParametersTimetableDaysOfWeekEnum `json:"days-of-week,omitempty"`
	HoursOfDay  []int64                                                         `json:"hours-of-day,omitempty"`
	Months      []UpdateScheduleUpdateScheduleParametersTimetableMonthsEnum     `json:"months,omitempty"`
	PerHour     *int64                                                          `json:"per-hour,omitempty"`
}

// UpdateScheduleUpdateScheduleParameters
// The parameters for an update schedule request
type UpdateScheduleUpdateScheduleParameters struct {
	AttributionActor *UpdateScheduleUpdateScheduleParametersAttributionActorEnum `json:"attribution-actor,omitempty"`
	Description      *string                                                     `json:"description,omitempty"`
	Name             *string                                                     `json:"name,omitempty"`
	Parameters       map[string]interface{}                                      `json:"parameters,omitempty"`
	Timetable        *UpdateScheduleUpdateScheduleParametersTimetable            `json:"timetable,omitempty"`
}

// UpdateScheduleScheduleUser
// The attribution actor who will run the scheduled pipeline.
type UpdateScheduleScheduleUser struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

type UpdateScheduleScheduleTimetable1DaysOfWeekEnum string

const (
	UpdateScheduleScheduleTimetable1DaysOfWeekEnumTue UpdateScheduleScheduleTimetable1DaysOfWeekEnum = "TUE"
	UpdateScheduleScheduleTimetable1DaysOfWeekEnumSat UpdateScheduleScheduleTimetable1DaysOfWeekEnum = "SAT"
	UpdateScheduleScheduleTimetable1DaysOfWeekEnumSun UpdateScheduleScheduleTimetable1DaysOfWeekEnum = "SUN"
	UpdateScheduleScheduleTimetable1DaysOfWeekEnumMon UpdateScheduleScheduleTimetable1DaysOfWeekEnum = "MON"
	UpdateScheduleScheduleTimetable1DaysOfWeekEnumThu UpdateScheduleScheduleTimetable1DaysOfWeekEnum = "THU"
	UpdateScheduleScheduleTimetable1DaysOfWeekEnumWed UpdateScheduleScheduleTimetable1DaysOfWeekEnum = "WED"
	UpdateScheduleScheduleTimetable1DaysOfWeekEnumFri UpdateScheduleScheduleTimetable1DaysOfWeekEnum = "FRI"
)

type UpdateScheduleScheduleTimetable1MonthsEnum string

const (
	UpdateScheduleScheduleTimetable1MonthsEnumMar UpdateScheduleScheduleTimetable1MonthsEnum = "MAR"
	UpdateScheduleScheduleTimetable1MonthsEnumNov UpdateScheduleScheduleTimetable1MonthsEnum = "NOV"
	UpdateScheduleScheduleTimetable1MonthsEnumDec UpdateScheduleScheduleTimetable1MonthsEnum = "DEC"
	UpdateScheduleScheduleTimetable1MonthsEnumJun UpdateScheduleScheduleTimetable1MonthsEnum = "JUN"
	UpdateScheduleScheduleTimetable1MonthsEnumMay UpdateScheduleScheduleTimetable1MonthsEnum = "MAY"
	UpdateScheduleScheduleTimetable1MonthsEnumOct UpdateScheduleScheduleTimetable1MonthsEnum = "OCT"
	UpdateScheduleScheduleTimetable1MonthsEnumFeb UpdateScheduleScheduleTimetable1MonthsEnum = "FEB"
	UpdateScheduleScheduleTimetable1MonthsEnumApr UpdateScheduleScheduleTimetable1MonthsEnum = "APR"
	UpdateScheduleScheduleTimetable1MonthsEnumSep UpdateScheduleScheduleTimetable1MonthsEnum = "SEP"
	UpdateScheduleScheduleTimetable1MonthsEnumAug UpdateScheduleScheduleTimetable1MonthsEnum = "AUG"
	UpdateScheduleScheduleTimetable1MonthsEnumJan UpdateScheduleScheduleTimetable1MonthsEnum = "JAN"
	UpdateScheduleScheduleTimetable1MonthsEnumJul UpdateScheduleScheduleTimetable1MonthsEnum = "JUL"
)

type UpdateScheduleScheduleTimetable1 struct {
	DaysOfMonth []int64                                          `json:"days-of-month,omitempty"`
	DaysOfWeek  []UpdateScheduleScheduleTimetable1DaysOfWeekEnum `json:"days-of-week"`
	HoursOfDay  []int64                                          `json:"hours-of-day"`
	Months      []UpdateScheduleScheduleTimetable1MonthsEnum     `json:"months,omitempty"`
	PerHour     int64                                            `json:"per-hour"`
}

type UpdateScheduleScheduleTimetable2DaysOfWeekEnum string

const (
	UpdateScheduleScheduleTimetable2DaysOfWeekEnumTue UpdateScheduleScheduleTimetable2DaysOfWeekEnum = "TUE"
	UpdateScheduleScheduleTimetable2DaysOfWeekEnumSat UpdateScheduleScheduleTimetable2DaysOfWeekEnum = "SAT"
	UpdateScheduleScheduleTimetable2DaysOfWeekEnumSun UpdateScheduleScheduleTimetable2DaysOfWeekEnum = "SUN"
	UpdateScheduleScheduleTimetable2DaysOfWeekEnumMon UpdateScheduleScheduleTimetable2DaysOfWeekEnum = "MON"
	UpdateScheduleScheduleTimetable2DaysOfWeekEnumThu UpdateScheduleScheduleTimetable2DaysOfWeekEnum = "THU"
	UpdateScheduleScheduleTimetable2DaysOfWeekEnumWed UpdateScheduleScheduleTimetable2DaysOfWeekEnum = "WED"
	UpdateScheduleScheduleTimetable2DaysOfWeekEnumFri UpdateScheduleScheduleTimetable2DaysOfWeekEnum = "FRI"
)

type UpdateScheduleScheduleTimetable2MonthsEnum string

const (
	UpdateScheduleScheduleTimetable2MonthsEnumMar UpdateScheduleScheduleTimetable2MonthsEnum = "MAR"
	UpdateScheduleScheduleTimetable2MonthsEnumNov UpdateScheduleScheduleTimetable2MonthsEnum = "NOV"
	UpdateScheduleScheduleTimetable2MonthsEnumDec UpdateScheduleScheduleTimetable2MonthsEnum = "DEC"
	UpdateScheduleScheduleTimetable2MonthsEnumJun UpdateScheduleScheduleTimetable2MonthsEnum = "JUN"
	UpdateScheduleScheduleTimetable2MonthsEnumMay UpdateScheduleScheduleTimetable2MonthsEnum = "MAY"
	UpdateScheduleScheduleTimetable2MonthsEnumOct UpdateScheduleScheduleTimetable2MonthsEnum = "OCT"
	UpdateScheduleScheduleTimetable2MonthsEnumFeb UpdateScheduleScheduleTimetable2MonthsEnum = "FEB"
	UpdateScheduleScheduleTimetable2MonthsEnumApr UpdateScheduleScheduleTimetable2MonthsEnum = "APR"
	UpdateScheduleScheduleTimetable2MonthsEnumSep UpdateScheduleScheduleTimetable2MonthsEnum = "SEP"
	UpdateScheduleScheduleTimetable2MonthsEnumAug UpdateScheduleScheduleTimetable2MonthsEnum = "AUG"
	UpdateScheduleScheduleTimetable2MonthsEnumJan UpdateScheduleScheduleTimetable2MonthsEnum = "JAN"
	UpdateScheduleScheduleTimetable2MonthsEnumJul UpdateScheduleScheduleTimetable2MonthsEnum = "JUL"
)

type UpdateScheduleScheduleTimetable2 struct {
	DaysOfMonth []int64                                          `json:"days-of-month"`
	DaysOfWeek  []UpdateScheduleScheduleTimetable2DaysOfWeekEnum `json:"days-of-week,omitempty"`
	HoursOfDay  []int64                                          `json:"hours-of-day"`
	Months      []UpdateScheduleScheduleTimetable2MonthsEnum     `json:"months,omitempty"`
	PerHour     int64                                            `json:"per-hour"`
}

// UpdateScheduleSchedule
// A schedule response
type UpdateScheduleSchedule struct {
	Actor       UpdateScheduleScheduleUser `json:"actor"`
	CreatedAt   time.Time                  `json:"created-at"`
	Description string                     `json:"description"`
	ID          string                     `json:"id"`
	Name        string                     `json:"name"`
	Parameters  map[string]interface{}     `json:"parameters"`
	ProjectSlug string                     `json:"project-slug"`
	Timetable   interface{}                `json:"timetable"`
	UpdatedAt   time.Time                  `json:"updated-at"`
}

type UpdateScheduleDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateScheduleRequest struct {
	PathParams UpdateSchedulePathParams
	Request    *UpdateScheduleUpdateScheduleParameters `request:"mediaType=application/json"`
}

type UpdateScheduleResponse struct {
	ContentType                                string
	Schedule                                   *UpdateScheduleSchedule
	StatusCode                                 int64
	UpdateScheduleDefaultApplicationJSONObject *UpdateScheduleDefaultApplicationJSON
}
