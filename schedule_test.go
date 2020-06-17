package main

import (
	"fmt"
	"testing"
)

func TestScheduleReturnOneGiantAvailabilityBlock(t *testing.T) {
	meetingTime := 30
	sean := Schedule{
		name:   RandomName(2),
		events: []CalendarRange{},
		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromRawMinValue(0),
			NewCalendarTimeFromRawMinValue(MaxMinsInDay),
		),
	}

	seanAvailability := sean.findAvailability(meetingTime)
	fmt.Printf("print dat: %v", seanAvailability)

	expectedArr := []CalendarRange{NewCalendarRangeAllDay()}

	for index, expected := range expectedArr {
		if !expected.similarTo(seanAvailability[index]) {
			t.Error("test failed expected availability:", expected, "to be equal to :", seanAvailability[index])
		}
	}
}

func TestScheduleReturnBeforeNoon(t *testing.T) {
	meetingTime := 30
	sean := Schedule{
		name:   RandomName(2),
		events: []CalendarRange{},
		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromRawMinValue(0),
			NewCalendarTimeFromRawMinValue(NoonInMins),
		),
	}
	expectedArr := []CalendarRange{NewCalendarRangeBeforeNoon()}
	availabilityArr := sean.findAvailability(meetingTime)

	for index, expected := range expectedArr {
		if !expected.similarTo(availabilityArr[index]) {
			t.Error("test failed expected availability:", expected, "to be equal to :", availabilityArr[index])
		}
	}
}

func TestScheduleReturnAfterNoon(t *testing.T) {
	meetingTime := 30
	sean := Schedule{
		name:   RandomName(2),
		events: []CalendarRange{},
		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromRawMinValue(NoonInMins),
			NewCalendarTimeFromRawMinValue(MaxMinsInDay-1),
		),
	}
	expectedArr := []CalendarRange{NewCalendarRangeAfterNoon()}
	availabilityArr := sean.findAvailability(meetingTime)

	for index, expected := range expectedArr {
		if !expected.similarTo(availabilityArr[index]) {
			t.Error("test failed expected availability:", expected, "to be equal to :", availabilityArr[index])
		}
	}
}

func TestScheduleOneEvent(t *testing.T) {
	meetingTime := 30

	sean := Schedule{
		name:   RandomName(2),
		events: []CalendarRange{NewCalendarRangeLunchBreak()},
		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromRawMinValue(NoonInMins),
			NewCalendarTimeFromRawMinValue(MaxMinsInDay-1),
		),
	}
	expectedArr := []CalendarRange{
		NewCalendarRangeBeforeNoon(),
		NewCalendarRangeFor2Times(NewCalendarTimeFromMilitaryStr("13:00"), NewCalendarTimeFromRawMinValue(MaxMinsInDay-1)),
	}
	availabilityArr := sean.findAvailability(meetingTime)

	for index, expected := range expectedArr {
		if !expected.similarTo(availabilityArr[index]) {
			t.Error("test expected:", expected, "but received:", availabilityArr[index])
		}
	}
}

func TestScheduleTwoEvents(t *testing.T) {
	meetingTime := 30
	events := []CalendarRange{
		NewCalendarRangeLunchBreak(),
		NewCalendarRangeFromMilitaryStrings("14:30", "17:55"),
	}

	sean := Schedule{
		name:   RandomName(2),
		events: events,
		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromRawMinValue(NoonInMins),
			NewCalendarTimeFromRawMinValue(MaxMinsInDay-1),
		),
	}
	expectedArr := []CalendarRange{
		NewCalendarRangeBeforeNoon(),
		NewCalendarRangeFor2Times(NewCalendarTimeFromMilitaryStr("13:00"), NewCalendarTimeFromMilitaryStr("14:30")),
		NewCalendarRangeFor2Times(NewCalendarTimeFromMilitaryStr("17:55"), NewCalendarTimeFromRawMinValue(MaxMinsInDay-1)),
	}
	availabilityArr := sean.findAvailability(meetingTime)

	for index, expected := range expectedArr {
		if !expected.similarTo(availabilityArr[index]) {
			t.Error("test expected:", expected, "but received:", availabilityArr[index])
		}
	}
}

func TestCommonAvailabilitySimpleExample(t *testing.T) {
	meetingTime := 30
	seanEvents := []CalendarRange{
		NewCalendarRangeFromMilitaryStrings("9:00", "13:00"),
	}
	ashlyEvents := []CalendarRange{
		NewCalendarRangeFromMilitaryStrings("9:00", "9:30"),
		NewCalendarRangeFromMilitaryStrings("9:30", "10:30"),
		NewCalendarRangeFromMilitaryStrings("10:30", "10:30"),
		NewCalendarRangeFromMilitaryStrings("11:30", "10:30"),
		NewCalendarRangeFromMilitaryStrings("12:30", "13:00"),
	}
	expectedArr := []CalendarRange{
		NewCalendarRangeFromMilitaryStrings("0:00", "9:00"),
		NewCalendarRangeFromMilitaryStrings("13:00", "23:59"),
	}
	sean := Schedule{
		name:   RandomName(2),
		events: seanEvents,
		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromRawMinValue(NoonInMins),
			NewCalendarTimeFromRawMinValue(MaxMinsInDay-1),
		),
	}
	ashly := Schedule{
		name:   RandomName(2),
		events: ashlyEvents,
		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromRawMinValue(NoonInMins),
			NewCalendarTimeFromRawMinValue(MaxMinsInDay-1),
		),
	}
	availabilityArr := FindCommonAvailability(meetingTime, sean, ashly)

	for index, expected := range expectedArr {
		if !expected.similarTo(availabilityArr[index]) {
			t.Error("test expected:", expected, "but received:", availabilityArr[index])
		}
	}
}

func TestCommonAvailability(t *testing.T) {
	meetingTime := 30
	maxChunks := 3
	seanEvents := []CalendarRange{
		NewCalendarRangeLunchBreak(),
		NewCalendarRangeFromMilitaryStrings("14:30", "17:55"),
	}
	ashlyEvents := []CalendarRange{
		NewCalendarRangeFromMilitaryStrings("9:00", "9:30"),
		NewCalendarRangeLunchBreak(),
		NewCalendarRangeFromMilitaryStrings("15:30", "18:55"),
	}
	expectedArr := []CalendarRange{
		NewCalendarRangeFromMilitaryStrings("0:00", "9:00"),
		NewCalendarRangeFromMilitaryStrings("9:30", "12:00"),
		NewCalendarRangeFromMilitaryStrings("13:00", "14:30"),
		NewCalendarRangeFromMilitaryStrings("18:55", "23:59"),
	}
	sean := Schedule{
		name:   RandomName(2),
		events: seanEvents,
		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromRawMinValue(NoonInMins),
			NewCalendarTimeFromRawMinValue(MaxMinsInDay-1),
		),
	}
	ashly := Schedule{
		name:   RandomName(2),
		events: ashlyEvents,
		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromRawMinValue(NoonInMins),
			NewCalendarTimeFromRawMinValue(MaxMinsInDay-1),
		),
	}
	availabilityArr := FindCommonAvailability(meetingTime, sean, ashly)

	for index, expected := range expectedArr {
		if !expected.similarTo(availabilityArr[index]) {
			t.Error("test expected:", expected, "but received:", availabilityArr[index])
		}
	}

	expectedChunks := []CalendarRange{
		NewCalendarRangeFromMilitaryStrings("0:00", "0:30"),
		NewCalendarRangeFromMilitaryStrings("0:30", "1:00"),
		NewCalendarRangeFromMilitaryStrings("1:00", "1:30"),
	}
	chunks := ChunkAvailability(availabilityArr, meetingTime, maxChunks)

	for index, expected := range expectedChunks {
		if !expected.similarTo(chunks[index]) {
			t.Error("test expected:", expected, "but received:", availabilityArr[index])
		}
	}

}

func TestCommonAvailability2(t *testing.T) {
	meetingTime := 90
	maxChunks := 4
	seanEvents := []CalendarRange{
		NewCalendarRangeLunchBreak(),
		NewCalendarRangeFromMilitaryStrings("14:30", "17:55"),
	}
	ashlyEvents := []CalendarRange{
		NewCalendarRangeFromMilitaryStrings("9:00", "9:30"),
		NewCalendarRangeLunchBreak(),
		NewCalendarRangeFromMilitaryStrings("15:30", "18:55"),
	}
	sean := Schedule{
		name:   RandomName(2),
		events: seanEvents,

		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromMilitaryStr("9:00"),
			NewCalendarTimeFromMilitaryStr("19:00"),
		),
	}
	ashly := Schedule{
		name:   RandomName(2),
		events: ashlyEvents,
		availability: NewCalendarRangeFor2Times(
			NewCalendarTimeFromRawMinValue(NoonInMins),
			NewCalendarTimeFromRawMinValue(MaxMinsInDay-1),
		),
	}
	availabilityArr := FindCommonAvailability(meetingTime, sean, ashly)

	expectedChunks := []CalendarRange{
		NewCalendarRangeFromMilitaryStrings("9:30", "11:00"),
		NewCalendarRangeFromMilitaryStrings("13:00", "14:30"),
	}
	chunks := ChunkAvailability(availabilityArr, meetingTime, maxChunks)

	for index, expected := range expectedChunks {
		if !expected.similarTo(chunks[index]) {
			t.Error("test expected:", expected, "but received:", availabilityArr[index])
		}
	}

	fmt.Printf("Given %s's schedule:\n", sean.name)
	sean.print()

	fmt.Printf("Given %s's schedule:\n", ashly.name)
	ashly.print()

	fmt.Printf("I was able to find %v possible meeting times of %v mins long\n", len(chunks), meetingTime)
	for _, suggestedTime := range chunks {
		suggestedTime.print()
	}

	fmt.Println("fin")
}
