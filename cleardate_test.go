package cleardate

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	fixedTime := time.Date(2025, time.February, 9, 14, 30, 45, 987654321, time.UTC)

	tests := []struct {
		name   string
		layout string
		t      time.Time
		want   string
	}{
		{
			name:   "Simple Date (yyyy-MM-dd)",
			layout: "yyyy-MM-dd",
			t:      fixedTime,
			want:   "2025-02-09",
		},
		{
			name:   "24-hour Time (HH:mm:ss)",
			layout: "HH:mm:ss",
			t:      fixedTime,
			want:   "14:30:45",
		},
		{
			name:   "12-hour Time with AM/PM (hh:mm:ss a)",
			layout: "hh:mm:ss a",
			t:      fixedTime,
			want:   "02:30:45 PM",
		},
		{
			name:   "Fractional Seconds (ss.SSS)",
			layout: "ss.SSS",
			t:      fixedTime,
			want:   "45.987",
		},
		{
			name:   "Full and Abbreviated Month Names (MMMM MMM)",
			layout: "MMMM MMM",
			t:      fixedTime,
			want:   "February Feb",
		},
		{
			name:   "Day Tokens (d dd)",
			layout: "d dd",
			t:      fixedTime,
			want:   "9 09",
		},
		{
			name:   "Day of Week Tokens (EEEE EEE)",
			layout: "EEEE EEE",
			t:      fixedTime,
			want:   "Sunday Sun",
		},
		{
			name:   "Short Year (yy)",
			layout: "yy",
			t:      fixedTime,
			want:   "25",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Format(tc.t, tc.layout)
			if got != tc.want {
				t.Errorf("Format(%v, %q) = %q; want %q", tc.t, tc.layout, got, tc.want)
			}
		})
	}

	mst := time.FixedZone("MST", -7*3600)
	tZone := time.Date(2025, time.February, 9, 14, 30, 0, 0, mst)
	tzTests := []struct {
		name   string
		layout string
		t      time.Time
		want   string
	}{
		{
			name:   "Timezone Numeric without Colon (ZZ)",
			layout: "ZZ",
			t:      tZone,
			want:   "-0700",
		},
		{
			name:   "Timezone Numeric with Colon (Z)",
			layout: "Z",
			t:      tZone,
			want:   "-07:00",
		},
		{
			name:   "Timezone Abbreviation (z)",
			layout: "z",
			t:      tZone,
			want:   "MST",
		},
	}

	for _, tc := range tzTests {
		t.Run(tc.name, func(t *testing.T) {
			got := Format(tc.t, tc.layout)
			if got != tc.want {
				t.Errorf("Format(%v, %q) = %q; want %q", tc.t, tc.layout, got, tc.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	layout := "yyyy-MM-dd HH:mm:ss Z"
	input := "2025-02-09 14:30:00 +00:00"
	expected := time.Date(2025, time.February, 9, 14, 30, 0, 0, time.UTC)
	parsedTime, err := Parse(layout, input)
	if err != nil {
		t.Fatalf("Parse(%q, %q) returned error: %v", layout, input, err)
	}
	if !parsedTime.Equal(expected) {
		t.Errorf("Parse(%q, %q) = %v; want %v", layout, input, parsedTime, expected)
	}

	layout2 := "yyyy-MM-dd HH:mm:ss"
	input2 := "2025-02-09 14:30:00"

	expected2 := time.Date(2025, time.February, 9, 14, 30, 0, 0, time.UTC)
	parsedTime2, err := Parse(layout2, input2)
	if err != nil {
		t.Fatalf("Parse(%q, %q) returned error: %v", layout2, input2, err)
	}
	if !parsedTime2.Equal(expected2) {
		t.Errorf("Parse(%q, %q) = %v; want %v", layout2, input2, parsedTime2, expected2)
	}

	layout3 := "yyyy-MM-dd HH:mm:ss.SSS"
	input3 := "2025-02-09 14:30:45.987"
	expected3 := time.Date(2025, time.February, 9, 14, 30, 45, 987000000, time.UTC)
	parsedTime3, err := Parse(layout3, input3)
	if err != nil {
		t.Fatalf("Parse(%q, %q) returned error: %v", layout3, input3, err)
	}
	if !parsedTime3.Equal(expected3) {
		t.Errorf("Parse(%q, %q) = %v; want %v", layout3, input3, parsedTime3, expected3)
	}

	layout4 := "yyyy-MM-dd"
	input4 := "invalid-date"
	_, err = Parse(layout4, input4)
	if err == nil {
		t.Errorf("Parse(%q, %q) expected error, got nil", layout4, input4)
	}
}

func TestParseInLocation(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Fatalf("Failed to load location Europe/Berlin: %v", err)
	}

	layout := "yyyy-MM-dd HH:mm:ss"
	input := "2025-02-09 14:30:00"
	expected := time.Date(2025, time.February, 9, 14, 30, 0, 0, loc)
	parsedTime, err := ParseInLocation(layout, input, loc)
	if err != nil {
		t.Fatalf("ParseInLocation(%q, %q, loc) returned error: %v", layout, input, err)
	}
	if !parsedTime.Equal(expected) {
		t.Errorf("ParseInLocation(%q, %q, loc) = %v; want %v", layout, input, parsedTime, expected)
	}
}
