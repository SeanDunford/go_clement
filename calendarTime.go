package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// The number of mins in a day 60
const MaxMinsInDay = 60 * 24
const NoonInMins = MaxMinsInDay / 2

type CalendarTime struct {
	hour           int
	minute         int
	rawMinValue    int
	originalString string
}

func clampMinValue(minValue int) int {
	clamped := 0
	if minValue >= 0 && minValue < MaxMinsInDay {
		clamped = minValue
	} else if minValue <= MaxMinsInDay {
		clamped = MaxMinsInDay - 1
	}
	return clamped
}

func clampHour(hour int) int {
	clampedHour := 0
	if hour >= 0 && hour < 24 {
		clampedHour = hour
	}
	return clampedHour
}

func clampMin(min int) int {
	clampedMin := 0
	if min >= 0 && min < 60 {
		clampedMin = min
	}
	return clampedMin
}

func parseHour(strHour string) int {
	hour, err := strconv.Atoi(strHour)
	if err != nil {
		fmt.Errorf("ya goofed: %v", err)
		hour = 0
	}

	return clampHour(hour)
}

func parseMin(strMin string) int {
	min, err := strconv.Atoi(strMin)
	if err != nil {
		fmt.Errorf("ya goofed: %v", err)
		min = 0
	}
	return clampMin(min)
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func (c CalendarTime) absTimeDiff(c2 CalendarTime) (int, bool) {
	timediff := Abs(c.rawMinValue - c2.rawMinValue)
	overflow := timediff > MaxMinsInDay

	return timediff, overflow
}

// calendaraTime.addTime really should mutate itself or be called something different

func (c CalendarTime) addTime(mins int) CalendarTime {
	newRawValue := c.rawMinValue + mins
	return NewCalendarTimeFromRawMinValue(newRawValue)
}

func calculateRawminValue(hour int, min int) int {
	return (60 * hour) + (min)
}

func minCalendarTime(c1, c2 CalendarTime) CalendarTime {
	if c1.rawMinValue < c2.rawMinValue {
		return c1
	}

	return c2
}

func maxCalendarTime(c1, c2 CalendarTime) CalendarTime {
	if c1.rawMinValue > c2.rawMinValue {
		return c1
	}

	return c2
}

func NewEmptyCalendarTime() CalendarTime {
	return CalendarTime{
		hour:           0,
		minute:         0,
		originalString: "",
		rawMinValue:    0,
	}
}

func NewMaxCalendarTime() CalendarTime {
	return NewCalendarTimeFromMilitaryStr("23:59")
}

func NewRandomCalendarTime() CalendarTime {
	rand.Seed(time.Now().UnixNano())
	return NewCalendarTimeFromRawMinValue(rand.Intn(MaxMinsInDay))
}

func NewRandomCalendarTimeWithClamp(min, max int) CalendarTime {
	rand.Seed(time.Now().UnixNano())
	return NewCalendarTimeFromRawMinValue(rand.Intn(max-min+1) + min)
}

func NewCalendarTime(hour int, min int, originalString string) CalendarTime {
	return CalendarTime{
		hour:           clampHour(hour),
		minute:         clampMin(min),
		originalString: originalString,
		rawMinValue:    calculateRawminValue(hour, min),
	}
}

func NewCalendarTimeFromRawMinValue(rawMinValue int) CalendarTime {
	minValue := clampMinValue(rawMinValue)
	hour := clampHour(minValue / 60)
	min := clampMin(minValue % 60)
	return NewCalendarTime(hour, min, "")
}

func NewCalendarTimeFromMilitaryStr(miltaryTime string) CalendarTime {
	time := strings.Split(miltaryTime, ":")
	hour := parseHour(time[0])
	min := parseMin(time[1])

	return NewCalendarTime(hour, min, miltaryTime)
}

func (c CalendarTime) print() {
	fmt.Printf("CalendarTime: %v\n", c)
}

func (c CalendarTime) similarTo(c2 CalendarTime) bool {
	return c.rawMinValue == c2.rawMinValue
}
