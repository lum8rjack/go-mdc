package main

import (
	"testing"
)

// TestGetAreaByKey
func TestGetAreaByName(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
	}{
		{name: "B.K. Leach", expected: "8514"},
		{name: "Ten Mile Pond", expected: "8241"},
		{name: "Otter Slough", expected: "5004"},
		{name: "Schell-Osage", expected: "5701"},
		{name: "Four Rivers", expected: "8238"},
		{name: "Montrose", expected: "5604"},
		{name: "Duck Creek", expected: "5001"},
		{name: "Columbia Bottoms", expected: "9736"},
		{name: "Marais Temps Clair", expected: "7902"},
		{name: "Ted Shanks", expected: "7011"},
		{name: "Eagle Bluffs", expected: "8931"},
		{name: "Grand Pass", expected: "8010"},
		{name: "Fountain Grove", expected: "4601"},
		{name: "Bob Brown", expected: "8142"},
		{name: "Nodaway Valley", expected: "9134"},
		{name: "invalid name", expected: ""},
		{name: "ted", expected: "7011"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			area, _ := GetAreaByName(tt.name)
			if area.Key != tt.expected {
				t.Errorf("GetAreaByName(%s) = %s, should equal %s\n", tt.name, area.Name, tt.expected)
			}
		})
	}
}

// TestCheckYear
func TestCheckYear(t *testing.T) {
	var tests = []struct {
		name     string
		year     int
		expected string
	}{
		{name: "too far back", year: 2015, expected: ""},
		{name: "past year", year: 2016, expected: "2016-2017"},
		{name: "future year", year: 3000, expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			season, _ := CheckYear(tt.year)
			if season != tt.expected {
				t.Errorf("CheckYear(%d) = %s, should equal %s\n", tt.year, season, tt.expected)
			}
		})
	}
}
