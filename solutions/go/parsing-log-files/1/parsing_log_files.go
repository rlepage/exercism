package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`).MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[\*~=\-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		re := regexp.MustCompile(`\"[^"]*password[^"]*\"`)
		match := re.FindAllStringIndex(strings.ToLower(line), -1)
		count += len(match)
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	reUser := regexp.MustCompile(`User\s+(?P<username>[^\s]+)\s`)
	res := []string{}
	for _, line := range lines {
		match := reUser.FindStringSubmatch(line)
		if len(match) > 1 {
			username := match[1]
			line = "[USR] " + username + " " + line
		}
		res = append(res, line)
	}
	return res
}
