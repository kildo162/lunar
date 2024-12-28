package calendar

// ...other imports if needed...

type LunarDate struct {
	Year  int
	Month int
	Day   int
}

func NewLunarDate(year, month, day int) *LunarDate {
	return &LunarDate{
		Year:  year,
		Month: month,
		Day:   day,
	}
}
