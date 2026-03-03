"use client";

import { useState, useRef, useEffect, useCallback } from "react";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import {
  Plus,
  Send,
  X,
  Loader2,
  Circle,
  CircleCheckBig,
} from "lucide-react";
import { cn } from "@/lib/utils";
import { useCreateTask, useUpdateTaskStatus } from "@/hooks/use-tasks";
import { useCreateNote } from "@/hooks/use-notes";
import type { TaskPriority } from "@/types";

/* ─── TaskItem: toggleable status ──────────────────────────────── */

export function TaskItem({
  task,
}: {
  task: { id: string; title: string; description?: string; status: string; priority: string };
}) {
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
      <div className="flex-1 min-w-0">
        <span className={cn("truncate block", isDone && "line-through")}>{task.title}</span>
        {task.description && (
          <p className="text-[10px] text-muted-foreground line-clamp-1">{task.description}</p>
        )}
      </div>
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

export function QuickAddTask({ dateStr }: { dateStr: string }) {
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

export function QuickAddNote({ dateStr }: { dateStr: string }) {
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
