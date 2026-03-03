export interface APIResponse<T = unknown> {
  status: string;
  data: T;
  message?: string;
}

export interface APIError {
  status: string;
  error: {
    code: string;
    message: string;
    details?: { field: string; message: string }[];
  };
}

export interface PaginatedResponse<T = unknown> {
  status: string;
  data: T[];
  meta: {
    page: number;
    limit: number;
    total_items: number;
    total_pages: number;
    has_next: boolean;
    has_prev: boolean;
  };
}
