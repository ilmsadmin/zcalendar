"use client";

import Link from "next/link";
import { useAuthStore } from "@/stores/auth-store";
import { LogIn, UserPlus } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export function AuthGuard({ children }: { children: React.ReactNode }) {
  const isAuthenticated = useAuthStore((s) => s.isAuthenticated);

  if (!isAuthenticated) {
    return (
      <div className="flex flex-1 items-center justify-center p-6">
        <Card className="w-full max-w-md text-center">
          <CardHeader>
            <CardTitle className="text-xl">Yêu cầu đăng nhập</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-sm text-muted-foreground">
              Bạn cần đăng nhập để sử dụng tính năng này.
            </p>
            <div className="flex items-center justify-center gap-3">
              <Button asChild>
                <Link href="/login">
                  <LogIn className="mr-2 h-4 w-4" />
                  Đăng nhập
                </Link>
              </Button>
              <Button variant="outline" asChild>
                <Link href="/register">
                  <UserPlus className="mr-2 h-4 w-4" />
                  Đăng ký
                </Link>
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  return <>{children}</>;
}
