package main

import (
	"flag"
	"fmt"
	"strings"
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

func parseMilitaryPairStr(militaryStrPair string) CalendarRange {
	eventTimes := strings.Split(militaryStrPair, "-")
	return NewCalendarRangeFromMilitaryStrings(eventTimes[0], eventTimes[1])
}

func parseScheduleFlag(scheduleStr string) []CalendarRange {
	scheduleStrArr := strings.Split(scheduleStr, ",")
	result := make([]CalendarRange, len(scheduleStrArr))

	for index, strCalendarRange := range scheduleStrArr {
		result[index] = parseMilitaryPairStr(strCalendarRange)
	}

	return result
}

func main() {
	meetingTime := *flag.Int("meetingTime", 30, "Duration of the meeting time we hope to find")
	maxChunks := *flag.Int("maxSuggestions", 3, "Maximum number of suggestions you would like")
	s1Name := *flag.String("s1Name", "s1", "Name of first schedule")
	s2Name := *flag.String("s2Name", "s2", "Name of second schedule")
	schedule1Str := *flag.String("schedule1", "09:30-10:00,12:00-13:00,14:00-15:00", "Schedule for first person")
	schedule2Str := *flag.String("schedule2", "09:00-10:00,11:00-12:00,17:00-18:00", "Schedule for second person")
	workingHours1 := *flag.String("workingHours1", "05:30-20:00", "Working Hours for first person")
	workingHours2 := *flag.String("workingHours2", "09:00-17:00", "Woeking Hours for second person")
	flag.Parse()

	s1Events := parseScheduleFlag(schedule1Str)
	s2Events := parseScheduleFlag(schedule2Str)
	wh1 := parseMilitaryPairStr(workingHours1)
	wh2 := parseMilitaryPairStr(workingHours2)

	s1 := Schedule{
		name:         s1Name,
		events:       s1Events,
		availability: wh1,
	}
	s2 := Schedule{
		name:         s2Name,
		events:       s2Events,
		availability: wh2,
	}

	availabilityArr := FindCommonAvailability(meetingTime, s1, s2)

	chunks := ChunkAvailability(availabilityArr, meetingTime, maxChunks)

	s1.print()
	s2.print()

	fmt.Printf("I was able to find %v possible meeting times of %v mins long\n", len(chunks), meetingTime)
	for _, suggestedTime := range chunks {
		fmt.Printf("\t%v\n", suggestedTime.humanReadable())
	}

	fmt.Println("\n\nfin")
}
