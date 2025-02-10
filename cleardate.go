package cleardate

import (
	"strings"
	"time"
)

var replacer = strings.NewReplacer(
	// Fractional seconds
	"SSSSSSSSS", "000000000",
	"SSSSSSSS", "00000000",
	"SSSSSSS", "0000000",
	"SSSSSS", "000000",
	"SSSSS", "00000",
	"SSSS", "0000",
	"SSS", "000",
	"SS", "00",
	"S", "0",

	// Year tokens
	"yyyy", "2006",
	"yy", "06",

	// Month tokens
	"MMMM", "January",
	"MMM", "Jan",
	"MM", "01",
	"M", "1",

	// Day of week tokens
	"EEEE", "Monday",
	"EEE", "Mon",

	// Day tokens
	"dd", "02",
	"d", "2",

	// Hour tokens
	"HH", "15",
	"H", "15",
	"hh", "03",
	"h", "3",

	// Minute tokens
	"mm", "04",
	"m", "4",

	// Second tokens
	"ss", "05",
	"s", "5",

	// Time zone tokens
	"ZZ", "Z0700",
	"Z", "Z07:00",
	"z", "MST",

	// AM/PM tokens
	"a", "PM",
	"A", "PM",
)

func convertLayout(layout string) string {
	return replacer.Replace(layout)
}

func Format(t time.Time, layout string) string {
	goLayout := convertLayout(layout)
	return t.Format(goLayout)
}

func Parse(layout, value string) (time.Time, error) {
	goLayout := convertLayout(layout)
	return time.Parse(goLayout, value)
}

func ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	goLayout := convertLayout(layout)
	return time.ParseInLocation(goLayout, value, loc)
}
