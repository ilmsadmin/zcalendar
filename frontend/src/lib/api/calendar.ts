import api from "@/lib/api-client";
import type {
  APIResponse,
  DayInfo,
  TodayInfo,
  MonthInfo,
  ConvertResult,
  HolidaysYear,
} from "@/types";

export const calendarApi = {
  getToday: () => api.get<APIResponse<TodayInfo>>("/calendar/today"),

  getDay: (day: number, month: number, year: number) =>
    api.get<APIResponse<DayInfo>>("/calendar/day", {
      params: { day, month, year },
    }),

  getMonth: (year: number, month: number) =>
    api.get<APIResponse<MonthInfo>>(`/calendar/month/${year}/${month}`),

  solarToLunar: (day: number, month: number, year: number) =>
    api.get<APIResponse<ConvertResult>>("/calendar/convert/solar-to-lunar", {
      params: { day, month, year },
    }),

  lunarToSolar: (
    day: number,
    month: number,
    year: number,
    leap: boolean = false
  ) =>
    api.get<APIResponse<ConvertResult>>("/calendar/convert/lunar-to-solar", {
      params: { day, month, year, leap },
    }),

  getHolidays: (year: number) =>
    api.get<APIResponse<HolidaysYear>>(`/calendar/holidays/${year}`),
};
