package calendar

import (
	"fmt"
)

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

func (ld *LunarDate) Format() string {
	return fmt.Sprintf("%02d-%02d-%d", ld.Day, ld.Month, ld.Year)
}

func (ld *LunarDate) YearInfo() string {
	return fmt.Sprintf("Year: %d", ld.Year)
}
