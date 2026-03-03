"use client";

import { useMemo } from "react";
import { useCalendarStore, getWeekStart } from "@/stores/calendar-store";
import { useDayInfo } from "@/hooks/use-calendar";
import { useTasksByDateRange } from "@/hooks/use-tasks";
import { useNotesByDateRange } from "@/hooks/use-notes";
import { cn } from "@/lib/utils";
import { Badge } from "@/components/ui/badge";
import { useRouter } from "next/navigation";
import { useDoubleTap } from "@/hooks/use-double-tap";
import type { Task, Note } from "@/types";

const dayNamesShort = ["CN", "T2", "T3", "T4", "T5", "T6", "T7"];

function WeekDayCard({ date, isSelected, onSelect, onNavigate, tasks, notes }: {
  date: Date;
  isSelected: boolean;
  onSelect: () => void;
  onNavigate: () => void;
  tasks?: Task[];
  notes?: Note[];
}) {
  const day = date.getDate();
  const month = date.getMonth() + 1;
  const year = date.getFullYear();
  const dow = date.getDay();

  const { data: dayInfo, isLoading } = useDayInfo(day, month, year);

  const isToday = useMemo(() => {
    const now = new Date();
    return day === now.getDate() && month === now.getMonth() + 1 && year === now.getFullYear();
  }, [day, month, year]);

  const dateStr = `${year}-${String(month).padStart(2, "0")}-${String(day).padStart(2, "0")}`;

  const { onClick: handleTap } = useDoubleTap({
    onSingleTap: onSelect,
    onDoubleTap: onNavigate,
  });

  return (
    <button
      className={cn(
        "flex flex-col items-center rounded-2xl p-2.5 min-w-[80px] transition-all",
        "active:scale-95 shrink-0 snap-center",
        isToday && !isSelected && "ring-1.5 ring-primary/50 bg-primary/5",
        isSelected
          ? "bg-primary text-primary-foreground shadow-lg shadow-primary/25 scale-[1.02]"
          : "bg-card hover:bg-accent/50",
      )}
      onClick={handleTap}
    >
      {/* Day of week */}
      <span className={cn(
        "text-[10px] font-semibold uppercase tracking-wider",
        isSelected ? "text-primary-foreground/80" : dow === 0 ? "text-red-500" : dow === 6 ? "text-blue-500" : "text-muted-foreground"
      )}>
        {dayNamesShort[dow]}
      </span>

      {/* Day number */}
      <span className={cn(
        "text-2xl font-bold leading-tight mt-0.5",
        isSelected ? "text-primary-foreground" : isToday ? "text-primary" : dow === 0 ? "text-red-500" : dow === 6 ? "text-blue-500" : ""
      )}>
        {day}
      </span>

      {/* Lunar info */}
      {dayInfo && (
        <span className={cn(
          "text-[9px] mt-0.5",
          isSelected ? "text-primary-foreground/70" : "text-amber-600 dark:text-amber-400"
        )}>
          {dayInfo.lunar_date.day}/{dayInfo.lunar_date.month} ÂL
        </span>
      )}

      {/* Quality badge */}
      {dayInfo && (
        <div className={cn(
          "mt-1.5 rounded-full px-2 py-0.5 text-[8px] font-medium",
          isSelected
            ? "bg-primary-foreground/20 text-primary-foreground"
            : dayInfo.good_bad.quality === "good"
            ? "bg-green-100 dark:bg-green-950/40 text-green-700 dark:text-green-400"
            : dayInfo.good_bad.quality === "bad"
            ? "bg-red-100 dark:bg-red-950/40 text-red-700 dark:text-red-400"
            : "bg-gray-100 dark:bg-gray-800 text-muted-foreground"
        )}>
          {dayInfo.good_bad.quality_name}
        </div>
      )}

      {/* Holiday indicator */}
      {dayInfo && dayInfo.holidays.length > 0 && (
        <span className={cn(
          "text-[8px] font-medium mt-1 truncate max-w-[72px]",
          isSelected ? "text-primary-foreground/80" : "text-red-500"
        )}>
          🎉 {dayInfo.holidays[0].name}
        </span>
      )}

      {/* Task/note dots */}
      {((tasks && tasks.length > 0) || (notes && notes.length > 0)) && (
        <div className="flex gap-1 mt-1.5">
          {tasks && tasks.length > 0 && (
            <span className={cn("h-1.5 w-1.5 rounded-full", isSelected ? "bg-primary-foreground/60" : "bg-blue-500")} />
          )}
          {notes && notes.length > 0 && (
            <span className={cn("h-1.5 w-1.5 rounded-full", isSelected ? "bg-primary-foreground/60" : "bg-emerald-500")} />
          )}
        </div>
      )}

      {isLoading && (
        <div className="mt-1">
          <div className="h-2 w-2 animate-spin rounded-full border border-current border-t-transparent opacity-40" />
        </div>
      )}
    </button>
  );
}

export function CalendarWeek() {
  const { selectedDate, setSelectedDate } = useCalendarStore();
  const router = useRouter();

  const weekDates = useMemo(() => {
    const start = getWeekStart(selectedDate);
    return Array.from({ length: 7 }, (_, i) => {
      const d = new Date(start);
      d.setDate(d.getDate() + i);
      return d;
    });
  }, [selectedDate]);

  // Date range for the week
  const { dateFrom, dateTo } = useMemo(() => {
    const fmt = (d: Date) =>
      `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, "0")}-${String(d.getDate()).padStart(2, "0")}`;
    return { dateFrom: fmt(weekDates[0]), dateTo: fmt(weekDates[6]) };
  }, [weekDates]);

  const { data: taskMap } = useTasksByDateRange(dateFrom, dateTo);
  const { data: noteMap } = useNotesByDateRange(dateFrom, dateTo);

  return (
    <div className="flex gap-2 overflow-x-auto pb-2 snap-x snap-mandatory scrollbar-none -mx-1 px-1">
      {weekDates.map((date, i) => {
        const dateKey = `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, "0")}-${String(date.getDate()).padStart(2, "0")}`;
        return (
          <WeekDayCard
            key={i}
            date={date}
            isSelected={
              date.getDate() === selectedDate.getDate() &&
              date.getMonth() === selectedDate.getMonth() &&
              date.getFullYear() === selectedDate.getFullYear()
            }
            onSelect={() => setSelectedDate(new Date(date))}
            onNavigate={() => router.push(`/calendar/day/${dateKey}`)}
            tasks={taskMap?.[dateKey]}
            notes={noteMap?.[dateKey]}
          />
        );
      })}
    </div>
  );
}
