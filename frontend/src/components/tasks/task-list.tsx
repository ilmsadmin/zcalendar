"use client";

import { TaskItem } from "./task-item";
import { Skeleton } from "@/components/ui/skeleton";
import { CheckSquare } from "lucide-react";
import type { Task } from "@/types";

interface TaskListProps {
  tasks: Task[];
  isLoading: boolean;
  onEdit?: (task: Task) => void;
}

export function TaskList({ tasks, isLoading, onEdit }: TaskListProps) {
  if (isLoading) {
    return (
      <div className="space-y-2">
        {Array.from({ length: 5 }).map((_, i) => (
          <Skeleton key={i} className="h-20 w-full rounded-lg" />
        ))}
      </div>
    );
  }

  if (tasks.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center py-12 text-center">
        <CheckSquare className="h-12 w-12 text-muted-foreground/50 mb-3" />
        <h3 className="text-lg font-medium text-muted-foreground">
          Chưa có công việc
        </h3>
        <p className="text-sm text-muted-foreground/70">
          Tạo công việc đầu tiên bằng nút &quot;Thêm mới&quot;
        </p>
      </div>
    );
  }

  return (
    <div className="space-y-2">
      {tasks.map((task) => (
        <TaskItem key={task.id} task={task} onEdit={onEdit} />
      ))}
    </div>
  );
}
