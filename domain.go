package domain

import (
	"bytes"
	"os"
	"strings"
)

func ProduceAsciiArt(art string, file string) string {
	var buf bytes.Buffer
	// Replacing the "\n" text into an actual new line.
	newLine := strings.ReplaceAll(art, "\\n", "\n")

	// Splitting the input "string" after the "\n" being converted into an actual new line into multiple lines.
	String := strings.Split(newLine, "\n")

	// Removing the empty string from the end starting a new line.
	for ; String[len(String)-1] == ""; String = String[:len(String)-1] {
		defer os.Stdout.WriteString("\n")
	}

	// Reading the text from the "standard" file.

	// switch style
	Text, _ := os.ReadFile(file)

	// Replacing the carriage return into new lines to unify the LF and CRLF.
	UniText := strings.ReplaceAll(string(Text), "\r\n", "\n")

	// Splitting the Text into slice of strings.
	alphabet := strings.Split(UniText, "\n")

	buf.WriteString("\n")
	for _, line := range String { //Going through the string we input.
		for i := 1; i < 9; i++ { //Due to each character in the ASCII table is 8 lines long.
			for _, char := range line { //Going through each characther of the input string.
				buf.WriteString(alphabet[(char-32)*9+rune(i)])
			}
			buf.WriteString("\n")
		}
	}

	return buf.String()
}
