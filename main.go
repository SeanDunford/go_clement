package main

import (
	"fmt"
)

func clampMeetingTime(meetingTimesMins int) int { // better way to do this?
	const MaxMins = 60 * 24 // 24 hours * 60mins
	if meetingTimesMins < 0 {
		meetingTimesMins = 0
	}
	if meetingTimesMins > MaxMins {
		meetingTimesMins = MaxMins
	}

	return meetingTimesMins
}

func find3Times(meetingTimeMins int,
	s1 Schedule,
	s2 Schedule,
) []string {
	meetingTimeMins = clampMeetingTime(meetingTimeMins) // How to avoid manipulating input var?
	fmt.Println("Finding a ", meetingTimeMins, " mins long meeting")

	return []string{"a", "b", "c"}

}

func main() {
	sean := NewSchedule(
		"sean",
		[]CalendarRange{
			NewCalendarRangeFromMilitaryStrings("10:00", "11:30"),
			NewCalendarRangeFromMilitaryStrings("12:30", "14:30"),
			NewCalendarRangeFromMilitaryStrings("14:30", "15:00"),
		},
		NewCalendarRangeFromMilitaryStrings("9:00", "20:00"),
	)
	sean.print()
	clement := NewSchedule(
		"clement",
		[]CalendarRange{
			NewCalendarRangeFromMilitaryStrings("10:00", "11:30"),
			NewCalendarRangeFromMilitaryStrings("12:30", "14:30"),
			NewCalendarRangeFromMilitaryStrings("14:30", "15:00"),
		},
		NewCalendarRangeFromMilitaryStrings("9:00", "20:00"),
	)
	clement.print()

	fmt.Println("We did it")

	// seanMeetingBlocks = sean.findMeetingTimes(30)
	// clementMeetingBlocks = clement.findMeetingTimes(30)

	// overlap = seanMeetingBlocks.findOverlap(clementMeetingBlocks)
	// slice to 3 items
	// return
	// format output
}
