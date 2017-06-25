package date_range_format

var formatTemplates = map[string]map[string]string{

	// Ukrainian:
	//
	// same-day:        14 лютого 2017
	// same month:      14 - 21 березня 2015
	// same year:       30 лютого - 6 березня 2015
	// different-years: 31 грудня 2014 - 6 січня 2015
	"uk": map[string]string{
		"same-day":        "{{.D1Date}} {{.D1Month}} {{.D1Year}}",
		"same-month":      "{{.D1Date}} - {{.D2Date}} {{.D2Month}} {{.D2Year}}",
		"same-year":       "{{.D1Date}} {{.D1Month}} - {{.D2Date}} {{.D2Month}} {{.D2Year}}",
		"different-years": "{{.D1Date}} {{.D1Month}} {{.D1Year}} - {{.D2Date}} {{.D2Month}} {{.D2Year}}",
	},

	// Russian:
	//
	// same-day:        14 февраля 2017
	// same month:      14 - 21 марта 2015
	// same year:       30 января - 6 марта 2015
	// different-years: 31 декабря 2014 - 6 января 2015
	"ru": map[string]string{
		"same-day":        "{{.D1Date}} {{.D1Month}} {{.D1Year}}",
		"same-month":      "{{.D1Date}} - {{.D2Date}} {{.D2Month}} {{.D2Year}}",
		"same-year":       "{{.D1Date}} {{.D1Month}} - {{.D2Date}} {{.D2Month}} {{.D2Year}}",
		"different-years": "{{.D1Date}} {{.D1Month}} {{.D1Year}} - {{.D2Date}} {{.D2Month}} {{.D2Year}}",
	},

	// English:
	//
	// same-day:        December, 31 2014
	// same month:      September, 1-7 2015
	// same year:       August, 1 - September 6 2015
	// different-years: December, 31 2014 - January, 6 2015
	"en": map[string]string{
		"same-day":        "{{.D1Month}}, {{.D1Date}} {{.D1Year}}",
		"same-month":      "{{.D1Month}}, {{.D1Date}}-{{.D2Date}} {{.D1Year}}",
		"same-year":       "{{.D1Month}}, {{.D1Date}} - {{.D2Month}}, {{.D2Date}} {{.D1Year}}",
		"different-years": "{{.D1Month}}, {{.D1Date}} {{.D1Year}} - {{.D2Month}}, {{.D2Date}} {{.D2Year}}",
	},
}