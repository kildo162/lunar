package calendar

// LunarCalendarRules contains traditional Vietnamese lunar calendar rules for feng shui
type LunarCalendarRules struct{}

// Good and bad days based on lunar calendar principles
var (
	// Ngày tốt trong tháng âm lịch
	GoodLunarDays = []int{1, 3, 5, 8, 13, 15, 18, 21, 23, 25, 28}

	// Ngày xấu trong tháng âm lịch
	BadLunarDays = []int{7, 14, 17, 22, 27}

	// Can tốt
	GoodCan = []string{"Giáp", "Ất", "Canh", "Tân"}

	// Can xấu
	BadCan = []string{"Bính", "Đinh", "Mậu", "Kỷ", "Nhâm", "Quý"}

	// Chi tốt
	GoodChi = []string{"Tý", "Dần", "Mão", "Thìn", "Ngọ", "Mùi", "Dậu", "Tuất"}

	// Chi xấu
	BadChi = []string{"Sửu", "Tỵ", "Thân", "Hợi"}

	// Tháng âm lịch tốt
	GoodLunarMonths = []int{1, 3, 5, 7, 9, 11}

	// Tháng âm lịch xấu
	BadLunarMonths = []int{4, 6, 8, 10, 12}
)

// GetDayQualityScore calculates the quality score of a day based on lunar calendar rules
func (r *LunarCalendarRules) GetDayQualityScore(lunar *LunarDate, purpose string) int {
	score := 5 // Base score

	// Analyze lunar day
	if contains(GoodLunarDays, lunar.LunarDay) {
		score += 2
	} else if contains(BadLunarDays, lunar.LunarDay) {
		score -= 2
	}

	// Analyze Can Chi day
	canChiDay := lunar.GetCanDay()
	can := getCan(canChiDay)
	chi := getChi(canChiDay)

	if containsString(GoodCan, can) {
		score += 1
	} else if containsString(BadCan, can) {
		score -= 1
	}

	if containsString(GoodChi, chi) {
		score += 1
	} else if containsString(BadChi, chi) {
		score -= 1
	}

	// Analyze lunar month
	if contains(GoodLunarMonths, lunar.LunarMonth) {
		score += 1
	} else if contains(BadLunarMonths, lunar.LunarMonth) {
		score -= 1
	}

	// Special analysis for new moon (sóc) and full moon (vọng)
	if lunar.LunarDay == 1 || lunar.LunarDay == 15 {
		score += 2
	}

	// Purpose-specific scoring
	score += r.getPurposeScore(lunar, purpose)

	// Ensure score is within bounds
	if score > 10 {
		score = 10
	} else if score < 1 {
		score = 1
	}

	return score
}

// getPurposeScore returns additional score based on purpose
func (r *LunarCalendarRules) getPurposeScore(lunar *LunarDate, purpose string) int {
	score := 0

	switch purpose {
	case "wedding", "cưới hỏi":
		// Odd lunar days are better for weddings
		if lunar.LunarDay%2 == 1 {
			score += 2
		}
		// Specific good days for wedding
		if contains([]int{3, 9, 15, 18, 21, 25}, lunar.LunarDay) {
			score += 1
		}

	case "business", "kinh doanh":
		// Days with good business energy
		if contains([]int{1, 5, 8, 13, 18, 21, 28}, lunar.LunarDay) {
			score += 1
		}

	case "travel", "du lịch":
		// Good days for starting journeys
		if contains([]int{3, 8, 13, 18, 23}, lunar.LunarDay) {
			score += 1
		}

	case "construction", "xây dựng":
		// Good days for construction
		if contains([]int{1, 5, 8, 15, 21, 28}, lunar.LunarDay) {
			score += 1
		}

	case "moving", "chuyển nhà":
		// Good days for moving
		if contains([]int{3, 8, 13, 18, 25}, lunar.LunarDay) {
			score += 1
		}
	}

	return score
}

// GetDayDescription returns a description of the day's quality
func (r *LunarCalendarRules) GetDayDescription(lunar *LunarDate, score int) string {
	switch {
	case score >= 8:
		return "Đại cát - Rất tốt"
	case score >= 6:
		return "Cát - Tốt"
	case score == 5:
		return "Bình - Trung bình"
	case score >= 3:
		return "Hung - Không tốt"
	default:
		return "Đại hung - Rất xấu"
	}
}

// GetDetailedAnalysis returns detailed analysis of the day
func (r *LunarCalendarRules) GetDetailedAnalysis(lunar *LunarDate, purpose string) []string {
	var reasons []string

	// Lunar day analysis
	if contains(GoodLunarDays, lunar.LunarDay) {
		reasons = append(reasons, "Ngày âm lịch tốt")
	} else if contains(BadLunarDays, lunar.LunarDay) {
		reasons = append(reasons, "Ngày âm lịch xấu")
	}

	// Special days
	if lunar.LunarDay == 1 {
		reasons = append(reasons, "Ngày sóc (trăng mới) - rất tốt")
	} else if lunar.LunarDay == 15 {
		reasons = append(reasons, "Ngày vọng (trăng tròn) - rất tốt")
	}

	// Can Chi analysis
	canChiDay := lunar.GetCanDay()
	can := getCan(canChiDay)
	chi := getChi(canChiDay)

	if containsString(GoodCan, can) {
		reasons = append(reasons, "Can "+can+" tốt")
	} else if containsString(BadCan, can) {
		reasons = append(reasons, "Can "+can+" xấu")
	}

	if containsString(GoodChi, chi) {
		reasons = append(reasons, "Chi "+chi+" tốt")
	} else if containsString(BadChi, chi) {
		reasons = append(reasons, "Chi "+chi+" xấu")
	}

	// Purpose-specific analysis
	switch purpose {
	case "wedding", "cưới hỏi":
		if lunar.LunarDay%2 == 1 {
			reasons = append(reasons, "Ngày lẻ tốt cho cưới hỏi")
		}
	case "business", "kinh doanh":
		reasons = append(reasons, "Phù hợp kinh doanh")
	case "travel", "du lịch":
		reasons = append(reasons, "Phù hợp du lịch")
	}

	if len(reasons) == 0 {
		reasons = append(reasons, "Ngày bình thường")
	}

	return reasons
}

// Helper functions
func contains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func containsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func getCan(canChi string) string {
	// Extract Can from Can Chi string (first part before space)
	for _, can := range CAN {
		if len(canChi) > 0 && canChi[0:len(can)] == can {
			return can
		}
	}
	return ""
}

func getChi(canChi string) string {
	// Extract Chi from Can Chi string (second part after space)
	for _, chi := range CHI {
		if len(canChi) >= len(chi) && canChi[len(canChi)-len(chi):] == chi {
			return chi
		}
	}
	return ""
}
