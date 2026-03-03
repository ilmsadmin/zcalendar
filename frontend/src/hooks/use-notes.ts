"use client";

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { noteApi } from "@/lib/api";
import { useAuthStore } from "@/stores/auth-store";
import type {
  APIError,
  CreateNoteRequest,
  UpdateNoteRequest,
  NoteListParams,
  Note,
} from "@/types";
import { toast } from "sonner";

export function useNotes(params?: NoteListParams) {
  return useQuery({
    queryKey: ["notes", params],
    queryFn: () => noteApi.list(params).then((r) => r.data),
    staleTime: 30 * 1000,
  });
}

export function useNote(id: string) {
  return useQuery({
    queryKey: ["notes", id],
    queryFn: () => noteApi.get(id).then((r) => r.data.data),
    enabled: !!id,
  });
}

export function useNotesByDate(date: string) {
  return useQuery({
    queryKey: ["notes", "by-date", date],
    queryFn: () => noteApi.getByDate(date).then((r) => r.data.data),
    enabled: !!date,
  });
}

/**
 * Fetch notes within a date range and index them by date string (YYYY-MM-DD).
 * Only runs when authenticated.
 */
export function useNotesByDateRange(dateFrom: string, dateTo: string) {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  return useQuery({
    queryKey: ["notes", "range", dateFrom, dateTo],
    queryFn: async () => {
      const resp = await noteApi.list({
        date_from: dateFrom,
        date_to: dateTo,
        limit: 200,
      });
      const notes: Note[] = resp.data.data ?? [];
      const map: Record<string, Note[]> = {};
      for (const n of notes) {
        const key = n.date?.slice(0, 10);
        if (key) {
          (map[key] ??= []).push(n);
        }
      }
      return map;
    },
    enabled: isAuthenticated && !!dateFrom && !!dateTo,
    staleTime: 30 * 1000,
  });
}

export function useCreateNote() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateNoteRequest) => noteApi.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["notes"] });
      toast.success("Tạo ghi chú thành công!");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg =
        error.response?.data?.error?.message || "Tạo ghi chú thất bại";
      toast.error(msg);
    },
  });
}

export function useUpdateNote() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateNoteRequest }) =>
      noteApi.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["notes"] });
      toast.success("Cập nhật ghi chú thành công!");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg =
        error.response?.data?.error?.message || "Cập nhật ghi chú thất bại";
      toast.error(msg);
    },
  });
}

export function useDeleteNote() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => noteApi.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["notes"] });
      toast.success("Xoá ghi chú thành công!");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg =
        error.response?.data?.error?.message || "Xoá ghi chú thất bại";
      toast.error(msg);
    },
  });
}
