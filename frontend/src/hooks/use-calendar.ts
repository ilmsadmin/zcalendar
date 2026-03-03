"use client";

import { useQuery } from "@tanstack/react-query";
import { calendarApi } from "@/lib/api";
import { useCalendarStore } from "@/stores/calendar-store";

export function useToday() {
  return useQuery({
    queryKey: ["calendar", "today"],
    queryFn: () => calendarApi.getToday().then((r) => r.data.data),
    staleTime: 60 * 1000, // 1 min
  });
}

export function useDayInfo(day: number, month: number, year: number) {
  return useQuery({
    queryKey: ["calendar", "day", day, month, year],
    queryFn: () => calendarApi.getDay(day, month, year).then((r) => r.data.data),
    enabled: day > 0 && month > 0 && year > 0,
    staleTime: 5 * 60 * 1000, // 5 min
  });
}

export function useMonthInfo(year?: number, month?: number) {
  const store = useCalendarStore();
  const y = year ?? store.currentYear;
  const m = month ?? store.currentMonth;

  return useQuery({
    queryKey: ["calendar", "month", y, m],
    queryFn: () => calendarApi.getMonth(y, m).then((r) => r.data.data),
    staleTime: 5 * 60 * 1000,
  });
}

export function useHolidays(year: number) {
  return useQuery({
    queryKey: ["calendar", "holidays", year],
    queryFn: () => calendarApi.getHolidays(year).then((r) => r.data.data),
    staleTime: 24 * 60 * 60 * 1000, // 1 day
  });
}
