package calendar

import (
	"time"
)

type TodayInfo struct {
	SolarDate string
	LunarDate string
	CanChi    string
	GoodHours []string
	BadHours  []string
}

func GetToday() TodayInfo {
	now := time.Now()
	year, month, day := now.Year(), int(now.Month()), now.Day()
	// Tạo SolarDate
	solar := NewSolarDate(year, month, day)
	// Tạo LunarDate
	lunar := NewLunarDate(year, month, day, now.Hour(), now.Minute(), now.Second(), 7.0)
	// Can Chi
	canChi := lunar.GetCanChiYear() + " " + lunar.GetCanChiMonth() + " " + lunar.GetCanDay()
	// Giờ hoàng đạo/hắc đạo (giả lập, cần sửa lại nếu có logic)
	goodHours := []string{"Tý", "Sửu", "Thìn", "Tỵ", "Mùi", "Tuất"}
	badHours := []string{"Dần", "Mão", "Ngọ", "Thân", "Dậu", "Hợi"}
	return TodayInfo{
		SolarDate: solar.Detail(),
		LunarDate: lunar.Detail(),
		CanChi:    canChi,
		GoodHours: goodHours,
		BadHours:  badHours,
	}
}


type Calendar struct {
	SoalrDate *SolarDate
	LunarDate *LunarDate
}

func (c *Calendar) ToSolar() *SolarDate {
	return c.SoalrDate
}

func (c *Calendar) ToLunar() *LunarDate {
	return c.LunarDate
}

func NewCalendar(date CalendarDate) *Calendar {
	c := Calendar{
		SoalrDate: NewSolarDate(date.Year, date.Month, date.Day),
		LunarDate: NewLunarDate(date.Year, date.Month, date.Day, date.Hour, date.Min, date.Second, date.TimeZone),
	}
	return &c
}
