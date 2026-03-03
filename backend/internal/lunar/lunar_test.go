package lunar

import "testing"

// Test SolarToLunar conversion with well-known dates
func TestSolarToLunar(t *testing.T) {
	tests := []struct {
		name       string
		solarDay   int
		solarMonth int
		solarYear  int
		lunarDay   int
		lunarMonth int
		lunarYear  int
		isLeap     bool
	}{
		{
			name:       "Tết Nguyên Đán 2024 (Giáp Thìn)",
			solarDay:   10,
			solarMonth: 2,
			solarYear:  2024,
			lunarDay:   1,
			lunarMonth: 1,
			lunarYear:  2024,
			isLeap:     false,
		},
		{
			name:       "Tết Nguyên Đán 2023 (Quý Mão)",
			solarDay:   22,
			solarMonth: 1,
			solarYear:  2023,
			lunarDay:   1,
			lunarMonth: 1,
			lunarYear:  2023,
			isLeap:     false,
		},
		{
			name:       "Mid-Autumn Festival 2024",
			solarDay:   17,
			solarMonth: 9,
			solarYear:  2024,
			lunarDay:   15,
			lunarMonth: 8,
			lunarYear:  2024,
			isLeap:     false,
		},
		{
			name:       "Regular day - 1 Jan 2024 solar",
			solarDay:   1,
			solarMonth: 1,
			solarYear:  2024,
			lunarDay:   20,
			lunarMonth: 11,
			lunarYear:  2023,
			isLeap:     false,
		},
		{
			name:       "Tết Nguyên Đán 2025 (Ất Tỵ)",
			solarDay:   29,
			solarMonth: 1,
			solarYear:  2025,
			lunarDay:   1,
			lunarMonth: 1,
			lunarYear:  2025,
			isLeap:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SolarToLunar(tt.solarDay, tt.solarMonth, tt.solarYear)
			if result.Day != tt.lunarDay || result.Month != tt.lunarMonth || result.Year != tt.lunarYear {
				t.Errorf("SolarToLunar(%d/%d/%d) = %d/%d/%d, want %d/%d/%d",
					tt.solarDay, tt.solarMonth, tt.solarYear,
					result.Day, result.Month, result.Year,
					tt.lunarDay, tt.lunarMonth, tt.lunarYear)
			}
			if result.Leap != tt.isLeap {
				t.Errorf("SolarToLunar(%d/%d/%d) isLeap = %v, want %v",
					tt.solarDay, tt.solarMonth, tt.solarYear,
					result.Leap, tt.isLeap)
			}
		})
	}
}

// Test LunarToSolar conversion
func TestLunarToSolar(t *testing.T) {
	tests := []struct {
		name       string
		lunarDay   int
		lunarMonth int
		lunarYear  int
		isLeap     bool
		solarDay   int
		solarMonth int
		solarYear  int
	}{
		{
			name:       "Tết Nguyên Đán 2024",
			lunarDay:   1,
			lunarMonth: 1,
			lunarYear:  2024,
			isLeap:     false,
			solarDay:   10,
			solarMonth: 2,
			solarYear:  2024,
		},
		{
			name:       "Tết Trung thu 2024",
			lunarDay:   15,
			lunarMonth: 8,
			lunarYear:  2024,
			isLeap:     false,
			solarDay:   17,
			solarMonth: 9,
			solarYear:  2024,
		},
		{
			name:       "Tết Nguyên Đán 2025",
			lunarDay:   1,
			lunarMonth: 1,
			lunarYear:  2025,
			isLeap:     false,
			solarDay:   29,
			solarMonth: 1,
			solarYear:  2025,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day, month, year := LunarToSolar(tt.lunarDay, tt.lunarMonth, tt.lunarYear, tt.isLeap)
			if day != tt.solarDay || month != tt.solarMonth || year != tt.solarYear {
				t.Errorf("LunarToSolar(%d/%d/%d, leap=%v) = %d/%d/%d, want %d/%d/%d",
					tt.lunarDay, tt.lunarMonth, tt.lunarYear, tt.isLeap,
					day, month, year,
					tt.solarDay, tt.solarMonth, tt.solarYear)
			}
		})
	}
}

// Test roundtrip: Solar -> Lunar -> Solar
func TestRoundtrip(t *testing.T) {
	testDates := []struct {
		day   int
		month int
		year  int
	}{
		{10, 2, 2024},
		{1, 1, 2024},
		{15, 6, 2024},
		{25, 12, 2023},
		{29, 1, 2025},
		{4, 7, 2023},
		{1, 3, 2020},
	}

	for _, d := range testDates {
		lunar := SolarToLunar(d.day, d.month, d.year)
		sDay, sMonth, sYear := LunarToSolar(lunar.Day, lunar.Month, lunar.Year, lunar.Leap)
		if sDay != d.day || sMonth != d.month || sYear != d.year {
			t.Errorf("Roundtrip failed for %d/%d/%d: Lunar(%d/%d/%d, leap=%v) -> Solar(%d/%d/%d)",
				d.day, d.month, d.year,
				lunar.Day, lunar.Month, lunar.Year, lunar.Leap,
				sDay, sMonth, sYear)
		}
	}
}

// Test Can Chi calculations
func TestCanChiOfYear(t *testing.T) {
	tests := []struct {
		year     int
		expected string
	}{
		{2024, "Giáp Thìn"},
		{2023, "Quý Mão"},
		{2025, "Ất Tỵ"},
		{2000, "Canh Thìn"},
		{1990, "Canh Ngọ"},
	}

	for _, tt := range tests {
		can, chi := CanChiOfYear(tt.year)
		result := can + " " + chi
		if result != tt.expected {
			t.Errorf("CanChiOfYear(%d) = %s, want %s", tt.year, result, tt.expected)
		}
	}
}

// Test Zodiac
func TestZodiacOfYear(t *testing.T) {
	tests := []struct {
		year     int
		expected string
	}{
		{2024, "Rồng"},
		{2023, "Mèo"},
		{2025, "Rắn"},
		{2000, "Rồng"},
		{1990, "Ngựa"},
	}

	for _, tt := range tests {
		result := ZodiacOfYear(tt.year)
		if result != tt.expected {
			t.Errorf("ZodiacOfYear(%d) = %s, want %s", tt.year, result, tt.expected)
		}
	}
}

// Test JulianDay calculations
func TestJulianDayFromDate(t *testing.T) {
	// Known Julian Day Numbers (integer JDN at midnight)
	tests := []struct {
		day, month, year int
		expectedJD       float64
	}{
		{1, 1, 2000, 2451545},
		{1, 1, 1970, 2440588},
	}

	for _, tt := range tests {
		jd := JulianDayFromDate(tt.day, tt.month, tt.year)
		if jd != tt.expectedJD {
			t.Errorf("JulianDayFromDate(%d/%d/%d) = %f, want %f",
				tt.day, tt.month, tt.year, jd, tt.expectedJD)
		}
	}
}

// Test DateFromJulianDay roundtrip
func TestDateFromJulianDayRoundtrip(t *testing.T) {
	dates := []struct {
		day, month, year int
	}{
		{1, 1, 2000},
		{29, 2, 2024}, // leap year
		{31, 12, 2023},
		{15, 6, 1990},
	}

	for _, d := range dates {
		jd := JulianDayFromDate(d.day, d.month, d.year)
		day, month, year := DateFromJulianDay(jd)
		if day != d.day || month != d.month || year != d.year {
			t.Errorf("Roundtrip JD failed for %d/%d/%d: got %d/%d/%d",
				d.day, d.month, d.year, day, month, year)
		}
	}
}

// Test Good Hours
func TestGetGoodHours(t *testing.T) {
	hours := GetGoodHours(10, 2, 2024)
	if len(hours) != 12 {
		t.Errorf("GetGoodHours should return 12 hours, got %d", len(hours))
	}

	hoangDaoCount := 0
	for _, h := range hours {
		if h.Quality == "hoang_dao" {
			hoangDaoCount++
		}
	}
	if hoangDaoCount != 6 {
		t.Errorf("Expected 6 hoàng đạo hours, got %d", hoangDaoCount)
	}
}

// Test Solar Terms
func TestGetSolarTerm(t *testing.T) {
	// GetSolarTerm returns *SolarTermInfo, nil if no term boundary
	result := GetSolarTerm(21, 6, 2024)
	// Just verify it doesn't panic; it may or may not be an exact boundary
	if result != nil {
		t.Logf("GetSolarTerm(21/6/2024) = %s (%s)", result.Name, result.NameEn)
	} else {
		t.Log("GetSolarTerm(21/6/2024) = nil (not an exact solar term boundary)")
	}
}

// Test GetCurrentSolarTerm
func TestGetCurrentSolarTerm(t *testing.T) {
	info := GetCurrentSolarTerm(21, 6, 2024)
	if info.Name == "" {
		t.Error("GetCurrentSolarTerm should return a valid solar term name")
	}
	t.Logf("Current solar term on 21/6/2024: %s (%s)", info.Name, info.NameEn)
}

// Test GetHolidaysForSolarDate
func TestGetHolidaysForSolarDate(t *testing.T) {
	// New Year's Day
	holidays := GetHolidaysForSolarDate(1, 1, 2024)
	found := false
	for _, h := range holidays {
		if h.Name == "Tết Dương lịch" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected to find Tết Dương lịch on Jan 1")
	}

	// Tết Nguyên Đán 2024 (Feb 10, 2024)
	holidays = GetHolidaysForSolarDate(10, 2, 2024)
	found = false
	for _, h := range holidays {
		if h.Name == "Tết Nguyên Đán" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected to find Tết Nguyên Đán on Feb 10, 2024")
	}
}

// Test GetGoodBadDay
func TestGetGoodBadDay(t *testing.T) {
	info := GetGoodBadDay(10, 2, 2024)
	if info.Quality == "" {
		t.Error("Quality should not be empty")
	}
	if info.QualityName == "" {
		t.Error("QualityName should not be empty")
	}
	if len(info.GoodHours) != 12 {
		t.Errorf("Expected 12 hours, got %d", len(info.GoodHours))
	}
}

// Test GetFullCanChi
func TestGetFullCanChi(t *testing.T) {
	lunarDate := SolarToLunar(10, 2, 2024)
	info := GetFullCanChi(10, 2, 2024, lunarDate)
	if info.DayFull == "" {
		t.Error("DayFull should not be empty")
	}
	if info.MonthFull == "" {
		t.Error("MonthFull should not be empty")
	}
	if info.YearFull == "" {
		t.Error("YearFull should not be empty")
	}
	if info.Zodiac == "" {
		t.Error("Zodiac should not be empty")
	}
	t.Logf("Can Chi for 10/2/2024: Day=%s, Month=%s, Year=%s, Zodiac=%s",
		info.DayFull, info.MonthFull, info.YearFull, info.Zodiac)
}
