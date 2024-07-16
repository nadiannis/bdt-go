package utils

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"time"
)

func GetInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func FormatDate(datetime time.Time) string {
	format := "02 Jan 2006"
	return datetime.Format(format)
}

func BoolToStatusDisplayText(isPresent bool) string {
	if isPresent {
		return "present"
	}
	return "absent"
}

func StatusInputTextToBool(presenceStatus string) (bool, error) {
	if strings.ToLower(presenceStatus) == "y" {
		return true, nil
	} else if strings.ToLower(presenceStatus) == "n" || presenceStatus == "" {
		return false, nil
	} else {
		return false, errors.New("presence status should be 'y' or 'n'")
	}
}
