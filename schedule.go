package main

import (
	"fmt"
	"sort"
)

// Schedule is a representatoin of a person's google calendar
type Schedule struct {
	name         string
	events       []CalendarRange
	availability CalendarRange
}

const MinEventSize = 10

// NewSchedule returns Schedule
// this is used as the public constructor in order to validate params
func NewSchedule(name string, events []CalendarRange, workingHours CalendarRange) Schedule {
	return Schedule{
		name:         name,
		events:       events,
		availability: workingHours,
	}
}

func NewRandomSchedule(numEvents int) Schedule {
	return NewSchedule(
		RandomName(2),
		NewRandomCalendarRangeCollection(numEvents),
		NewRandomWorkingHoursCalendarRange(),
	)
}

func (s *Schedule) sortEvents() {
	sort.Slice(s.events, func(i, j int) bool {
		return s.events[i].begin.rawMinValue < s.events[j].begin.rawMinValue
	})
}

func (s Schedule) findAvailability(duration int) []CalendarRange {
	if len(s.events) == 0 {
		return []CalendarRange{s.availability}
	}

	result := []CalendarRange{}
	runner := NewCalendarTimeFromRawMinValue(s.availability.begin.rawMinValue)

	for _, event := range s.events {
		diff, overflow := runner.absTimeDiff(event.begin)
		if overflow {
			fmt.Errorf("Seems like you overflowed into the next day")
		}
		if diff >= duration && event.begin.rawMinValue > runner.rawMinValue {
			result = append(result, NewCalendarRangeFor2Times(runner, event.begin))
		}
		if event.end.rawMinValue > runner.rawMinValue {
			runner = event.end
		}
	}
	diff, overflow := runner.absTimeDiff(s.availability.end)
	if overflow {
		fmt.Errorf("Seems like you overflowed into the next day")
	}

	if runner.rawMinValue < s.availability.end.rawMinValue && diff > duration {
		result = append(result, NewCalendarRangeFor2Times(runner, s.availability.end))
	}

	return result
}

func FindCommonAvailability(duration int, s1 Schedule, s2 Schedule) []CalendarRange {
	begin := maxCalendarTime(s1.availability.begin, s2.availability.begin)
	end := minCalendarTime(s1.availability.end, s2.availability.end)

	combinedSchedule := NewSchedule(
		"combined",
		append(s1.events, s2.events...),
		NewCalendarRangeFor2Times(begin, end),
	)
	combinedSchedule.sortEvents()
	return combinedSchedule.findAvailability(duration)
}

func (s Schedule) print() {
	fmt.Printf("%v's Schedule:\n", s.name)
	fmt.Printf("Availability: %v\n", s.availability.humanReadable())
	for index, event := range s.events {
		fmt.Printf("\t%d) %v\n", index, event.humanReadable())
	}
	fmt.Println()
}
