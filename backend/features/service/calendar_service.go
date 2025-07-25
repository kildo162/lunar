package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/features/calendar"
	"backend/features/models"
)

// Chuyển đổi dương lịch sang âm lịch
func ConvertToLunarService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req models.ConvertToLunarRequest
	if r.Method == "GET" {
		// Parse query params
		q := r.URL.Query()
		req.Day = parseInt(q.Get("day"))
		req.Month = parseInt(q.Get("month"))
		req.Year = parseInt(q.Get("year"))
		req.Hour = parseInt(q.Get("hour"))
		req.Minute = parseInt(q.Get("minute"))
		req.Second = parseInt(q.Get("second"))
		req.TimeZone = parseFloat(q.Get("timezone"), 7.0)
	} else {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, "Invalid JSON", err.Error())
			return
		}
	}

	// Validate
	if req.Day < 1 || req.Day > 31 || req.Month < 1 || req.Month > 12 || req.Year < 1 {
		writeError(w, "Invalid date", "Please provide valid day (1-31), month (1-12), and year")
		return
	}

	// Chuyển đổi
	cal := calendar.NewCalendar(calendar.CalendarDate{
		Day:      req.Day,
		Month:    req.Month,
		Year:     req.Year,
		Hour:     req.Hour,
		Min:      req.Minute,
		Second:   req.Second,
		TimeZone: req.TimeZone,
	})

	resp := models.CalendarResponse{
		Solar: buildSolarResponse(cal.SoalrDate),
		Lunar: buildLunarResponse(cal.LunarDate),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func parseFloat(s string, def float64) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return def
	}
	return v
}

func writeError(w http.ResponseWriter, err, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(models.ErrorResponse{Error: err, Message: msg})
}

func buildSolarResponse(solar *calendar.SolarDate) models.SolarResponse {
	return models.SolarResponse{
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

func buildLunarResponse(lunar *calendar.LunarDate) models.LunarResponse {
	monthName := ""
	if lunar.LunarMonth > 0 && lunar.LunarMonth <= 12 {
		monthName = calendar.THANG[lunar.LunarMonth-1]
	}
	return models.LunarResponse{
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

// Chuyển đổi âm lịch sang dương lịch
func ConvertToSolarService(w http.ResponseWriter, r *http.Request) {
	// ... business logic ...
	w.Write([]byte("ConvertToSolarService called"))
}

// Lấy thông tin ngày hôm nay
func GetTodayService(w http.ResponseWriter, r *http.Request) {
	// Lấy ngày hiện tại
	today := calendar.GetToday()
	resp := struct {
		SolarDate string   `json:"solarDate"`
		LunarDate string   `json:"lunarDate"`
		CanChi    string   `json:"canChi"`
		GoodHours []string `json:"goodHours"`
		BadHours  []string `json:"badHours"`
	}{
		SolarDate: today.SolarDate,
		LunarDate: today.LunarDate,
		CanChi:    today.CanChi,
		GoodHours: today.GoodHours,
		BadHours:  today.BadHours,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Lấy lịch tháng
func GetMonthCalendarService(w http.ResponseWriter, r *http.Request) {
	// ... business logic ...
	w.Write([]byte("GetMonthCalendarService called"))
}

// Tìm ngày tốt/xấu
func SearchGoodDaysService(w http.ResponseWriter, r *http.Request) {
	// ... business logic ...
	w.Write([]byte("SearchGoodDaysService called"))
}
