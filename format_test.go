package date_range_format

import (
	"testing"
	"time"
)

type TestCase struct {
	key    string
	lang   string
	d1     time.Time
	d2     time.Time
	result string
}

var testCases = []TestCase{
	{
		"uk-same-date",
		"uk",
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local),
		"1 січня 2017",
	},
	{
		"uk-same-month",
		"uk",
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2017, 1, 5, 0, 0, 0, 0, time.Local),
		"1 - 5 січня 2017",
	},
	{
		"uk-same-year",
		"uk",
		time.Date(2017, 1, 30, 0, 0, 0, 0, time.Local),
		time.Date(2017, 2, 7, 0, 0, 0, 0, time.Local),
		"30 січня - 7 лютого 2017",
	},
	{
		"uk-different-years",
		"uk",
		time.Date(2016, 12, 30, 0, 0, 0, 0, time.Local),
		time.Date(2017, 1, 6, 0, 0, 0, 0, time.Local),
		"30 грудня 2016 - 6 січня 2017",
	},
	{
		"ru-same-date",
		"ru",
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local),
		"1 января 2017",
	},
	{
		"ru-same-month",
		"ru",
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2017, 1, 5, 0, 0, 0, 0, time.Local),
		"1 - 5 января 2017",
	},
	{
		"ru-same-year",
		"ru",
		time.Date(2017, 1, 30, 0, 0, 0, 0, time.Local),
		time.Date(2017, 2, 7, 0, 0, 0, 0, time.Local),
		"30 января - 7 февраля 2017",
	},
	{
		"ru-different-years",
		"ru",
		time.Date(2016, 12, 30, 0, 0, 0, 0, time.Local),
		time.Date(2017, 1, 6, 0, 0, 0, 0, time.Local),
		"30 декабря 2016 - 6 января 2017",
	},
	{
		"en-same-date",
		"en",
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local),
		"January, 1 2017",
	},
	{
		"en-same-month",
		"en",
		time.Date(2017, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2017, 1, 5, 0, 0, 0, 0, time.Local),
		"January, 1-5 2017",
	},
	{
		"en-same-year",
		"en",
		time.Date(2017, 1, 30, 0, 0, 0, 0, time.Local),
		time.Date(2017, 2, 7, 0, 0, 0, 0, time.Local),
		"January, 30 - February, 7 2017",
	},
	{
		"en-different-years",
		"en",
		time.Date(2016, 12, 30, 0, 0, 0, 0, time.Local),
		time.Date(2017, 1, 6, 0, 0, 0, 0, time.Local),
		"December, 30 2016 - January, 6 2017",
	},
}

func TestDateRangeFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.key, func(t *testing.T){
			f := DateRangeFormat(tc.lang, tc.d1, tc.d2)
			if f != tc.result {
				t.Errorf("got %s; want %s", f, tc.result)
			}
		})
	}
}
