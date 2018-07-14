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

func ExampleParse() {
	date, err := desamber.Parse("2018N07")
	if err != nil {
		panic(err)
	}
	fmt.Println(date.WithCentury())
	// output: 2018N07
}

func TestParse(t *testing.T) {
	tt := []struct {
		Name           string
		Input          string
		ExpectedString string
		ExpectedError  string
	}{
		{
			Name:           "January 1st, no year",
			Input:          "A01",
			ExpectedString: "A01",
			ExpectedError:  "",
		},
		{
			Name:           "First month",
			Input:          "A",
			ExpectedString: "A",
			ExpectedError:  "",
		},
		{
			Name:           "14th day",
			Input:          "14",
			ExpectedString: "14",
			ExpectedError:  "",
		},
		{
			Name:           "Year 18 of month N",
			Input:          "18N",
			ExpectedString: "18N",
			ExpectedError:  "",
		},
		{
			Name:           "Invalid + Date",
			Input:          "2016+03",
			ExpectedString: "",
			ExpectedError:  "Invalid day for leap year",
		},
		{
			Name:           "Invalid Date",
			Input:          "ASDFEFW2321",
			ExpectedString: "",
			ExpectedError:  "String did not match format",
		},
		{
			Name:           "Invalid Date bad year 3 digits",
			Input:          "018A01",
			ExpectedString: "",
			ExpectedError:  "String did not match format",
		},
		{
			Name:           "Invalid Date bad year 1 digit",
			Input:          "8A01",
			ExpectedString: "",
			ExpectedError:  "String did not match format",
		},
		{
			Name:           "Full date",
			Input:          "2018A01",
			ExpectedString: "2018A01",
			ExpectedError:  "",
		},
		{
			Name:           "Year Date",
			Input:          "18A01",
			ExpectedString: "18A01",
			ExpectedError:  "",
		},
		{
			Name:           "Year Day",
			Input:          "18+01",
			ExpectedString: "18+01",
			ExpectedError:  "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := desamber.Parse(tc.Input)
			if result != nil {
				result = result.WithCentury()
			}
			if err == nil && result.String() != tc.ExpectedString {
				t.Errorf("got %s %q, expected %s", result.String(), result.String(), tc.ExpectedString)
			}
			if (len(tc.ExpectedError) != 0 || err != nil) && err.Error() != tc.ExpectedError {
				t.Errorf("got unexpected error: %s \ngot instead: %s", err, tc.ExpectedError)
			}
		})
	}

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
				Day:     01,
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
				Day:     01,
			},
		},
		{
			Name:  "last month start",
			Input: time.Date(2018, 12, 17, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    18,
				Month:   'Z',
				Day:     01,
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
				Day:     01,
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
				Day:     02,
			},
		},
		{
			Name:  "year day leap day",
			Input: time.Date(2016, 12, 30, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    16,
				Month:   '+',
				Day:     01,
			},
		},
		{
			Name:  "year day",
			Input: time.Date(2018, 12, 31, 0, 0, 0, 0, time.UTC),
			Expected: desamber.Date{
				Century: 20,
				Year:    18,
				Month:   '+',
				Day:     01,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := desamber.New(tc.Input)
			if result.Century != tc.Expected.Century {
				t.Errorf("Expected century to be %v, but got %v", tc.Expected.Century, result.Century)
			}
			if result.Year != tc.Expected.Year {
				t.Errorf("Expected year to be %v, but got %v", tc.Expected.Year, result.Year)
			}
			if result.Month != tc.Expected.Month {
				t.Errorf("Expected month to be %v, but got %v", string(tc.Expected.Month), string(result.Month))
			}
			if result.Day != tc.Expected.Day {
				t.Errorf("Expected day to be %v, but got %v", tc.Expected.Day, result.Day)
			}
		})
	}
}
