package calendar

import (
	"math"
)

type CalendarDate struct {
	Day   int
	Month int
	Year  int
}

const PI = math.Pi

func INT(d float64) int {
	return int(math.Floor(d))
}

type Calendar struct {
	Day      int
	Month    int
	Year     int
	JD       *float64 // Julian Date, used for astronomical calculations
	LeapYear *bool    // Indicates if the year is a leap year

	SoalrDate *SolarDate
	LunarDate *LunarDate
}

func NewCalendar(date CalendarDate) *Calendar {
	c := &Calendar{
		Day:   date.Day,
		Month: date.Month,
		Year:  date.Year,
	}
	jd := c.calculateJD()
	c.JD = &jd
	leapYear := c.isLeapYear()
	c.LeapYear = &leapYear

	c.SoalrDate = NewSolarDate(c.Year, c.Month, c.Day)
	c.LunarDate = NewLunarDate(c.Year, c.Month, c.Day)
	return c
}

func (c *Calendar) SetDate(date CalendarDate) {
	c.Day = date.Day
	c.Month = date.Month
	c.Year = date.Year
	jd := c.calculateJD()
	c.JD = &jd
	leapYear := c.isLeapYear()
	c.LeapYear = &leapYear
}

func (c *Calendar) isLeapYear() bool {
	year := c.Year
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func (c *Calendar) calculateJD() float64 {
	year := c.Year
	month := c.Month
	day := c.Day
	if month <= 2 {
		year--
		month += 12
	}
	A := INT(float64(year) / 100)
	B := 2 - A + INT(float64(A)/4)
	return float64(INT(365.25*float64(year))) + float64(INT(30.6001*float64(month+1))) + float64(day) + 1720994.5 + float64(B)
}

func (c *Calendar) ToSolar() *SolarDate {
	return c.SoalrDate
}

func (c *Calendar) ToLunar() *LunarDate {
	return c.LunarDate
}
