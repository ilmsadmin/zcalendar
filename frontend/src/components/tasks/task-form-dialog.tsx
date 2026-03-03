"use client";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { useEffect } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useCreateTask, useUpdateTask } from "@/hooks/use-tasks";
import type { Task } from "@/types";

const taskSchema = z.object({
  title: z.string().min(1, "Tiêu đề không được để trống"),
  description: z.string().optional(),
  due_date: z.string().optional(),
  priority: z.enum(["low", "medium", "high", "urgent"]),
});

type TaskFormValues = z.infer<typeof taskSchema>;

interface TaskFormDialogProps {
  open: boolean;
  onClose: () => void;
  editTask?: Task | null;
}

export function TaskFormDialog({ open, onClose, editTask }: TaskFormDialogProps) {
  const createTask = useCreateTask();
  const updateTask = useUpdateTask();
  const isEditing = !!editTask;

  const {
    register,
    handleSubmit,
    setValue,
    watch,
    reset,
    formState: { errors },
  } = useForm<TaskFormValues>({
    resolver: zodResolver(taskSchema),
    defaultValues: {
      title: editTask?.title || "",
      description: editTask?.description || "",
      due_date: editTask?.due_date || "",
      priority: editTask?.priority || "medium",
    },
  });

  // Reset form values when editTask changes or dialog opens
  useEffect(() => {
    if (open) {
      reset({
        title: editTask?.title || "",
        description: editTask?.description || "",
        due_date: editTask?.due_date || "",
        priority: editTask?.priority || "medium",
      });
    }
  }, [open, editTask, reset]);

  const onSubmit = (data: TaskFormValues) => {
    if (isEditing && editTask) {
      updateTask.mutate(
        { id: editTask.id, data },
        {
          onSuccess: () => {
            reset();
            onClose();
          },
        }
      );
    } else {
      createTask.mutate(data, {
        onSuccess: () => {
          reset();
          onClose();
        },
      });
    }
  };

  return (
    <Dialog open={open} onOpenChange={(v) => !v && onClose()}>
      <DialogContent className="sm:max-w-lg">
        <DialogHeader>
          <DialogTitle>
            {isEditing ? "Chỉnh sửa công việc" : "Thêm công việc mới"}
          </DialogTitle>
        </DialogHeader>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="title">Tiêu đề</Label>
            <Input
              id="title"
              placeholder="Nhập tiêu đề..."
              {...register("title")}
              className={errors.title ? "border-red-500" : ""}
            />
            {errors.title && (
              <p className="text-xs text-red-500">{errors.title.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="description">Mô tả</Label>
            <textarea
              id="description"
              placeholder="Nhập mô tả..."
              {...register("description")}
              rows={3}
              className="flex w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring"
            />
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="due_date">Ngày hạn</Label>
              <Input id="due_date" type="date" {...register("due_date")} />
            </div>

            <div className="space-y-2">
              <Label>Ưu tiên</Label>
              <Select
                value={watch("priority")}
                onValueChange={(v) =>
                  setValue("priority", v as TaskFormValues["priority"])
                }
              >
                <SelectTrigger>
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="low">🟢 Thấp</SelectItem>
                  <SelectItem value="medium">🔵 Trung bình</SelectItem>
                  <SelectItem value="high">🟠 Cao</SelectItem>
                  <SelectItem value="urgent">🔴 Khẩn cấp</SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>

          <div className="flex justify-end gap-2 pt-2">
            <Button type="button" variant="outline" onClick={onClose}>
              Huỷ
            </Button>
            <Button
              type="submit"
              disabled={createTask.isPending || updateTask.isPending}
              className="bg-gradient-to-r from-amber-500 to-red-500 text-white"
            >
              {createTask.isPending || updateTask.isPending
                ? "Đang lưu..."
                : isEditing
                ? "Cập nhật"
                : "Tạo mới"}
            </Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
}
