import { create } from "zustand";

export type CalendarView = "month" | "week" | "day";

interface CalendarState {
  selectedDate: Date;
  currentMonth: number;
  currentYear: number;
  view: CalendarView;
  setSelectedDate: (date: Date) => void;
  setMonth: (month: number, year: number) => void;
  setView: (view: CalendarView) => void;
  nextMonth: () => void;
  prevMonth: () => void;
  nextPeriod: () => void;
  prevPeriod: () => void;
  goToToday: () => void;
}

function getWeekStart(date: Date): Date {
  const d = new Date(date);
  const day = d.getDay(); // 0=CN
  d.setDate(d.getDate() - day);
  d.setHours(0, 0, 0, 0);
  return d;
}

export const useCalendarStore = create<CalendarState>((set, get) => {
  const now = new Date();
  return {
    selectedDate: now,
    currentMonth: now.getMonth() + 1,
    currentYear: now.getFullYear(),
    view: "month",

    setSelectedDate: (date) =>
      set({
        selectedDate: date,
        currentMonth: date.getMonth() + 1,
        currentYear: date.getFullYear(),
      }),

    setMonth: (month, year) => set({ currentMonth: month, currentYear: year }),

    setView: (view) => set({ view }),

    nextMonth: () => {
      const { currentMonth, currentYear } = get();
      if (currentMonth === 12) {
        set({ currentMonth: 1, currentYear: currentYear + 1 });
      } else {
        set({ currentMonth: currentMonth + 1 });
      }
    },

    prevMonth: () => {
      const { currentMonth, currentYear } = get();
      if (currentMonth === 1) {
        set({ currentMonth: 12, currentYear: currentYear - 1 });
      } else {
        set({ currentMonth: currentMonth - 1 });
      }
    },

    nextPeriod: () => {
      const { view, selectedDate } = get();
      if (view === "month") {
        get().nextMonth();
      } else if (view === "week") {
        const next = new Date(selectedDate);
        next.setDate(next.getDate() + 7);
        set({
          selectedDate: next,
          currentMonth: next.getMonth() + 1,
          currentYear: next.getFullYear(),
        });
      } else {
        const next = new Date(selectedDate);
        next.setDate(next.getDate() + 1);
        set({
          selectedDate: next,
          currentMonth: next.getMonth() + 1,
          currentYear: next.getFullYear(),
        });
      }
    },

    prevPeriod: () => {
      const { view, selectedDate } = get();
      if (view === "month") {
        get().prevMonth();
      } else if (view === "week") {
        const prev = new Date(selectedDate);
        prev.setDate(prev.getDate() - 7);
        set({
          selectedDate: prev,
          currentMonth: prev.getMonth() + 1,
          currentYear: prev.getFullYear(),
        });
      } else {
        const prev = new Date(selectedDate);
        prev.setDate(prev.getDate() - 1);
        set({
          selectedDate: prev,
          currentMonth: prev.getMonth() + 1,
          currentYear: prev.getFullYear(),
        });
      }
    },

    goToToday: () => {
      const now = new Date();
      set({
        selectedDate: now,
        currentMonth: now.getMonth() + 1,
        currentYear: now.getFullYear(),
      });
    },
  };
});

export { getWeekStart };
