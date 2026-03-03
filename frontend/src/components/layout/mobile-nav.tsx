"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { Calendar, CheckSquare, FileText, Settings } from "lucide-react";
import { cn } from "@/lib/utils";

const navItems = [
  { label: "Lịch", href: "/calendar", icon: Calendar },
  { label: "Công việc", href: "/tasks", icon: CheckSquare },
  { label: "Ghi chú", href: "/notes", icon: FileText },
  { label: "Cài đặt", href: "/settings", icon: Settings },
];

export function MobileNav() {
  const pathname = usePathname();

  return (
    <nav className="fixed bottom-0 left-0 right-0 z-50 border-t bg-card/95 backdrop-blur-md md:hidden pb-safe">
      <div className="flex items-stretch justify-around">
        {navItems.map((item) => {
          const isActive =
            pathname === item.href || pathname.startsWith(item.href + "/");
          return (
            <Link
              key={item.href}
              href={item.href}
              className={cn(
                "flex flex-col items-center justify-center gap-0.5 py-2 px-4 min-h-[56px] min-w-[64px] active:scale-95 transition-transform",
                isActive ? "text-primary" : "text-muted-foreground"
              )}
            >
              <div className={cn(
                "flex items-center justify-center rounded-full transition-colors",
                isActive ? "bg-primary/10 w-10 h-7" : "w-10 h-7"
              )}>
                <item.icon className={cn("h-[18px] w-[18px]", isActive && "stroke-[2.5px]")} />
              </div>
              <span className={cn("text-[10px]", isActive ? "font-semibold" : "font-medium")}>
                {item.label}
              </span>
            </Link>
          );
        })}
      </div>
    </nav>
  );
}
