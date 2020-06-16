package availability

// Schedule 9am to 3pm

// meeting from 1pm to 2:30pm

// meeting time 30mins

// Should retuen
//     [
//         [9, 9:30],
//         [9:30, 10],
//         [10:30, 11],
//         [11, 11:30],
//         [11:30, 12],
//         [12: 12:30]
//         [12:30, 1]
//         [2:30, 3]
//     ]

func echo(shiftStart, shiftEnd, meetingStart, meetingEnd string) []string {
	return []string{shiftStart, shiftEnd, meetingStart, meetingEnd}
}

func findAvailability(shiftStart, shiftEnd, meetingStart, meetingEnd string) [][]string {
	CalendarTime
	return [][]string{{"hello, GO!"}}
}
