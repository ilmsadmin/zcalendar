"use client";

import { useState } from "react";
import { useTheme } from "next-themes";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Separator } from "@/components/ui/separator";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Switch } from "@/components/ui/switch";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { useAuthStore } from "@/stores/auth-store";
import { useCalendarStore, CalendarView } from "@/stores/calendar-store";
import { AuthGuard } from "@/components/auth-guard";
import {
  useLogout,
  useChangePassword,
  useUpdateProfile,
  useDeleteAccount,
} from "@/hooks/use-auth";
import {
  Moon,
  Sun,
  Monitor,
  User,
  Lock,
  Palette,
  CalendarDays,
  LogOut,
  Trash2,
  Loader2,
  Eye,
  EyeOff,
  Save,
  Shield,
} from "lucide-react";

export default function SettingsPage() {
  return (
    <AuthGuard>
      <SettingsContent />
    </AuthGuard>
  );
}

function SettingsContent() {
  const { theme, setTheme } = useTheme();
  const user = useAuthStore((s) => s.user);
  const { view, setView } = useCalendarStore();
  const logoutMutation = useLogout();

  return (
    <div className="p-4 lg:p-6 space-y-6 max-w-2xl pb-32">
      <div>
        <h1 className="text-2xl font-bold">Cài đặt</h1>
        <p className="text-sm text-muted-foreground">
          Tùy chỉnh ứng dụng theo ý bạn
        </p>
      </div>

      {/* Profile */}
      <ProfileSection />

      {/* Appearance */}
      <Card>
        <CardHeader className="pb-3">
          <div className="flex items-center gap-2">
            <Palette className="h-5 w-5 text-muted-foreground" />
            <CardTitle className="text-base">Giao diện</CardTitle>
          </div>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex items-center justify-between">
            <div>
              <Label>Chế độ hiển thị</Label>
              <p className="text-xs text-muted-foreground mt-0.5">Chọn giao diện sáng, tối hoặc theo hệ thống</p>
            </div>
            <Select value={theme} onValueChange={setTheme}>
              <SelectTrigger className="w-40">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="light">
                  <span className="flex items-center gap-2">
                    <Sun className="h-4 w-4" /> Sáng
                  </span>
                </SelectItem>
                <SelectItem value="dark">
                  <span className="flex items-center gap-2">
                    <Moon className="h-4 w-4" /> Tối
                  </span>
                </SelectItem>
                <SelectItem value="system">
                  <span className="flex items-center gap-2">
                    <Monitor className="h-4 w-4" /> Hệ thống
                  </span>
                </SelectItem>
              </SelectContent>
            </Select>
          </div>
        </CardContent>
      </Card>

      {/* Calendar preferences */}
      <Card>
        <CardHeader className="pb-3">
          <div className="flex items-center gap-2">
            <CalendarDays className="h-5 w-5 text-muted-foreground" />
            <CardTitle className="text-base">Lịch</CardTitle>
          </div>
          <CardDescription>Tuỳ chỉnh hiển thị lịch</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex items-center justify-between">
            <div>
              <Label>Chế độ xem mặc định</Label>
              <p className="text-xs text-muted-foreground mt-0.5">Chọn chế độ xem khi mở ứng dụng</p>
            </div>
            <Select value={view} onValueChange={(v) => setView(v as CalendarView)}>
              <SelectTrigger className="w-32">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="month">Tháng</SelectItem>
                <SelectItem value="week">Tuần</SelectItem>
                <SelectItem value="day">Ngày</SelectItem>
              </SelectContent>
            </Select>
          </div>
        </CardContent>
      </Card>

      {/* Security: Change password */}
      <ChangePasswordSection />

      {/* Account actions */}
      <Card>
        <CardHeader className="pb-3">
          <div className="flex items-center gap-2">
            <LogOut className="h-5 w-5 text-muted-foreground" />
            <CardTitle className="text-base">Phiên đăng nhập</CardTitle>
          </div>
        </CardHeader>
        <CardContent>
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm font-medium">Đăng xuất</p>
              <p className="text-xs text-muted-foreground">Đăng xuất khỏi thiết bị này</p>
            </div>
            <Button
              variant="outline"
              size="sm"
              onClick={() => logoutMutation.mutate()}
              disabled={logoutMutation.isPending}
            >
              {logoutMutation.isPending ? (
                <Loader2 className="h-4 w-4 mr-1.5 animate-spin" />
              ) : (
                <LogOut className="h-4 w-4 mr-1.5" />
              )}
              Đăng xuất
            </Button>
          </div>
        </CardContent>
      </Card>

      {/* Danger zone */}
      <DeleteAccountSection />
    </div>
  );
}

/* ─── Profile Section ────────────────────────────────────────────── */

function ProfileSection() {
  const user = useAuthStore((s) => s.user);
  const updateProfile = useUpdateProfile();

  const [isEditing, setIsEditing] = useState(false);
  const [fullName, setFullName] = useState(user?.full_name || "");
  const [username, setUsername] = useState(user?.username || "");

  const handleSave = () => {
    const changes: Record<string, string> = {};
    if (fullName !== user?.full_name) changes.full_name = fullName;
    if (username !== user?.username) changes.username = username;

    if (Object.keys(changes).length === 0) {
      setIsEditing(false);
      return;
    }

    updateProfile.mutate(changes, {
      onSuccess: () => setIsEditing(false),
    });
  };

  const handleCancel = () => {
    setFullName(user?.full_name || "");
    setUsername(user?.username || "");
    setIsEditing(false);
  };

  return (
    <Card>
      <CardHeader className="pb-3">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <User className="h-5 w-5 text-muted-foreground" />
            <CardTitle className="text-base">Hồ sơ</CardTitle>
          </div>
          {!isEditing && (
            <Button variant="ghost" size="sm" onClick={() => setIsEditing(true)} className="text-xs">
              Chỉnh sửa
            </Button>
          )}
        </div>
      </CardHeader>
      <CardContent className="space-y-3">
        <div className="flex justify-between items-center">
          <Label className="text-muted-foreground">Họ tên</Label>
          {isEditing ? (
            <Input
              value={fullName}
              onChange={(e) => setFullName(e.target.value)}
              className="w-48 h-8 text-sm"
              placeholder="Nhập họ tên"
            />
          ) : (
            <span className="text-sm font-medium">{user?.full_name || "—"}</span>
          )}
        </div>
        <Separator />
        <div className="flex justify-between items-center">
          <Label className="text-muted-foreground">Username</Label>
          {isEditing ? (
            <Input
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-48 h-8 text-sm"
              placeholder="Nhập username"
            />
          ) : (
            <span className="text-sm font-medium">{user?.username || "—"}</span>
          )}
        </div>
        <Separator />
        <div className="flex justify-between items-center">
          <Label className="text-muted-foreground">Email</Label>
          <span className="text-sm font-medium text-muted-foreground">{user?.email}</span>
        </div>
        <Separator />
        <div className="flex justify-between items-center">
          <Label className="text-muted-foreground">Ngày tham gia</Label>
          <span className="text-sm text-muted-foreground">
            {user?.created_at
              ? new Date(user.created_at).toLocaleDateString("vi-VN", {
                  day: "2-digit",
                  month: "2-digit",
                  year: "numeric",
                })
              : "—"}
          </span>
        </div>

        {isEditing && (
          <>
            <Separator />
            <div className="flex justify-end gap-2 pt-1">
              <Button variant="ghost" size="sm" onClick={handleCancel} className="text-xs">
                Huỷ
              </Button>
              <Button
                size="sm"
                onClick={handleSave}
                disabled={updateProfile.isPending}
                className="text-xs"
              >
                {updateProfile.isPending ? (
                  <Loader2 className="h-3.5 w-3.5 mr-1 animate-spin" />
                ) : (
                  <Save className="h-3.5 w-3.5 mr-1" />
                )}
                Lưu
              </Button>
            </div>
          </>
        )}
      </CardContent>
    </Card>
  );
}

/* ─── Change Password Section ────────────────────────────────────── */

function ChangePasswordSection() {
  const changePassword = useChangePassword();

  const [showForm, setShowForm] = useState(false);
  const [oldPassword, setOldPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [showOld, setShowOld] = useState(false);
  const [showNew, setShowNew] = useState(false);

  const passwordMismatch = confirmPassword !== "" && newPassword !== confirmPassword;
  const tooShort = newPassword !== "" && newPassword.length < 8;
  const canSubmit =
    oldPassword.length > 0 &&
    newPassword.length >= 8 &&
    newPassword === confirmPassword &&
    !changePassword.isPending;

  const handleSubmit = () => {
    if (!canSubmit) return;
    changePassword.mutate(
      { old_password: oldPassword, new_password: newPassword },
      {
        onSuccess: () => {
          setShowForm(false);
          setOldPassword("");
          setNewPassword("");
          setConfirmPassword("");
        },
      }
    );
  };

  const handleCancel = () => {
    setShowForm(false);
    setOldPassword("");
    setNewPassword("");
    setConfirmPassword("");
  };

  return (
    <Card>
      <CardHeader className="pb-3">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <Shield className="h-5 w-5 text-muted-foreground" />
            <CardTitle className="text-base">Bảo mật</CardTitle>
          </div>
          {!showForm && (
            <Button variant="ghost" size="sm" onClick={() => setShowForm(true)} className="text-xs">
              <Lock className="h-3.5 w-3.5 mr-1" />
              Đổi mật khẩu
            </Button>
          )}
        </div>
      </CardHeader>
      {showForm && (
        <CardContent className="space-y-3">
          <div className="space-y-1.5">
            <Label className="text-xs">Mật khẩu hiện tại</Label>
            <div className="relative">
              <Input
                type={showOld ? "text" : "password"}
                value={oldPassword}
                onChange={(e) => setOldPassword(e.target.value)}
                placeholder="Nhập mật khẩu hiện tại"
                className="h-10 pr-10 text-sm"
              />
              <button
                type="button"
                className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground"
                onClick={() => setShowOld(!showOld)}
              >
                {showOld ? <EyeOff className="h-4 w-4" /> : <Eye className="h-4 w-4" />}
              </button>
            </div>
          </div>

          <div className="space-y-1.5">
            <Label className="text-xs">Mật khẩu mới</Label>
            <div className="relative">
              <Input
                type={showNew ? "text" : "password"}
                value={newPassword}
                onChange={(e) => setNewPassword(e.target.value)}
                placeholder="Tối thiểu 8 ký tự"
                className="h-10 pr-10 text-sm"
              />
              <button
                type="button"
                className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground"
                onClick={() => setShowNew(!showNew)}
              >
                {showNew ? <EyeOff className="h-4 w-4" /> : <Eye className="h-4 w-4" />}
              </button>
            </div>
            {tooShort && (
              <p className="text-xs text-destructive">Mật khẩu phải có ít nhất 8 ký tự</p>
            )}
          </div>

          <div className="space-y-1.5">
            <Label className="text-xs">Xác nhận mật khẩu mới</Label>
            <Input
              type="password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              placeholder="Nhập lại mật khẩu mới"
              className="h-10 text-sm"
            />
            {passwordMismatch && (
              <p className="text-xs text-destructive">Mật khẩu xác nhận không khớp</p>
            )}
          </div>

          <div className="flex justify-end gap-2 pt-1">
            <Button variant="ghost" size="sm" onClick={handleCancel} className="text-xs">
              Huỷ
            </Button>
            <Button
              size="sm"
              onClick={handleSubmit}
              disabled={!canSubmit}
              className="text-xs"
            >
              {changePassword.isPending ? (
                <Loader2 className="h-3.5 w-3.5 mr-1 animate-spin" />
              ) : (
                <Lock className="h-3.5 w-3.5 mr-1" />
              )}
              Đổi mật khẩu
            </Button>
          </div>
        </CardContent>
      )}
      {!showForm && (
        <CardContent>
          <p className="text-xs text-muted-foreground">
            Thay đổi mật khẩu để bảo vệ tài khoản của bạn.
          </p>
        </CardContent>
      )}
    </Card>
  );
}

/* ─── Delete Account Section ─────────────────────────────────────── */

function DeleteAccountSection() {
  const deleteAccount = useDeleteAccount();
  const [password, setPassword] = useState("");
  const [dialogOpen, setDialogOpen] = useState(false);

  const handleDelete = () => {
    if (!password) return;
    deleteAccount.mutate(password, {
      onSuccess: () => setDialogOpen(false),
    });
  };

  return (
    <Card className="border-destructive/30">
      <CardHeader className="pb-3">
        <div className="flex items-center gap-2">
          <Trash2 className="h-5 w-5 text-destructive" />
          <CardTitle className="text-base text-destructive">Khu vực nguy hiểm</CardTitle>
        </div>
      </CardHeader>
      <CardContent className="space-y-3">
        <p className="text-sm text-muted-foreground">
          Xoá vĩnh viễn tài khoản và tất cả dữ liệu. Hành động này không thể hoàn tác.
        </p>

        <AlertDialog open={dialogOpen} onOpenChange={setDialogOpen}>
          <AlertDialogTrigger asChild>
            <Button variant="destructive" size="sm">
              <Trash2 className="h-4 w-4 mr-1.5" />
              Xoá tài khoản
            </Button>
          </AlertDialogTrigger>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Xoá tài khoản?</AlertDialogTitle>
              <AlertDialogDescription>
                Tất cả dữ liệu của bạn (công việc, ghi chú, danh mục…) sẽ bị xoá vĩnh viễn.
                Nhập mật khẩu để xác nhận.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <div className="py-2">
              <Input
                type="password"
                placeholder="Nhập mật khẩu xác nhận"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="h-10"
              />
            </div>
            <AlertDialogFooter>
              <AlertDialogCancel onClick={() => setPassword("")}>
                Huỷ
              </AlertDialogCancel>
              <AlertDialogAction
                onClick={handleDelete}
                disabled={!password || deleteAccount.isPending}
                className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
              >
                {deleteAccount.isPending ? (
                  <Loader2 className="h-4 w-4 mr-1.5 animate-spin" />
                ) : (
                  <Trash2 className="h-4 w-4 mr-1.5" />
                )}
                Xoá tài khoản
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </CardContent>
    </Card>
  );
}
