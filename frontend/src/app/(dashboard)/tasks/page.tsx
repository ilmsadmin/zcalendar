"use client";

import { useState, useMemo } from "react";
import { Plus } from "lucide-react";
import { Button } from "@/components/ui/button";
import { TaskList } from "@/components/tasks/task-list";
import { TaskFilters } from "@/components/tasks/task-filters";
import { TaskFormDialog } from "@/components/tasks/task-form-dialog";
import { useTasks } from "@/hooks/use-tasks";
import { AuthGuard } from "@/components/auth-guard";
import type { Task, TaskStatus, TaskPriority } from "@/types";

export default function TasksPage() {
  return (
    <AuthGuard>
      <TasksContent />
    </AuthGuard>
  );
}

function TasksContent() {
  const [dialogOpen, setDialogOpen] = useState(false);
  const [editTask, setEditTask] = useState<Task | null>(null);
  const [status, setStatus] = useState("all");
  const [priority, setPriority] = useState("all");
  const [search, setSearch] = useState("");

  const params = useMemo(
    () => ({
      status: status !== "all" ? (status as TaskStatus) : undefined,
      priority: priority !== "all" ? (priority as TaskPriority) : undefined,
      search: search || undefined,
      limit: 50,
    }),
    [status, priority, search]
  );

  const { data, isLoading } = useTasks(params);

  const tasks = data?.data || [];

  const handleEdit = (task: Task) => {
    setEditTask(task);
    setDialogOpen(true);
  };

  const handleClose = () => {
    setDialogOpen(false);
    setEditTask(null);
  };

  return (
    <div className="p-4 lg:p-6 space-y-4">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold">Công việc</h1>
          <p className="text-sm text-muted-foreground">
            Quản lý danh sách công việc của bạn
          </p>
        </div>
        <Button
          onClick={() => setDialogOpen(true)}
          className="bg-gradient-to-r from-amber-500 to-red-500 text-white"
        >
          <Plus className="mr-2 h-4 w-4" />
          Thêm mới
        </Button>
      </div>

      <TaskFilters
        status={status}
        priority={priority}
        search={search}
        onStatusChange={setStatus}
        onPriorityChange={setPriority}
        onSearchChange={setSearch}
      />

      <TaskList tasks={tasks} isLoading={isLoading} onEdit={handleEdit} />

      <TaskFormDialog
        open={dialogOpen}
        onClose={handleClose}
        editTask={editTask}
      />
    </div>
  );
}
