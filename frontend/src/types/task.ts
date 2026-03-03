export interface Task {
  id: string;
  user_id: string;
  category_id?: string;
  parent_task_id?: string;
  title: string;
  description: string;
  due_date?: string;
  due_time?: string;
  start_date?: string;
  start_time?: string;
  lunar_day?: number;
  lunar_month?: number;
  lunar_year?: number;
  is_lunar_based: boolean;
  status: TaskStatus;
  priority: TaskPriority;
  is_recurring: boolean;
  recurrence_rule?: string;
  recurrence_end_date?: string;
  progress: number;
  completed_at?: string;
  sort_order: number;
  estimated_hours?: number;
  actual_hours?: number;
  category?: Category;
  tags: Tag[];
  subtasks: Task[];
  created_at: string;
  updated_at: string;
}

export type TaskStatus = "todo" | "in_progress" | "done" | "cancelled";
export type TaskPriority = "low" | "medium" | "high" | "urgent";

export interface CreateTaskRequest {
  title: string;
  description?: string;
  category_id?: string;
  parent_task_id?: string;
  due_date?: string;
  due_time?: string;
  start_date?: string;
  start_time?: string;
  is_lunar_based?: boolean;
  lunar_day?: number;
  lunar_month?: number;
  lunar_year?: number;
  priority?: TaskPriority;
  is_recurring?: boolean;
  recurrence_rule?: string;
  recurrence_end_date?: string;
  estimated_hours?: number;
  tag_ids?: string[];
}

export interface UpdateTaskRequest {
  title?: string;
  description?: string;
  category_id?: string;
  parent_task_id?: string;
  due_date?: string;
  due_time?: string;
  start_date?: string;
  start_time?: string;
  is_lunar_based?: boolean;
  lunar_day?: number;
  lunar_month?: number;
  lunar_year?: number;
  priority?: TaskPriority;
  is_recurring?: boolean;
  recurrence_rule?: string;
  recurrence_end_date?: string;
  progress?: number;
  estimated_hours?: number;
  actual_hours?: number;
  sort_order?: number;
  tag_ids?: string[];
}

export interface Category {
  id: string;
  user_id: string;
  name: string;
  color: string;
  icon: string;
  type: string;
  sort_order: number;
  created_at: string;
}

export interface Tag {
  id: string;
  user_id: string;
  name: string;
  color: string;
  created_at: string;
}

export interface CreateCategoryRequest {
  name: string;
  color?: string;
  icon?: string;
  type: string;
  sort_order?: number;
}

export interface CreateTagRequest {
  name: string;
  color?: string;
}

export interface TaskListParams {
  page?: number;
  limit?: number;
  status?: TaskStatus;
  priority?: TaskPriority;
  category_id?: string;
  tag_id?: string;
  date_from?: string;
  date_to?: string;
  search?: string;
}
