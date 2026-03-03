"use client";

import { useState, useRef, useEffect, useCallback } from "react";
import { useRouter } from "next/navigation";
import { useCalendarStore } from "@/stores/calendar-store";
import { useDayInfo } from "@/hooks/use-calendar";
import { useTasksByDate, useCreateTask, useUpdateTaskStatus } from "@/hooks/use-tasks";
import { useNotesByDate, useCreateNote } from "@/hooks/use-notes";
import { useAuthStore } from "@/stores/auth-store";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
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
  Plus,
  Send,
  X,
  Loader2,
  Circle,
  CircleCheckBig,
} from "lucide-react";
import { cn } from "@/lib/utils";
import type { TaskPriority } from "@/types";

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

/* ─── TaskItem: toggleable status ──────────────────────────────── */

function TaskItem({ task }: { task: { id: string; title: string; status: string; priority: string } }) {
  const updateStatus = useUpdateTaskStatus();
  const isDone = task.status === "done";

  const toggle = useCallback(() => {
    updateStatus.mutate({
      id: task.id,
      status: isDone ? "todo" : "done",
    });
  }, [task.id, isDone, updateStatus]);

  return (
    <div
      className={cn(
        "flex items-center gap-2 rounded-xl px-2.5 py-2 text-xs group",
        isDone
          ? "bg-green-50 dark:bg-green-950/20 text-muted-foreground"
          : task.priority === "urgent" || task.priority === "high"
          ? "bg-red-50 dark:bg-red-950/20"
          : "bg-blue-50 dark:bg-blue-950/20"
      )}
    >
      <button
        onClick={toggle}
        disabled={updateStatus.isPending}
        className="shrink-0 text-muted-foreground hover:text-foreground transition-colors active:scale-90"
      >
        {updateStatus.isPending ? (
          <Loader2 className="h-3.5 w-3.5 animate-spin" />
        ) : isDone ? (
          <CircleCheckBig className="h-3.5 w-3.5 text-green-500" />
        ) : (
          <Circle className="h-3.5 w-3.5" />
        )}
      </button>
      <span className={cn("truncate flex-1", isDone && "line-through")}>{task.title}</span>
      <Badge variant="outline" className="text-[9px] shrink-0 h-4">
        {isDone ? "Xong" : task.status === "in_progress" ? "Đang làm" : "Chờ"}
      </Badge>
    </div>
  );
}

/* ─── Quick Add Task ───────────────────────────────────────────── */

const priorities: { value: TaskPriority; label: string; color: string }[] = [
  { value: "low", label: "Thấp", color: "bg-gray-400" },
  { value: "medium", label: "TB", color: "bg-blue-500" },
  { value: "high", label: "Cao", color: "bg-orange-500" },
  { value: "urgent", label: "GẤP", color: "bg-red-500" },
];

function QuickAddTask({ dateStr }: { dateStr: string }) {
  const [expanded, setExpanded] = useState(false);
  const [title, setTitle] = useState("");
  const [priority, setPriority] = useState<TaskPriority>("medium");
  const createTask = useCreateTask();
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (expanded) inputRef.current?.focus();
  }, [expanded]);

  const handleSubmit = () => {
    const trimmed = title.trim();
    if (!trimmed) return;
    createTask.mutate(
      { title: trimmed, due_date: dateStr, priority },
      {
        onSuccess: () => {
          setTitle("");
          setPriority("medium");
          // keep form open for rapid entry
        },
      }
    );
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      handleSubmit();
    }
    if (e.key === "Escape") {
      setTitle("");
      setExpanded(false);
    }
  };

  const handleClose = () => {
    setTitle("");
    setPriority("medium");
    setExpanded(false);
  };

  if (!expanded) {
    return (
      <button
        onClick={() => setExpanded(true)}
        className="w-full flex items-center gap-1.5 py-2 px-2.5 rounded-xl text-xs text-muted-foreground hover:text-foreground hover:bg-blue-50 dark:hover:bg-blue-950/20 transition-colors active:scale-[0.98]"
      >
        <Plus className="h-3.5 w-3.5" />
        Thêm task nhanh…
      </button>
    );
  }

  return (
    <div className="space-y-2 rounded-xl bg-blue-50/50 dark:bg-blue-950/10 p-2.5 border border-blue-200/50 dark:border-blue-800/30">
      <div className="flex items-center gap-1.5">
        <Input
          ref={inputRef}
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          onKeyDown={handleKeyDown}
          placeholder="Tiêu đề task…"
          className="h-8 text-xs rounded-lg flex-1 bg-background"
          disabled={createTask.isPending}
        />
        <Button
          size="icon"
          variant="ghost"
          className="h-8 w-8 shrink-0 text-muted-foreground hover:text-foreground"
          onClick={handleClose}
        >
          <X className="h-3.5 w-3.5" />
        </Button>
      </div>

      {/* Priority chips */}
      <div className="flex items-center gap-1">
        <span className="text-[10px] text-muted-foreground mr-1">Ưu tiên:</span>
        {priorities.map((p) => (
          <button
            key={p.value}
            onClick={() => setPriority(p.value)}
            className={cn(
              "px-2 py-0.5 rounded-full text-[10px] font-medium transition-all active:scale-95",
              priority === p.value
                ? `${p.color} text-white shadow-sm`
                : "bg-muted text-muted-foreground hover:bg-accent"
            )}
          >
            {p.label}
          </button>
        ))}
      </div>

      <Button
        size="sm"
        className="w-full h-8 rounded-lg text-xs"
        disabled={!title.trim() || createTask.isPending}
        onClick={handleSubmit}
      >
        {createTask.isPending ? (
          <Loader2 className="h-3.5 w-3.5 mr-1 animate-spin" />
        ) : (
          <Send className="h-3.5 w-3.5 mr-1" />
        )}
        Tạo task
      </Button>
    </div>
  );
}

/* ─── Quick Add Note ───────────────────────────────────────────── */

const noteColors: { value: string; css: string }[] = [
  { value: "", css: "bg-gray-300 dark:bg-gray-600" },
  { value: "red", css: "bg-red-400" },
  { value: "orange", css: "bg-orange-400" },
  { value: "yellow", css: "bg-yellow-400" },
  { value: "green", css: "bg-green-400" },
  { value: "blue", css: "bg-blue-400" },
  { value: "purple", css: "bg-purple-400" },
  { value: "pink", css: "bg-pink-400" },
];

function QuickAddNote({ dateStr }: { dateStr: string }) {
  const [expanded, setExpanded] = useState(false);
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [color, setColor] = useState("");
  const createNote = useCreateNote();
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (expanded) inputRef.current?.focus();
  }, [expanded]);

  const handleSubmit = () => {
    const trimmedTitle = title.trim();
    if (!trimmedTitle) return;
    createNote.mutate(
      {
        title: trimmedTitle,
        content: content.trim() || undefined,
        date: dateStr,
        color: color || undefined,
      },
      {
        onSuccess: () => {
          setTitle("");
          setContent("");
          setColor("");
          // keep form open for rapid entry
        },
      }
    );
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" && (e.metaKey || e.ctrlKey)) {
      e.preventDefault();
      handleSubmit();
    }
    if (e.key === "Escape") {
      setTitle("");
      setContent("");
      setExpanded(false);
    }
  };

  const handleClose = () => {
    setTitle("");
    setContent("");
    setColor("");
    setExpanded(false);
  };

  if (!expanded) {
    return (
      <button
        onClick={() => setExpanded(true)}
        className="w-full flex items-center gap-1.5 py-2 px-2.5 rounded-xl text-xs text-muted-foreground hover:text-foreground hover:bg-emerald-50 dark:hover:bg-emerald-950/20 transition-colors active:scale-[0.98]"
      >
        <Plus className="h-3.5 w-3.5" />
        Thêm ghi chú nhanh…
      </button>
    );
  }

  return (
    <div className="space-y-2 rounded-xl bg-emerald-50/50 dark:bg-emerald-950/10 p-2.5 border border-emerald-200/50 dark:border-emerald-800/30">
      <div className="flex items-center gap-1.5">
        <Input
          ref={inputRef}
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          onKeyDown={handleKeyDown}
          placeholder="Tiêu đề ghi chú…"
          className="h-8 text-xs rounded-lg flex-1 bg-background"
          disabled={createNote.isPending}
        />
        <Button
          size="icon"
          variant="ghost"
          className="h-8 w-8 shrink-0 text-muted-foreground hover:text-foreground"
          onClick={handleClose}
        >
          <X className="h-3.5 w-3.5" />
        </Button>
      </div>

      <Textarea
        value={content}
        onChange={(e) => setContent(e.target.value)}
        onKeyDown={handleKeyDown}
        placeholder="Nội dung (tuỳ chọn)… ⌘+Enter để lưu"
        className="text-xs rounded-lg bg-background resize-none min-h-[60px]"
        rows={2}
        disabled={createNote.isPending}
      />

      {/* Color chips */}
      <div className="flex items-center gap-1">
        <span className="text-[10px] text-muted-foreground mr-1">Màu:</span>
        {noteColors.map((c) => (
          <button
            key={c.value}
            onClick={() => setColor(c.value)}
            className={cn(
              "h-5 w-5 rounded-full transition-all active:scale-90",
              c.css,
              color === c.value
                ? "ring-2 ring-offset-1 ring-foreground/40 scale-110"
                : "opacity-60 hover:opacity-100"
            )}
          />
        ))}
      </div>

      <Button
        size="sm"
        className="w-full h-8 rounded-lg text-xs bg-emerald-600 hover:bg-emerald-700 text-white"
        disabled={!title.trim() || createNote.isPending}
        onClick={handleSubmit}
      >
        {createNote.isPending ? (
          <Loader2 className="h-3.5 w-3.5 mr-1 animate-spin" />
        ) : (
          <Send className="h-3.5 w-3.5 mr-1" />
        )}
        Tạo ghi chú
      </Button>
    </div>
  );
}