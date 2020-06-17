package main

import (
	"fmt"
)

type CalendarRange struct {
	begin    CalendarTime
	end      CalendarTime
	duration int
}

func NewCalendarRangeFromMilitaryStrings(beginStr string, endStr string) CalendarRange {
	begin := NewCalendarTimeFromMilitaryStr(beginStr)
	end := NewCalendarTimeFromMilitaryStr(endStr)

	return NewCalendarRangeFor2Times(begin, end)
}

func NewCalendarRangeFor2Times(begin, end CalendarTime) CalendarRange {
	duration, overflow := begin.absTimeDiff(end)
	if overflow {
		fmt.Printf("Begin: %v\n", begin)
		fmt.Printf("End: %v\n", end)
		fmt.Errorf("Seems like you overflowed into the next day")
	}

	return CalendarRange{
		begin:    begin,
		end:      end,
		duration: duration,
	}
}

func NewCalendarRangeLunchBreak() CalendarRange {
	return NewCalendarRangeForTimeAndDuration(NewCalendarTimeFromMilitaryStr("12:00"), 60)
}

func NewRandomCalendarRange() CalendarRange {
	begin := NewRandomCalendarTime()
	end := NewRandomCalendarTimeWithClamp(begin.addTime(MinEventSize).rawMinValue, MaxMinsInDay)
	duration, overflow := end.absTimeDiff(begin)
	if overflow {
		fmt.Printf("Begin: %v\n", begin)
		fmt.Printf("End: %v\n", end)
		fmt.Errorf("Seems like you overflowed into the next day")
	}
	return CalendarRange{
		begin:    begin,
		end:      end,
		duration: duration,
	}
}

func NewRandomCalendarRangeCollection(len int) []CalendarRange {
	cr := make([]CalendarRange, len)
	for i := 0; i < len; i++ {
		cr[i] = NewRandomCalendarRange()
	}
	return cr
}

func NewRandomWorkingHoursCalendarRange() CalendarRange {
	begin := NewRandomCalendarTimeWithClamp(0, NoonInMins)
	end := NewRandomCalendarTimeWithClamp(NoonInMins, MaxMinsInDay)
	duration, overflow := end.absTimeDiff(begin)
	if overflow {
		fmt.Printf("Begin: %v\n", begin)
		fmt.Printf("End: %v\n", end)
		fmt.Errorf("Seems like you overflowed into the next day")
	}
	return CalendarRange{
		begin:    begin,
		end:      end,
		duration: duration,
	}
}

func NewCalendarRangeForMilitaryTimeAndDuration(militaryTime string, duration int) CalendarRange {
	begin := NewCalendarTimeFromMilitaryStr(militaryTime)
	return CalendarRange{
		begin:    begin,
		end:      begin.addTime(duration),
		duration: duration,
	}
}

func NewCalendarRangeForTimeAndDuration(begin CalendarTime, duration int) CalendarRange {
	return CalendarRange{
		begin:    begin,
		end:      begin.addTime(duration),
		duration: duration,
	}
}

func NewCalendarRangeAllDay() CalendarRange {
	return NewCalendarRangeForTimeAndDuration(NewCalendarTimeFromRawMinValue(0), MaxMinsInDay-1)
}

func NewCalendarRangeBeforeNoon() CalendarRange {
	return NewCalendarRangeFor2Times(
		NewCalendarTimeFromRawMinValue(0),
		NewCalendarTimeFromRawMinValue(NoonInMins),
	)
}
func NewCalendarRangeAfterNoon() CalendarRange {
	return NewCalendarRangeFor2Times(
		NewCalendarTimeFromRawMinValue(NoonInMins),
		NewCalendarTimeFromRawMinValue(MaxMinsInDay-1),
	)
}

func (c CalendarRange) humanReadable() string {
	return fmt.Sprintf("%v-%v", c.begin.humanReadable(), c.end.humanReadable())
}

func (c CalendarRange) print() {
	fmt.Printf(c.humanReadable())
}

func (c CalendarRange) similarTo(c2 CalendarRange) bool {
	return c.begin.similarTo(c2.begin) && c.end.similarTo(c2.end) && c.duration == c2.duration
}

func mergeEventArrays(c1 []CalendarRange, c2 []CalendarRange) []CalendarRange {
	return []CalendarRange{}
}

func ChunkAvailability(availiability []CalendarRange, duration, maxChunks int) []CalendarRange {
	// TODO: I want to use make([]CalendarRange, maxChunks) but i need it to be intialized with undefined? Is that a thing?
	chunkedAvailability := []CalendarRange{}
	counter := 0

	for _, currentEvent := range availiability {
		if counter >= maxChunks {
			break
		}
		runner := currentEvent.begin
		diff, overflow := runner.absTimeDiff(currentEvent.end)
		if overflow {
			fmt.Errorf("Seems like you overflowed into the next day")
		}

		for diff >= duration {
			if counter >= maxChunks {
				break
			}
			end := runner.addTime(duration)
			chunkedAvailability = append(chunkedAvailability, NewCalendarRangeFor2Times(runner, end))

			// re-initialize for next iteration
			runner = end
			counter++
			diff, overflow = runner.absTimeDiff(currentEvent.end)
			if overflow {
				fmt.Errorf("Seems like you overflowed into the next day")
			}
		}

	}

	return chunkedAvailability
}
