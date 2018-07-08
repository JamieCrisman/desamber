package desamber_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jamiecrisman/desamber"
)

func ExampleNew() {
	date := desamber.New(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC))
	fmt.Println(date)
	// output: A01
}

func ExampleDate_WithCentury() {
	date := desamber.New(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC))
	date = date.WithCentury()
	fmt.Println(date)
	// output: 2018A01
}

func ExampleDate_WithYear() {
	date := desamber.New(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC))
	date = date.WithYear()
	fmt.Println(date)
	// output: 18A01
}

func TestFromDate(t *testing.T) {
	tt := []struct {
		Name     string
		Input    time.Time
		Expected desamber.Date
	}{
		{
			Name:  "first month start",
			Input: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    18,
				Month:   'A',
				Day:     1,
			},
		},
		{
			Name:  "first month end",
			Input: time.Date(2018, 1, 14, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    18,
				Month:   'A',
				Day:     14,
			},
		},
		{
			Name:  "second month start",
			Input: time.Date(2018, 1, 15, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    18,
				Month:   'B',
				Day:     1,
			},
		},
		{
			Name:  "second month end",
			Input: time.Date(2018, 1, 28, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    18,
				Month:   'B',
				Day:     14,
			},
		},
		{
			Name:  "third month start",
			Input: time.Date(2018, 1, 29, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    18,
				Month:   'C',
				Day:     1,
			},
		},
		{
			Name:  "last month start",
			Input: time.Date(2018, 12, 17, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    18,
				Month:   'Z',
				Day:     1,
			},
		},
		{
			Name:  "last month end",
			Input: time.Date(2018, 12, 30, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    18,
				Month:   'Z',
				Day:     14,
			},
		},
		{
			Name:  "century",
			Input: time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 19,
				Year:    99,
				Month:   '+',
				Day:     1,
			},
		},
		{
			Name:  "year",
			Input: time.Unix(1551000000, 0),
			Expected: desamber.Date{
				Century: 20,
				Year:    19,
				Month:   'D',
				Day:     13,
			},
		},
		{
			Name:  "month overflow year bug",
			Input: time.Unix(1483000000, 0),
			Expected: desamber.Date{
				Century: 20,
				Year:    16,
				Month:   'Z',
				Day:     14,
			},
		},
		{
			Name:  "year day",
			Input: time.Date(2016, 12, 31, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    16,
				Month:   '+',
				Day:     2,
			},
		},
		{
			Name:  "year day leap day",
			Input: time.Date(2016, 12, 30, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    16,
				Month:   '+',
				Day:     1,
			},
		},
		{
			Name:  "year day",
			Input: time.Date(2018, 12, 31, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    18,
				Month:   '+',
				Day:     1,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := desamber.New(tc.Input)
			if result.Century != tc.Expected.Century {
				t.Errorf("Expected century to be %d, but got %d", tc.Expected.Century, result.Century)
			}
			if result.Year != tc.Expected.Year {
				t.Errorf("Expected year to be %d, but got %d", tc.Expected.Year, result.Year)
			}
			if result.Month != tc.Expected.Month {
				t.Errorf("Expected month to be %s, but got %s", string(tc.Expected.Month), string(result.Month))
			}
			if result.Day != tc.Expected.Day {
				t.Errorf("Expected day to be %d, but got %d", tc.Expected.Day, result.Day)
			}
		})
	}
}
