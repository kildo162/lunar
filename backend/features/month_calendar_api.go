package features

import (
	"backend/features/calendar"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// MonthCalendarRequest represents the request for month calendar
type MonthCalendarRequest struct {
	Month    int     `json:"month"`
	Year     int     `json:"year"`
	TimeZone float64 `json:"timezone,omitempty"`
}

// DayInfo represents information for a single day
type DayInfo struct {
	SolarDay   int    `json:"solar_day"`
	SolarMonth int    `json:"solar_month"`
	SolarYear  int    `json:"solar_year"`
	LunarDay   int    `json:"lunar_day"`
	LunarMonth int    `json:"lunar_month"`
	LunarYear  int    `json:"lunar_year"`
	LeapMonth  int    `json:"leap_month"`
	Weekday    string `json:"weekday"`
	CanChiDay  string `json:"can_chi_day"`
	TietKhi    string `json:"tiet_khi,omitempty"`
	IsToday    bool   `json:"is_today"`
	IsWeekend  bool   `json:"is_weekend"`
}

// MonthCalendarResponse represents the month calendar response
type MonthCalendarResponse struct {
	Month     int       `json:"month"`
	Year      int       `json:"year"`
	MonthName string    `json:"month_name"`
	Days      []DayInfo `json:"days"`
	TotalDays int       `json:"total_days"`
}

// HandleGetMonthCalendar returns calendar information for a month
func HandleGetMonthCalendar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		handleGetMonthCalendarByQuery(w, r)
		return
	}

	if r.Method != "POST" {
		writeErrorResponse(w, "Method not allowed", "Only GET and POST methods are allowed", http.StatusMethodNotAllowed)
		return
	}

	var req MonthCalendarRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, "Invalid JSON", err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Month < 1 || req.Month > 12 || req.Year < 1 {
		writeErrorResponse(w, "Invalid date", "Please provide valid month (1-12) and year", http.StatusBadRequest)
		return
	}

	// Set default timezone
	if req.TimeZone == 0 {
		req.TimeZone = 7.0 // Vietnam timezone
	}

	response := buildMonthCalendarResponse(req.Month, req.Year, req.TimeZone)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// handleGetMonthCalendarByQuery handles GET request for month calendar
func handleGetMonthCalendarByQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	monthStr := query.Get("month")
	yearStr := query.Get("year")
	timezoneStr := query.Get("timezone")

	// Use current month/year if not provided
	now := time.Now()
	month := int(now.Month())
	year := now.Year()

	if monthStr != "" {
		if m, err := strconv.Atoi(monthStr); err == nil {
			month = m
		}
	}

	if yearStr != "" {
		if y, err := strconv.Atoi(yearStr); err == nil {
			year = y
		}
	}

	timezone := 7.0
	if timezoneStr != "" {
		if tz, err := strconv.ParseFloat(timezoneStr, 64); err == nil {
			timezone = tz
		}
	}

	// Validate input
	if month < 1 || month > 12 || year < 1 {
		writeErrorResponse(w, "Invalid date", "Please provide valid month (1-12) and year", http.StatusBadRequest)
		return
	}

	response := buildMonthCalendarResponse(month, year, timezone)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// buildMonthCalendarResponse builds the month calendar response
func buildMonthCalendarResponse(month, year int, timezone float64) MonthCalendarResponse {
	// Get days in month
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)
	totalDays := lastDay.Day()

	// Today for comparison
	today := time.Now()
	todayDay := today.Day()
	todayMonth := int(today.Month())
	todayYear := today.Year()

	var days []DayInfo

	// Generate days for the month
	for day := 1; day <= totalDays; day++ {
		cal := calendar.NewCalendar(calendar.CalendarDate{
			Day:      day,
			Month:    month,
			Year:     year,
			Hour:     0,
			Min:      0,
			Second:   0,
			TimeZone: timezone,
		})

		currentDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		weekday := currentDate.Weekday()

		// Check if it's today
		isToday := (day == todayDay && month == todayMonth && year == todayYear)

		// Check if it's weekend (Saturday or Sunday)
		isWeekend := (weekday == time.Saturday || weekday == time.Sunday)

		// Get Tiet Khi (only for the first day of the solar term)
		tietKhi := ""
		if day == 1 || shouldShowTietKhi(day, month, year, timezone) {
			tietKhi = cal.LunarDate.GetTietKhi()
		}

		dayInfo := DayInfo{
			SolarDay:   day,
			SolarMonth: month,
			SolarYear:  year,
			LunarDay:   cal.LunarDate.LunarDay,
			LunarMonth: cal.LunarDate.LunarMonth,
			LunarYear:  cal.LunarDate.LunarYear,
			LeapMonth:  cal.LunarDate.LeapMonth,
			Weekday:    calendar.TUAN[weekday],
			CanChiDay:  cal.LunarDate.GetCanDay(),
			TietKhi:    tietKhi,
			IsToday:    isToday,
			IsWeekend:  isWeekend,
		}

		days = append(days, dayInfo)
	}

	monthNames := []string{
		"", "Tháng Một", "Tháng Hai", "Tháng Ba", "Tháng Tư", "Tháng Năm", "Tháng Sáu",
		"Tháng Bảy", "Tháng Tám", "Tháng Chín", "Tháng Mười", "Tháng Mười Một", "Tháng Mười Hai",
	}

	return MonthCalendarResponse{
		Month:     month,
		Year:      year,
		MonthName: monthNames[month],
		Days:      days,
		TotalDays: totalDays,
	}
}

// shouldShowTietKhi determines if we should show Tiet Khi for this day
func shouldShowTietKhi(day, month, year int, timezone float64) bool {
	// For simplicity, let's show Tiet Khi every 15 days approximately
	// In a real implementation, you would calculate the exact solar terms
	return day%15 == 1
}
