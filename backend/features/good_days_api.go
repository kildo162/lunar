package features

import (
	"backend/features/calendar"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// GoodDaySearchRequest represents the request for searching good/bad days
type GoodDaySearchRequest struct {
	StartMonth int     `json:"start_month"`
	StartYear  int     `json:"start_year"`
	EndMonth   int     `json:"end_month"`
	EndYear    int     `json:"end_year"`
	Purpose    string  `json:"purpose,omitempty"` // wedding, business, travel, etc.
	TimeZone   float64 `json:"timezone,omitempty"`
}

// DayAnalysis represents the analysis of a day
type DayAnalysis struct {
	Date        string `json:"date"`
	SolarDay    int    `json:"solar_day"`
	SolarMonth  int    `json:"solar_month"`
	SolarYear   int    `json:"solar_year"`
	LunarDay    int    `json:"lunar_day"`
	LunarMonth  int    `json:"lunar_month"`
	LunarYear   int    `json:"lunar_year"`
	CanChiDay   string `json:"can_chi_day"`
	Score       int    `json:"score"` // 1-10, higher is better
	IsGoodDay   bool   `json:"is_good_day"`
	Reason      string `json:"reason"`
	GioHoangDao string `json:"gio_hoang_dao"`
	TietKhi     string `json:"tiet_khi"`
}

// GoodDaySearchResponse represents the response for good day search
type GoodDaySearchResponse struct {
	GoodDays []DayAnalysis `json:"good_days"`
	BadDays  []DayAnalysis `json:"bad_days"`
	Total    int           `json:"total"`
}

// HandleSearchGoodDays searches for good/bad days in a date range
func HandleSearchGoodDays(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		handleGetSearchGoodDays(w, r)
		return
	}

	if r.Method != "POST" {
		writeErrorResponse(w, "Method not allowed", "Only GET and POST methods are allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GoodDaySearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, "Invalid JSON", err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if req.StartMonth < 1 || req.StartMonth > 12 || req.StartYear < 1 ||
		req.EndMonth < 1 || req.EndMonth > 12 || req.EndYear < 1 {
		writeErrorResponse(w, "Invalid date range", "Please provide valid months (1-12) and years", http.StatusBadRequest)
		return
	}

	// Set default timezone
	if req.TimeZone == 0 {
		req.TimeZone = 7.0 // Vietnam timezone
	}

	response := searchGoodBadDays(req.StartMonth, req.StartYear, req.EndMonth, req.EndYear, req.Purpose, req.TimeZone)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// handleGetSearchGoodDays handles GET request for searching good days
func handleGetSearchGoodDays(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	startMonthStr := query.Get("start_month")
	startYearStr := query.Get("start_year")
	endMonthStr := query.Get("end_month")
	endYearStr := query.Get("end_year")
	purpose := query.Get("purpose")
	timezoneStr := query.Get("timezone")

	// Use current month/year if not provided
	now := time.Now()
	startMonth := int(now.Month())
	startYear := now.Year()
	endMonth := startMonth
	endYear := startYear

	if startMonthStr != "" {
		if m, err := strconv.Atoi(startMonthStr); err == nil {
			startMonth = m
		}
	}

	if startYearStr != "" {
		if y, err := strconv.Atoi(startYearStr); err == nil {
			startYear = y
		}
	}

	if endMonthStr != "" {
		if m, err := strconv.Atoi(endMonthStr); err == nil {
			endMonth = m
		}
	}

	if endYearStr != "" {
		if y, err := strconv.Atoi(endYearStr); err == nil {
			endYear = y
		}
	}

	timezone := 7.0
	if timezoneStr != "" {
		if tz, err := strconv.ParseFloat(timezoneStr, 64); err == nil {
			timezone = tz
		}
	}

	// Validate input
	if startMonth < 1 || startMonth > 12 || startYear < 1 ||
		endMonth < 1 || endMonth > 12 || endYear < 1 {
		writeErrorResponse(w, "Invalid date range", "Please provide valid months (1-12) and years", http.StatusBadRequest)
		return
	}

	response := searchGoodBadDays(startMonth, startYear, endMonth, endYear, purpose, timezone)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// searchGoodBadDays performs the actual search for good/bad days
func searchGoodBadDays(startMonth, startYear, endMonth, endYear int, purpose string, timezone float64) GoodDaySearchResponse {
	var goodDays []DayAnalysis
	var badDays []DayAnalysis
	total := 0

	// Convert to time for easier iteration
	startDate := time.Date(startYear, time.Month(startMonth), 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth+1), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1)

	// Limit search to maximum 3 months to avoid performance issues
	if endDate.Sub(startDate).Hours() > 24*90 {
		endDate = startDate.AddDate(0, 3, -1)
	}

	current := startDate
	for current.Before(endDate) || current.Equal(endDate) {
		cal := calendar.NewCalendar(calendar.CalendarDate{
			Day:      current.Day(),
			Month:    int(current.Month()),
			Year:     current.Year(),
			Hour:     0,
			Min:      0,
			Second:   0,
			TimeZone: timezone,
		})

		analysis := analyzeDayQuality(cal, purpose)
		total++

		if analysis.IsGoodDay {
			goodDays = append(goodDays, analysis)
		} else {
			badDays = append(badDays, analysis)
		}

		current = current.AddDate(0, 0, 1)
	}

	return GoodDaySearchResponse{
		GoodDays: goodDays,
		BadDays:  badDays,
		Total:    total,
	}
}

// analyzeDayQuality analyzes the quality of a day based on lunar calendar principles
func analyzeDayQuality(cal *calendar.Calendar, purpose string) DayAnalysis {
	solar := cal.SoalrDate
	lunar := cal.LunarDate

	// Use improved lunar rules
	rules := &calendar.LunarCalendarRules{}
	score := rules.GetDayQualityScore(lunar, purpose)

	// Determine if it's a good day
	isGoodDay := score >= 6

	// Format date
	dateStr := time.Date(solar.Year, time.Month(solar.Month), solar.Day, 0, 0, 0, 0, time.UTC).Format("02/01/2006")

	// Get detailed analysis
	reasons := rules.GetDetailedAnalysis(lunar, purpose)
	reasonStr := rules.GetDayDescription(lunar, score)
	if len(reasons) > 0 {
		reasonStr = reasons[0] // Take the first reason for API response
	}

	return DayAnalysis{
		Date:        dateStr,
		SolarDay:    solar.Day,
		SolarMonth:  solar.Month,
		SolarYear:   solar.Year,
		LunarDay:    lunar.LunarDay,
		LunarMonth:  lunar.LunarMonth,
		LunarYear:   lunar.LunarYear,
		CanChiDay:   lunar.GetCanDay(),
		Score:       score,
		IsGoodDay:   isGoodDay,
		Reason:      reasonStr,
		GioHoangDao: lunar.GetGioHoangDao(),
		TietKhi:     lunar.GetTietKhi(),
	}
}
