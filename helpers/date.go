package helpers

import (
	"fmt"
	"time"
)

func Str2Date(data string) (ret time.Time) {

	loc, _ := time.LoadLocation("America/Sao_Paulo")

	ret, err := time.ParseInLocation("20060102150405", data, loc)

	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func Hoje() (ret time.Time) {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	hoje := Date2_html(time.Now())

	ret, err := time.ParseInLocation("2006-01-02 15:04", hoje+" 00:00", loc)

	if err != nil {
		fmt.Println(err.Error())
	}

	return
}

func Date2_html(t time.Time) (ret string) {
	if t.IsZero() {
		ret = ""
	} else {
		ret = fmt.Sprintf("%04d-%02d-%02d",
			t.Year(),
			t.Month(),
			t.Day())
	}
	return ret
}

func main() {
	start := time.Date(1990, 2, 1, 3, 30, 0, 0, time.UTC)

	fmt.Println(start)
	fmt.Println(time.Now())

	// calculate years, month, days and time betwen dates
	Ano, Mes, Dia, Hora, Min, Sec := Diff(start, time.Now())

	fmt.Printf("difference %d years, %d months, %d days, %d hours, %d mins and %d seconds.", Ano, Mes, Dia, Hora, Min, Sec)
	fmt.Printf("")

	// calculate total number of days
	duration := time.Now().Sub(start)
	fmt.Printf("difference %d days", int(duration.Hours()/24))
}

func Diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
