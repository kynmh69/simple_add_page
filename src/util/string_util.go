package util

import (
	"bufio"
	"os"
)

func SplitLines(content os.File) []string {
	var lines []string

	scanner := bufio.NewScanner(&content)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
