"use client";

import { cn } from "@/lib/utils";
import { useDoubleTap } from "@/hooks/use-double-tap";
import type { MonthDay } from "@/types";

const lunarDayNames: Record<number, string> = {
  1: "Mùng 1",
  15: "Rằm",
};

interface CalendarDayProps {
  day: MonthDay;
  isSelected: boolean;
  taskCount?: number;
  noteCount?: number;
  onClick: () => void;
  onNavigate?: () => void;
}

export function CalendarDay({ day, isSelected, taskCount = 0, noteCount = 0, onClick, onNavigate }: CalendarDayProps) {
  const { onClick: handleTap } = useDoubleTap({
    onSingleTap: onClick,
    onDoubleTap: onNavigate,
  });
  const isLunarFirst = day.lunar_day === 1;
  const isLunarFull = day.lunar_day === 15;
  const isSunday = day.day_of_week === 0;
  const isSaturday = day.day_of_week === 6;

  const lunarLabel =
    lunarDayNames[day.lunar_day] ||
    day.lunar_day_name ||
    `${day.lunar_day}`;

  return (
    <button
      onClick={handleTap}
      className={cn(
        "relative flex flex-col items-center justify-start gap-0.5 rounded-xl p-1 text-center transition-all",
        "min-h-[64px] sm:min-h-[80px]",
        "active:scale-95 active:bg-accent/70",
        !day.is_current_month && "opacity-30",
        day.is_today && !isSelected &&
          "bg-primary/8 ring-1.5 ring-primary/50",
        isSelected &&
          "bg-primary/12 ring-2 ring-primary shadow-sm",
        day.has_holiday && !isSelected && !day.is_today && "bg-red-50/60 dark:bg-red-950/20"
      )}
    >
      {/* Solar day */}
      <span
        className={cn(
          "text-sm font-semibold leading-tight sm:text-base",
          isSunday && "text-red-500",
          isSaturday && "text-blue-500",
          day.is_today && "text-primary font-bold",
          isSelected && "text-primary"
        )}
      >
        {day.solar_day}
      </span>

      {/* Lunar day */}
      <span
        className={cn(
          "text-[9px] leading-tight sm:text-[11px]",
          isLunarFirst
            ? "font-semibold text-red-500"
            : isLunarFull
            ? "font-medium text-amber-600"
            : "text-muted-foreground/70"
        )}
      >
        {isLunarFirst ? `${day.lunar_day}/${day.lunar_month}` : lunarLabel}
      </span>

      {/* Holiday name (desktop only) */}
      {day.has_holiday && (
        <span className="text-[7px] text-red-500 font-medium truncate max-w-full px-0.5 hidden sm:inline-block leading-tight">
          {day.holiday_name}
        </span>
      )}

      {/* Dot indicators */}
      {(day.has_holiday || taskCount > 0 || noteCount > 0) && (
        <div className="flex gap-[3px] items-center mt-auto pb-0.5">
          {day.has_holiday && (
            <span className="h-1 w-1 rounded-full bg-red-500 sm:hidden" />
          )}
          {taskCount > 0 && (
            <span className="h-1 w-1 rounded-full bg-blue-500" />
          )}
          {noteCount > 0 && (
            <span className="h-1 w-1 rounded-full bg-emerald-500" />
          )}
        </div>
      )}
    </button>
  );
}
