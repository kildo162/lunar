package calendar

import (
	"fmt"
)

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

func (ld *LunarDate) FormatDetailed() string {
	yearCanChi := fmt.Sprintf("%s %s", CAN[ld.Year%10], CHI[ld.Year%12])
	monthCanChi := fmt.Sprintf("%s %s", CAN[(ld.Year*12+ld.Month+3)%10], CHI[(ld.Month+1)%12])
	dayCanChi := fmt.Sprintf("%s %s", CAN[(ld.Year*360+ld.Month*30+ld.Day+6)%10], CHI[(ld.Day+1)%12])
	return fmt.Sprintf("Ngày %02d Tháng %02d - Ngày %s, Tháng %s, Năm %s", ld.Day, ld.Month, dayCanChi, monthCanChi, yearCanChi)
}
