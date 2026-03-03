"use client";

import { TodayBanner } from "@/components/calendar/today-banner";
import { CalendarHeader } from "@/components/calendar/calendar-header";
import { CalendarGrid } from "@/components/calendar/calendar-grid";
import { CalendarWeek } from "@/components/calendar/calendar-week";
import { CalendarDayView } from "@/components/calendar/calendar-day-view";
import { LunarInfoPanel } from "@/components/calendar/lunar-info-panel";
import { useCalendarStore } from "@/stores/calendar-store";

export default function CalendarPage() {
  const view = useCalendarStore((s) => s.view);

  return (
    <div className="px-3 py-3 md:px-4 md:py-4 lg:px-6 space-y-3">
      <TodayBanner />

      <div className="flex flex-col lg:flex-row gap-3 lg:gap-4">
        {/* Main Calendar */}
        <div className="flex-1 min-w-0 space-y-3">
          <CalendarHeader />

          {view === "month" && <CalendarGrid />}
          {view === "week" && <CalendarWeek />}
          {view === "day" && <CalendarDayView />}
        </div>

        {/* Lunar Info — inline below calendar on mobile, sidebar on lg */}
        {view !== "day" && (
          <div className="w-full lg:w-72 xl:w-80">
            <LunarInfoPanel />
          </div>
        )}
      </div>
    </div>
  );
}
