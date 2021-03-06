package date_range_format

import (
	"bytes"
	"strconv"
	"text/template"
	"time"
)

func DateRangeFormat(lang string, d1, d2 time.Time) string {
	var (
		key    string
		writer = new(bytes.Buffer)

		sameDayCondition   = d1.Month() == d2.Month() && d1.Year() == d2.Year() && d1.Day() == d2.Day()
		sameMonthCondition = d1.Month() == d2.Month() && d1.Year() == d2.Year()
		sameYearCondition  = d1.Year() == d2.Year()
	)

	if sameDayCondition {
		key = "same-day"
	} else if sameMonthCondition {
		key = "same-month"
	} else if sameYearCondition {
		key = "same-year"
	} else {
		key = "different-years"
	}

	var data = struct {
		D1Date,
		D1Month,
		D1Year,
		D2Date,
		D2Month,
		D2Year string
	}{
		strconv.Itoa(d1.Day()),
		monthGenitives[lang][int(d1.Month())],
		strconv.Itoa(d1.Year()),
		strconv.Itoa(d2.Day()),
		monthGenitives[lang][int(d2.Month())],
		strconv.Itoa(d2.Year()),
	}

	var t = template.Must(template.New("d").Parse(formatTemplates[lang][key]))
	t.Execute(writer, data)

	return writer.String()
}
