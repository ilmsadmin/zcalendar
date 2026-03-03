import api from "@/lib/api-client";
import type {
  APIResponse,
  Category,
  Tag,
  CreateCategoryRequest,
  CreateTagRequest,
} from "@/types";

export const categoryApi = {
  create: (data: CreateCategoryRequest) =>
    api.post<APIResponse<Category>>("/categories", data),

  list: (type?: string) =>
    api.get<APIResponse<Category[]>>("/categories", { params: { type } }),

  update: (id: string, data: Partial<CreateCategoryRequest>) =>
    api.put<APIResponse<Category>>(`/categories/${id}`, data),

  delete: (id: string) => api.delete<APIResponse>(`/categories/${id}`),
};

export const tagApi = {
  create: (data: CreateTagRequest) =>
    api.post<APIResponse<Tag>>("/tags", data),

  list: () => api.get<APIResponse<Tag[]>>("/tags"),

  update: (id: string, data: Partial<CreateTagRequest>) =>
    api.put<APIResponse<Tag>>(`/tags/${id}`, data),

  delete: (id: string) => api.delete<APIResponse>(`/tags/${id}`),
};
