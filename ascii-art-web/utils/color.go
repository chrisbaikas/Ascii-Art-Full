package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type ColorTarget struct {
	ColorCode string // hex format, e.g. "#ff0000"
	Substring string // substring to apply the color to ("" = global color)
}

// Converts a hex color "#rrggbb" into an ANSI 256-color escape code
func parseColorCode(code string) string {
	if strings.HasPrefix(code, "#") && len(code) == 7 {
		r, err1 := strconv.ParseInt(code[1:3], 16, 0)
		g, err2 := strconv.ParseInt(code[3:5], 16, 0)
		b, err3 := strconv.ParseInt(code[5:7], 16, 0)
		if err1 == nil && err2 == nil && err3 == nil {
			return rgbToAnsi(int(r), int(g), int(b))
		}
	}
	return ""
}

// Converts RGB values to ANSI 256-color escape sequence
func rgbToAnsi(r, g, b int) string {
	r6 := r * 6 / 256
	g6 := g * 6 / 256
	b6 := b * 6 / 256
	code := 16 + (36 * r6) + (6 * g6) + b6
	return "\033[38;5;" + strconv.Itoa(code) + "m"
}

// Converts ANSI 256-color escape sequences into <span style="color:#hex">
var ansi256Pattern = regexp.MustCompile(`\x1b\[38;5;(\d{1,3})m`)

func AnsiToHTML256(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input string is empty")
	}

	defer func() {
		if r := recover(); r != nil {
			// This catches any panic from regex or string ops
			log.Printf("panic in AnsiToHTML256: %v", r)
		}
	}()

	result := ansi256Pattern.ReplaceAllStringFunc(input, func(match string) string {
		matches := ansi256Pattern.FindStringSubmatch(match)
		if len(matches) != 2 {
			log.Printf("invalid ANSI match: %q", match)
			return "" // fail silently or return safe fallback
		}
		hex := ansi256ToHex(matches[1])
		return `<span style="color:` + hex + `">`
	})

	result = strings.ReplaceAll(result, "\033[0m", "</span>")
	return result, nil
}

// Converts an ANSI 256-color index to its closest web-safe hex color
func ansi256ToHex(code string) string {
	n, err := strconv.Atoi(code)
	if err != nil || n < 16 || n > 255 {
		return "#000000" // fallback to black
	}
	c := n - 16
	r := (c / 36) % 6 * 51
	g := (c / 6 % 6) * 51
	b := (c % 6) * 51
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// Removes ANSI escape codes from the string
func stripANSICodes(s string) string {
	result := ""
	inEscape := false
	for i := 0; i < len(s); i++ {
		if s[i] == '\033' {
			inEscape = true
			continue
		}
		if inEscape {
			if s[i] == 'm' {
				inEscape = false
			}
			continue
		}
		result += string(s[i])
	}
	return result
}

// Returns the visible length of a string (ignoring ANSI codes)
func visibleLength(s string) int {
	return len(stripANSICodes(s))
}

// Builds ASCII art rows for a line of text and applies coloring rules
func buildAsciiRowsWithColor(line string, banner BannerType, colorTargets []ColorTarget) ([]string, error) {
	var globalColor string
	for _, t := range colorTargets {
		if t.Substring == "" {
			globalColor = parseColorCode(t.ColorCode)
			break
		}
	}

	out := make([]string, blockLines)
	i := 0

	for i < len(line) {
		matched := false
		for _, t := range colorTargets {
			if t.Substring == "" {
				continue
			}
			if strings.HasPrefix(line[i:], t.Substring) {
				segRows, err := buildAsciiRows(t.Substring, banner)
				if err != nil {
					return nil, err
				}
				colorCode := parseColorCode(t.ColorCode)
				for r := 0; r < blockLines; r++ {
					out[r] += colorCode + segRows[r] + "\033[0m"
				}
				i += len(t.Substring)
				matched = true
				break
			}
		}

		if matched {
			continue
		}

		rows, err := buildAsciiRows(string(line[i]), banner)
		if err != nil {
			return nil, err
		}
		for r := 0; r < blockLines; r++ {
			if globalColor != "" {
				out[r] += globalColor + rows[r] + "\033[0m"
			} else {
				out[r] += rows[r]
			}
		}
		i++
	}

	return out, nil
}
