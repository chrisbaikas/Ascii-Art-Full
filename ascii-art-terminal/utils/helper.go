package utils

import (
	"fmt"
	"strings"
)

// Help message displayed when usage is incorrect
const UsageMsg = `Usage:
  go run . [--output=banner.txt] [--align=<left|center|right|justify>] <"text"> [banner]

Notes:
  - The text to convert is required.
  - If a banner file is not provided, the default "standard.txt" is used.
  - Optional flags (output and align) must appear in the first positions.
Examples:
  go run . "hello"
  go run . "hello" thinkertoy
  go run . --output=banner.txt "hello" thinkertoy
  go run . --align=center "hello" standard

Additional color usage:
  go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring> [--color=<color> <substring>] [--align=...] [--output=...] "text"

Reverse mode (drops color, outputs raw ASCII art in reverse order of lines):
  go run . --reverse=example04.txt`

// ParseArgs parses CLI arguments and returns all relevant fields
func ParseArgs(args []string) (outputFile, alignType, inputText, bannerFile string, colorTargets []ColorTarget, err error) {
	alignType = "left"
	bannerFile = "standard"
	colorTargets = []ColorTarget{}

	// Parse flags
	for len(args) > 0 && (strings.HasPrefix(args[0], "--output=") ||
		strings.HasPrefix(args[0], "--align=") ||
		strings.HasPrefix(args[0], "--color=")) {

		switch {
		case strings.HasPrefix(args[0], "--output="):
			outputFile = strings.TrimPrefix(args[0], "--output=")
			args = args[1:]

		case strings.HasPrefix(args[0], "--align="):
			alignType = strings.TrimPrefix(args[0], "--align=")
			args = args[1:]

		case strings.HasPrefix(args[0], "--color="):
			colorCode := strings.TrimPrefix(args[0], "--color=")
			args = args[1:]

			substring := ""
			if len(args) > 1 && !strings.HasPrefix(args[0], "--") {
				substring = args[0]
				args = args[1:]
			}
			colorTargets = append(colorTargets, ColorTarget{ColorCode: colorCode, Substring: substring})

		default:
			return "", "", "", "", nil, fmt.Errorf("unrecognized option: %q\n\n%s", args[0], UsageMsg)
		}
	}

	// Validate alignment
	validAligns := map[string]bool{"left": true, "right": true, "center": true, "justify": true}
	if !validAligns[alignType] {
		return "", "", "", "", nil, fmt.Errorf("invalid alignment option: %q\n\n%s", alignType, UsageMsg)
	}

	// Require at least one argument for text
	if len(args) < 1 {
		return "", "", "", "", nil, fmt.Errorf("missing required text argument\n\n%s", UsageMsg)
	}

	text := strings.ReplaceAll(args[0], "\\n", "\n")
	if len(args) >= 2 {
		bannerFile = args[1]
		if !strings.HasSuffix(bannerFile, ".txt") {
			bannerFile += ".txt"
		}
	} else {
		bannerFile += ".txt"
	}

	return outputFile, alignType, text, bannerFile, colorTargets, nil
}
