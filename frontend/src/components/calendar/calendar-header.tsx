"use client";

import { ChevronLeft, ChevronRight } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useCalendarStore, CalendarView, getWeekStart } from "@/stores/calendar-store";
import { DatePickerPopover } from "./date-picker-popover";

const monthNames = [
	"Tháng 1", "Tháng 2", "Tháng 3", "Tháng 4", "Tháng 5", "Tháng 6",
	"Tháng 7", "Tháng 8", "Tháng 9", "Tháng 10", "Tháng 11", "Tháng 12",
];

const dayOfWeekShort = ["CN", "T2", "T3", "T4", "T5", "T6", "T7"];

const viewOptions: { label: string; value: CalendarView }[] = [
	{ label: "Tháng", value: "month" },
	{ label: "Tuần", value: "week" },
	{ label: "Ngày", value: "day" },
];

function formatHeaderTitle(view: CalendarView, selectedDate: Date, currentMonth: number, currentYear: number): string {
	if (view === "month") {
		return `${monthNames[currentMonth - 1]} ${currentYear}`;
	}
	if (view === "week") {
		const ws = getWeekStart(selectedDate);
		const we = new Date(ws);
		we.setDate(we.getDate() + 6);
		const fmtD = (d: Date) => `${d.getDate()}/${d.getMonth() + 1}`;
		return `${fmtD(ws)} – ${fmtD(we)}, ${we.getFullYear()}`;
	}
	// day
	const dow = dayOfWeekShort[selectedDate.getDay()];
	return `${dow}, ${selectedDate.getDate()}/${selectedDate.getMonth() + 1}/${selectedDate.getFullYear()}`;
}

export function CalendarHeader() {
	const { currentMonth, currentYear, view, selectedDate, nextPeriod, prevPeriod, goToToday, setView, setMonth, setSelectedDate } =
		useCalendarStore();

	const title = formatHeaderTitle(view, selectedDate, currentMonth, currentYear);

	const handlePickerSelect = (month: number, year: number) => {
		setMonth(month, year);
		// Also update selected date to first of that month
		setSelectedDate(new Date(year, month - 1, 1));
	};

	const handleJumpToDate = (day: number, month: number, year: number) => {
		const d = new Date(year, month - 1, day);
		setSelectedDate(d);
		setMonth(d.getMonth() + 1, d.getFullYear());
	};

	return (
		<div className="space-y-2">
			{/* Row 1: Navigation + Title + Today */}
			<div className="flex items-center justify-between">
				<div className="flex items-center gap-0.5">
					<Button
						variant="ghost"
						size="icon"
						className="h-9 w-9 shrink-0 active:scale-90 transition-transform"
						onClick={prevPeriod}
					>
						<ChevronLeft className="h-5 w-5" />
					</Button>

					<DatePickerPopover
						currentMonth={currentMonth}
						currentYear={currentYear}
						onSelect={handlePickerSelect}
						onJumpToDate={handleJumpToDate}
						triggerLabel={
							<Button
								variant="ghost"
								className="h-9 px-2 text-base md:text-lg font-bold hover:bg-accent active:scale-95 transition-transform"
							>
								{title}
							</Button>
						}
					/>

					<Button
						variant="ghost"
						size="icon"
						className="h-9 w-9 shrink-0 active:scale-90 transition-transform"
						onClick={nextPeriod}
					>
						<ChevronRight className="h-5 w-5" />
					</Button>
				</div>

				<Button
					variant="outline"
					size="sm"
					onClick={goToToday}
					className="h-8 text-xs rounded-full px-3 shrink-0 active:scale-95 transition-transform"
				>
					Hôm nay
				</Button>
			</div>

			{/* Row 2: View tabs — iOS-style segmented control */}
			<div className="flex gap-1 bg-muted/50 rounded-xl p-1">
				{viewOptions.map((opt) => (
					<button
						key={opt.value}
						onClick={() => setView(opt.value)}
						className={`flex-1 py-1.5 text-xs font-medium rounded-lg transition-all active:scale-95 ${
							view === opt.value
								? "bg-background text-foreground shadow-sm"
								: "text-muted-foreground hover:text-foreground"
						}`}
					>
						{opt.label}
					</button>
				))}
			</div>
		</div>
	);
}
