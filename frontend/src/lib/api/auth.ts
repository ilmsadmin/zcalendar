import api from "@/lib/api-client";
import type {
  APIResponse,
  AuthResponse,
  LoginRequest,
  RegisterRequest,
  TokenResponse,
  ChangePasswordRequest,
  UpdateProfileRequest,
  UpdatePreferencesRequest,
  User,
  DeleteAccountRequest,
} from "@/types";

export const authApi = {
  register: (data: RegisterRequest) =>
    api.post<APIResponse<AuthResponse>>("/auth/register", data),

  login: (data: LoginRequest) =>
    api.post<APIResponse<AuthResponse>>("/auth/login", data),

  logout: () => api.post<APIResponse>("/auth/logout"),

  refreshToken: (refreshToken: string) =>
    api.post<APIResponse<TokenResponse>>("/auth/refresh", {
      refresh_token: refreshToken,
    }),

  changePassword: (data: ChangePasswordRequest) =>
    api.post<APIResponse>("/auth/change-password", data),

  getProfile: () =>
    api.get<APIResponse<User>>("/auth/profile"),

  updateProfile: (data: UpdateProfileRequest) =>
    api.put<APIResponse<User>>("/auth/profile", data),

  updatePreferences: (data: UpdatePreferencesRequest) =>
    api.put<APIResponse<User>>("/auth/preferences", data),

  deleteAccount: (data: DeleteAccountRequest) =>
    api.delete<APIResponse>("/auth/account", { data }),
};
