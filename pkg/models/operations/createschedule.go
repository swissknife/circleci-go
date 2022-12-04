package operations

import (
	"time"
)

type CreateSchedulePathParams struct {
	ProjectSlug string `pathParam:"style=simple,explode=false,name=project-slug"`
}

type CreateScheduleCreateScheduleParametersAttributionActorEnum string

const (
	CreateScheduleCreateScheduleParametersAttributionActorEnumCurrent CreateScheduleCreateScheduleParametersAttributionActorEnum = "current"
	CreateScheduleCreateScheduleParametersAttributionActorEnumSystem  CreateScheduleCreateScheduleParametersAttributionActorEnum = "system"
)

type CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnum string

const (
	CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnumTue CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnum = "TUE"
	CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnumSat CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnum = "SAT"
	CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnumSun CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnum = "SUN"
	CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnumMon CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnum = "MON"
	CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnumThu CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnum = "THU"
	CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnumWed CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnum = "WED"
	CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnumFri CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnum = "FRI"
)

type CreateScheduleCreateScheduleParametersTimetable1MonthsEnum string

const (
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumMar CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "MAR"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumNov CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "NOV"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumDec CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "DEC"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumJun CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "JUN"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumMay CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "MAY"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumOct CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "OCT"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumFeb CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "FEB"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumApr CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "APR"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumSep CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "SEP"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumAug CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "AUG"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumJan CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "JAN"
	CreateScheduleCreateScheduleParametersTimetable1MonthsEnumJul CreateScheduleCreateScheduleParametersTimetable1MonthsEnum = "JUL"
)

type CreateScheduleCreateScheduleParametersTimetable1 struct {
	DaysOfMonth []int64                                                          `json:"days-of-month,omitempty"`
	DaysOfWeek  []CreateScheduleCreateScheduleParametersTimetable1DaysOfWeekEnum `json:"days-of-week"`
	HoursOfDay  []int64                                                          `json:"hours-of-day"`
	Months      []CreateScheduleCreateScheduleParametersTimetable1MonthsEnum     `json:"months,omitempty"`
	PerHour     int64                                                            `json:"per-hour"`
}

type CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnum string

const (
	CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnumTue CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnum = "TUE"
	CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnumSat CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnum = "SAT"
	CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnumSun CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnum = "SUN"
	CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnumMon CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnum = "MON"
	CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnumThu CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnum = "THU"
	CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnumWed CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnum = "WED"
	CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnumFri CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnum = "FRI"
)

type CreateScheduleCreateScheduleParametersTimetable2MonthsEnum string

const (
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumMar CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "MAR"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumNov CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "NOV"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumDec CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "DEC"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumJun CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "JUN"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumMay CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "MAY"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumOct CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "OCT"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumFeb CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "FEB"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumApr CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "APR"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumSep CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "SEP"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumAug CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "AUG"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumJan CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "JAN"
	CreateScheduleCreateScheduleParametersTimetable2MonthsEnumJul CreateScheduleCreateScheduleParametersTimetable2MonthsEnum = "JUL"
)

type CreateScheduleCreateScheduleParametersTimetable2 struct {
	DaysOfMonth []int64                                                          `json:"days-of-month"`
	DaysOfWeek  []CreateScheduleCreateScheduleParametersTimetable2DaysOfWeekEnum `json:"days-of-week,omitempty"`
	HoursOfDay  []int64                                                          `json:"hours-of-day"`
	Months      []CreateScheduleCreateScheduleParametersTimetable2MonthsEnum     `json:"months,omitempty"`
	PerHour     int64                                                            `json:"per-hour"`
}

// CreateScheduleCreateScheduleParameters
// The parameters for a create schedule request
type CreateScheduleCreateScheduleParameters struct {
	AttributionActor CreateScheduleCreateScheduleParametersAttributionActorEnum `json:"attribution-actor"`
	Description      *string                                                    `json:"description,omitempty"`
	Name             string                                                     `json:"name"`
	Parameters       map[string]interface{}                                     `json:"parameters"`
	Timetable        interface{}                                                `json:"timetable"`
}

// CreateScheduleScheduleUser
// The attribution actor who will run the scheduled pipeline.
type CreateScheduleScheduleUser struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

type CreateScheduleScheduleTimetable1DaysOfWeekEnum string

const (
	CreateScheduleScheduleTimetable1DaysOfWeekEnumTue CreateScheduleScheduleTimetable1DaysOfWeekEnum = "TUE"
	CreateScheduleScheduleTimetable1DaysOfWeekEnumSat CreateScheduleScheduleTimetable1DaysOfWeekEnum = "SAT"
	CreateScheduleScheduleTimetable1DaysOfWeekEnumSun CreateScheduleScheduleTimetable1DaysOfWeekEnum = "SUN"
	CreateScheduleScheduleTimetable1DaysOfWeekEnumMon CreateScheduleScheduleTimetable1DaysOfWeekEnum = "MON"
	CreateScheduleScheduleTimetable1DaysOfWeekEnumThu CreateScheduleScheduleTimetable1DaysOfWeekEnum = "THU"
	CreateScheduleScheduleTimetable1DaysOfWeekEnumWed CreateScheduleScheduleTimetable1DaysOfWeekEnum = "WED"
	CreateScheduleScheduleTimetable1DaysOfWeekEnumFri CreateScheduleScheduleTimetable1DaysOfWeekEnum = "FRI"
)

type CreateScheduleScheduleTimetable1MonthsEnum string

const (
	CreateScheduleScheduleTimetable1MonthsEnumMar CreateScheduleScheduleTimetable1MonthsEnum = "MAR"
	CreateScheduleScheduleTimetable1MonthsEnumNov CreateScheduleScheduleTimetable1MonthsEnum = "NOV"
	CreateScheduleScheduleTimetable1MonthsEnumDec CreateScheduleScheduleTimetable1MonthsEnum = "DEC"
	CreateScheduleScheduleTimetable1MonthsEnumJun CreateScheduleScheduleTimetable1MonthsEnum = "JUN"
	CreateScheduleScheduleTimetable1MonthsEnumMay CreateScheduleScheduleTimetable1MonthsEnum = "MAY"
	CreateScheduleScheduleTimetable1MonthsEnumOct CreateScheduleScheduleTimetable1MonthsEnum = "OCT"
	CreateScheduleScheduleTimetable1MonthsEnumFeb CreateScheduleScheduleTimetable1MonthsEnum = "FEB"
	CreateScheduleScheduleTimetable1MonthsEnumApr CreateScheduleScheduleTimetable1MonthsEnum = "APR"
	CreateScheduleScheduleTimetable1MonthsEnumSep CreateScheduleScheduleTimetable1MonthsEnum = "SEP"
	CreateScheduleScheduleTimetable1MonthsEnumAug CreateScheduleScheduleTimetable1MonthsEnum = "AUG"
	CreateScheduleScheduleTimetable1MonthsEnumJan CreateScheduleScheduleTimetable1MonthsEnum = "JAN"
	CreateScheduleScheduleTimetable1MonthsEnumJul CreateScheduleScheduleTimetable1MonthsEnum = "JUL"
)

type CreateScheduleScheduleTimetable1 struct {
	DaysOfMonth []int64                                          `json:"days-of-month,omitempty"`
	DaysOfWeek  []CreateScheduleScheduleTimetable1DaysOfWeekEnum `json:"days-of-week"`
	HoursOfDay  []int64                                          `json:"hours-of-day"`
	Months      []CreateScheduleScheduleTimetable1MonthsEnum     `json:"months,omitempty"`
	PerHour     int64                                            `json:"per-hour"`
}

type CreateScheduleScheduleTimetable2DaysOfWeekEnum string

const (
	CreateScheduleScheduleTimetable2DaysOfWeekEnumTue CreateScheduleScheduleTimetable2DaysOfWeekEnum = "TUE"
	CreateScheduleScheduleTimetable2DaysOfWeekEnumSat CreateScheduleScheduleTimetable2DaysOfWeekEnum = "SAT"
	CreateScheduleScheduleTimetable2DaysOfWeekEnumSun CreateScheduleScheduleTimetable2DaysOfWeekEnum = "SUN"
	CreateScheduleScheduleTimetable2DaysOfWeekEnumMon CreateScheduleScheduleTimetable2DaysOfWeekEnum = "MON"
	CreateScheduleScheduleTimetable2DaysOfWeekEnumThu CreateScheduleScheduleTimetable2DaysOfWeekEnum = "THU"
	CreateScheduleScheduleTimetable2DaysOfWeekEnumWed CreateScheduleScheduleTimetable2DaysOfWeekEnum = "WED"
	CreateScheduleScheduleTimetable2DaysOfWeekEnumFri CreateScheduleScheduleTimetable2DaysOfWeekEnum = "FRI"
)

type CreateScheduleScheduleTimetable2MonthsEnum string

const (
	CreateScheduleScheduleTimetable2MonthsEnumMar CreateScheduleScheduleTimetable2MonthsEnum = "MAR"
	CreateScheduleScheduleTimetable2MonthsEnumNov CreateScheduleScheduleTimetable2MonthsEnum = "NOV"
	CreateScheduleScheduleTimetable2MonthsEnumDec CreateScheduleScheduleTimetable2MonthsEnum = "DEC"
	CreateScheduleScheduleTimetable2MonthsEnumJun CreateScheduleScheduleTimetable2MonthsEnum = "JUN"
	CreateScheduleScheduleTimetable2MonthsEnumMay CreateScheduleScheduleTimetable2MonthsEnum = "MAY"
	CreateScheduleScheduleTimetable2MonthsEnumOct CreateScheduleScheduleTimetable2MonthsEnum = "OCT"
	CreateScheduleScheduleTimetable2MonthsEnumFeb CreateScheduleScheduleTimetable2MonthsEnum = "FEB"
	CreateScheduleScheduleTimetable2MonthsEnumApr CreateScheduleScheduleTimetable2MonthsEnum = "APR"
	CreateScheduleScheduleTimetable2MonthsEnumSep CreateScheduleScheduleTimetable2MonthsEnum = "SEP"
	CreateScheduleScheduleTimetable2MonthsEnumAug CreateScheduleScheduleTimetable2MonthsEnum = "AUG"
	CreateScheduleScheduleTimetable2MonthsEnumJan CreateScheduleScheduleTimetable2MonthsEnum = "JAN"
	CreateScheduleScheduleTimetable2MonthsEnumJul CreateScheduleScheduleTimetable2MonthsEnum = "JUL"
)

type CreateScheduleScheduleTimetable2 struct {
	DaysOfMonth []int64                                          `json:"days-of-month"`
	DaysOfWeek  []CreateScheduleScheduleTimetable2DaysOfWeekEnum `json:"days-of-week,omitempty"`
	HoursOfDay  []int64                                          `json:"hours-of-day"`
	Months      []CreateScheduleScheduleTimetable2MonthsEnum     `json:"months,omitempty"`
	PerHour     int64                                            `json:"per-hour"`
}

// CreateScheduleSchedule
// A schedule response
type CreateScheduleSchedule struct {
	Actor       CreateScheduleScheduleUser `json:"actor"`
	CreatedAt   time.Time                  `json:"created-at"`
	Description string                     `json:"description"`
	ID          string                     `json:"id"`
	Name        string                     `json:"name"`
	Parameters  map[string]interface{}     `json:"parameters"`
	ProjectSlug string                     `json:"project-slug"`
	Timetable   interface{}                `json:"timetable"`
	UpdatedAt   time.Time                  `json:"updated-at"`
}

type CreateScheduleDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateScheduleRequest struct {
	PathParams CreateSchedulePathParams
	Request    *CreateScheduleCreateScheduleParameters `request:"mediaType=application/json"`
}

type CreateScheduleResponse struct {
	ContentType                                string
	Schedule                                   *CreateScheduleSchedule
	StatusCode                                 int64
	CreateScheduleDefaultApplicationJSONObject *CreateScheduleDefaultApplicationJSON
}
