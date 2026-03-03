"use client";

import { useMemo } from "react";
import { useRouter } from "next/navigation";
import { useMonthInfo } from "@/hooks/use-calendar";
import { useTasksByDateRange } from "@/hooks/use-tasks";
import { useNotesByDateRange } from "@/hooks/use-notes";
import { useCalendarStore } from "@/stores/calendar-store";
import { CalendarDay } from "./calendar-day";
import { Skeleton } from "@/components/ui/skeleton";

const dayNames = ["CN", "T2", "T3", "T4", "T5", "T6", "T7"];

export function CalendarGrid() {
  const { currentYear, currentMonth, selectedDate, setSelectedDate } =
    useCalendarStore();
  const { data: monthInfo, isLoading } = useMonthInfo(currentYear, currentMonth);
  const router = useRouter();

  // Date range for the visible month (include overflow days)
  const { dateFrom, dateTo } = useMemo(() => {
    const from = new Date(currentYear, currentMonth - 2, 20);
    const to = new Date(currentYear, currentMonth, 12);
    const fmt = (d: Date) =>
      `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, "0")}-${String(d.getDate()).padStart(2, "0")}`;
    return { dateFrom: fmt(from), dateTo: fmt(to) };
  }, [currentYear, currentMonth]);

  const { data: taskMap } = useTasksByDateRange(dateFrom, dateTo);
  const { data: noteMap } = useNotesByDateRange(dateFrom, dateTo);

  if (isLoading) {
    return (
      <div className="space-y-1">
        <div className="grid grid-cols-7">
          {dayNames.map((d) => (
            <div key={d} className="py-2 text-center text-[11px] font-medium text-muted-foreground">
              {d}
            </div>
          ))}
        </div>
        <div className="grid grid-cols-7 gap-px">
          {Array.from({ length: 42 }).map((_, i) => (
            <Skeleton key={i} className="h-[64px] sm:h-[80px] rounded-xl" />
          ))}
        </div>
      </div>
    );
  }

  if (!monthInfo) return null;

  const selectedDay = selectedDate.getDate();
  const selectedMonth = selectedDate.getMonth() + 1;
  const selectedYear = selectedDate.getFullYear();

  return (
    <div className="space-y-0">
      {/* Day names header */}
      <div className="grid grid-cols-7">
        {dayNames.map((d, i) => (
          <div
            key={d}
            className={`py-2 text-center text-[11px] font-semibold ${
              i === 0 ? "text-red-500" : i === 6 ? "text-blue-500" : "text-muted-foreground"
            }`}
          >
            {d}
          </div>
        ))}
      </div>

      {/* Calendar cells — tight gap for edge-to-edge feel */}
      <div className="grid grid-cols-7 gap-px">
        {monthInfo.days.map((day, index) => {
          const y = day.solar_month > currentMonth && day.solar_day < 8 ? currentYear - 1
            : day.solar_month < currentMonth && day.solar_day > 20 ? currentYear + 1
            : currentYear;
          const dateKey = `${y}-${String(day.solar_month).padStart(2, "0")}-${String(day.solar_day).padStart(2, "0")}`;
          const tasksForDay = taskMap?.[dateKey];
          const notesForDay = noteMap?.[dateKey];
          return (
            <CalendarDay
              key={index}
              day={day}
              isSelected={
                day.solar_day === selectedDay &&
                day.solar_month === selectedMonth &&
                selectedYear === currentYear &&
                day.is_current_month
              }
              taskCount={tasksForDay?.length ?? 0}
              noteCount={notesForDay?.length ?? 0}
              onClick={() => {
                setSelectedDate(
                  new Date(currentYear, day.solar_month - 1, day.solar_day)
                );
              }}
              onNavigate={() => {
                router.push(`/calendar/day/${dateKey}`);
              }}
            />
          );
        })}
      </div>
    </div>
  );
}
