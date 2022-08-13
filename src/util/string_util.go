package util

import (
	"bufio"
	"os"
	"strings"
)

func SplitLines(content os.File) []string {
	var lines []string

	scanner := bufio.NewScanner(&content)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.ReplaceAll(text, ".jpg", "")
		lines = append(lines, text)
	}
	return lines
}
