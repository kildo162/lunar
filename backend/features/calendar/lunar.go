package calendar

import (
	"fmt"
	"math"
)

var THANG = []string{"Giêng", "Hai", "Ba", "Tư", "Năm", "Sáu", "Bảy", "Tám", "Chín", "Mười", "Một", "Chạp"}
var CAN = []string{"Giáp", "Ất", "Bính", "Đinh", "Mậu", "Kỷ", "Canh", "Tân", "Nhâm", "Quý"}
var CHI = []string{"Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ", "Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi"}
var DAY = []string{"Chủ nhật", "Thứ hai", "Thứ ba", "Thứ tư", "Thứ năm", "Thứ sáu", "Thứ bảy"}
var GIO_HD = []string{"110100101100", "001101001011", "110011010010", "101100110100", "001011001101", "010010110011"}
var TIETKHI = []string{
	"Xuân phân", "Thanh minh", "Cốc vũ", "Lập hạ", "Tiểu mãn", "Mang chủng", "Hạ chí", "Tiểu thử", "Đại thử",
	"Lập thu", "Xử thử", "Bạch lộ", "Thu phân", "Hàn lộ", "Sương giáng", "Lập đông", "Tiểu tuyết", "Đại tuyết",
	"Đông chí", "Tiểu hàn", "Đại hàn", "Lập xuân", "Vũ Thủy", "Kinh trập"}

type LunarDate struct {
	LunarYear   int
	LunarMonth  int
	LunarDay    int
	LunarHour   int
	LunarMin    int
	LunarSecond int
	Timezone    float32
	LeapMonth   int
	JD          int
}

func NewLunarDate(solarYear, solarMonth, solarDay, solarHour, solarMin, solarSecond int, timeZone float64) *LunarDate {
	l := LunarDate{
		JD:          JDFromDate(solarDay, solarMonth, solarYear),
		LunarHour:   solarHour,
		LunarMin:    solarMin,
		LunarSecond: solarSecond,
	}
	l.LunarDay, l.LunarMonth, l.LunarYear, l.LeapMonth = convertSolar2Lunar(solarDay, solarMonth, solarYear, timeZone)
	return &l
}

func INT(d float64) int {
	return int(math.Floor(d))
}

func MOD(x, y int) int {
	z := x - int(float64(y)*math.Floor(float64(x)/float64(y)))
	if z == 0 {
		z = y
	}
	return z
}

// JDFromDate computes the Julian day number of a given date (dd/mm/yyyy).
// It returns the number of days between 1/1/4713 BC (Julian calendar) and the given date.
func JDFromDate(dd, mm, yy int) int {
	a := INT(float64(14-mm) / 12)
	y := yy + 4800 - a
	m := mm + 12*a - 3
	jd := dd + INT(float64(153*m+2)/5) + 365*y + INT(float64(y)/4) - INT(float64(y)/100) + INT(float64(y)/400) - 32045
	if jd < 2299161 {
		jd = dd + INT(float64(153*m+2)/5) + 365*y + INT(float64(y)/4) - 32083
	}
	return jd
}

// JDToDate converts a Julian day number to a date (day/month/year).
// It returns the day, month, and year corresponding to the given Julian day number.
func JDToDate(jd int) (int, int, int) {
	var a, b, c, d, e, m int
	if jd > 2299160 {
		a = jd + 32044
		b = INT(float64(4*a+3) / 146097)
		c = a - INT(float64(b*146097)/4)
	} else {
		b = 0
		c = jd + 32082
	}
	d = INT(float64(4*c+3) / 1461)
	e = c - INT(float64(1461*d)/4)
	m = INT(float64(5*e+2) / 153)
	day := e - INT(float64(153*m+2)/5) + 1
	month := m + 3 - 12*INT(float64(m)/10)
	year := b*100 + d - 4800 + INT(float64(m)/10)
	return day, month, year
}

/* Compute the time of the k-th new moon after the new moon of 1/1/1900 13:52 UCT
 * (measured as the number of days since 1/1/4713 BC noon UCT, e.g., 2451545.125 is 1/1/2000 15:00 UTC).
 * Returns a floating number, e.g., 2415079.9758617813 for k=2 or 2414961.935157746 for k=-2
 * Algorithm from: "Astronomical Algorithms" by Jean Meeus, 1998
 */
func NewMoon(k int) float64 {
	T := float64(k) / 1236.85
	T2 := T * T
	T3 := T2 * T
	dr := math.Pi / 180
	Jd1 := 2415020.75933 + 29.53058868*float64(k) + 0.0001178*T2 - 0.000000155*T3
	Jd1 += 0.00033 * math.Sin((166.56+132.87*T-0.009173*T2)*dr)
	M := 359.2242 + 29.10535608*float64(k) - 0.0000333*T2 - 0.00000347*T3
	Mpr := 306.0253 + 385.81691806*float64(k) + 0.0107306*T2 + 0.00001236*T3
	F := 21.2964 + 390.67050646*float64(k) - 0.0016528*T2 - 0.00000239*T3
	C1 := (0.1734-0.000393*T)*math.Sin(M*dr) + 0.0021*math.Sin(2*dr*M)
	C1 -= 0.4068 * math.Sin(Mpr*dr)
	C1 += 0.0161 * math.Sin(dr*2*Mpr)
	C1 -= 0.0004 * math.Sin(dr*3*Mpr)
	C1 += 0.0104 * math.Sin(dr*2*F)
	C1 -= 0.0051 * math.Sin(dr*(M+Mpr))
	C1 -= 0.0074 * math.Sin(dr*(M-Mpr))
	C1 += 0.0004 * math.Sin(dr*(2*F+M))
	C1 -= 0.0004 * math.Sin(dr*(2*F-M))
	C1 -= 0.0006 * math.Sin(dr*(2*F+Mpr))
	C1 += 0.0010 * math.Sin(dr*(2*F-Mpr))
	C1 += 0.0005 * math.Sin(dr*(2*Mpr+M))
	var deltat float64
	if T < -11 {
		deltat = 0.001 + 0.000839*T + 0.0002261*T2 - 0.00000845*T3 - 0.000000081*T*T3
	} else {
		deltat = -0.000278 + 0.000265*T + 0.000262*T2
	}
	JdNew := Jd1 + C1 - deltat
	return JdNew
}

/* Compute the longitude of the sun at any time.
 * Parameter: floating number jdn, the number of days since 1/1/4713 BC noon
 * Algorithm from: "Astronomical Algorithms" by Jean Meeus, 1998
 */
func SunLongitude(jdn float64) float64 {
	T := (jdn - 2451545.0) / 36525
	T2 := T * T
	dr := math.Pi / 180
	M := 357.52910 + 35999.05030*T - 0.0001559*T2 - 0.00000048*T*T2
	L0 := 280.46645 + 36000.76983*T + 0.0003032*T2
	DL := (1.914600-0.004817*T-0.000014*T2)*math.Sin(dr*M) +
		(0.019993-0.000101*T)*math.Sin(dr*2*M) +
		0.000290*math.Sin(dr*3*M)
	L := L0 + DL
	L = L * dr
	L = L - math.Pi*2*float64(INT(L/(math.Pi*2)))
	return L
}

/* Compute sun position at midnight of the day with the given Julian day number.
 * The time zone is the time difference between local time and UTC: 7.0 for UTC+7:00.
 * Returns a number between 0 and 11.
 * From the day after March equinox and the 1st major term after March equinox, 0 is returned.
 * After that, return 1, 2, 3 ...
 */
func getSunLongitude(dayNumber float64, timeZone float64) int {
	return INT(SunLongitude(dayNumber-0.5-timeZone/24) / math.Pi * 6)
}

/* Compute the day of the k-th new moon in the given time zone.
 * The time zone is the time difference between local time and UTC: 7.0 for UTC+7:00
 */
func getNewMoonDay(k int, timeZone float64) int {
	return INT(NewMoon(k) + 0.5 + timeZone/24)
}

/* Find the day that starts the lunar month 11 of the given year for the given time zone */
func getLunarMonth11(yy int, timeZone float64) int {
	off := JDFromDate(31, 12, yy) - 2415021
	k := INT(float64(off) / 29.530588853)
	nm := getNewMoonDay(k, timeZone)
	sunLong := getSunLongitude(float64(nm), timeZone)
	if sunLong >= 9 {
		nm = getNewMoonDay(k-1, timeZone)
	}
	return nm
}

/* Find the index of the leap month after the month starting on the day a11. */
func getLeapMonthOffset(a11 int, timeZone float64) int {
	k := INT((float64(a11) - 2415021.076998695) / 29.530588853)
	last := 0
	i := 1
	arc := getSunLongitude(float64(getNewMoonDay(k+i, timeZone)), timeZone)
	for arc != last && i < 14 {
		last = arc
		i++
		arc = getSunLongitude(float64(getNewMoonDay(k+i, timeZone)), timeZone)
	}
	return i - 1
}

/* Convert solar date dd/mm/yyyy to the corresponding lunar date */
func convertSolar2Lunar(dd, mm, yy int, timeZone float64) (int, int, int, int) {
	dayNumber := JDFromDate(dd, mm, yy)
	k := INT((float64(dayNumber) - 2415021.076998695) / 29.530588853)
	monthStart := getNewMoonDay(k+1, timeZone)
	if monthStart > dayNumber {
		monthStart = getNewMoonDay(k, timeZone)
	}
	a11 := getLunarMonth11(yy, timeZone)
	b11 := a11
	var lunarYear int
	if a11 >= monthStart {
		lunarYear = yy
		a11 = getLunarMonth11(yy-1, timeZone)
	} else {
		lunarYear = yy + 1
		b11 = getLunarMonth11(yy+1, timeZone)
	}
	lunarDay := dayNumber - monthStart + 1
	diff := INT(float64(monthStart-a11) / 29)
	lunarLeap := 0
	lunarMonth := diff + 11
	if b11-a11 > 365 {
		leapMonthDiff := getLeapMonthOffset(a11, timeZone)
		if diff >= leapMonthDiff {
			lunarMonth = diff + 10
			if diff == leapMonthDiff {
				lunarLeap = 1
			}
		}
	}
	if lunarMonth > 12 {
		lunarMonth -= 12
	}
	if lunarMonth >= 11 && diff < 4 {
		lunarYear -= 1
	}
	return lunarDay, lunarMonth, lunarYear, lunarLeap
}

/* Convert a lunar date to the corresponding solar date */
func convertLunar2Solar(lunarDay, lunarMonth, lunarYear, lunarLeap int, timeZone float64) (int, int, int) {
	var k, a11, b11 int
	if lunarMonth < 11 {
		a11 = getLunarMonth11(lunarYear-1, timeZone)
		b11 = getLunarMonth11(lunarYear, timeZone)
	} else {
		a11 = getLunarMonth11(lunarYear, timeZone)
		b11 = getLunarMonth11(lunarYear+1, timeZone)
	}
	k = INT(0.5 + (float64(a11)-2415021.076998695)/29.530588853)
	off := lunarMonth - 11
	if off < 0 {
		off += 12
	}
	if b11-a11 > 365 {
		leapOff := getLeapMonthOffset(a11, timeZone)
		leapMonth := leapOff - 2
		if leapMonth < 0 {
			leapMonth += 12
		}
		if lunarLeap != 0 && lunarMonth != leapMonth {
			return 0, 0, 0
		} else if lunarLeap != 0 || off >= leapOff {
			off += 1
		}
	}
	monthStart := getNewMoonDay(k+off, timeZone)
	return JDToDate(monthStart + lunarDay - 1)
}

func (l *LunarDate) GetCanChiYear() string {
	return CAN[(l.LunarYear+6)%10] + " " + CHI[(l.LunarYear+8)%12]
}

func (l *LunarDate) GetCanChiMonth() string {
	return CAN[(l.LunarYear*12+l.LunarMonth+3)%10] + " " + CHI[(l.LunarMonth+1)%12]
}

func (l *LunarDate) GetCanDay() string {
	dayName := CAN[(l.JD+9)%10] + " " + CHI[(l.JD+1)%12]
	return dayName
}

func (l *LunarDate) GetCanHour() string {
	chiGio := 0
	if l.LunarHour >= 23 {
		chiGio = 0
	} else {
		chiGio = INT(float64(l.LunarHour+1) / 2)
	}
	canGio := ((l.JD-1)*2 + chiGio) % 10
	return CAN[canGio] + " " + CHI[chiGio]
}

func (l *LunarDate) GetGioHoangDao() string {
	chiOfDay := (l.JD + 1) % 12
	gioHD := GIO_HD[chiOfDay%6] // same values for Ty' (1) and Ngo. (6), for Suu and Mui etc.
	ret := ""
	count := 0
	for i := 0; i < 12; i++ {
		if gioHD[i:i+1] == "1" {
			ret += CHI[i]
			ret += fmt.Sprintf(" (%02d-%02d)", (i*2+23)%24, (i*2+1)%24)
			if count++; count <= 5 {
				ret += ", "
			}
		}
	}
	return ret
}

func (l *LunarDate) GetTietKhi() string {
	return TIETKHI[getSunLongitude(float64(l.JD+1), 7.0)]
}

func (l *LunarDate) GetBeginHour() string {
	return CAN[(l.JD-1)*2%10] + " " + CHI[0]
}

func (l *LunarDate) Detail() string {
	return fmt.Sprintf(
		"Âm Lịch: %02d:%02d - Giờ %s, Ngày %s, Tháng %s, Năm %s - Giờ tốt: %s - Tiết Khí: %s",
		l.LunarHour, l.LunarMin,
		l.GetCanHour(),
		l.GetCanDay(),
		l.GetCanChiMonth(),
		l.GetCanChiYear(),
		l.GetGioHoangDao(),
		l.GetTietKhi(),
	)
}
