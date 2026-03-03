package service

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"zplus-calendar/internal/dto"
	"zplus-calendar/internal/lunar"
)

// CalendarService defines the interface for calendar operations
type CalendarService interface {
	GetDayInfo(day, month, year int) (*dto.DayInfoResponse, error)
	GetMonthInfo(year, month int) (*dto.MonthInfoResponse, error)
	GetTodayInfo() (*dto.TodayInfoResponse, error)
	SolarToLunar(day, month, year int) (*dto.ConvertResponse, error)
	LunarToSolar(day, month, year int, leap bool) (*dto.ConvertResponse, error)
	GetHolidaysByYear(year int) (*dto.HolidaysYearResponse, error)
}

type calendarService struct {
	logger *zap.Logger
}

// NewCalendarService creates a new CalendarService
func NewCalendarService(logger *zap.Logger) CalendarService {
	return &calendarService{logger: logger}
}

func (s *calendarService) GetDayInfo(day, month, year int) (*dto.DayInfoResponse, error) {
	// Validate date
	if !isValidDate(day, month, year) {
		return nil, fmt.Errorf("invalid date: %d/%d/%d", day, month, year)
	}

	// Solar date info
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	dow := int(t.Weekday())
	dowNames := []string{"Chủ nhật", "Thứ hai", "Thứ ba", "Thứ tư", "Thứ năm", "Thứ sáu", "Thứ bảy"}

	solarDate := dto.SolarDateResponse{
		Day:           day,
		Month:         month,
		Year:          year,
		DayOfWeek:     dow,
		DayOfWeekName: dowNames[dow],
	}

	// Lunar date
	lunarDate := lunar.SolarToLunar(day, month, year)
	lunarResp := dto.LunarDateResponse{
		Day:       lunarDate.Day,
		Month:     lunarDate.Month,
		Year:      lunarDate.Year,
		Leap:      lunarDate.Leap,
		DayName:   lunarDate.DayName,
		MonthName: lunarDate.MonthName,
	}

	// Can Chi
	canChiInfo := lunar.GetFullCanChi(day, month, year, lunarDate)
	canChiResp := dto.CanChiResponse{
		DayCanChi:   canChiInfo.DayFull,
		MonthCanChi: canChiInfo.MonthFull,
		YearCanChi:  canChiInfo.YearFull,
		Zodiac:      canChiInfo.Zodiac,
	}

	// Solar term
	var solarTermResp *dto.SolarTermResponse
	if st := lunar.GetSolarTerm(day, month, year); st != nil {
		solarTermResp = &dto.SolarTermResponse{
			Name:   st.Name,
			NameEn: st.NameEn,
		}
	}

	// Holidays
	holidays := lunar.GetHolidaysForSolarDate(day, month, year)
	holidayResps := make([]dto.HolidayResponse, 0, len(holidays))
	for _, h := range holidays {
		holidayResps = append(holidayResps, dto.ToHolidayResponse(h))
	}

	// Good/Bad day
	goodBad := lunar.GetGoodBadDay(day, month, year)
	goodHours := make([]dto.GoodHourResponse, 0, len(goodBad.GoodHours))
	for _, h := range goodBad.GoodHours {
		goodHours = append(goodHours, dto.GoodHourResponse{
			Chi:     h.Chi,
			Time:    h.Time,
			Quality: h.Quality,
		})
	}

	return &dto.DayInfoResponse{
		SolarDate: solarDate,
		LunarDate: lunarResp,
		CanChi:    canChiResp,
		SolarTerm: solarTermResp,
		Holidays:  holidayResps,
		GoodBad: dto.GoodBadResponse{
			Quality:     goodBad.Quality,
			QualityName: goodBad.QualityName,
			GoodFor:     goodBad.GoodFor,
			BadFor:      goodBad.BadFor,
			GoodHours:   goodHours,
		},
	}, nil
}

func (s *calendarService) GetMonthInfo(year, month int) (*dto.MonthInfoResponse, error) {
	if month < 1 || month > 12 || year < 1900 || year > 2100 {
		return nil, fmt.Errorf("invalid month: %d/%d", month, year)
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	// Get the first day of the month
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	// Get the last day of the month
	lastDay := firstDay.AddDate(0, 1, -1)

	// Calculate start of calendar (fill previous month days to align to week)
	startDOW := int(firstDay.Weekday())
	calStart := firstDay.AddDate(0, 0, -startDOW)

	// Calculate end of calendar (fill next month days)
	endDOW := int(lastDay.Weekday())
	calEnd := lastDay.AddDate(0, 0, 6-endDOW)

	var days []dto.MonthDayResponse
	for d := calStart; !d.After(calEnd); d = d.AddDate(0, 0, 1) {
		sDay := d.Day()
		sMonth := int(d.Month())
		sYear := d.Year()

		lunarDate := lunar.SolarToLunar(sDay, sMonth, sYear)
		jd := lunar.JulianDayFromDate(sDay, sMonth, sYear)
		dayCan, dayChi := lunar.CanChiOfDay(jd)

		holidays := lunar.GetHolidaysForSolarDate(sDay, sMonth, sYear)
		hasHoliday := len(holidays) > 0
		holidayName := ""
		if hasHoliday {
			holidayName = holidays[0].Name
		}

		days = append(days, dto.MonthDayResponse{
			SolarDay:       sDay,
			SolarMonth:     sMonth,
			LunarDay:       lunarDate.Day,
			LunarMonth:     lunarDate.Month,
			LunarLeap:      lunarDate.Leap,
			LunarDayName:   lunarDate.DayName,
			DayOfWeek:      int(d.Weekday()),
			IsToday:        d.Equal(today),
			IsCurrentMonth: sMonth == month,
			HasHoliday:     hasHoliday,
			HolidayName:    holidayName,
			CanChiDay:      dayCan + " " + dayChi,
		})
	}

	return &dto.MonthInfoResponse{
		Year:  year,
		Month: month,
		Days:  days,
	}, nil
}

func (s *calendarService) GetTodayInfo() (*dto.TodayInfoResponse, error) {
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	vnNow := now.In(loc)

	dayInfo, err := s.GetDayInfo(vnNow.Day(), int(vnNow.Month()), vnNow.Year())
	if err != nil {
		return nil, err
	}

	hour := vnNow.Hour()
	var greeting string
	switch {
	case hour < 6:
		greeting = "Khuya rồi, nghỉ ngơi thôi! 🌙"
	case hour < 11:
		greeting = "Chào buổi sáng! ☀️"
	case hour < 13:
		greeting = "Chào buổi trưa! 🌤"
	case hour < 18:
		greeting = "Chào buổi chiều! 🌇"
	default:
		greeting = "Chào buổi tối! 🌙"
	}

	return &dto.TodayInfoResponse{
		DayInfoResponse: *dayInfo,
		Greeting:        greeting,
	}, nil
}

func (s *calendarService) SolarToLunar(day, month, year int) (*dto.ConvertResponse, error) {
	if !isValidDate(day, month, year) {
		return nil, fmt.Errorf("invalid date: %d/%d/%d", day, month, year)
	}

	lunarDate := lunar.SolarToLunar(day, month, year)
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	dow := int(t.Weekday())
	dowNames := []string{"Chủ nhật", "Thứ hai", "Thứ ba", "Thứ tư", "Thứ năm", "Thứ sáu", "Thứ bảy"}

	return &dto.ConvertResponse{
		Solar: dto.SolarDateResponse{
			Day:           day,
			Month:         month,
			Year:          year,
			DayOfWeek:     dow,
			DayOfWeekName: dowNames[dow],
		},
		Lunar: dto.LunarDateResponse{
			Day:       lunarDate.Day,
			Month:     lunarDate.Month,
			Year:      lunarDate.Year,
			Leap:      lunarDate.Leap,
			DayName:   lunarDate.DayName,
			MonthName: lunarDate.MonthName,
		},
	}, nil
}

func (s *calendarService) LunarToSolar(day, month, year int, leap bool) (*dto.ConvertResponse, error) {
	sDay, sMonth, sYear := lunar.LunarToSolar(day, month, year, leap)
	if sDay == 0 {
		return nil, fmt.Errorf("invalid lunar date: %d/%d/%d (leap=%v)", day, month, year, leap)
	}

	t := time.Date(sYear, time.Month(sMonth), sDay, 0, 0, 0, 0, time.UTC)
	dow := int(t.Weekday())
	dowNames := []string{"Chủ nhật", "Thứ hai", "Thứ ba", "Thứ tư", "Thứ năm", "Thứ sáu", "Thứ bảy"}

	monthName := lunar.SolarToLunar(sDay, sMonth, sYear).MonthName

	return &dto.ConvertResponse{
		Solar: dto.SolarDateResponse{
			Day:           sDay,
			Month:         sMonth,
			Year:          sYear,
			DayOfWeek:     dow,
			DayOfWeekName: dowNames[dow],
		},
		Lunar: dto.LunarDateResponse{
			Day:       day,
			Month:     month,
			Year:      year,
			Leap:      leap,
			MonthName: monthName,
		},
	}, nil
}

func (s *calendarService) GetHolidaysByYear(year int) (*dto.HolidaysYearResponse, error) {
	allHolidays := lunar.GetAllHolidaysInYear(year)

	holidays := make([]dto.HolidayDateResponse, 0, len(allHolidays))
	for _, h := range allHolidays {
		holidays = append(holidays, dto.HolidayDateResponse{
			SolarDay:    h.SolarDay,
			SolarMonth:  h.SolarMonth,
			Name:        h.Holiday.Name,
			Description: h.Holiday.Description,
			IsLunar:     h.Holiday.IsLunar,
			IsOfficial:  h.Holiday.IsOfficial,
		})
	}

	return &dto.HolidaysYearResponse{
		Year:     year,
		Holidays: holidays,
	}, nil
}

// isValidDate checks if a date is valid
func isValidDate(day, month, year int) bool {
	if year < 1900 || year > 2100 || month < 1 || month > 12 || day < 1 {
		return false
	}
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return t.Day() == day && int(t.Month()) == month && t.Year() == year
}
