"use client";

import Link from "next/link";
import { Moon as MoonIcon, Sun, LogOut, LogIn, User } from "lucide-react";
import { useTheme } from "next-themes";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useAuthStore } from "@/stores/auth-store";
import { useLogout } from "@/hooks/use-auth";

export function Header() {
  const { theme, setTheme } = useTheme();
  const user = useAuthStore((s) => s.user);
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);
  const logoutMutation = useLogout();

  const initials = user?.full_name
    ? user.full_name
        .split(" ")
        .map((n) => n[0])
        .join("")
        .toUpperCase()
        .slice(0, 2)
    : "U";

  return (
    <header className="flex h-12 md:h-14 items-center justify-between border-b bg-card/80 backdrop-blur-md px-3 md:px-6 sticky top-0 z-40">
      {/* Left: Brand */}
      <Link href="/calendar" className="flex items-center gap-1.5">
        <MoonIcon className="h-5 w-5 text-amber-500" />
        <span className="text-sm md:text-base font-bold bg-gradient-to-r from-amber-500 to-red-500 bg-clip-text text-transparent">
          Zplus Calendar
        </span>
      </Link>

      {/* Right: Actions */}
      <div className="flex items-center gap-1">
        {/* Theme toggle */}
        <Button
          variant="ghost"
          size="icon"
          className="h-8 w-8"
          onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
        >
          <Sun className="h-4 w-4 rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
          <MoonIcon className="absolute h-4 w-4 rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
        </Button>

        {/* User menu or Login button */}
        {isAuthenticated ? (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" size="icon" className="h-8 w-8">
                <Avatar className="h-7 w-7">
                  <AvatarFallback className="text-[10px] bg-primary/10 text-primary font-semibold">
                    {initials}
                  </AvatarFallback>
                </Avatar>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-52">
              <div className="px-3 py-2">
                <p className="text-sm font-semibold">{user?.full_name}</p>
                <p className="text-xs text-muted-foreground">{user?.email}</p>
              </div>
              <DropdownMenuSeparator />
              <DropdownMenuItem asChild>
                <Link href="/settings" className="gap-2">
                  <User className="h-4 w-4" /> Cài đặt
                </Link>
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={() => logoutMutation.mutate()}
                className="text-red-600 gap-2"
              >
                <LogOut className="h-4 w-4" />
                Đăng xuất
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        ) : (
          <Button
            variant="ghost"
            size="sm"
            className="h-8 text-xs gap-1"
            asChild
          >
            <Link href="/login">
              <LogIn className="h-3.5 w-3.5" />
              <span className="hidden sm:inline">Đăng nhập</span>
            </Link>
          </Button>
        )}
      </div>
    </header>
  );
}
