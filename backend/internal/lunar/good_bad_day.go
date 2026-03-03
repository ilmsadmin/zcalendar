package lunar

// GoodBadInfo contains information about good/bad qualities of a day
type GoodBadInfo struct {
	Quality     string     `json:"quality"` // good, neutral, bad
	QualityName string     `json:"quality_name"`
	GoodFor     []string   `json:"good_for"`
	BadFor      []string   `json:"bad_for"`
	GoodHours   []HourInfo `json:"good_hours"`
}

// HourInfo represents information about a Chi hour
type HourInfo struct {
	Chi     string `json:"chi"`
	Time    string `json:"time"`
	Quality string `json:"quality"` // hoang_dao, hac_dao
}

// Hoàng đạo patterns based on Địa Chi of the day
// Each day chi maps to which hours are hoàng đạo
var hoangDaoPatterns = map[string][]int{
	"Tý":   {0, 1, 4, 5, 8, 9},   // Tý, Sửu, Thìn, Tỵ, Thân, Dậu
	"Sửu":  {2, 3, 6, 7, 10, 11}, // Dần, Mão, Ngọ, Mùi, Tuất, Hợi
	"Dần":  {0, 1, 4, 5, 8, 9},
	"Mão":  {2, 3, 6, 7, 10, 11},
	"Thìn": {0, 1, 4, 5, 8, 9},
	"Tỵ":   {2, 3, 6, 7, 10, 11},
	"Ngọ":  {0, 1, 4, 5, 8, 9},
	"Mùi":  {2, 3, 6, 7, 10, 11},
	"Thân": {0, 1, 4, 5, 8, 9},
	"Dậu":  {2, 3, 6, 7, 10, 11},
	"Tuất": {0, 1, 4, 5, 8, 9},
	"Hợi":  {2, 3, 6, 7, 10, 11},
}

// Chi hour time ranges
var chiHourRanges = [12]string{
	"23:00-01:00", // Tý
	"01:00-03:00", // Sửu
	"03:00-05:00", // Dần
	"05:00-07:00", // Mão
	"07:00-09:00", // Thìn
	"09:00-11:00", // Tỵ
	"11:00-13:00", // Ngọ
	"13:00-15:00", // Mùi
	"15:00-17:00", // Thân
	"17:00-19:00", // Dậu
	"19:00-21:00", // Tuất
	"21:00-23:00", // Hợi
}

// GetGoodHours returns the hoàng đạo/hắc đạo hours for a given day
func GetGoodHours(solarDay, solarMonth, solarYear int) []HourInfo {
	jd := JulianDayFromDate(solarDay, solarMonth, solarYear)
	_, dayChi := CanChiOfDay(jd)

	hours := make([]HourInfo, 12)
	hdPattern := hoangDaoPatterns[dayChi]

	for i := 0; i < 12; i++ {
		quality := "hac_dao"
		for _, hdIdx := range hdPattern {
			if hdIdx == i {
				quality = "hoang_dao"
				break
			}
		}

		hours[i] = HourInfo{
			Chi:     DiaChi[i],
			Time:    chiHourRanges[i],
			Quality: quality,
		}
	}

	return hours
}

// Good/Bad activities based on Can Chi combinations
var goodActivities = map[string][]string{
	"Giáp": {"Khai trương", "Xuất hành", "Cưới hỏi"},
	"Ất":   {"Cầu tài", "Giao dịch", "Trồng cây"},
	"Bính": {"Cưới hỏi", "Khai trương", "Nhập trạch"},
	"Đinh": {"Cầu tài", "Xuất hành", "Động thổ"},
	"Mậu":  {"Xây dựng", "Nhập trạch", "Khai trương"},
	"Kỷ":   {"Cưới hỏi", "Xuất hành", "Khai trương"},
	"Canh": {"Giao dịch", "Ký hợp đồng", "Xuất hành"},
	"Tân":  {"Cưới hỏi", "Cầu tài", "Khai trương"},
	"Nhâm": {"Xuất hành", "Nhập trạch", "Cầu tài"},
	"Quý":  {"Cưới hỏi", "Khai trương", "Giao dịch"},
}

var badActivities = map[string][]string{
	"Giáp": {"An táng", "Phá thổ"},
	"Ất":   {"Kiện tụng", "An táng"},
	"Bính": {"Phá thổ", "Đào giếng"},
	"Đinh": {"An táng", "Kiện tụng"},
	"Mậu":  {"Cưới hỏi", "Xuất hành xa"},
	"Kỷ":   {"Phá thổ", "An táng"},
	"Canh": {"Cưới hỏi", "An táng"},
	"Tân":  {"Phá thổ", "Kiện tụng"},
	"Nhâm": {"An táng", "Động thổ"},
	"Quý":  {"Phá thổ", "Đào giếng"},
}

// GetGoodBadDay returns good/bad information for a given day
func GetGoodBadDay(solarDay, solarMonth, solarYear int) GoodBadInfo {
	jd := JulianDayFromDate(solarDay, solarMonth, solarYear)
	dayCan, _ := CanChiOfDay(jd)

	goodFor := goodActivities[dayCan]
	badFor := badActivities[dayCan]
	goodHours := GetGoodHours(solarDay, solarMonth, solarYear)

	// Count hoàng đạo hours
	hoangDaoCount := 0
	for _, h := range goodHours {
		if h.Quality == "hoang_dao" {
			hoangDaoCount++
		}
	}

	quality := "neutral"
	qualityName := "Ngày Bình Thường"
	if hoangDaoCount >= 7 {
		quality = "good"
		qualityName = "Ngày Tốt"
	} else if hoangDaoCount <= 4 {
		quality = "bad"
		qualityName = "Ngày Xấu"
	}

	return GoodBadInfo{
		Quality:     quality,
		QualityName: qualityName,
		GoodFor:     goodFor,
		BadFor:      badFor,
		GoodHours:   goodHours,
	}
}
