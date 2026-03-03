"use client";

import { cn } from "@/lib/utils";
import { Checkbox } from "@/components/ui/checkbox";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { MoreHorizontal, Pencil, Trash2, Calendar } from "lucide-react";
import type { Task, TaskPriority } from "@/types";
import { useUpdateTaskStatus, useDeleteTask } from "@/hooks/use-tasks";

const priorityColors: Record<TaskPriority, string> = {
  low: "bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-400",
  medium:
    "bg-blue-100 text-blue-700 dark:bg-blue-900/30 dark:text-blue-400",
  high: "bg-orange-100 text-orange-700 dark:bg-orange-900/30 dark:text-orange-400",
  urgent: "bg-red-100 text-red-700 dark:bg-red-900/30 dark:text-red-400",
};

const priorityLabels: Record<TaskPriority, string> = {
  low: "Thấp",
  medium: "Trung bình",
  high: "Cao",
  urgent: "Khẩn cấp",
};

const statusLabels: Record<string, string> = {
  todo: "Cần làm",
  in_progress: "Đang làm",
  done: "Hoàn thành",
  cancelled: "Đã hủy",
};

interface TaskItemProps {
  task: Task;
  onEdit?: (task: Task) => void;
}

export function TaskItem({ task, onEdit }: TaskItemProps) {
  const updateStatus = useUpdateTaskStatus();
  const deleteTask = useDeleteTask();

  const isDone = task.status === "done";

  const handleToggle = () => {
    const newStatus = isDone ? "todo" : "done";
    updateStatus.mutate({ id: task.id, status: newStatus });
  };

  return (
    <div
      className={cn(
        "group flex items-start gap-3 rounded-lg border p-3 transition-colors hover:bg-accent/30 active:bg-accent/50 cursor-pointer",
        isDone && "opacity-60"
      )}
      onClick={() => onEdit?.(task)}
    >
      <Checkbox
        checked={isDone}
        onCheckedChange={handleToggle}
        className="mt-0.5"
        onClick={(e) => e.stopPropagation()}
      />

      <div className="flex-1 min-w-0">
        <div className="flex items-center gap-2">
          <h4
            className={cn(
              "text-sm font-medium truncate",
              isDone && "line-through text-muted-foreground"
            )}
          >
            {task.title}
          </h4>
        </div>

        {task.description && (
          <p className="text-xs text-muted-foreground mt-0.5 line-clamp-1">
            {task.description}
          </p>
        )}

        <div className="flex items-center gap-2 mt-2 flex-wrap">
          <Badge
            variant="secondary"
            className={cn("text-[10px]", priorityColors[task.priority])}
          >
            {priorityLabels[task.priority]}
          </Badge>

          <Badge variant="outline" className="text-[10px]">
            {statusLabels[task.status]}
          </Badge>

          {task.due_date && (
            <span className="flex items-center gap-0.5 text-[10px] text-muted-foreground">
              <Calendar className="h-3 w-3" />
              {task.due_date}
            </span>
          )}

          {task.category && (
            <Badge
              variant="secondary"
              className="text-[10px]"
              style={{ backgroundColor: task.category.color + "20", color: task.category.color }}
            >
              {task.category.name}
            </Badge>
          )}

          {task.tags?.map((tag) => (
            <Badge key={tag.id} variant="outline" className="text-[10px]">
              {tag.name}
            </Badge>
          ))}
        </div>
      </div>

      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            variant="ghost"
            size="icon"
            className="h-8 w-8 shrink-0 opacity-100 sm:opacity-0 sm:group-hover:opacity-100"
            onClick={(e) => e.stopPropagation()}
          >
            <MoreHorizontal className="h-4 w-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuItem onClick={(e) => { e.stopPropagation(); onEdit?.(task); }}>
            <Pencil className="mr-2 h-4 w-4" />
            Chỉnh sửa
          </DropdownMenuItem>
          <DropdownMenuItem
            onClick={(e) => { e.stopPropagation(); deleteTask.mutate(task.id); }}
            className="text-red-600"
          >
            <Trash2 className="mr-2 h-4 w-4" />
            Xoá
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
}
