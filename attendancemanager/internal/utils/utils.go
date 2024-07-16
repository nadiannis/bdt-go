package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func getInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
