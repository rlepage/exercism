package booking

import (
	"fmt"
	"time"
)

const (
	// Allowed values for date format: https://pkg.go.dev/time#pkg-constants
	scheduleDateFormat string = "1/2/2006 15:04:05"
	hasPassedDateFormat string =  "January 2, 2006 15:04:05"
	isAfternoonDateFormat string = "Monday, January 2, 2006 15:04:05"
	descriptionDateFormat string = "Monday, January 2, 2006, at 15:04."
	anniversary string = "09/15/%d 00:00:00"
	noon int = 12
)

func parseDate(dateFormat string, date string) time.Time {
	res, err := time.Parse(dateFormat, date)
	if err != nil {
		panic(err)
	}
	return res
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	return parseDate(scheduleDateFormat, date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return parseDate(hasPassedDateFormat, date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	var hour int = parseDate(isAfternoonDateFormat, date).Hour()
	return hour >= noon && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	return "You have an appointment on " + parseDate(scheduleDateFormat, date).Format(descriptionDateFormat)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	var currentYear int = time.Now().Year()
	return parseDate(scheduleDateFormat, fmt.Sprintf(anniversary, currentYear))
}
