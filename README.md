# Desamber

Calendar format based on [Neauoire's Calendar](https://wiki.xxiivv.com/#desamber)

> The Desamber Calendar has 26 months of 14 days each. The 365th day of the year is the Year Day, preceded by the Leap Day on leap years.

> Each month has 2 weeks of 7 days, and each month's name is one of the 26 letters of the alphabet.

``` golang
func ExampleDate_WithYear() {
	date := desamber.New(time.Date(2018, 7, 8, 0, 0, 0, 0, time.UTC))
	date = date.WithYear()
	fmt.Println(date)
	// output: 18N07
}
```