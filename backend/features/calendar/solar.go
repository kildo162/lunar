package calendar

import (
	"fmt"
	"time"
)

var TUAN = []string{"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"}

type SolarDate struct {
	Year     int
	Month    int
	Day      int
	LeapYear bool
}

func NewSolarDate(year, month, day int) *SolarDate {
	s := SolarDate{
		Year:     year,
		Month:    month,
		Day:      day,
		LeapYear: false,
	}
	s.LeapYear = s.IsLeapYear()
	return &s
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

func (sd *SolarDate) Detail() string {
	weekday := TUAN[sd.GetWeekday()]
	return fmt.Sprintf("Dương Lịch: %s, %02d/%02d/%d", weekday, sd.Day, sd.Month, sd.Year)
}
