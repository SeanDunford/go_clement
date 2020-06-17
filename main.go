package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
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

func cliActionFmt(meetingTime int, maxChunks int, s1 Schedule, s2 Schedule) {
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

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V, v"},
		Usage:   "print only the version",
	}
	app := &cli.App{
		UseShortOptionHandling: true,
		Name:                   "clement",
		Version:                "v0.1.1",
		EnableBashCompletion:   true,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "meetingTime",
				Value: 30,
				Usage: "Duration of the meeting time we hope to find",
			},
			&cli.IntFlag{
				Name:  "maxSuggestions",
				Value: 3,
				Usage: "Maximum number of suggestions you would like",
			},
			&cli.StringFlag{Name: "s1Name", Value: "s1", Usage: "Name of first schedule"},
			&cli.StringFlag{Name: "s2Name", Value: "s2", Usage: "Name of second schedule"},
			&cli.StringFlag{Name: "schedule1", Value: "09:30-10:00,12:00-13:00,14:00-15:00", Usage: "Schedule for first person"},
			&cli.StringFlag{Name: "schedule2", Value: "09:00-10:00,11:00-12:00,17:00-18:00", Usage: "Schedule for second person"},
			&cli.StringFlag{Name: "workingHours1", Value: "05:30-20:00", Usage: "Working Hours for first person"},
			&cli.StringFlag{Name: "workingHours2", Value: "09:00-17:00", Usage: "Working Hours for second person"},
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "findMeetingTime",
			Aliases: []string{"fmt"},
			Usage:   "Find a meeting time for 2 different schedules",
			Action: func(c *cli.Context) error {
				s1 := Schedule{
					name:         c.String("s1Name"),
					events:       parseScheduleFlag(c.String("schedule1")),
					availability: parseMilitaryPairStr(c.String("workingHours1")),
				}
				s2 := Schedule{
					name:         c.String("s2Name"),
					events:       parseScheduleFlag(c.String("schedule2")),
					availability: parseMilitaryPairStr(c.String("workingHours2")),
				}
				cliActionFmt(c.Int("meetingTime"), c.Int("maxSuggestions"), s1, s2)
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
