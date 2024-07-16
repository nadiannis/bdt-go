package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func GetInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func GenerateStatusText(isPresent bool) string {
	if isPresent {
		return "present"
	}
	return "absent"
}
