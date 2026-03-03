"use client";

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { taskApi } from "@/lib/api";
import { useAuthStore } from "@/stores/auth-store";
import type {
  APIError,
  CreateTaskRequest,
  UpdateTaskRequest,
  TaskListParams,
  Task,
} from "@/types";
import { toast } from "sonner";

export function useTasks(params?: TaskListParams) {
  return useQuery({
    queryKey: ["tasks", params],
    queryFn: () => taskApi.list(params).then((r) => r.data),
    staleTime: 30 * 1000,
  });
}

export function useTask(id: string) {
  return useQuery({
    queryKey: ["tasks", id],
    queryFn: () => taskApi.get(id).then((r) => r.data.data),
    enabled: !!id,
  });
}

export function useTasksByDate(date: string) {
  return useQuery({
    queryKey: ["tasks", "by-date", date],
    queryFn: () => taskApi.getByDate(date).then((r) => r.data.data),
    enabled: !!date,
  });
}

/**
 * Fetch tasks within a date range and index them by date string (YYYY-MM-DD).
 * Only runs when authenticated.
 */
export function useTasksByDateRange(dateFrom: string, dateTo: string) {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  return useQuery({
    queryKey: ["tasks", "range", dateFrom, dateTo],
    queryFn: async () => {
      const resp = await taskApi.list({
        date_from: dateFrom,
        date_to: dateTo,
        limit: 200,
      });
      const tasks: Task[] = resp.data.data ?? [];
      const map: Record<string, Task[]> = {};
      for (const t of tasks) {
        // use due_date or start_date to group
        const key = t.due_date?.slice(0, 10) || t.start_date?.slice(0, 10);
        if (key) {
          (map[key] ??= []).push(t);
        }
      }
      return map;
    },
    enabled: isAuthenticated && !!dateFrom && !!dateTo,
    staleTime: 30 * 1000,
  });
}

export function useCreateTask() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateTaskRequest) => taskApi.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["tasks"] });
      toast.success("Tạo task thành công!");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg = error.response?.data?.error?.message || "Tạo task thất bại";
      toast.error(msg);
    },
  });
}

export function useUpdateTask() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateTaskRequest }) =>
      taskApi.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["tasks"] });
      toast.success("Cập nhật task thành công!");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg =
        error.response?.data?.error?.message || "Cập nhật task thất bại";
      toast.error(msg);
    },
  });
}

export function useDeleteTask() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => taskApi.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["tasks"] });
      toast.success("Xoá task thành công!");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg = error.response?.data?.error?.message || "Xoá task thất bại";
      toast.error(msg);
    },
  });
}

export function useUpdateTaskStatus() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, status }: { id: string; status: string }) =>
      taskApi.updateStatus(id, status),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["tasks"] });
    },
    onError: (error: AxiosError<APIError>) => {
      const msg =
        error.response?.data?.error?.message ||
        "Cập nhật trạng thái thất bại";
      toast.error(msg);
    },
  });
}
