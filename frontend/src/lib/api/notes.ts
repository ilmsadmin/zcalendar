import api from "@/lib/api-client";
import type {
  APIResponse,
  PaginatedResponse,
  Note,
  CreateNoteRequest,
  UpdateNoteRequest,
  NoteListParams,
} from "@/types";

export const noteApi = {
  create: (data: CreateNoteRequest) =>
    api.post<APIResponse<Note>>("/notes", data),

  get: (id: string) => api.get<APIResponse<Note>>(`/notes/${id}`),

  list: (params?: NoteListParams) =>
    api.get<PaginatedResponse<Note>>("/notes", { params }),

  update: (id: string, data: UpdateNoteRequest) =>
    api.put<APIResponse<Note>>(`/notes/${id}`, data),

  delete: (id: string) => api.delete<APIResponse>(`/notes/${id}`),

  getByDate: (date: string) =>
    api.get<APIResponse<Note[]>>(`/notes/by-date/${date}`),

  search: (query: string, limit?: number, offset?: number) =>
    api.get<PaginatedResponse<Note>>("/notes/search", {
      params: { q: query, limit, offset },
    }),
};
