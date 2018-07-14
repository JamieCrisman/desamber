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
	Year    int
	yearSet bool

	Century    int
	centurySet bool

	Day    int
	daySet bool

	Month rune

	useCentry bool
	useYear   bool
}

// New takes in a time.Time and converts it to desamber.Date object
func New(date time.Time) *Date {

	doty := date.YearDay()
	m := 'A'
	if doty == 365 || doty == 366 {
		m = '+'
	} else {
		mval := (float64(doty-1) / float64(364)) * 26
		m = rune(65 + mval)
	}
	day := doty % 14
	// we do range of 1 to 14
	if day == 0 {
		day = 14
	}

	d := Date{
		Century:    date.Year() / 100,
		centurySet: true,
		Year:       date.Year() % 100,
		yearSet:    true,
		Month:      m,
		Day:        day,
		daySet:     true,
	}
	return &d
}

var (
	// Splits the string into the parts
	reg = regexp.MustCompile(`^((\d{2})?(\d{2}))?([A-Z\+])?(\d{2})?[^\s]*$`)
	// ensures a valid string
	match = regexp.MustCompile(`^(\d\d(\d\d)?)?[A-Z\+](\d\d)?$|^[A-Z\+](\d\d)|^[A-Z\+]$|^(\d\d)$`)
)

// Parse will accept a string representation of a desamber date
func Parse(s string) (*Date, error) {
	if !match.MatchString(s) {
		return nil, errors.New("String did not match format")
	}
	ss := reg.FindStringSubmatch(strings.ToUpper(s))
	var month rune
	if len(ss[4]) != 0 {
		month = rune(ss[4][0])
	}
	var parsedYear, parsedCentury, parsedDay int64
	var centurySet, yearSet, daySet bool
	var err error
	if len(ss[1]) == 4 {
		fullYear, err := strconv.ParseInt(ss[1], 10, 0)
		if err != nil {
			return nil, err
		}
		if isLeap(fullYear) && month == '+' && ss[5] != "01" && ss[5] != "02" {
			return nil, errors.New("Invalid day for leap year")
		}

		parsedCentury, err = strconv.ParseInt(ss[2], 10, 0)
		if err != nil {
			return nil, err
		}
	}
	if len(ss[3]) != 0 {
		parsedYear, err = strconv.ParseInt(ss[3], 10, 0)
		if err != nil {
			return nil, err
		}
	}
	if len(ss[5]) != 0 {
		parsedDay, err = strconv.ParseInt(ss[5], 10, 0)
		if err != nil {
			return nil, err
		}
	}

	centurySet = len(ss[2]) != 0
	yearSet = len(ss[3]) != 0
	daySet = len(ss[5]) != 0
	result := Date{
		Day:        int(parsedDay),
		daySet:     daySet,
		Month:      month,
		Year:       int(parsedYear),
		yearSet:    yearSet,
		Century:    int(parsedCentury),
		centurySet: centurySet,
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
	if d.useYear && d.yearSet {
		if d.useCentry && d.centurySet {
			b.WriteString(fmt.Sprintf("%02d", d.Century))
		}
		b.WriteString(fmt.Sprintf("%02d", d.Year))
	}
	if d.Month != 0 {
		b.WriteString(string(d.Month))
	}
	if d.daySet {
		b.WriteString(fmt.Sprintf("%02d", d.Day))
	}
	return fmt.Sprint(b.String())
}

func isLeap(year int64) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}
