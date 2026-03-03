"use client";

import { useDayInfo } from "@/hooks/use-calendar";
import { useCalendarStore } from "@/stores/calendar-store";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Moon,
  Clock,
  ThumbsUp,
  ThumbsDown,
  Sparkles,
  ChevronDown,
} from "lucide-react";
import { useState } from "react";
import { cn } from "@/lib/utils";

export function LunarInfoPanel() {
  const selectedDate = useCalendarStore((s) => s.selectedDate);
  const day = selectedDate.getDate();
  const month = selectedDate.getMonth() + 1;
  const year = selectedDate.getFullYear();
  const [expanded, setExpanded] = useState(false);

  const { data: dayInfo, isLoading } = useDayInfo(day, month, year);

  if (isLoading) {
    return (
      <div className="rounded-2xl border bg-card p-3 space-y-2">
        <Skeleton className="h-5 w-40" />
        <Skeleton className="h-4 w-full" />
        <Skeleton className="h-4 w-3/4" />
      </div>
    );
  }

  if (!dayInfo) return null;

  const { solar_date, lunar_date, can_chi, solar_term, holidays, good_bad } =
    dayInfo;

  return (
    <div className="rounded-2xl border bg-card overflow-hidden">
      {/* Header — always visible, tappable on mobile */}
      <button
        className="w-full p-3 flex items-center justify-between text-left lg:cursor-default active:bg-accent/50 lg:active:bg-transparent transition-colors"
        onClick={() => setExpanded(!expanded)}
      >
        <div className="flex items-center gap-2 min-w-0">
          <Moon className="h-4 w-4 text-amber-500 shrink-0" />
          <div className="min-w-0">
            <p className="text-sm font-semibold truncate">
              {solar_date.day_of_week_name}, {solar_date.day}/{solar_date.month}/
              {solar_date.year}
            </p>
            <p className="text-xs text-amber-600 dark:text-amber-400">
              ÂL: {lunar_date.day_name || `${lunar_date.day}`}/
              {lunar_date.month_name}
              {lunar_date.leap ? " (Nhuận)" : ""}
            </p>
          </div>
        </div>
        <div className="flex items-center gap-2 shrink-0">
          <Badge
            variant={
              good_bad.quality === "good"
                ? "default"
                : good_bad.quality === "bad"
                ? "destructive"
                : "secondary"
            }
            className="text-[10px] h-5"
          >
            {good_bad.quality_name}
          </Badge>
          <ChevronDown
            className={cn(
              "h-4 w-4 text-muted-foreground transition-transform lg:hidden",
              expanded && "rotate-180"
            )}
          />
        </div>
      </button>

      {/* Expandable content on mobile, always visible on lg */}
      <div
        className={cn(
          "overflow-hidden transition-all duration-200",
          "lg:max-h-none lg:opacity-100",
          expanded
            ? "max-h-[800px] opacity-100"
            : "max-h-0 opacity-0 lg:max-h-none lg:opacity-100"
        )}
      >
        <div className="px-3 pb-3 space-y-3 border-t pt-3">
          {/* Can Chi */}
          <div>
            <div className="flex items-center gap-1.5 mb-1.5">
              <Sparkles className="h-3 w-3 text-purple-500" />
              <span className="text-[10px] font-semibold uppercase text-muted-foreground">
                Can Chi
              </span>
            </div>
            <div className="grid grid-cols-3 gap-1.5 text-xs">
              <div className="flex justify-between rounded-lg bg-muted/50 px-2 py-1.5">
                <span className="text-muted-foreground">Ngày</span>
                <span className="font-semibold">{can_chi.day}</span>
              </div>
              <div className="flex justify-between rounded-lg bg-muted/50 px-2 py-1.5">
                <span className="text-muted-foreground">Tháng</span>
                <span className="font-semibold">{can_chi.month}</span>
              </div>
              <div className="flex justify-between rounded-lg bg-muted/50 px-2 py-1.5">
                <span className="text-muted-foreground">Năm</span>
                <span className="font-semibold">{can_chi.year}</span>
              </div>
            </div>
            <p className="text-[10px] text-center text-muted-foreground mt-1">
              🐾 {can_chi.zodiac}
            </p>
          </div>

          {/* Solar term */}
          {solar_term && (
            <div className="flex items-center gap-2 rounded-lg bg-green-50 dark:bg-green-950/20 px-2.5 py-2">
              <span className="text-sm">🌿</span>
              <div>
                <p className="text-xs font-medium">{solar_term.name}</p>
                <p className="text-[10px] text-muted-foreground">
                  {solar_term.name_en}
                </p>
              </div>
            </div>
          )}

          {/* Holidays */}
          {holidays.length > 0 && (
            <div className="space-y-1">
              {holidays.map((h, i) => (
                <div key={i} className="flex items-center gap-2">
                  <span className="text-xs">🎉</span>
                  <Badge
                    variant={h.is_official ? "destructive" : "secondary"}
                    className="text-[10px] h-5"
                  >
                    {h.name}
                  </Badge>
                </div>
              ))}
            </div>
          )}

          {/* Good/Bad */}
          {good_bad.good_for.length > 0 && (
            <div>
              <p className="text-[10px] font-semibold text-green-600 dark:text-green-400 flex items-center gap-1 mb-1">
                <ThumbsUp className="h-2.5 w-2.5" /> Nên làm
              </p>
              <div className="flex flex-wrap gap-1">
                {good_bad.good_for.map((g, i) => (
                  <Badge
                    key={i}
                    variant="outline"
                    className="text-[9px] h-4.5 border-green-200"
                  >
                    {g}
                  </Badge>
                ))}
              </div>
            </div>
          )}

          {good_bad.bad_for.length > 0 && (
            <div>
              <p className="text-[10px] font-semibold text-red-600 dark:text-red-400 flex items-center gap-1 mb-1">
                <ThumbsDown className="h-2.5 w-2.5" /> Không nên
              </p>
              <div className="flex flex-wrap gap-1">
                {good_bad.bad_for.map((b, i) => (
                  <Badge
                    key={i}
                    variant="outline"
                    className="text-[9px] h-4.5 border-red-200"
                  >
                    {b}
                  </Badge>
                ))}
              </div>
            </div>
          )}

          {/* Good Hours - compact */}
          {good_bad.good_hours.filter((h) => h.quality === "hoang_dao")
            .length > 0 && (
            <div>
              <div className="flex items-center gap-1.5 mb-1.5">
                <Clock className="h-3 w-3 text-green-500" />
                <span className="text-[10px] font-semibold uppercase text-muted-foreground">
                  Giờ hoàng đạo
                </span>
              </div>
              <div className="flex flex-wrap gap-1.5">
                {good_bad.good_hours
                  .filter((h) => h.quality === "hoang_dao")
                  .map((h, i) => (
                    <div
                      key={i}
                      className="rounded-lg bg-green-50 dark:bg-green-950/30 px-2 py-1 text-center"
                    >
                      <span className="text-[10px] font-bold text-green-700 dark:text-green-400">
                        {h.chi}
                      </span>
                      <span className="text-[9px] text-green-600 dark:text-green-500 ml-1">
                        {h.time}
                      </span>
                    </div>
                  ))}
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
