"use client";

import { useState } from "react";
import Link from "next/link";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Moon, Eye, EyeOff } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useLogin } from "@/hooks/use-auth";

const loginSchema = z.object({
  email: z.string().email("Email không hợp lệ"),
  password: z.string().min(6, "Mật khẩu tối thiểu 6 ký tự"),
});

type LoginForm = z.infer<typeof loginSchema>;

export default function LoginPage() {
  const [showPassword, setShowPassword] = useState(false);
  const loginMutation = useLogin();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginForm>({
    resolver: zodResolver(loginSchema),
  });

  const onSubmit = (data: LoginForm) => {
    loginMutation.mutate(data);
  };

  return (
    <div className="rounded-2xl border-0 bg-white/80 dark:bg-gray-900/80 backdrop-blur-lg shadow-xl p-6">
      {/* Brand */}
      <div className="text-center mb-6">
        <div className="flex justify-center mb-3">
          <div className="flex items-center gap-2">
            <Moon className="h-7 w-7 text-amber-500" />
            <span className="text-xl font-bold bg-gradient-to-r from-amber-500 to-red-500 bg-clip-text text-transparent">
              Zplus Calendar
            </span>
          </div>
        </div>
        <h1 className="text-lg font-semibold">Đăng nhập</h1>
        <p className="text-xs text-muted-foreground mt-0.5">
          Chào mừng bạn trở lại!
        </p>
      </div>

      <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
        <div className="space-y-1.5">
          <Label htmlFor="email" className="text-xs">
            Email
          </Label>
          <Input
            id="email"
            type="email"
            placeholder="email@example.com"
            inputMode="email"
            autoComplete="email"
            {...register("email")}
            className={`h-11 rounded-xl text-sm ${
              errors.email ? "border-red-500" : ""
            }`}
          />
          {errors.email && (
            <p className="text-[11px] text-red-500">{errors.email.message}</p>
          )}
        </div>

        <div className="space-y-1.5">
          <Label htmlFor="password" className="text-xs">
            Mật khẩu
          </Label>
          <div className="relative">
            <Input
              id="password"
              type={showPassword ? "text" : "password"}
              placeholder="••••••"
              autoComplete="current-password"
              {...register("password")}
              className={`h-11 rounded-xl text-sm pr-11 ${
                errors.password ? "border-red-500" : ""
              }`}
            />
            <Button
              type="button"
              variant="ghost"
              size="icon"
              className="absolute right-0 top-0 h-11 w-11"
              onClick={() => setShowPassword(!showPassword)}
            >
              {showPassword ? (
                <EyeOff className="h-4 w-4" />
              ) : (
                <Eye className="h-4 w-4" />
              )}
            </Button>
          </div>
          {errors.password && (
            <p className="text-[11px] text-red-500">{errors.password.message}</p>
          )}
        </div>

        <Button
          type="submit"
          className="w-full h-11 rounded-xl bg-gradient-to-r from-amber-500 to-red-500 hover:from-amber-600 hover:to-red-600 text-white text-sm font-medium active:scale-[0.98] transition-transform"
          disabled={loginMutation.isPending}
        >
          {loginMutation.isPending ? "Đang đăng nhập..." : "Đăng nhập"}
        </Button>
      </form>

      <div className="mt-5 text-center text-xs">
        <span className="text-muted-foreground">Chưa có tài khoản? </span>
        <Link
          href="/register"
          className="font-medium text-primary hover:underline"
        >
          Đăng ký
        </Link>
      </div>
    </div>
  );
}
