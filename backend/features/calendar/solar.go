package calendar

import (
	"fmt"
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

func (sd *SolarDate) Format() string {
	return fmt.Sprintf("%02d-%02d-%d", sd.Day, sd.Month, sd.Year)
}

func (sd *SolarDate) Detail() string {
	return fmt.Sprintf(
		"Year: %d, Month: %d, Day: %d, Week of Year: %d, Day of Year: %d",
		sd.Year, sd.Month, sd.Day, sd.GetWeekOfYear(), sd.GetDayOfYear(),
	)
}

func (sd *SolarDate) YearInfo() string {
	return fmt.Sprintf("Year: %d, Leap Year: %t", sd.Year, sd.IsLeapYear())
}
