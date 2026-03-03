"use client";

import { use, useMemo } from "react";
import { useRouter } from "next/navigation";
import { useDayInfo } from "@/hooks/use-calendar";
import { useTasksByDate } from "@/hooks/use-tasks";
import { useNotesByDate } from "@/hooks/use-notes";
import { useAuthStore } from "@/stores/auth-store";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import { cn } from "@/lib/utils";
import {
  ArrowLeft,
  ChevronLeft,
  ChevronRight,
  Moon,
  Sparkles,
  Star,
  Clock,
  ThumbsUp,
  ThumbsDown,
  SunMedium,
  CheckCircle2,
  StickyNote,
} from "lucide-react";
import { TaskItem, QuickAddTask, QuickAddNote } from "@/components/calendar/quick-add-panels";

function parseDate(dateStr: string) {
  const [y, m, d] = dateStr.split("-").map(Number);
  return { day: d, month: m, year: y };
}

function formatDate(d: number, m: number, y: number) {
  return `${y}-${String(m).padStart(2, "0")}-${String(d).padStart(2, "0")}`;
}

function addDays(day: number, month: number, year: number, offset: number) {
  const d = new Date(year, month - 1, day + offset);
  return { day: d.getDate(), month: d.getMonth() + 1, year: d.getFullYear() };
}

export default function DayDetailPage({
  params,
}: {
  params: Promise<{ date: string }>;
}) {
  const resolvedParams = use(params);
  const router = useRouter();
  const { day, month, year } = parseDate(resolvedParams.date);
  const { data: dayInfo, isLoading } = useDayInfo(day, month, year);
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  const { data: tasks } = useTasksByDate(isAuthenticated ? resolvedParams.date : "");
  const { data: notes } = useNotesByDate(isAuthenticated ? resolvedParams.date : "");

  const prev = addDays(day, month, year, -1);
  const next = addDays(day, month, year, 1);

  const isToday = useMemo(() => {
    const now = new Date();
    return day === now.getDate() && month === now.getMonth() + 1 && year === now.getFullYear();
  }, [day, month, year]);

  if (isLoading) {
    return (
      <div className="px-3 py-3 space-y-3">
        <Skeleton className="h-9 w-32 rounded-xl" />
        <Skeleton className="h-36 rounded-2xl" />
        <Skeleton className="h-24 rounded-2xl" />
        <Skeleton className="h-24 rounded-2xl" />
      </div>
    );
  }

  if (!dayInfo) {
    return (
      <div className="px-3 py-3">
        <Button variant="ghost" size="sm" onClick={() => router.push("/calendar")} className="active:scale-95">
          <ArrowLeft className="mr-2 h-4 w-4" /> Quay lại
        </Button>
        <p className="text-muted-foreground mt-4 text-sm">Không tìm thấy thông tin ngày này.</p>
      </div>
    );
  }

  const { solar_date, lunar_date, can_chi, solar_term, holidays, good_bad } = dayInfo;

  return (
    <div className="px-3 py-3 md:px-4 lg:px-6 space-y-3 max-w-3xl mx-auto pb-24">
      {/* Top navigation — sticky */}
      <div className="flex items-center justify-between">
        <Button
          variant="ghost"
          size="sm"
          onClick={() => router.push("/calendar")}
          className="h-9 px-2 active:scale-95 transition-transform"
        >
          <ArrowLeft className="mr-1.5 h-4 w-4" /> Lịch
        </Button>
        <div className="flex items-center gap-1">
          <Button
            variant="ghost"
            size="icon"
            className="h-9 w-9 active:scale-90 transition-transform"
            onClick={() => router.push(`/calendar/day/${formatDate(prev.day, prev.month, prev.year)}`)}
          >
            <ChevronLeft className="h-5 w-5" />
          </Button>
          {!isToday && (
            <Button
              variant="outline"
              size="sm"
              className="h-8 text-xs rounded-full px-3 active:scale-95 transition-transform"
              onClick={() => {
                const now = new Date();
                router.push(`/calendar/day/${formatDate(now.getDate(), now.getMonth() + 1, now.getFullYear())}`);
              }}
            >
              Hôm nay
            </Button>
          )}
          <Button
            variant="ghost"
            size="icon"
            className="h-9 w-9 active:scale-90 transition-transform"
            onClick={() => router.push(`/calendar/day/${formatDate(next.day, next.month, next.year)}`)}
          >
            <ChevronRight className="h-5 w-5" />
          </Button>
        </div>
      </div>

      {/* Hero banner */}
      <div className={cn(
        "rounded-2xl overflow-hidden p-4 md:p-6",
        isToday
          ? "bg-gradient-to-br from-amber-500 via-orange-500 to-red-500 text-white"
          : "bg-gradient-to-br from-amber-50 to-orange-50 dark:from-amber-950/30 dark:to-orange-950/20"
      )}>
        <div className="flex items-center gap-4">
          {/* Large date */}
          <div className="text-center shrink-0">
            <p className={cn(
              "text-6xl md:text-7xl font-black leading-none",
              isToday ? "text-white" : "text-amber-600 dark:text-amber-400"
            )}>
              {solar_date.day}
            </p>
            <p className={cn(
              "text-xs font-semibold mt-1",
              isToday ? "text-white/80" : "text-muted-foreground"
            )}>
              {solar_date.day_of_week_name}
            </p>
          </div>

          <div className="flex-1 min-w-0 space-y-1.5">
            <p className={cn(
              "text-base font-bold",
              isToday ? "text-white" : "text-foreground"
            )}>
              Tháng {solar_date.month}, {solar_date.year}
            </p>

            <div className={cn(
              "flex items-center gap-1.5 text-xs",
              isToday ? "text-white/85" : "text-muted-foreground"
            )}>
              <Moon className="h-3.5 w-3.5 shrink-0" />
              <span className="truncate">
                ÂL: {lunar_date.day_name || `${lunar_date.day}`} {lunar_date.month_name}
                {lunar_date.leap ? " (Nhuận)" : ""}
              </span>
            </div>

            <div className="flex flex-wrap gap-1.5 mt-1">
              {isToday && (
                <Badge className="bg-white/20 text-white border-white/30 text-[10px] h-5">
                  📍 Hôm nay
                </Badge>
              )}
              <Badge
                variant={
                  good_bad.quality === "good" ? "default" :
                  good_bad.quality === "bad" ? "destructive" : "secondary"
                }
                className={cn("text-[10px] h-5", isToday && "bg-white/20 text-white border-white/30")}
              >
                {good_bad.quality_name}
              </Badge>
              {solar_term && (
                <Badge variant="outline" className={cn(
                  "text-[10px] h-5",
                  isToday ? "border-white/30 text-white" : "border-green-200 text-green-700 dark:text-green-400"
                )}>
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
          <div className="flex items-center gap-1.5 mb-1">
            <span className="text-sm">🎉</span>
            <span className="text-xs font-semibold">Ngày lễ</span>
          </div>
          {holidays.map((h, i) => (
            <div key={i} className="flex items-start gap-2">
              <Badge variant={h.is_official ? "destructive" : "secondary"} className="text-[10px] h-5 shrink-0">
                {h.is_official ? "Nghỉ lễ" : "Văn hoá"}
              </Badge>
              <div className="min-w-0">
                <p className="text-xs font-medium">{h.name}</p>
                {h.description && (
                  <p className="text-[10px] text-muted-foreground">{h.description}</p>
                )}
              </div>
            </div>
          ))}
        </div>
      )}

      {/* Can Chi */}
      <div className="rounded-2xl bg-card border p-3">
        <div className="flex items-center gap-1.5 mb-2">
          <Sparkles className="h-3.5 w-3.5 text-purple-500" />
          <span className="text-xs font-semibold">Can Chi</span>
        </div>
        <div className="grid grid-cols-3 gap-2">
          <div className="text-center rounded-xl bg-purple-50 dark:bg-purple-950/30 py-2.5 px-1.5">
            <p className="text-[10px] text-muted-foreground">Ngày</p>
            <p className="text-sm font-bold text-purple-600 dark:text-purple-400">{can_chi.day}</p>
          </div>
          <div className="text-center rounded-xl bg-blue-50 dark:bg-blue-950/30 py-2.5 px-1.5">
            <p className="text-[10px] text-muted-foreground">Tháng</p>
            <p className="text-sm font-bold text-blue-600 dark:text-blue-400">{can_chi.month}</p>
          </div>
          <div className="text-center rounded-xl bg-amber-50 dark:bg-amber-950/30 py-2.5 px-1.5">
            <p className="text-[10px] text-muted-foreground">Năm</p>
            <p className="text-sm font-bold text-amber-600 dark:text-amber-400">{can_chi.year}</p>
          </div>
        </div>
        <p className="text-center text-[10px] text-muted-foreground mt-2">
          🐾 Tuổi {can_chi.zodiac}
        </p>
      </div>

      {/* Good/Bad */}
      <div className="rounded-2xl bg-card border p-3 space-y-2.5">
        <div className="flex items-center gap-1.5">
          <SunMedium className="h-3.5 w-3.5 text-amber-500" />
          <span className="text-xs font-semibold">Ngày tốt / xấu</span>
        </div>

        {good_bad.good_for.length > 0 && (
          <div>
            <p className="text-[10px] font-semibold text-green-600 dark:text-green-400 flex items-center gap-1 mb-1.5">
              <ThumbsUp className="h-3 w-3" /> Nên làm
            </p>
            <div className="flex flex-wrap gap-1.5">
              {good_bad.good_for.map((g, i) => (
                <Badge key={i} variant="outline" className="text-[10px] h-5 border-green-200 dark:border-green-800 bg-green-50/50 dark:bg-green-950/20">
                  ✅ {g}
                </Badge>
              ))}
            </div>
          </div>
        )}

        {good_bad.bad_for.length > 0 && (
          <div>
            <p className="text-[10px] font-semibold text-red-600 dark:text-red-400 flex items-center gap-1 mb-1.5">
              <ThumbsDown className="h-3 w-3" /> Không nên
            </p>
            <div className="flex flex-wrap gap-1.5">
              {good_bad.bad_for.map((b, i) => (
                <Badge key={i} variant="outline" className="text-[10px] h-5 border-red-200 dark:border-red-800 bg-red-50/50 dark:bg-red-950/20">
                  ❌ {b}
                </Badge>
              ))}
            </div>
          </div>
        )}
      </div>

      {/* Good Hours */}
      <div className="rounded-2xl bg-card border p-3">
        <div className="flex items-center gap-1.5 mb-2">
          <Clock className="h-3.5 w-3.5 text-green-500" />
          <span className="text-xs font-semibold">12 Giờ trong ngày</span>
        </div>
        <div className="grid grid-cols-3 sm:grid-cols-4 md:grid-cols-6 gap-1.5">
          {good_bad.good_hours.map((h, i) => (
            <div
              key={i}
              className={cn(
                "rounded-xl py-2 px-1.5 text-center transition-all",
                h.quality === "hoang_dao"
                  ? "bg-green-50 dark:bg-green-950/30 border border-green-200 dark:border-green-800"
                  : "bg-muted/30"
              )}
            >
              <p className={cn(
                "text-sm font-bold",
                h.quality === "hoang_dao"
                  ? "text-green-600 dark:text-green-400"
                  : "text-muted-foreground"
              )}>
                {h.chi}
              </p>
              <p className={cn(
                "text-[9px] mt-0.5",
                h.quality === "hoang_dao"
                  ? "text-green-500"
                  : "text-muted-foreground/60"
              )}>
                {h.time}
              </p>
              <p className="text-[8px] mt-0.5">
                {h.quality === "hoang_dao" ? "✨" : ""}
              </p>
            </div>
          ))}
        </div>
      </div>

      {/* Solar term mobile */}
      {solar_term && (
        <div className="rounded-2xl bg-green-50 dark:bg-green-950/20 p-3 flex items-center gap-2.5 md:hidden">
          <Star className="h-4 w-4 text-green-500 shrink-0" />
          <div>
            <p className="text-xs font-medium">{solar_term.name}</p>
            <p className="text-[10px] text-muted-foreground">{solar_term.name_en}</p>
          </div>
        </div>
      )}

      {/* Tasks & Notes */}
      {isAuthenticated && (
        <div className="space-y-3">
          {/* Tasks */}
          <div className="rounded-2xl bg-card border p-3">
            <div className="flex items-center justify-between mb-2">
              <div className="flex items-center gap-1.5">
                <CheckCircle2 className="h-3.5 w-3.5 text-blue-500" />
                <span className="text-xs font-semibold">Task ({tasks?.length ?? 0})</span>
              </div>
            </div>
            <div className="space-y-1.5">
              {tasks && tasks.length > 0 ? (
                tasks.map((t) => (
                  <TaskItem key={t.id} task={t} />
                ))
              ) : (
                <p className="text-xs text-muted-foreground py-2 text-center">Không có task nào</p>
              )}
            </div>
            {/* Quick add task */}
            <QuickAddTask dateStr={resolvedParams.date} />
          </div>

          {/* Notes */}
          <div className="rounded-2xl bg-card border p-3">
            <div className="flex items-center justify-between mb-2">
              <div className="flex items-center gap-1.5">
                <StickyNote className="h-3.5 w-3.5 text-emerald-500" />
                <span className="text-xs font-semibold">Ghi chú ({notes?.length ?? 0})</span>
              </div>
            </div>
            <div className="space-y-1.5">
              {notes && notes.length > 0 ? (
                notes.map((n) => (
                  <div key={n.id} className="rounded-xl px-2.5 py-2 bg-emerald-50 dark:bg-emerald-950/20">
                    <p className="text-xs font-medium">{n.title}</p>
                    {n.content && (
                      <p className="text-[10px] text-muted-foreground line-clamp-2 mt-0.5">{n.content}</p>
                    )}
                  </div>
                ))
              ) : (
                <p className="text-xs text-muted-foreground py-2 text-center">Không có ghi chú nào</p>
              )}
            </div>
            {/* Quick add note */}
            <QuickAddNote dateStr={resolvedParams.date} />
          </div>
        </div>
      )}
    </div>
  );
}
