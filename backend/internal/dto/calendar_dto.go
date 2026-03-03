package dto

import (
	"zplus-calendar/internal/lunar"
)

// ─── Calendar Request DTOs ──────────────────────────────────────────

// SolarToLunarRequest is the query params for solar to lunar conversion
type SolarToLunarRequest struct {
	Day   int `query:"day" validate:"required,min=1,max=31"`
	Month int `query:"month" validate:"required,min=1,max=12"`
	Year  int `query:"year" validate:"required,min=1900,max=2100"`
}

// LunarToSolarRequest is the query params for lunar to solar conversion
type LunarToSolarRequest struct {
	Day   int  `query:"day" validate:"required,min=1,max=30"`
	Month int  `query:"month" validate:"required,min=1,max=12"`
	Year  int  `query:"year" validate:"required,min=1900,max=2100"`
	Leap  bool `query:"leap"`
}

// ─── Calendar Response DTOs ─────────────────────────────────────────

// DayInfoResponse contains full information for a single day
type DayInfoResponse struct {
	SolarDate SolarDateResponse  `json:"solar_date"`
	LunarDate LunarDateResponse  `json:"lunar_date"`
	CanChi    CanChiResponse     `json:"can_chi"`
	SolarTerm *SolarTermResponse `json:"solar_term,omitempty"`
	Holidays  []HolidayResponse  `json:"holidays"`
	GoodBad   GoodBadResponse    `json:"good_bad"`
}

// SolarDateResponse represents a solar date
type SolarDateResponse struct {
	Day           int    `json:"day"`
	Month         int    `json:"month"`
	Year          int    `json:"year"`
	DayOfWeek     int    `json:"day_of_week"`
	DayOfWeekName string `json:"day_of_week_name"`
}

// LunarDateResponse represents a lunar date
type LunarDateResponse struct {
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	Leap      bool   `json:"leap"`
	DayName   string `json:"day_name,omitempty"`
	MonthName string `json:"month_name"`
}

// CanChiResponse contains can chi information
type CanChiResponse struct {
	DayCanChi   string `json:"day"`
	MonthCanChi string `json:"month"`
	YearCanChi  string `json:"year"`
	HourCanChi  string `json:"hour,omitempty"`
	Zodiac      string `json:"zodiac"`
}

// SolarTermResponse represents a solar term
type SolarTermResponse struct {
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
}

// HolidayResponse represents a holiday
type HolidayResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsLunar     bool   `json:"is_lunar"`
	IsOfficial  bool   `json:"is_official"`
}

// GoodBadResponse contains good/bad day info
type GoodBadResponse struct {
	Quality     string             `json:"quality"`
	QualityName string             `json:"quality_name"`
	GoodFor     []string           `json:"good_for"`
	BadFor      []string           `json:"bad_for"`
	GoodHours   []GoodHourResponse `json:"good_hours"`
}

// GoodHourResponse represents a good/bad hour
type GoodHourResponse struct {
	Chi     string `json:"chi"`
	Time    string `json:"time"`
	Quality string `json:"quality"`
}

// MonthInfoResponse contains all days in a month
type MonthInfoResponse struct {
	Year  int                `json:"year"`
	Month int                `json:"month"`
	Days  []MonthDayResponse `json:"days"`
}

// MonthDayResponse is a condensed day info for month view
type MonthDayResponse struct {
	SolarDay       int    `json:"solar_day"`
	SolarMonth     int    `json:"solar_month"`
	LunarDay       int    `json:"lunar_day"`
	LunarMonth     int    `json:"lunar_month"`
	LunarLeap      bool   `json:"lunar_leap,omitempty"`
	LunarDayName   string `json:"lunar_day_name,omitempty"`
	DayOfWeek      int    `json:"day_of_week"`
	IsToday        bool   `json:"is_today,omitempty"`
	IsCurrentMonth bool   `json:"is_current_month"`
	HasHoliday     bool   `json:"has_holiday,omitempty"`
	HolidayName    string `json:"holiday_name,omitempty"`
	CanChiDay      string `json:"can_chi_day"`
}

// TodayInfoResponse is the response for today's date info
type TodayInfoResponse struct {
	DayInfoResponse
	Greeting string `json:"greeting"`
}

// ConvertResponse is the response for date conversion
type ConvertResponse struct {
	Solar SolarDateResponse `json:"solar"`
	Lunar LunarDateResponse `json:"lunar"`
}

// HolidaysYearResponse is the response for holidays in a year
type HolidaysYearResponse struct {
	Year     int                   `json:"year"`
	Holidays []HolidayDateResponse `json:"holidays"`
}

// HolidayDateResponse adds the solar date to a holiday
type HolidayDateResponse struct {
	SolarDay    int    `json:"solar_day"`
	SolarMonth  int    `json:"solar_month"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsLunar     bool   `json:"is_lunar"`
	IsOfficial  bool   `json:"is_official"`
}

// Helper to convert lunar HolidayInfo to DTO
func ToHolidayResponse(h lunar.VietnameseHolidayInfo) HolidayResponse {
	return HolidayResponse{
		Name:        h.Name,
		Description: h.Description,
		IsLunar:     h.IsLunar,
		IsOfficial:  h.IsOfficial,
	}
}
