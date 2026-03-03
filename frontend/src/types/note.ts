export interface Note {
  id: string;
  user_id: string;
  title: string;
  content: string;
  date: string;
  category_id?: string;
  tags: string[];
  is_pinned: boolean;
  color: string;
  lunar_day?: number;
  lunar_month?: number;
  lunar_year?: number;
  created_at: string;
  updated_at: string;
}

export interface CreateNoteRequest {
  title: string;
  content?: string;
  date: string;
  category_id?: string;
  tags?: string[];
  is_pinned?: boolean;
  color?: string;
  lunar_day?: number;
  lunar_month?: number;
  lunar_year?: number;
}

export interface UpdateNoteRequest {
  title?: string;
  content?: string;
  date?: string;
  category_id?: string;
  tags?: string[];
  is_pinned?: boolean;
  color?: string;
  lunar_day?: number;
  lunar_month?: number;
  lunar_year?: number;
}

export interface NoteListParams {
  page?: number;
  limit?: number;
  date?: string;
  date_from?: string;
  date_to?: string;
  search?: string;
  pinned?: boolean;
}
