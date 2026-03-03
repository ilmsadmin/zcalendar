"use client";

import { useRouter } from "next/navigation";
import { useCalendarStore } from "@/stores/calendar-store";
import { useDayInfo } from "@/hooks/use-calendar";
import { useTasksByDate } from "@/hooks/use-tasks";
import { useNotesByDate } from "@/hooks/use-notes";
import { useAuthStore } from "@/stores/auth-store";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Moon,
  Sparkles,
  Clock,
  ThumbsUp,
  ThumbsDown,
  ExternalLink,
  CheckCircle2,
  StickyNote,
} from "lucide-react";
import { cn } from "@/lib/utils";
import { TaskItem, QuickAddTask, QuickAddNote } from "./quick-add-panels";

export function CalendarDayView() {
  const { selectedDate } = useCalendarStore();
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  const router = useRouter();
  const day = selectedDate.getDate();
  const month = selectedDate.getMonth() + 1;
  const year = selectedDate.getFullYear();

  const { data: dayInfo, isLoading } = useDayInfo(day, month, year);
  const dateStr = `${year}-${String(month).padStart(2, "0")}-${String(day).padStart(2, "0")}`;
  const { data: tasks } = useTasksByDate(isAuthenticated ? dateStr : "");
  const { data: notes } = useNotesByDate(isAuthenticated ? dateStr : "");

  if (isLoading) {
    return (
      <div className="space-y-3">
        <Skeleton className="h-28 rounded-2xl" />
        <Skeleton className="h-20 rounded-2xl" />
        <Skeleton className="h-24 rounded-2xl" />
      </div>
    );
  }

  if (!dayInfo) return null;

  const { solar_date, lunar_date, can_chi, solar_term, holidays, good_bad } = dayInfo;

  return (
    <div className="space-y-3">
      {/* Hero date card */}
      <div className="rounded-2xl bg-gradient-to-br from-amber-50 to-orange-50 dark:from-amber-950/30 dark:to-orange-950/20 p-4">
        <div className="flex items-center gap-4">
          {/* Big day number */}
          <div className="text-center shrink-0">
            <p className="text-5xl md:text-6xl font-black text-amber-600 dark:text-amber-400 leading-none">
              {solar_date.day}
            </p>
            <p className="text-xs font-medium text-muted-foreground mt-1">
              {solar_date.day_of_week_name}
            </p>
          </div>
          <div className="flex-1 min-w-0 space-y-1">
            <p className="text-sm font-semibold">
              Tháng {solar_date.month}, {solar_date.year}
            </p>
            <div className="flex items-center gap-1.5 text-xs text-muted-foreground">
              <Moon className="h-3 w-3 text-amber-500 shrink-0" />
              <span className="truncate">
                ÂL: {lunar_date.day_name || `${lunar_date.day}`}/{lunar_date.month_name}
                {lunar_date.leap ? " (Nhuận)" : ""}
              </span>
            </div>
            <div className="flex flex-wrap gap-1 mt-1">
              <Badge
                variant={
                  good_bad.quality === "good" ? "default" :
                  good_bad.quality === "bad" ? "destructive" : "secondary"
                }
                className="text-[10px] h-5"
              >
                {good_bad.quality_name}
              </Badge>
              {solar_term && (
                <Badge variant="outline" className="text-[10px] h-5 border-green-200 text-green-700 dark:text-green-400">
                  🌿 {solar_term.name}
                </Badge>
              )}
            </div>
          </div>
        </div>
      </div>

      {/* Holidays */}
      {holidays.length > 0 && (
        <div className="rounded-2xl bg-red-50/60 dark:bg-red-950/20 p-3 space-y-1.5">
          {holidays.map((h, i) => (
            <div key={i} className="flex items-center gap-2">
              <span className="text-sm">🎉</span>
              <Badge variant={h.is_official ? "destructive" : "secondary"} className="text-[10px] h-5">
                {h.name}
              </Badge>
            </div>
          ))}
        </div>
      )}

      {/* Can Chi - compact horizontal */}
      <div className="rounded-2xl bg-card border p-3">
        <div className="flex items-center gap-1.5 mb-2">
          <Sparkles className="h-3.5 w-3.5 text-purple-500" />
          <span className="text-xs font-semibold">Can Chi</span>
        </div>
        <div className="grid grid-cols-3 gap-2">
          <div className="text-center rounded-xl bg-purple-50 dark:bg-purple-950/30 py-2 px-1">
            <p className="text-[9px] text-muted-foreground">Ngày</p>
            <p className="text-xs font-bold text-purple-600 dark:text-purple-400">{can_chi.day}</p>
          </div>
          <div className="text-center rounded-xl bg-blue-50 dark:bg-blue-950/30 py-2 px-1">
            <p className="text-[9px] text-muted-foreground">Tháng</p>
            <p className="text-xs font-bold text-blue-600 dark:text-blue-400">{can_chi.month}</p>
          </div>
          <div className="text-center rounded-xl bg-amber-50 dark:bg-amber-950/30 py-2 px-1">
            <p className="text-[9px] text-muted-foreground">Năm</p>
            <p className="text-xs font-bold text-amber-600 dark:text-amber-400">{can_chi.year}</p>
          </div>
        </div>
        <p className="text-center text-[10px] text-muted-foreground mt-1.5">
          🐾 Tuổi {can_chi.zodiac}
        </p>
      </div>

      {/* Good/Bad - compact */}
      <div className="rounded-2xl bg-card border p-3 space-y-2">
        {good_bad.good_for.length > 0 && (
          <div>
            <p className="text-[10px] font-semibold text-green-600 dark:text-green-400 flex items-center gap-1 mb-1">
              <ThumbsUp className="h-3 w-3" /> Nên làm
            </p>
            <div className="flex flex-wrap gap-1">
              {good_bad.good_for.map((g, i) => (
                <Badge key={i} variant="outline" className="text-[9px] h-5 border-green-200">{g}</Badge>
              ))}
            </div>
          </div>
        )}
        {good_bad.bad_for.length > 0 && (
          <div>
            <p className="text-[10px] font-semibold text-red-600 dark:text-red-400 flex items-center gap-1 mb-1">
              <ThumbsDown className="h-3 w-3" /> Không nên
            </p>
            <div className="flex flex-wrap gap-1">
              {good_bad.bad_for.map((b, i) => (
                <Badge key={i} variant="outline" className="text-[9px] h-5 border-red-200">{b}</Badge>
              ))}
            </div>
          </div>
        )}
      </div>

      {/* Good Hours - compact 2-col */}
      <div className="rounded-2xl bg-card border p-3">
        <div className="flex items-center gap-1.5 mb-2">
          <Clock className="h-3.5 w-3.5 text-green-500" />
          <span className="text-xs font-semibold">Giờ hoàng đạo</span>
        </div>
        <div className="grid grid-cols-3 gap-1.5">
          {good_bad.good_hours
            .filter((h) => h.quality === "hoang_dao")
            .map((h, i) => (
              <div
                key={i}
                className="text-center rounded-lg bg-green-50 dark:bg-green-950/30 py-1.5 px-1"
              >
                <p className="text-xs font-bold text-green-700 dark:text-green-400">{h.chi}</p>
                <p className="text-[9px] text-green-600 dark:text-green-500">{h.time}</p>
              </div>
            ))}
        </div>
      </div>

      {/* Tasks & Notes */}
      {isAuthenticated && (
        <>
          {/* Tasks */}
          <div className="rounded-2xl bg-card border p-3">
            <div className="flex items-center justify-between mb-2">
              <div className="flex items-center gap-1.5">
                <CheckCircle2 className="h-3.5 w-3.5 text-blue-500" />
                <span className="text-xs font-semibold">
                  Task {tasks && tasks.length > 0 ? `(${tasks.length})` : ""}
                </span>
              </div>
            </div>

            {/* Existing tasks */}
            {tasks && tasks.length > 0 && (
              <div className="space-y-1.5 mb-2">
                {tasks.map((t) => (
                  <TaskItem key={t.id} task={t} />
                ))}
              </div>
            )}

            {/* Quick add task */}
            <QuickAddTask dateStr={dateStr} />
          </div>

          {/* Notes */}
          <div className="rounded-2xl bg-card border p-3">
            <div className="flex items-center justify-between mb-2">
              <div className="flex items-center gap-1.5">
                <StickyNote className="h-3.5 w-3.5 text-emerald-500" />
                <span className="text-xs font-semibold">
                  Ghi chú {notes && notes.length > 0 ? `(${notes.length})` : ""}
                </span>
              </div>
            </div>

            {/* Existing notes */}
            {notes && notes.length > 0 && (
              <div className="space-y-1.5 mb-2">
                {notes.map((n) => (
                  <div key={n.id} className="rounded-xl px-2.5 py-2 bg-emerald-50 dark:bg-emerald-950/20">
                    <p className="text-xs font-medium">{n.title}</p>
                    {n.content && (
                      <p className="text-[10px] text-muted-foreground line-clamp-2 mt-0.5">{n.content}</p>
                    )}
                  </div>
                ))}
              </div>
            )}

            {/* Quick add note */}
            <QuickAddNote dateStr={dateStr} />
          </div>
        </>
      )}

      {/* View full detail button */}
      <Button
        variant="outline"
        className="w-full h-10 rounded-xl active:scale-[0.98] transition-transform text-sm"
        onClick={() => router.push(`/calendar/day/${dateStr}`)}
      >
        <ExternalLink className="mr-2 h-4 w-4" />
        Xem trang chi tiết ngày
      </Button>
    </div>
  );
}

