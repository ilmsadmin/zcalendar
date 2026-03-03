import api from "@/lib/api-client";
import type {
  APIResponse,
  PaginatedResponse,
  Task,
  CreateTaskRequest,
  UpdateTaskRequest,
  TaskListParams,
} from "@/types";

export const taskApi = {
  create: (data: CreateTaskRequest) =>
    api.post<APIResponse<Task>>("/tasks", data),

  get: (id: string) => api.get<APIResponse<Task>>(`/tasks/${id}`),

  list: (params?: TaskListParams) =>
    api.get<PaginatedResponse<Task>>("/tasks", { params }),

  update: (id: string, data: UpdateTaskRequest) =>
    api.put<APIResponse<Task>>(`/tasks/${id}`, data),

  delete: (id: string) => api.delete<APIResponse>(`/tasks/${id}`),

  updateStatus: (id: string, status: string) =>
    api.patch<APIResponse>(`/tasks/${id}/status`, { status }),

  getByDate: (date: string) =>
    api.get<APIResponse<Task[]>>(`/tasks/by-date/${date}`),

  search: (query: string, limit?: number, offset?: number) =>
    api.get<PaginatedResponse<Task>>("/tasks/search", {
      params: { q: query, limit, offset },
    }),
};
