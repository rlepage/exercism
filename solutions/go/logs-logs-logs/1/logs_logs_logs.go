package logs

import "unicode/utf8"

var applications map[rune]string = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		if r == '❗' || r == '🔍' || r == '☀' {
			return applications[r]
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var modified string
	for _, r := range log {
		if r == oldRune {
			modified += string(newRune)
		} else {
			modified += string(r)
		}
	}
	return modified
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
