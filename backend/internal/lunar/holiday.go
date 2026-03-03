package lunar

// VietnameseHolidayInfo represents a Vietnamese holiday
type VietnameseHolidayInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsLunar     bool   `json:"is_lunar"`
	Month       int    `json:"month"`
	Day         int    `json:"day"`
	IsOfficial  bool   `json:"is_official"` // Ngày nghỉ lễ chính thức
}

// SolarHolidays - Vietnamese holidays based on Solar calendar
var SolarHolidays = []VietnameseHolidayInfo{
	{Name: "Tết Dương lịch", Description: "Ngày đầu năm mới dương lịch", IsLunar: false, Month: 1, Day: 1, IsOfficial: true},
	{Name: "Ngày Thầy thuốc Việt Nam", Description: "Kỷ niệm ngày thầy thuốc Việt Nam", IsLunar: false, Month: 2, Day: 27, IsOfficial: false},
	{Name: "Quốc tế Phụ nữ", Description: "Ngày Quốc tế Phụ nữ", IsLunar: false, Month: 3, Day: 8, IsOfficial: false},
	{Name: "Ngày Giỗ Tổ Hùng Vương", Description: "Giỗ Tổ Hùng Vương - Mùng 10 tháng 3 âm lịch (also tracked in lunar)", IsLunar: false, Month: 0, Day: 0, IsOfficial: true},
	{Name: "Ngày Giải phóng miền Nam", Description: "Kỷ niệm ngày Giải phóng miền Nam, thống nhất đất nước", IsLunar: false, Month: 4, Day: 30, IsOfficial: true},
	{Name: "Ngày Quốc tế Lao động", Description: "Ngày Quốc tế Lao động", IsLunar: false, Month: 5, Day: 1, IsOfficial: true},
	{Name: "Ngày sinh Chủ tịch Hồ Chí Minh", Description: "Kỷ niệm ngày sinh Chủ tịch Hồ Chí Minh", IsLunar: false, Month: 5, Day: 19, IsOfficial: false},
	{Name: "Quốc tế Thiếu nhi", Description: "Ngày Quốc tế Thiếu nhi", IsLunar: false, Month: 6, Day: 1, IsOfficial: false},
	{Name: "Ngày Thương binh Liệt sĩ", Description: "Kỷ niệm ngày Thương binh Liệt sĩ", IsLunar: false, Month: 7, Day: 27, IsOfficial: false},
	{Name: "Quốc khánh", Description: "Ngày Quốc khánh nước CHXHCN Việt Nam", IsLunar: false, Month: 9, Day: 2, IsOfficial: true},
	{Name: "Ngày Phụ nữ Việt Nam", Description: "Ngày Phụ nữ Việt Nam", IsLunar: false, Month: 10, Day: 20, IsOfficial: false},
	{Name: "Ngày Nhà giáo Việt Nam", Description: "Kỷ niệm ngày Nhà giáo Việt Nam", IsLunar: false, Month: 11, Day: 20, IsOfficial: false},
	{Name: "Ngày thành lập Quân đội Nhân dân Việt Nam", Description: "Kỷ niệm ngày thành lập QĐND Việt Nam", IsLunar: false, Month: 12, Day: 22, IsOfficial: false},
}

// LunarHolidays - Vietnamese holidays based on Lunar calendar
var LunarHolidays = []VietnameseHolidayInfo{
	{Name: "Tết Nguyên Đán", Description: "Mùng 1 Tết Nguyên Đán", IsLunar: true, Month: 1, Day: 1, IsOfficial: true},
	{Name: "Mùng 2 Tết", Description: "Mùng 2 Tết Nguyên Đán", IsLunar: true, Month: 1, Day: 2, IsOfficial: true},
	{Name: "Mùng 3 Tết", Description: "Mùng 3 Tết Nguyên Đán", IsLunar: true, Month: 1, Day: 3, IsOfficial: true},
	{Name: "Tất Niên", Description: "Ngày cuối năm âm lịch (30 hoặc 29 tháng Chạp)", IsLunar: true, Month: 12, Day: 30, IsOfficial: false},
	{Name: "Ông Công Ông Táo", Description: "Tiễn ông Công ông Táo về trời", IsLunar: true, Month: 12, Day: 23, IsOfficial: false},
	{Name: "Tết Nguyên tiêu", Description: "Rằm tháng Giêng - Tết Nguyên tiêu", IsLunar: true, Month: 1, Day: 15, IsOfficial: false},
	{Name: "Tết Hàn thực", Description: "Tết Hàn thực - Bánh trôi bánh chay", IsLunar: true, Month: 3, Day: 3, IsOfficial: false},
	{Name: "Giỗ Tổ Hùng Vương", Description: "Giỗ Tổ Hùng Vương - Quốc lễ", IsLunar: true, Month: 3, Day: 10, IsOfficial: true},
	{Name: "Lễ Phật Đản", Description: "Lễ Phật Đản", IsLunar: true, Month: 4, Day: 15, IsOfficial: false},
	{Name: "Tết Đoan Ngọ", Description: "Tết Đoan Ngọ - Diệt sâu bọ", IsLunar: true, Month: 5, Day: 5, IsOfficial: false},
	{Name: "Lễ Vu Lan", Description: "Lễ Vu Lan - Báo hiếu", IsLunar: true, Month: 7, Day: 15, IsOfficial: false},
	{Name: "Tết Trung thu", Description: "Tết Trung thu - Tết thiếu nhi", IsLunar: true, Month: 8, Day: 15, IsOfficial: false},
	{Name: "Tết Trùng Cửu", Description: "Tết Trùng Cửu - Mùng 9 tháng 9", IsLunar: true, Month: 9, Day: 9, IsOfficial: false},
	{Name: "Tết Trùng Thập", Description: "Tết Trùng Thập - Mùng 10 tháng 10", IsLunar: true, Month: 10, Day: 10, IsOfficial: false},
}

// GetSolarHoliday returns the holiday name if the solar date is a holiday, empty string otherwise
func GetSolarHoliday(day, month int) *VietnameseHolidayInfo {
	for _, h := range SolarHolidays {
		if h.Month == month && h.Day == day {
			return &h
		}
	}
	return nil
}

// GetLunarHoliday returns the holiday name if the lunar date is a holiday, empty string otherwise
func GetLunarHoliday(lunarDay, lunarMonth int) *VietnameseHolidayInfo {
	for _, h := range LunarHolidays {
		if h.Month == lunarMonth && h.Day == lunarDay {
			return &h
		}
	}
	return nil
}

// GetHolidaysForSolarDate returns all holidays (solar + lunar) for a given solar date
func GetHolidaysForSolarDate(solarDay, solarMonth, solarYear int) []VietnameseHolidayInfo {
	var holidays []VietnameseHolidayInfo

	// Check solar holidays
	if h := GetSolarHoliday(solarDay, solarMonth); h != nil {
		holidays = append(holidays, *h)
	}

	// Convert to lunar and check lunar holidays
	lunarDate := SolarToLunar(solarDay, solarMonth, solarYear)
	if h := GetLunarHoliday(lunarDate.Day, lunarDate.Month); h != nil {
		holidays = append(holidays, *h)
	}

	return holidays
}

// HolidayEntry represents a holiday with its solar date
type HolidayEntry struct {
	SolarDay   int                   `json:"solar_day"`
	SolarMonth int                   `json:"solar_month"`
	Holiday    VietnameseHolidayInfo `json:"holiday"`
}

// GetAllHolidaysInYear returns all holidays (solar + lunar) for a given year
func GetAllHolidaysInYear(year int) []HolidayEntry {
	var result []HolidayEntry

	// Add solar holidays
	for _, h := range SolarHolidays {
		if h.Month > 0 && h.Day > 0 {
			result = append(result, HolidayEntry{
				SolarDay:   h.Day,
				SolarMonth: h.Month,
				Holiday:    h,
			})
		}
	}

	// Add lunar holidays with their solar equivalents
	for _, h := range LunarHolidays {
		sDay, sMonth, _ := LunarToSolar(h.Day, h.Month, year, false)
		if sDay > 0 {
			result = append(result, HolidayEntry{
				SolarDay:   sDay,
				SolarMonth: sMonth,
				Holiday:    h,
			})
		}
	}

	return result
}
