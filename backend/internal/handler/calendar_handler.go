package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"zplus-calendar/internal/common"
	"zplus-calendar/internal/service"
)

// CalendarHandler handles calendar HTTP requests
type CalendarHandler struct {
	calendarService service.CalendarService
	logger          *zap.Logger
}

// NewCalendarHandler creates a new CalendarHandler
func NewCalendarHandler(calendarService service.CalendarService, logger *zap.Logger) *CalendarHandler {
	return &CalendarHandler{
		calendarService: calendarService,
		logger:          logger,
	}
}

// GetDayInfo returns detailed info for a specific date
// @Summary Get day info
// @Tags calendar
// @Produce json
// @Param day query int true "Day"
// @Param month query int true "Month"
// @Param year query int true "Year"
// @Success 200 {object} common.APIResponse{data=dto.DayInfoResponse}
// @Router /calendar/day [get]
func (h *CalendarHandler) GetDayInfo(c *fiber.Ctx) error {
	day, err := strconv.Atoi(c.Query("day"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "day không hợp lệ")
	}
	month, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "month không hợp lệ")
	}
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "year không hợp lệ")
	}

	resp, err := h.calendarService.GetDayInfo(day, month, year)
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_DATE", err.Error())
	}

	return common.Success(c, resp)
}

// GetMonthInfo returns calendar data for an entire month
// @Summary Get month info
// @Tags calendar
// @Produce json
// @Param year path int true "Year"
// @Param month path int true "Month"
// @Success 200 {object} common.APIResponse{data=dto.MonthInfoResponse}
// @Router /calendar/month/{year}/{month} [get]
func (h *CalendarHandler) GetMonthInfo(c *fiber.Ctx) error {
	year, err := strconv.Atoi(c.Params("year"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "year không hợp lệ")
	}
	month, err := strconv.Atoi(c.Params("month"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "month không hợp lệ")
	}

	resp, err := h.calendarService.GetMonthInfo(year, month)
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_MONTH", err.Error())
	}

	return common.Success(c, resp)
}

// GetTodayInfo returns today's calendar info
// @Summary Get today info
// @Tags calendar
// @Produce json
// @Success 200 {object} common.APIResponse{data=dto.TodayInfoResponse}
// @Router /calendar/today [get]
func (h *CalendarHandler) GetTodayInfo(c *fiber.Ctx) error {
	resp, err := h.calendarService.GetTodayInfo()
	if err != nil {
		return common.InternalError(c, "Không thể lấy thông tin hôm nay")
	}

	return common.Success(c, resp)
}

// SolarToLunar converts solar date to lunar date
// @Summary Convert solar to lunar
// @Tags calendar
// @Produce json
// @Param day query int true "Day"
// @Param month query int true "Month"
// @Param year query int true "Year"
// @Success 200 {object} common.APIResponse{data=dto.ConvertResponse}
// @Router /calendar/convert/solar-to-lunar [get]
func (h *CalendarHandler) SolarToLunar(c *fiber.Ctx) error {
	day, _ := strconv.Atoi(c.Query("day"))
	month, _ := strconv.Atoi(c.Query("month"))
	year, _ := strconv.Atoi(c.Query("year"))

	resp, err := h.calendarService.SolarToLunar(day, month, year)
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_DATE", err.Error())
	}

	return common.Success(c, resp)
}

// LunarToSolar converts lunar date to solar date
// @Summary Convert lunar to solar
// @Tags calendar
// @Produce json
// @Param day query int true "Day"
// @Param month query int true "Month"
// @Param year query int true "Year"
// @Param leap query bool false "Is leap month"
// @Success 200 {object} common.APIResponse{data=dto.ConvertResponse}
// @Router /calendar/convert/lunar-to-solar [get]
func (h *CalendarHandler) LunarToSolar(c *fiber.Ctx) error {
	day, _ := strconv.Atoi(c.Query("day"))
	month, _ := strconv.Atoi(c.Query("month"))
	year, _ := strconv.Atoi(c.Query("year"))
	leap := c.Query("leap") == "true"

	resp, err := h.calendarService.LunarToSolar(day, month, year, leap)
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_DATE", err.Error())
	}

	return common.Success(c, resp)
}

// GetHolidays returns all holidays for a given year
// @Summary Get holidays by year
// @Tags calendar
// @Produce json
// @Param year path int true "Year"
// @Success 200 {object} common.APIResponse{data=dto.HolidaysYearResponse}
// @Router /calendar/holidays/{year} [get]
func (h *CalendarHandler) GetHolidays(c *fiber.Ctx) error {
	year, err := strconv.Atoi(c.Params("year"))
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "year không hợp lệ")
	}

	resp, err := h.calendarService.GetHolidaysByYear(year)
	if err != nil {
		return common.Error(c, fiber.StatusBadRequest, "INVALID_YEAR", err.Error())
	}

	return common.Success(c, resp)
}
