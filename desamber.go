package desamber

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Date is the parsed date
type Date struct {
	Year      string
	Century   string
	Month     rune
	Day       string
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
		Century: fmt.Sprintf("%02v", c),
		Year:    fmt.Sprintf("%02v", y),
		Month:   m,
		Day:     fmt.Sprintf("%02v", day),
	}
	return &d
}

var (
	reg = regexp.MustCompile(`^((\d{2})?(\d{2}))?([A-Z\+])?(\d{2})?[^\s]*$`)
	// match = regexp.MustCompile(`^(\d\d(\d\d)?)?[A-Z\+](\d\d)|^[A-Z\+](\d\d)|^[A-Z\+]$|^(\d\d)$`)
	match = regexp.MustCompile(`^(\d\d(\d\d)?)?[A-Z\+](\d\d)?$|^[A-Z\+](\d\d)|^[A-Z\+]$|^(\d\d)$`)
)

// Parse will accept a string representation of a desamber date
func Parse(s string) (*Date, error) {
	if !match.MatchString(s) {
		return nil, errors.New("String did not match format")
	}
	ss := reg.FindStringSubmatch(strings.ToUpper(s))
	var month rune
	var year, century string
	if len(ss[4]) != 0 {
		month = rune(ss[4][0])
	}

	if len(ss[1]) == 4 {
		fullYear, err := strconv.ParseInt(ss[1], 10, 0)
		if err != nil {
			return nil, err
		}
		if isLeap(fullYear) && month == '+' && ss[5] != "01" && ss[5] != "02" {
			return nil, errors.New("Invalid day for leap year")
		}
		year = ss[3]
		century = ss[2]
	} else {
		year = ss[3]
	}

	result := Date{
		Day:     ss[5],
		Month:   month,
		Year:    year,
		Century: century,
	}

	return &result, nil
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
	if d.useYear && len(d.Year) != 0 {
		if d.useCentry && len(d.Century) != 0 {
			b.WriteString(d.Century)
		}
		b.WriteString(d.Year)
	}
	if d.Month != 0 {
		b.WriteString(string(d.Month))
	}
	if len(d.Day) != 0 {
		b.WriteString(d.Day)
	}
	return fmt.Sprint(b.String())
}

func isLeap(year int64) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}
