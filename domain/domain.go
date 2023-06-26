package domain

import (
	"bytes"
	"strings"
	"unicode"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func splitString(sequence string, max_len int) string {

	if len(sequence) < max_len {
		return sequence
	}

	var parts []string

	for i := 0; i < len(sequence); i += min(max_len, len(sequence)-i) {
		var str string = string(sequence[i : i+min(max_len, len(sequence)-i)])
		parts = append(parts, str)
	}

	return strings.Join(parts, "\n")
}

func ProduceAsciiArt(art string, file string) string {

	var buf bytes.Buffer

	// Replacing the "\n" text into an actual new line.
	art = strings.ReplaceAll(art, "\\n", "\n")

	// Splitting the input "string" after the "\n" being converted into an actual new line into multiple lines.
	String := strings.Split(splitString(art, 25), "\n")

	// pick banner type
	Text, _ := Banner.ReadFile(file)

	// Replacing the carriage return into new lines to unify the LF and CRLF.
	UniText := strings.ReplaceAll(string(Text), "\r\n", "\n")

	// Splitting the Text into slice of strings.
	alphabet := strings.Split(UniText, "\n")

	buf.WriteString("\n")
	for _, line := range String { //Going through the string we input.
		line = strings.ReplaceAll(string(line), "\r\n", "\n")
		if len(line) == 0 {
			continue
		}
		for i := 1; i < 9; i++ { //Due to each character in the ASCII table is 8 lines long.
			for _, char := range line { //Going through each characther of the input string.
				if !unicode.IsPrint(char) || char > 0x7E {
					continue
				}
				buf.WriteString(alphabet[(char-32)*9+rune(i)])
			}
			buf.WriteString("\n")
		}
	}

	return buf.String()
}
