package lunar

// 10 Thiên Can (Heavenly Stems)
var ThienCan = [10]string{
	"Giáp", "Ất", "Bính", "Đinh", "Mậu",
	"Kỷ", "Canh", "Tân", "Nhâm", "Quý",
}

// 12 Địa Chi (Earthly Branches)
var DiaChi = [12]string{
	"Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ",
	"Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi",
}

// 12 con giáp
var ConGiap = [12]string{
	"Chuột", "Trâu", "Hổ", "Mèo", "Rồng", "Rắn",
	"Ngựa", "Dê", "Khỉ", "Gà", "Chó", "Lợn",
}

// CanChiInfo contains Can Chi information for a date
type CanChiInfo struct {
	DayCan    string `json:"day_can"`
	DayChi    string `json:"day_chi"`
	DayFull   string `json:"day_full"`
	MonthCan  string `json:"month_can"`
	MonthChi  string `json:"month_chi"`
	MonthFull string `json:"month_full"`
	YearCan   string `json:"year_can"`
	YearChi   string `json:"year_chi"`
	YearFull  string `json:"year_full"`
	HourCan   string `json:"hour_can,omitempty"`
	HourChi   string `json:"hour_chi,omitempty"`
	HourFull  string `json:"hour_full,omitempty"`
	Zodiac    string `json:"zodiac"`
}

// CanChiOfDay calculates Can Chi for a day based on Julian Day Number
func CanChiOfDay(jd float64) (string, string) {
	jdn := int(jd + 0.5)
	canIndex := (jdn + 9) % 10
	chiIndex := (jdn + 1) % 12
	return ThienCan[canIndex], DiaChi[chiIndex]
}

// CanChiOfMonth calculates Can Chi for a lunar month
func CanChiOfMonth(lunarMonth, lunarYear int) (string, string) {
	// Thiên Can of the year determines the starting Can for month
	yearCan := (lunarYear + 6) % 10
	// Each year pair starts with a different Can for month 1
	monthCanStart := (yearCan % 5) * 2
	canIndex := (monthCanStart + lunarMonth + 1) % 10

	// Địa Chi of month: tháng Giêng = Dần(2), tháng 2 = Mão(3), ...
	chiIndex := (lunarMonth + 1) % 12

	return ThienCan[canIndex], DiaChi[chiIndex]
}

// CanChiOfYear calculates Can Chi for a lunar year
func CanChiOfYear(lunarYear int) (string, string) {
	canIndex := (lunarYear + 6) % 10
	chiIndex := (lunarYear + 8) % 12
	return ThienCan[canIndex], DiaChi[chiIndex]
}

// CanChiOfHour calculates Can Chi for an hour
// hour is 0-23
func CanChiOfHour(hour int, dayJD float64) (string, string) {
	// Convert 24h to 12 Chi periods
	chiIndex := ((hour + 1) / 2) % 12

	// Day's Can determines the starting Can for the hour
	dayCan := (int(dayJD+0.5) + 9) % 10
	hourCanStart := (dayCan % 5) * 2
	canIndex := (hourCanStart + chiIndex) % 10

	return ThienCan[canIndex], DiaChi[chiIndex]
}

// ZodiacOfYear returns the zodiac animal for a lunar year
func ZodiacOfYear(lunarYear int) string {
	chiIndex := (lunarYear + 8) % 12
	return ConGiap[chiIndex]
}

// GetFullCanChi returns full Can Chi info for a given date
func GetFullCanChi(solarDay, solarMonth, solarYear int, lunar LunarDate) CanChiInfo {
	jd := JulianDayFromDate(solarDay, solarMonth, solarYear)

	dayCan, dayChi := CanChiOfDay(jd)
	monthCan, monthChi := CanChiOfMonth(lunar.Month, lunar.Year)
	yearCan, yearChi := CanChiOfYear(lunar.Year)

	return CanChiInfo{
		DayCan:    dayCan,
		DayChi:    dayChi,
		DayFull:   dayCan + " " + dayChi,
		MonthCan:  monthCan,
		MonthChi:  monthChi,
		MonthFull: monthCan + " " + monthChi,
		YearCan:   yearCan,
		YearChi:   yearChi,
		YearFull:  yearCan + " " + yearChi,
		Zodiac:    ZodiacOfYear(lunar.Year),
	}
}
