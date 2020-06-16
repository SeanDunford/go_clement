package main

import (
	"testing"
)

func TestCalendarTimeConstruction(t *testing.T) {
	CalendarTime{
		hour:           0,
		minute:         0,
		rawMinValue:    0,
		originalString: "",
	}.print()
	NewEmptyCalendarTime().print()
	NewRandomCalendarTime().print()
	NewCalendarTime(0, 0, "").print()
}
