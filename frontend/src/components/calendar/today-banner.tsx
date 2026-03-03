"use client";

import { useToday } from "@/hooks/use-calendar";
import { Skeleton } from "@/components/ui/skeleton";
import { Moon, Sparkles } from "lucide-react";

export function TodayBanner() {
  const { data: today, isLoading } = useToday();

  if (isLoading) {
    return <Skeleton className="h-16 w-full rounded-2xl" />;
  }

  if (!today) return null;

  return (
    <div className="relative overflow-hidden rounded-2xl bg-gradient-to-r from-amber-500 via-orange-500 to-red-500 text-white p-3 md:p-4 shadow-lg">
      <div className="flex items-center gap-3">
        <div className="flex h-10 w-10 md:h-12 md:w-12 shrink-0 items-center justify-center rounded-xl bg-white/20 backdrop-blur-sm">
          <Moon className="h-5 w-5 md:h-6 md:w-6" />
        </div>
        <div className="flex-1 min-w-0">
          <p className="text-xs opacity-90 truncate">{today.greeting}</p>
          <p className="text-sm md:text-base font-bold truncate">
            {today.solar_date.day_of_week_name}, {today.solar_date.day}/
            {today.solar_date.month}/{today.solar_date.year}
          </p>
          <p className="text-[11px] md:text-xs opacity-85 truncate">
            ÂL: {today.lunar_date.day_name || `${today.lunar_date.day}`}{" "}
            {today.lunar_date.month_name} – {today.can_chi.year}
          </p>
        </div>
        {today.solar_term && (
          <div className="hidden sm:flex items-center gap-1 rounded-xl bg-white/15 px-2.5 py-1.5 shrink-0">
            <Sparkles className="h-3.5 w-3.5" />
            <span className="text-xs font-medium">{today.solar_term.name}</span>
          </div>
        )}
      </div>
      {/* Decorative circles */}
      <div className="absolute -right-4 -top-4 h-20 w-20 rounded-full bg-white/10" />
      <div className="absolute -right-2 -bottom-6 h-16 w-16 rounded-full bg-white/5" />
    </div>
  );
}
