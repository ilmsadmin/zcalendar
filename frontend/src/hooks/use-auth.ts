"use client";

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { AxiosError } from "axios";
import { authApi } from "@/lib/api";
import { useAuthStore } from "@/stores/auth-store";
import type {
  APIError,
  LoginRequest,
  RegisterRequest,
  UpdateProfileRequest,
  UpdatePreferencesRequest,
} from "@/types";
import { toast } from "sonner";

export function useLogin() {
  const router = useRouter();
  const setAuth = useAuthStore((s) => s.setAuth);

  return useMutation({
    mutationFn: (data: LoginRequest) => authApi.login(data),
    onSuccess: (response) => {
      const { user, access_token, refresh_token } = response.data.data;
      setAuth(user, access_token, refresh_token);
      toast.success("Đăng nhập thành công!");
      router.push("/calendar");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg =
        error.response?.data?.error?.message || "Đăng nhập thất bại";
      toast.error(msg);
    },
  });
}

export function useRegister() {
  const router = useRouter();
  const setAuth = useAuthStore((s) => s.setAuth);

  return useMutation({
    mutationFn: (data: RegisterRequest) => authApi.register(data),
    onSuccess: (response) => {
      const { user, access_token, refresh_token } = response.data.data;
      setAuth(user, access_token, refresh_token);
      toast.success("Đăng ký thành công!");
      router.push("/calendar");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg = error.response?.data?.error?.message || "Đăng ký thất bại";
      toast.error(msg);
    },
  });
}

export function useLogout() {
  const router = useRouter();
  const logout = useAuthStore((s) => s.logout);

  return useMutation({
    mutationFn: () => authApi.logout(),
    onSuccess: () => {
      logout();
      router.push("/login");
    },
    onError: () => {
      // Force logout even on error
      logout();
      router.push("/login");
    },
  });
}

export function useChangePassword() {
  return useMutation({
    mutationFn: (data: { old_password: string; new_password: string }) =>
      authApi.changePassword(data),
    onSuccess: () => {
      toast.success("Đổi mật khẩu thành công!");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg =
        error.response?.data?.error?.message || "Đổi mật khẩu thất bại";
      toast.error(msg);
    },
  });
}

export function useProfile() {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  return useQuery({
    queryKey: ["auth", "profile"],
    queryFn: () => authApi.getProfile().then((r) => r.data.data),
    enabled: isAuthenticated,
    staleTime: 5 * 60 * 1000,
  });
}

export function useUpdateProfile() {
  const queryClient = useQueryClient();
  const setUser = useAuthStore((s) => s.setUser);

  return useMutation({
    mutationFn: (data: UpdateProfileRequest) => authApi.updateProfile(data),
    onSuccess: (response) => {
      setUser(response.data.data);
      queryClient.invalidateQueries({ queryKey: ["auth", "profile"] });
      toast.success("Cập nhật hồ sơ thành công!");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg =
        error.response?.data?.error?.message || "Cập nhật hồ sơ thất bại";
      toast.error(msg);
    },
  });
}

export function useUpdatePreferences() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdatePreferencesRequest) =>
      authApi.updatePreferences(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["auth", "profile"] });
      toast.success("Cập nhật tuỳ chọn thành công!");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg =
        error.response?.data?.error?.message || "Cập nhật tuỳ chọn thất bại";
      toast.error(msg);
    },
  });
}

export function useDeleteAccount() {
  const router = useRouter();
  const logout = useAuthStore((s) => s.logout);

  return useMutation({
    mutationFn: (password: string) =>
      authApi.deleteAccount({ password }),
    onSuccess: () => {
      logout();
      toast.success("Tài khoản đã được xoá.");
      router.push("/login");
    },
    onError: (error: AxiosError<APIError>) => {
      const msg =
        error.response?.data?.error?.message || "Xoá tài khoản thất bại";
      toast.error(msg);
    },
  });
}
