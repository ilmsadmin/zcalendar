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
import { useRegister } from "@/hooks/use-auth";

const registerSchema = z
  .object({
    full_name: z.string().min(2, "Họ tên tối thiểu 2 ký tự"),
    username: z
      .string()
      .min(3, "Username tối thiểu 3 ký tự")
      .regex(/^[a-zA-Z0-9_]+$/, "Username chỉ chứa chữ, số và _"),
    email: z.string().email("Email không hợp lệ"),
    password: z.string().min(6, "Mật khẩu tối thiểu 6 ký tự"),
    confirmPassword: z.string(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Mật khẩu không khớp",
    path: ["confirmPassword"],
  });

type RegisterForm = z.infer<typeof registerSchema>;

export default function RegisterPage() {
  const [showPassword, setShowPassword] = useState(false);
  const registerMutation = useRegister();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterForm>({
    resolver: zodResolver(registerSchema),
  });

  const onSubmit = (data: RegisterForm) => {
    registerMutation.mutate({
      email: data.email,
      username: data.username,
      password: data.password,
      full_name: data.full_name,
    });
  };

  return (
    <div className="rounded-2xl border-0 bg-white/80 dark:bg-gray-900/80 backdrop-blur-lg shadow-xl p-6">
      {/* Brand */}
      <div className="text-center mb-5">
        <div className="flex justify-center mb-3">
          <div className="flex items-center gap-2">
            <Moon className="h-7 w-7 text-amber-500" />
            <span className="text-xl font-bold bg-gradient-to-r from-amber-500 to-red-500 bg-clip-text text-transparent">
              Zplus Calendar
            </span>
          </div>
        </div>
        <h1 className="text-lg font-semibold">Tạo tài khoản</h1>
        <p className="text-xs text-muted-foreground mt-0.5">
          Bắt đầu quản lý thời gian hiệu quả
        </p>
      </div>

      <form onSubmit={handleSubmit(onSubmit)} className="space-y-3.5">
        <div className="space-y-1.5">
          <Label htmlFor="full_name" className="text-xs">
            Họ và tên
          </Label>
          <Input
            id="full_name"
            placeholder="Nguyễn Văn A"
            autoComplete="name"
            {...register("full_name")}
            className={`h-11 rounded-xl text-sm ${
              errors.full_name ? "border-red-500" : ""
            }`}
          />
          {errors.full_name && (
            <p className="text-[11px] text-red-500">
              {errors.full_name.message}
            </p>
          )}
        </div>

        <div className="space-y-1.5">
          <Label htmlFor="username" className="text-xs">
            Username
          </Label>
          <Input
            id="username"
            placeholder="nguyenvana"
            autoComplete="username"
            autoCapitalize="none"
            {...register("username")}
            className={`h-11 rounded-xl text-sm ${
              errors.username ? "border-red-500" : ""
            }`}
          />
          {errors.username && (
            <p className="text-[11px] text-red-500">
              {errors.username.message}
            </p>
          )}
        </div>

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
              autoComplete="new-password"
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

        <div className="space-y-1.5">
          <Label htmlFor="confirmPassword" className="text-xs">
            Nhập lại mật khẩu
          </Label>
          <Input
            id="confirmPassword"
            type="password"
            placeholder="••••••"
            autoComplete="new-password"
            {...register("confirmPassword")}
            className={`h-11 rounded-xl text-sm ${
              errors.confirmPassword ? "border-red-500" : ""
            }`}
          />
          {errors.confirmPassword && (
            <p className="text-[11px] text-red-500">
              {errors.confirmPassword.message}
            </p>
          )}
        </div>

        <Button
          type="submit"
          className="w-full h-11 rounded-xl bg-gradient-to-r from-amber-500 to-red-500 hover:from-amber-600 hover:to-red-600 text-white text-sm font-medium active:scale-[0.98] transition-transform"
          disabled={registerMutation.isPending}
        >
          {registerMutation.isPending ? "Đang đăng ký..." : "Đăng ký"}
        </Button>
      </form>

      <div className="mt-5 text-center text-xs">
        <span className="text-muted-foreground">Đã có tài khoản? </span>
        <Link
          href="/login"
          className="font-medium text-primary hover:underline"
        >
          Đăng nhập
        </Link>
      </div>
    </div>
  );
}
