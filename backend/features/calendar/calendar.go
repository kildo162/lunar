package calendar

type CalendarDate struct {
	Day      int
	Month    int
	Year     int
	Hour     int
	Min      int
	Second   int
	TimeZone float64
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
