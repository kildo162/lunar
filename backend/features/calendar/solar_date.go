package calendar

import (
	"time"
)

type SolarDate struct {
	Year  int
	Month int
	Day   int
}

func NewSolarDate(year, month, day int) *SolarDate {
	return &SolarDate{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func (sd *SolarDate) ToLunar() *LunarDate {
	// Conversion logic from solar to lunar date
	// ...existing code...
	return &LunarDate{} // Placeholder
}

func FromLunar(lunarDate *LunarDate) *SolarDate {
	// Conversion logic from lunar to solar date
	// ...existing code...
	return &SolarDate{} // Placeholder
}

func (sd *SolarDate) IsLeapYear() bool {
	year := sd.Year
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func (sd *SolarDate) GetDayOfYear() int {
	startOfYear := time.Date(sd.Year, 1, 1, 0, 0, 0, 0, time.UTC)
	currentDate := time.Date(sd.Year, time.Month(sd.Month), sd.Day, 0, 0, 0, 0, time.UTC)
	return int(currentDate.Sub(startOfYear).Hours() / 24)
}

func (sd *SolarDate) GetWeekOfYear() int {
	currentDate := time.Date(sd.Year, time.Month(sd.Month), sd.Day, 0, 0, 0, 0, time.UTC)
	_, week := currentDate.ISOWeek()
	return week
}

func (sd *SolarDate) GetWeekday() time.Weekday {
	currentDate := time.Date(sd.Year, time.Month(sd.Month), sd.Day, 0, 0, 0, 0, time.UTC)
	return currentDate.Weekday()
}
