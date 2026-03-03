"use client";

import { useState, useCallback } from "react";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { ChevronLeft, ChevronRight, CalendarSearch, Moon, Sun, Loader2 } from "lucide-react";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { calendarApi } from "@/lib/api";
import { toast } from "sonner";

type PickerStep = "year" | "month" | "jump";
type DateType = "solar" | "lunar";

const monthNames = [
  "Th1", "Th2", "Th3", "Th4",
  "Th5", "Th6", "Th7", "Th8",
  "Th9", "Th10", "Th11", "Th12",
];

interface DatePickerPopoverProps {
  triggerLabel: React.ReactNode;
  currentMonth: number;
  currentYear: number;
  onSelect: (month: number, year: number) => void;
  onJumpToDate?: (day: number, month: number, year: number) => void;
}

export function DatePickerPopover({
  triggerLabel,
  currentMonth,
  currentYear,
  onSelect,
  onJumpToDate,
}: DatePickerPopoverProps) {
  const [open, setOpen] = useState(false);
  const [step, setStep] = useState<PickerStep>("year");
  const [selectedYear, setSelectedYear] = useState(currentYear);
  const [decadeStart, setDecadeStart] = useState(
    Math.floor(currentYear / 12) * 12
  );

  const [jumpDay, setJumpDay] = useState("");
  const [jumpMonth, setJumpMonth] = useState("");
  const [jumpYear, setJumpYear] = useState("");
  const [dateType, setDateType] = useState<DateType>("solar");
  const [isLeapMonth, setIsLeapMonth] = useState(false);
  const [isConverting, setIsConverting] = useState(false);

  const handleOpen = useCallback(
    (isOpen: boolean) => {
      setOpen(isOpen);
      if (isOpen) {
        setStep("year");
        setSelectedYear(currentYear);
        setDecadeStart(Math.floor(currentYear / 12) * 12);
        setJumpDay("");
        setJumpMonth("");
        setJumpYear("");
        setDateType("solar");
        setIsLeapMonth(false);
        setIsConverting(false);
      }
    },
    [currentYear]
  );

  const handleYearSelect = (year: number) => {
    setSelectedYear(year);
    setStep("month");
  };

  const handleMonthSelect = (month: number) => {
    onSelect(month, selectedYear);
    setOpen(false);
  };

  const handleJumpToDate = async () => {
    const d = parseInt(jumpDay);
    const m = parseInt(jumpMonth);
    const y = parseInt(jumpYear);
    if (!(d > 0 && d <= 31 && m > 0 && m <= 12 && y > 0 && y < 10000)) return;

    if (dateType === "solar") {
      onJumpToDate?.(d, m, y);
      setOpen(false);
    } else {
      // Lunar → convert to solar via API
      setIsConverting(true);
      try {
        const resp = await calendarApi.lunarToSolar(d, m, y, isLeapMonth);
        const solar = resp.data.data.solar;
        onJumpToDate?.(solar.day, solar.month, solar.year);
        setOpen(false);
      } catch {
        toast.error("Không thể chuyển đổi ngày âm lịch. Vui lòng kiểm tra lại.");
      } finally {
        setIsConverting(false);
      }
    }
  };

  const now = new Date();
  const todayYear = now.getFullYear();
  const todayMonth = now.getMonth() + 1;

  return (
    <Popover open={open} onOpenChange={handleOpen}>
      <PopoverTrigger asChild>{triggerLabel}</PopoverTrigger>
      <PopoverContent className="w-[280px] sm:w-[320px] p-0 rounded-2xl overflow-hidden" align="center">
        {/* Tabs */}
        <div className="flex border-b bg-muted/30">
          <button
            onClick={() => setStep("year")}
            className={cn(
              "flex-1 py-2.5 text-xs font-medium transition-colors active:bg-accent",
              step === "year" || step === "month"
                ? "text-primary border-b-2 border-primary"
                : "text-muted-foreground"
            )}
          >
            Chọn tháng/năm
          </button>
          {onJumpToDate && (
            <button
              onClick={() => setStep("jump")}
              className={cn(
                "flex-1 py-2.5 text-xs font-medium transition-colors active:bg-accent",
                step === "jump"
                  ? "text-primary border-b-2 border-primary"
                  : "text-muted-foreground"
              )}
            >
              <CalendarSearch className="inline h-3 w-3 mr-1" />
              Nhảy đến ngày
            </button>
          )}
        </div>

        {/* Year grid */}
        {step === "year" && (
          <div className="p-3 space-y-2.5">
            <div className="flex items-center justify-between">
              <Button
                variant="ghost"
                size="icon"
                className="h-9 w-9 active:scale-90 transition-transform"
                onClick={() => setDecadeStart((p) => p - 12)}
              >
                <ChevronLeft className="h-4 w-4" />
              </Button>
              <span className="text-sm font-semibold">
                {decadeStart} – {decadeStart + 11}
              </span>
              <Button
                variant="ghost"
                size="icon"
                className="h-9 w-9 active:scale-90 transition-transform"
                onClick={() => setDecadeStart((p) => p + 12)}
              >
                <ChevronRight className="h-4 w-4" />
              </Button>
            </div>
            <div className="grid grid-cols-4 gap-1.5">
              {Array.from({ length: 12 }, (_, i) => decadeStart + i).map(
                (year) => (
                  <button
                    key={year}
                    className={cn(
                      "h-10 rounded-xl text-xs font-medium transition-all active:scale-95",
                      year === currentYear
                        ? "bg-primary text-primary-foreground shadow-sm"
                        : "hover:bg-accent",
                      year === todayYear &&
                        year !== currentYear &&
                        "ring-1.5 ring-primary text-primary"
                    )}
                    onClick={() => handleYearSelect(year)}
                  >
                    {year}
                  </button>
                )
              )}
            </div>
          </div>
        )}

        {/* Month grid */}
        {step === "month" && (
          <div className="p-3 space-y-2.5">
            <div className="flex items-center justify-between">
              <button
                className="flex items-center text-xs font-medium text-muted-foreground hover:text-foreground active:scale-95 transition-transform h-9 px-2 rounded-lg"
                onClick={() => setStep("year")}
              >
                <ChevronLeft className="h-4 w-4 mr-0.5" />
                Năm {selectedYear}
              </button>
              <div className="flex gap-0.5">
                <Button
                  variant="ghost"
                  size="icon"
                  className="h-8 w-8 active:scale-90 transition-transform"
                  onClick={() => setSelectedYear((y) => y - 1)}
                >
                  <ChevronLeft className="h-3.5 w-3.5" />
                </Button>
                <Button
                  variant="ghost"
                  size="icon"
                  className="h-8 w-8 active:scale-90 transition-transform"
                  onClick={() => setSelectedYear((y) => y + 1)}
                >
                  <ChevronRight className="h-3.5 w-3.5" />
                </Button>
              </div>
            </div>
            <div className="grid grid-cols-3 gap-1.5">
              {monthNames.map((name, i) => {
                const m = i + 1;
                const isCurrent =
                  m === currentMonth && selectedYear === currentYear;
                const isToday =
                  m === todayMonth && selectedYear === todayYear;
                return (
                  <button
                    key={m}
                    className={cn(
                      "h-10 rounded-xl text-xs font-medium transition-all active:scale-95",
                      isCurrent
                        ? "bg-primary text-primary-foreground shadow-sm"
                        : "hover:bg-accent",
                      isToday &&
                        !isCurrent &&
                        "ring-1.5 ring-primary text-primary"
                    )}
                    onClick={() => handleMonthSelect(m)}
                  >
                    {name}
                  </button>
                );
              })}
            </div>
          </div>
        )}

        {/* Jump to date */}
        {step === "jump" && (
          <div className="p-3 space-y-3">
            <p className="text-xs text-muted-foreground">
              Nhập ngày muốn nhảy đến:
            </p>

            {/* Solar / Lunar toggle */}
            <div className="flex gap-1 bg-muted/50 rounded-xl p-1">
              <button
                onClick={() => { setDateType("solar"); setIsLeapMonth(false); }}
                className={cn(
                  "flex-1 flex items-center justify-center gap-1.5 py-2 text-xs font-medium rounded-lg transition-all active:scale-95",
                  dateType === "solar"
                    ? "bg-background text-foreground shadow-sm"
                    : "text-muted-foreground hover:text-foreground"
                )}
              >
                <Sun className="h-3.5 w-3.5" />
                Dương lịch
              </button>
              <button
                onClick={() => setDateType("lunar")}
                className={cn(
                  "flex-1 flex items-center justify-center gap-1.5 py-2 text-xs font-medium rounded-lg transition-all active:scale-95",
                  dateType === "lunar"
                    ? "bg-background text-foreground shadow-sm"
                    : "text-muted-foreground hover:text-foreground"
                )}
              >
                <Moon className="h-3.5 w-3.5" />
                Âm lịch
              </button>
            </div>

            <div className="grid grid-cols-3 gap-2">
              <div className="space-y-1">
                <Label className="text-[10px]">Ngày</Label>
                <Input
                  type="number"
                  min={1}
                  max={dateType === "lunar" ? 30 : 31}
                  placeholder="DD"
                  value={jumpDay}
                  onChange={(e) => setJumpDay(e.target.value)}
                  className="h-10 text-sm rounded-xl text-center"
                  inputMode="numeric"
                />
              </div>
              <div className="space-y-1">
                <Label className="text-[10px]">Tháng</Label>
                <Input
                  type="number"
                  min={1}
                  max={12}
                  placeholder="MM"
                  value={jumpMonth}
                  onChange={(e) => setJumpMonth(e.target.value)}
                  className="h-10 text-sm rounded-xl text-center"
                  inputMode="numeric"
                />
              </div>
              <div className="space-y-1">
                <Label className="text-[10px]">Năm</Label>
                <Input
                  type="number"
                  min={1}
                  max={9999}
                  placeholder="YYYY"
                  value={jumpYear}
                  onChange={(e) => setJumpYear(e.target.value)}
                  className="h-10 text-sm rounded-xl text-center"
                  inputMode="numeric"
                />
              </div>
            </div>

            {/* Leap month option for Lunar */}
            {dateType === "lunar" && (
              <label className="flex items-center gap-2 cursor-pointer select-none">
                <input
                  type="checkbox"
                  checked={isLeapMonth}
                  onChange={(e) => setIsLeapMonth(e.target.checked)}
                  className="h-4 w-4 rounded border-border accent-primary"
                />
                <span className="text-xs text-muted-foreground">Tháng nhuận</span>
              </label>
            )}

            <Button
              className="w-full h-10 rounded-xl text-sm active:scale-[0.98] transition-transform"
              onClick={handleJumpToDate}
              disabled={!jumpDay || !jumpMonth || !jumpYear || isConverting}
            >
              {isConverting ? (
                <Loader2 className="h-4 w-4 mr-1.5 animate-spin" />
              ) : (
                <CalendarSearch className="h-4 w-4 mr-1.5" />
              )}
              {dateType === "lunar" ? "Nhảy đến ngày (Âm lịch)" : "Nhảy đến ngày"}
            </Button>
            <Button
              variant="outline"
              className="w-full h-9 rounded-xl text-xs active:scale-[0.98] transition-transform"
              onClick={() => {
                const n = new Date();
                onJumpToDate?.(n.getDate(), n.getMonth() + 1, n.getFullYear());
                setOpen(false);
              }}
            >
              📍 Về hôm nay
            </Button>
          </div>
        )}
      </PopoverContent>
    </Popover>
  );
}
