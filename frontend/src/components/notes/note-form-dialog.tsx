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
import { cn } from "@/lib/utils";
import { useCreateNote, useUpdateNote } from "@/hooks/use-notes";
import type { Note } from "@/types";

const noteSchema = z.object({
  title: z.string().min(1, "Tiêu đề không được để trống"),
  content: z.string().optional(),
  date: z.string().min(1, "Chọn ngày"),
  color: z.string().optional(),
});

type NoteFormValues = z.infer<typeof noteSchema>;

const colors = [
  { value: "", label: "Mặc định" },
  { value: "red", label: "Đỏ" },
  { value: "orange", label: "Cam" },
  { value: "yellow", label: "Vàng" },
  { value: "green", label: "Xanh lá" },
  { value: "blue", label: "Xanh dương" },
  { value: "purple", label: "Tím" },
  { value: "pink", label: "Hồng" },
];

const colorClasses: Record<string, string> = {
  "": "bg-gray-200 dark:bg-gray-700",
  red: "bg-red-400",
  orange: "bg-orange-400",
  yellow: "bg-yellow-400",
  green: "bg-green-400",
  blue: "bg-blue-400",
  purple: "bg-purple-400",
  pink: "bg-pink-400",
};

interface NoteFormDialogProps {
  open: boolean;
  onClose: () => void;
  editNote?: Note | null;
}

export function NoteFormDialog({ open, onClose, editNote }: NoteFormDialogProps) {
  const createNote = useCreateNote();
  const updateNote = useUpdateNote();
  const isEditing = !!editNote;

  const {
    register,
    handleSubmit,
    setValue,
    watch,
    reset,
    formState: { errors },
  } = useForm<NoteFormValues>({
    resolver: zodResolver(noteSchema),
    defaultValues: {
      title: editNote?.title || "",
      content: editNote?.content || "",
      date: editNote?.date || new Date().toISOString().split("T")[0],
      color: editNote?.color || "",
    },
  });

  // Reset form values when editNote changes or dialog opens
  useEffect(() => {
    if (open) {
      reset({
        title: editNote?.title || "",
        content: editNote?.content || "",
        date: editNote?.date || new Date().toISOString().split("T")[0],
        color: editNote?.color || "",
      });
    }
  }, [open, editNote, reset]);

  const selectedColor = watch("color");

  const onSubmit = (data: NoteFormValues) => {
    if (isEditing && editNote) {
      updateNote.mutate(
        { id: editNote.id, data },
        {
          onSuccess: () => {
            reset();
            onClose();
          },
        }
      );
    } else {
      createNote.mutate(data, {
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
            {isEditing ? "Chỉnh sửa ghi chú" : "Ghi chú mới"}
          </DialogTitle>
        </DialogHeader>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="note-title">Tiêu đề</Label>
            <Input
              id="note-title"
              placeholder="Nhập tiêu đề..."
              {...register("title")}
              className={errors.title ? "border-red-500" : ""}
            />
            {errors.title && (
              <p className="text-xs text-red-500">{errors.title.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="note-content">Nội dung</Label>
            <textarea
              id="note-content"
              placeholder="Nhập nội dung..."
              {...register("content")}
              rows={6}
              className="flex w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring resize-none"
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="note-date">Ngày</Label>
            <Input id="note-date" type="date" {...register("date")} />
            {errors.date && (
              <p className="text-xs text-red-500">{errors.date.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label>Màu sắc</Label>
            <div className="flex gap-2">
              {colors.map((c) => (
                <button
                  key={c.value}
                  type="button"
                  onClick={() => setValue("color", c.value)}
                  className={cn(
                    "h-7 w-7 rounded-full transition-all",
                    colorClasses[c.value],
                    selectedColor === c.value
                      ? "ring-2 ring-offset-2 ring-primary"
                      : "hover:scale-110"
                  )}
                  title={c.label}
                />
              ))}
            </div>
          </div>

          <div className="flex justify-end gap-2 pt-2">
            <Button type="button" variant="outline" onClick={onClose}>
              Huỷ
            </Button>
            <Button
              type="submit"
              disabled={createNote.isPending || updateNote.isPending}
              className="bg-gradient-to-r from-amber-500 to-red-500 text-white"
            >
              {createNote.isPending || updateNote.isPending
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
