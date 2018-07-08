package desamber

import (
	"fmt"
	"strings"
	"time"
)

// Date is the parsed date
type Date struct {
	Year      int
	Century   int
	Month     rune
	Day       int
	useCentry bool
	useYear   bool
}

// New takes in a time.Time and converts it to desamber.Date object
func New(date time.Time) *Date {
	c := date.Year() / 100
	y := date.Year() % 100
	doty := date.YearDay()
	m := 'A'
	if doty == 365 || doty == 366 {
		m = '+'
	} else {
		mval := (float64(doty-1) / float64(364)) * 26
		m = rune(65 + mval)
	}
	day := doty % 14
	if day == 0 {
		day = 14
	}
	d := Date{
		Century: c,
		Year:    y,
		Month:   m,
		Day:     day,
	}
	return &d
}

// WithCentury will enable printing of the dates century, will automatically enable year as well.
func (d *Date) WithCentury() *Date {
	d.useCentry = true
	d.useYear = true
	return d
}

// WithYear will enable printing of the year (not the century)
func (d *Date) WithYear() *Date {
	d.useYear = true
	return d
}

func (d Date) String() string {
	var b strings.Builder
	if d.useYear {
		if d.useCentry {
			b.WriteString(fmt.Sprintf("%02v", d.Century))
		}
		b.WriteString(fmt.Sprintf("%02v", d.Year))
	}
	b.WriteString(string(d.Month))
	b.WriteString(fmt.Sprintf("%02v", d.Day))

	return fmt.Sprint(b.String())
}
