package models

// Định nghĩa các struct dùng chung cho calendar

type ConvertToLunarRequest struct {
	Day      int     `json:"day"`
	Month    int     `json:"month"`
	Year     int     `json:"year"`
	Hour     int     `json:"hour,omitempty"`
	Minute   int     `json:"minute,omitempty"`
	Second   int     `json:"second,omitempty"`
	TimeZone float64 `json:"timezone,omitempty"`
}

type ConvertToSolarRequest struct {
	LunarDay   int     `json:"lunar_day"`
	LunarMonth int     `json:"lunar_month"`
	LunarYear  int     `json:"lunar_year"`
	LeapMonth  int     `json:"leap_month,omitempty"`
	TimeZone   float64 `json:"timezone,omitempty"`
}

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

type CalendarResponse struct {
	Solar SolarResponse `json:"solar"`
	Lunar LunarResponse `json:"lunar"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
