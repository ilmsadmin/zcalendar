package lunar

// 24 Tiết Khí (Solar Terms)
var SolarTermNames = [24]string{
	"Xuân phân",   // 0° - Vernal Equinox
	"Thanh minh",  // 15°
	"Cốc vũ",      // 30°
	"Lập hạ",      // 45°
	"Tiểu mãn",    // 60°
	"Mang chủng",  // 75°
	"Hạ chí",      // 90° - Summer Solstice
	"Tiểu thử",    // 105°
	"Đại thử",     // 120°
	"Lập thu",     // 135°
	"Xử thử",      // 150°
	"Bạch lộ",     // 165°
	"Thu phân",    // 180° - Autumnal Equinox
	"Hàn lộ",      // 195°
	"Sương giáng", // 210°
	"Lập đông",    // 225°
	"Tiểu tuyết",  // 240°
	"Đại tuyết",   // 255°
	"Đông chí",    // 270° - Winter Solstice
	"Tiểu hàn",    // 285°
	"Đại hàn",     // 300°
	"Lập xuân",    // 315°
	"Vũ thủy",     // 330°
	"Kinh trập",   // 345°
}

var SolarTermNamesEn = [24]string{
	"Vernal Equinox",
	"Clear and Bright",
	"Grain Rain",
	"Start of Summer",
	"Grain Buds",
	"Grain in Ear",
	"Summer Solstice",
	"Minor Heat",
	"Major Heat",
	"Start of Autumn",
	"End of Heat",
	"White Dew",
	"Autumnal Equinox",
	"Cold Dew",
	"Frost's Descent",
	"Start of Winter",
	"Minor Snow",
	"Major Snow",
	"Winter Solstice",
	"Minor Cold",
	"Major Cold",
	"Start of Spring",
	"Rain Water",
	"Awakening of Insects",
}

type SolarTermInfo struct {
	Name        string `json:"name"`
	NameEn      string `json:"name_en"`
	Description string `json:"description,omitempty"`
	Degree      int    `json:"degree"`
}

// GetSolarTerm returns the solar term for a given solar date
// Returns nil if the date is not a solar term boundary
func GetSolarTerm(solarDay, solarMonth, solarYear int) *SolarTermInfo {
	jd := JulianDayFromDate(solarDay, solarMonth, solarYear)

	sunLongToday := SunLongitude(jd - 0.5 - TimeZoneVN/24.0)
	sunLongYesterday := SunLongitude(jd - 1.5 - TimeZoneVN/24.0)

	todayTerm := int(sunLongToday / 15)
	yesterdayTerm := int(sunLongYesterday / 15)

	if todayTerm != yesterdayTerm {
		index := todayTerm % 24
		return &SolarTermInfo{
			Name:   SolarTermNames[index],
			NameEn: SolarTermNamesEn[index],
			Degree: todayTerm * 15,
		}
	}

	return nil
}

// GetCurrentSolarTerm returns the current active solar term for a date
func GetCurrentSolarTerm(solarDay, solarMonth, solarYear int) SolarTermInfo {
	jd := JulianDayFromDate(solarDay, solarMonth, solarYear)
	sunLong := SunLongitude(jd - 0.5 - TimeZoneVN/24.0)

	index := int(sunLong/15) % 24
	return SolarTermInfo{
		Name:   SolarTermNames[index],
		NameEn: SolarTermNamesEn[index],
		Degree: int(sunLong/15) * 15,
	}
}
