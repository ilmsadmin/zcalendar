package lunar

import (
	"math"
)

// TimeZone for Vietnam (UTC+7)
const TimeZoneVN = 7.0

// LunarDate represents a Vietnamese lunar calendar date
type LunarDate struct {
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	Leap      bool   `json:"leap"` // tháng nhuận
	DayName   string `json:"day_name"`
	MonthName string `json:"month_name"`
}

// SolarDate represents a solar (Gregorian) calendar date
type SolarDate struct {
	Day           int    `json:"day"`
	Month         int    `json:"month"`
	Year          int    `json:"year"`
	DayOfWeek     int    `json:"day_of_week"`
	DayOfWeekName string `json:"day_of_week_name"`
}

// getNewMoonDay returns the Julian Day of the new moon with timezone offset
func getNewMoonDay(k int) float64 {
	return math.Floor(NewMoon(k) + 0.5 + TimeZoneVN/24.0)
}

// getSunLongitudeAtMidnight returns the sun's longitude at midnight local time
func getSunLongitudeAtMidnight(jd float64) float64 {
	return SunLongitude(jd - 0.5 - TimeZoneVN/24.0)
}

// getLeapMonthOffset determines which lunar month is the leap month
// in a given lunar year. Returns 0 if no leap month.
// a11 is the Julian Day of the 11th month's new moon
func getLeapMonthOffset(a11 float64) int {
	k := int(math.Floor((a11-2415021.076998695)/29.530588853 + 0.5))
	last := 0
	i := 1
	arc := getSolarTermIndex(getSunLongitudeAtMidnight(getNewMoonDay(k + i)))

	for {
		last = arc
		i++
		arc = getSolarTermIndex(getSunLongitudeAtMidnight(getNewMoonDay(k + i)))
		if arc == last || i >= 14 {
			break
		}
	}
	return i - 1
}

// SolarToLunar converts a Solar date to Vietnamese Lunar date
func SolarToLunar(solarDay, solarMonth, solarYear int) LunarDate {
	jd := JulianDayFromDate(solarDay, solarMonth, solarYear)

	k := int(math.Floor((jd - 2415021.076998695) / 29.530588853))
	monthStart := getNewMoonDay(k + 1)

	if monthStart > jd {
		monthStart = getNewMoonDay(k)
	}

	// Find the start of the lunar month
	a11 := getLunarMonth11(solarYear)
	b11 := a11

	var lunarYear int
	if a11 >= monthStart {
		lunarYear = solarYear
		a11 = getLunarMonth11(solarYear - 1)
	} else {
		lunarYear = solarYear + 1
		b11 = getLunarMonth11(solarYear + 1)
	}

	lunarDay := int(jd - monthStart + 1)

	diff := int(math.Floor((monthStart-a11)/29 + 0.5))
	lunarLeap := false
	lunarMonth := diff + 11

	if b11-a11 > 365 {
		leapMonthDiff := getLeapMonthOffset(a11)
		if diff >= leapMonthDiff {
			lunarMonth = diff + 10
			if diff == leapMonthDiff {
				lunarLeap = true
			}
		}
	}

	if lunarMonth > 12 {
		lunarMonth -= 12
	}
	if lunarMonth >= 11 && diff < 4 {
		lunarYear -= 1
	}

	result := LunarDate{
		Day:   lunarDay,
		Month: lunarMonth,
		Year:  lunarYear,
		Leap:  lunarLeap,
	}

	result.DayName = getLunarDayName(lunarDay)
	result.MonthName = getLunarMonthName(lunarMonth)

	return result
}

// LunarToSolar converts a Vietnamese Lunar date to Solar date
func LunarToSolar(lunarDay, lunarMonth, lunarYear int, lunarLeap bool) (int, int, int) {
	var a11, b11 float64

	if lunarMonth < 11 {
		a11 = getLunarMonth11(lunarYear - 1)
		b11 = getLunarMonth11(lunarYear)
	} else {
		a11 = getLunarMonth11(lunarYear)
		b11 = getLunarMonth11(lunarYear + 1)
	}

	k := int(math.Floor(0.5 + (a11-2415021.076998695)/29.530588853))
	off := lunarMonth - 11
	if off < 0 {
		off += 12
	}

	if b11-a11 > 365 {
		leapOff := getLeapMonthOffset(a11)
		leapMonth := leapOff - 2
		if leapMonth < 0 {
			leapMonth += 12
		}

		if lunarLeap && lunarMonth != leapMonth {
			return 0, 0, 0 // Invalid leap month
		}

		if lunarLeap || (off >= leapOff) {
			off += 1
		}
	}

	monthStart := getNewMoonDay(k + off)
	return DateFromJulianDay(monthStart + float64(lunarDay) - 1)
}

// getLunarMonth11 returns the Julian Day of the start of lunar month 11
func getLunarMonth11(year int) float64 {
	off := JulianDayFromDate(31, 12, year) - 2415021.076998695
	k := int(math.Floor(off / 29.530588853))
	nm := getNewMoonDay(k)
	sunLong := getSolarTermIndex(getSunLongitudeAtMidnight(nm))

	if sunLong >= 9 {
		nm = getNewMoonDay(k - 1)
	}

	return nm
}

// getLunarDayName returns a display name for special lunar days
func getLunarDayName(day int) string {
	switch day {
	case 1:
		return "Mùng 1"
	case 15:
		return "Rằm"
	case 30:
		return "30"
	default:
		return ""
	}
}

// getLunarMonthName returns Vietnamese month name
func getLunarMonthName(month int) string {
	names := map[int]string{
		1:  "Tháng Giêng",
		2:  "Tháng Hai",
		3:  "Tháng Ba",
		4:  "Tháng Tư",
		5:  "Tháng Năm",
		6:  "Tháng Sáu",
		7:  "Tháng Bảy",
		8:  "Tháng Tám",
		9:  "Tháng Chín",
		10: "Tháng Mười",
		11: "Tháng Một",
		12: "Tháng Chạp",
	}
	if name, ok := names[month]; ok {
		return name
	}
	return ""
}
