export interface User {
  id: string;
  email: string;
  username: string;
  full_name: string;
  avatar_url?: string;
  role: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  username: string;
  password: string;
  full_name: string;
}

export interface AuthResponse {
  user: User;
  access_token: string;
  refresh_token: string;
  expires_at: string;
}

export interface TokenResponse {
  access_token: string;
  refresh_token: string;
  expires_at: string;
}

export interface ChangePasswordRequest {
  old_password: string;
  new_password: string;
}

export interface UpdateProfileRequest {
  full_name?: string;
  username?: string;
}

export interface UpdatePreferencesRequest {
  theme?: string;
  locale?: string;
  calendar_view?: string;
  week_start_day?: number;
}

export interface UserPreferences {
  theme: string;
  locale: string;
  calendar_view: string;
  week_start_day: number;
}

export interface DeleteAccountRequest {
  password: string;
}
