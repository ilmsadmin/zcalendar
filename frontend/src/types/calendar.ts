export interface SolarDate {
  day: number;
  month: number;
  year: number;
  day_of_week: number;
  day_of_week_name: string;
}

export interface LunarDate {
  day: number;
  month: number;
  year: number;
  leap: boolean;
  day_name?: string;
  month_name: string;
}

export interface CanChi {
  day: string;
  month: string;
  year: string;
  hour?: string;
  zodiac: string;
}

export interface SolarTerm {
  name: string;
  name_en: string;
}

export interface Holiday {
  name: string;
  description: string;
  is_lunar: boolean;
  is_official: boolean;
}

export interface GoodHour {
  chi: string;
  time: string;
  quality: string;
}

export interface GoodBad {
  quality: string;
  quality_name: string;
  good_for: string[];
  bad_for: string[];
  good_hours: GoodHour[];
}

export interface DayInfo {
  solar_date: SolarDate;
  lunar_date: LunarDate;
  can_chi: CanChi;
  solar_term?: SolarTerm;
  holidays: Holiday[];
  good_bad: GoodBad;
}

export interface TodayInfo extends DayInfo {
  greeting: string;
}

export interface MonthDay {
  solar_day: number;
  solar_month: number;
  lunar_day: number;
  lunar_month: number;
  lunar_leap?: boolean;
  lunar_day_name?: string;
  day_of_week: number;
  is_today?: boolean;
  is_current_month: boolean;
  has_holiday?: boolean;
  holiday_name?: string;
  can_chi_day: string;
}

export interface MonthInfo {
  year: number;
  month: number;
  days: MonthDay[];
}

export interface ConvertResult {
  solar: SolarDate;
  lunar: LunarDate;
}

export interface HolidayDate {
  solar_day: number;
  solar_month: number;
  name: string;
  description: string;
  is_lunar: boolean;
  is_official: boolean;
}

export interface HolidaysYear {
  year: number;
  holidays: HolidayDate[];
}
