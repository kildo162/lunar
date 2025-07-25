package features

import (
	"backend/features/calendar"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// ConvertToLunarRequest represents the request for solar to lunar conversion
type ConvertToLunarRequest struct {
	Day      int     `json:"day"`
	Month    int     `json:"month"`
	Year     int     `json:"year"`
	Hour     int     `json:"hour,omitempty"`
	Minute   int     `json:"minute,omitempty"`
	Second   int     `json:"second,omitempty"`
	TimeZone float64 `json:"timezone,omitempty"`
}

// ConvertToSolarRequest represents the request for lunar to solar conversion
type ConvertToSolarRequest struct {
	LunarDay   int     `json:"lunar_day"`
	LunarMonth int     `json:"lunar_month"`
	LunarYear  int     `json:"lunar_year"`
	LeapMonth  int     `json:"leap_month,omitempty"`
	TimeZone   float64 `json:"timezone,omitempty"`
}

// LunarResponse represents the lunar date response
type LunarResponse struct {
	LunarDay    int    `json:"lunar_day"`
	LunarMonth  int    `json:"lunar_month"`
	LunarYear   int    `json:"lunar_year"`
	LeapMonth   int    `json:"leap_month"`
	MonthName   string `json:"month_name"`
	CanChiYear  string `json:"can_chi_year"`
	CanChiMonth string `json:"can_chi_month"`
	CanChiDay   string `json:"can_chi_day"`
	CanChiHour  string `json:"can_chi_hour"`
	GioHoangDao string `json:"gio_hoang_dao"`
	TietKhi     string `json:"tiet_khi"`
	Detail      string `json:"detail"`
}

// SolarResponse represents the solar date response
type SolarResponse struct {
	Day        int    `json:"day"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
	LeapYear   bool   `json:"leap_year"`
	DayOfYear  int    `json:"day_of_year"`
	WeekOfYear int    `json:"week_of_year"`
	Weekday    string `json:"weekday"`
	Detail     string `json:"detail"`
}

// CalendarResponse represents the combined calendar response
type CalendarResponse struct {
	Solar SolarResponse `json:"solar"`
	Lunar LunarResponse `json:"lunar"`
}

// ErrorResponse represents error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// HandleConvertToLunar converts solar date to lunar date
func HandleConvertToLunar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		// Handle GET request with query parameters
		handleGetConvertToLunar(w, r)
		return
	}

	if r.Method != "POST" {
		writeErrorResponse(w, "Method not allowed", "Only GET and POST methods are allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ConvertToLunarRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, "Invalid JSON", err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Day < 1 || req.Day > 31 || req.Month < 1 || req.Month > 12 || req.Year < 1 {
		writeErrorResponse(w, "Invalid date", "Please provide valid day (1-31), month (1-12), and year", http.StatusBadRequest)
		return
	}

	// Set default values
	if req.TimeZone == 0 {
		req.TimeZone = 7.0 // Vietnam timezone
	}

	// Create calendar
	cal := calendar.NewCalendar(calendar.CalendarDate{
		Day:      req.Day,
		Month:    req.Month,
		Year:     req.Year,
		Hour:     req.Hour,
		Min:      req.Minute,
		Second:   req.Second,
		TimeZone: req.TimeZone,
	})

	response := CalendarResponse{
		Solar: buildSolarResponse(cal.SoalrDate),
		Lunar: buildLunarResponse(cal.LunarDate),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// handleGetConvertToLunar handles GET request for solar to lunar conversion
func handleGetConvertToLunar(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Parse parameters
	dayStr := query.Get("day")
	monthStr := query.Get("month")
	yearStr := query.Get("year")
	hourStr := query.Get("hour")
	minuteStr := query.Get("minute")
	secondStr := query.Get("second")
	timezoneStr := query.Get("timezone")

	// Use current date if no parameters provided
	if dayStr == "" || monthStr == "" || yearStr == "" {
		now := time.Now()
		dayStr = strconv.Itoa(now.Day())
		monthStr = strconv.Itoa(int(now.Month()))
		yearStr = strconv.Itoa(now.Year())
		if hourStr == "" {
			hourStr = strconv.Itoa(now.Hour())
		}
		if minuteStr == "" {
			minuteStr = strconv.Itoa(now.Minute())
		}
		if secondStr == "" {
			secondStr = strconv.Itoa(now.Second())
		}
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		writeErrorResponse(w, "Invalid day", "Day must be a number", http.StatusBadRequest)
		return
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		writeErrorResponse(w, "Invalid month", "Month must be a number", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		writeErrorResponse(w, "Invalid year", "Year must be a number", http.StatusBadRequest)
		return
	}

	hour := 0
	if hourStr != "" {
		hour, _ = strconv.Atoi(hourStr)
	}

	minute := 0
	if minuteStr != "" {
		minute, _ = strconv.Atoi(minuteStr)
	}

	second := 0
	if secondStr != "" {
		second, _ = strconv.Atoi(secondStr)
	}

	timezone := 7.0
	if timezoneStr != "" {
		timezone, _ = strconv.ParseFloat(timezoneStr, 64)
	}

	// Validate input
	if day < 1 || day > 31 || month < 1 || month > 12 || year < 1 {
		writeErrorResponse(w, "Invalid date", "Please provide valid day (1-31), month (1-12), and year", http.StatusBadRequest)
		return
	}

	// Create calendar
	cal := calendar.NewCalendar(calendar.CalendarDate{
		Day:      day,
		Month:    month,
		Year:     year,
		Hour:     hour,
		Min:      minute,
		Second:   second,
		TimeZone: timezone,
	})

	response := CalendarResponse{
		Solar: buildSolarResponse(cal.SoalrDate),
		Lunar: buildLunarResponse(cal.LunarDate),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleConvertToSolar converts lunar date to solar date
func HandleConvertToSolar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		// Handle GET request with query parameters
		handleGetConvertToSolar(w, r)
		return
	}

	if r.Method != "POST" {
		writeErrorResponse(w, "Method not allowed", "Only GET and POST methods are allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ConvertToSolarRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, "Invalid JSON", err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if req.LunarDay < 1 || req.LunarDay > 30 || req.LunarMonth < 1 || req.LunarMonth > 12 || req.LunarYear < 1 {
		writeErrorResponse(w, "Invalid lunar date", "Please provide valid lunar day (1-30), month (1-12), and year", http.StatusBadRequest)
		return
	}

	// Set default values
	if req.TimeZone == 0 {
		req.TimeZone = 7.0 // Vietnam timezone
	}

	// Convert lunar to solar
	day, month, year := calendar.ConvertLunar2Solar(req.LunarDay, req.LunarMonth, req.LunarYear, req.LeapMonth, req.TimeZone)

	if day == 0 || month == 0 || year == 0 {
		writeErrorResponse(w, "Invalid conversion", "Unable to convert the provided lunar date", http.StatusBadRequest)
		return
	}

	// Create calendar with converted solar date
	cal := calendar.NewCalendar(calendar.CalendarDate{
		Day:      day,
		Month:    month,
		Year:     year,
		Hour:     0,
		Min:      0,
		Second:   0,
		TimeZone: req.TimeZone,
	})

	response := CalendarResponse{
		Solar: buildSolarResponse(cal.SoalrDate),
		Lunar: buildLunarResponse(cal.LunarDate),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// handleGetConvertToSolar handles GET request for lunar to solar conversion
func handleGetConvertToSolar(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	lunarDayStr := query.Get("lunar_day")
	lunarMonthStr := query.Get("lunar_month")
	lunarYearStr := query.Get("lunar_year")
	leapMonthStr := query.Get("leap_month")
	timezoneStr := query.Get("timezone")

	if lunarDayStr == "" || lunarMonthStr == "" || lunarYearStr == "" {
		writeErrorResponse(w, "Missing parameters", "lunar_day, lunar_month, and lunar_year are required", http.StatusBadRequest)
		return
	}

	lunarDay, err := strconv.Atoi(lunarDayStr)
	if err != nil {
		writeErrorResponse(w, "Invalid lunar_day", "lunar_day must be a number", http.StatusBadRequest)
		return
	}

	lunarMonth, err := strconv.Atoi(lunarMonthStr)
	if err != nil {
		writeErrorResponse(w, "Invalid lunar_month", "lunar_month must be a number", http.StatusBadRequest)
		return
	}

	lunarYear, err := strconv.Atoi(lunarYearStr)
	if err != nil {
		writeErrorResponse(w, "Invalid lunar_year", "lunar_year must be a number", http.StatusBadRequest)
		return
	}

	leapMonth := 0
	if leapMonthStr != "" {
		leapMonth, _ = strconv.Atoi(leapMonthStr)
	}

	timezone := 7.0
	if timezoneStr != "" {
		timezone, _ = strconv.ParseFloat(timezoneStr, 64)
	}

	// Validate input
	if lunarDay < 1 || lunarDay > 30 || lunarMonth < 1 || lunarMonth > 12 || lunarYear < 1 {
		writeErrorResponse(w, "Invalid lunar date", "Please provide valid lunar day (1-30), month (1-12), and year", http.StatusBadRequest)
		return
	}

	// Convert lunar to solar
	day, month, year := calendar.ConvertLunar2Solar(lunarDay, lunarMonth, lunarYear, leapMonth, timezone)

	if day == 0 || month == 0 || year == 0 {
		writeErrorResponse(w, "Invalid conversion", "Unable to convert the provided lunar date", http.StatusBadRequest)
		return
	}

	// Create calendar with converted solar date
	cal := calendar.NewCalendar(calendar.CalendarDate{
		Day:      day,
		Month:    month,
		Year:     year,
		Hour:     0,
		Min:      0,
		Second:   0,
		TimeZone: timezone,
	})

	response := CalendarResponse{
		Solar: buildSolarResponse(cal.SoalrDate),
		Lunar: buildLunarResponse(cal.LunarDate),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HandleGetToday returns today's calendar information
func HandleGetToday(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "GET" {
		writeErrorResponse(w, "Method not allowed", "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	now := time.Now()
	timezone := 7.0 // Vietnam timezone

	// Check for timezone parameter
	if tzStr := r.URL.Query().Get("timezone"); tzStr != "" {
		if tz, err := strconv.ParseFloat(tzStr, 64); err == nil {
			timezone = tz
		}
	}

	cal := calendar.NewCalendar(calendar.CalendarDate{
		Day:      now.Day(),
		Month:    int(now.Month()),
		Year:     now.Year(),
		Hour:     now.Hour(),
		Min:      now.Minute(),
		Second:   now.Second(),
		TimeZone: timezone,
	})

	response := CalendarResponse{
		Solar: buildSolarResponse(cal.SoalrDate),
		Lunar: buildLunarResponse(cal.LunarDate),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Helper functions
func buildSolarResponse(solar *calendar.SolarDate) SolarResponse {
	return SolarResponse{
		Day:        solar.Day,
		Month:      solar.Month,
		Year:       solar.Year,
		LeapYear:   solar.LeapYear,
		DayOfYear:  solar.GetDayOfYear(),
		WeekOfYear: solar.GetWeekOfYear(),
		Weekday:    calendar.TUAN[solar.GetWeekday()],
		Detail:     solar.Detail(),
	}
}

func buildLunarResponse(lunar *calendar.LunarDate) LunarResponse {
	monthName := ""
	if lunar.LunarMonth > 0 && lunar.LunarMonth <= 12 {
		monthName = calendar.THANG[lunar.LunarMonth-1]
	}

	return LunarResponse{
		LunarDay:    lunar.LunarDay,
		LunarMonth:  lunar.LunarMonth,
		LunarYear:   lunar.LunarYear,
		LeapMonth:   lunar.LeapMonth,
		MonthName:   monthName,
		CanChiYear:  lunar.GetCanChiYear(),
		CanChiMonth: lunar.GetCanChiMonth(),
		CanChiDay:   lunar.GetCanDay(),
		CanChiHour:  lunar.GetCanHour(),
		GioHoangDao: lunar.GetGioHoangDao(),
		TietKhi:     lunar.GetTietKhi(),
		Detail:      lunar.Detail(),
	}
}

func writeErrorResponse(w http.ResponseWriter, error, message string, statusCode int) {
	w.WriteHeader(statusCode)
	response := ErrorResponse{
		Error:   error,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}
