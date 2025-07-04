package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ColorTarget defines a color and the corresponding substring to color
type ColorTarget struct {
	ColorCode string
	Substring string
}

// parseColorCode parses a color string and returns the corresponding ANSI escape code
func parseColorCode(code string) string {
	switch strings.ToLower(code) {
	case "black":
		return "\033[30m"
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "magenta":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	case "orange":
		return "\033[38;5;208m" // Approximate orange in 256-color
	case "pink":
		return "\033[38;5;205m"
	case "purple":
		return "\033[38;5;93m"
	case "gray", "grey":
		return "\033[38;5;240m"
	case "brown":
		return "\033[38;5;94m"
	}

	// rgb(r, g, b)
	if strings.HasPrefix(code, "rgb(") && strings.HasSuffix(code, ")") {
		body := code[4 : len(code)-1] // trim "rgb(" … ")"
		parts := strings.Split(body, ",")
		if len(parts) == 3 {
			r, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
			g, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
			b, err3 := strconv.Atoi(strings.TrimSpace(parts[2]))
			if err1 == nil && err2 == nil && err3 == nil {
				return rgbToAnsi(r, g, b)
			}
		}
	}

	// #rrggbb
	if strings.HasPrefix(code, "#") && len(code) == 7 {
		r, err1 := strconv.ParseInt(code[1:3], 16, 0)
		g, err2 := strconv.ParseInt(code[3:5], 16, 0)
		b, err3 := strconv.ParseInt(code[5:7], 16, 0)
		if err1 == nil && err2 == nil && err3 == nil {
			return rgbToAnsi(int(r), int(g), int(b))
		}
	}

	// hsl(h, s%, l%)
	if strings.HasPrefix(code, "hsl(") && strings.HasSuffix(code, ")") {
		body := code[4 : len(code)-1] // trim "hsl(" … ")"
		parts := strings.Split(body, ",")
		if len(parts) == 3 {
			trim := func(s string) string {
				s = strings.TrimSpace(s)
				return strings.TrimSuffix(s, "%")
			}
			h, err1 := strconv.Atoi(trim(parts[0]))
			sv, err2 := strconv.Atoi(trim(parts[1]))
			l, err3 := strconv.Atoi(trim(parts[2]))
			if err1 == nil && err2 == nil && err3 == nil {
				r, g, b := hslToRgb(float64(h), float64(sv)/100, float64(l)/100)
				return rgbToAnsi(r, g, b)
			}
		}
	}

	// Unsupported / invalid → empty string
	return ""
}

// rgbToAnsi converts RGB values to the closest ANSI 256‑color code
func rgbToAnsi(r, g, b int) string {
	r6 := r * 6 / 256
	g6 := g * 6 / 256
	b6 := b * 6 / 256
	code := 16 + (36 * r6) + (6 * g6) + b6
	return "\033[38;5;" + strconv.Itoa(code) + "m"
}

// hslToRgb converts HSL color values to RGB
func hslToRgb(h, s, l float64) (int, int, int) {
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var r1, g1, b1 float64
	switch {
	case h < 60:
		r1, g1, b1 = c, x, 0
	case h < 120:
		r1, g1, b1 = x, c, 0
	case h < 180:
		r1, g1, b1 = 0, c, x
	case h < 240:
		r1, g1, b1 = 0, x, c
	case h < 300:
		r1, g1, b1 = x, 0, c
	default:
		r1, g1, b1 = c, 0, x
	}

	return int((r1+m)*255 + 0.5), int((g1+m)*255 + 0.5), int((b1+m)*255 + 0.5)
}

// buildAsciiRowsWithColor renders a line, colouring any configured substrings.
// It relies exclusively on buildAsciiRows for glyph rendering to avoid duplicated logic.
func buildAsciiRowsWithColor(line string, banner map[rune][]string, colorTargets []ColorTarget) ([]string, error) {
	// Fast path: only one target and paints the entire line.
	if len(colorTargets) == 1 && colorTargets[0].Substring == "" {
		rows, err := buildAsciiRows(line, banner)
		if err != nil {
			return nil, err
		}
		cc := parseColorCode(colorTargets[0].ColorCode)
		for i := range rows {
			rows[i] = cc + rows[i] + "\033[0m"
		}
		return rows, nil
	}

	// Prepare ANSI code for red warnings
	redWarn := parseColorCode("red")
	reset := "\033[0m"

	// 1) Handle default-only flags (Substring==""): last one wins
	var defaultCC string
	defaultCount := 0
	for _, t := range colorTargets {
		if t.Substring == "" {
			defaultCC = parseColorCode(t.ColorCode)
			defaultCount++
		}
	}
	if defaultCount > 1 {
		fmt.Fprintf(os.Stderr,
			"%swarning: %d default colors specified, using the last one%s\n",
			redWarn, defaultCount, reset)
	}

	// 2) Handle duplicate substring targets: keep only the last for each substring
	lastIdx := make(map[string]int)
	dupCount := make(map[string]int)
	for idx, t := range colorTargets {
		if t.Substring != "" {
			lastIdx[t.Substring] = idx
			dupCount[t.Substring]++
		}
	}
	for substr, cnt := range dupCount {
		if cnt > 1 {
			fmt.Fprintf(os.Stderr,
				"%swarning: %d color rules for substring %q; using the last one%s\n",
				redWarn, cnt, substr, reset)
		}
	}
	var substrTargets []ColorTarget
	for idx, t := range colorTargets {
		if t.Substring == "" {
			continue
		}
		if lastIdx[t.Substring] == idx {
			substrTargets = append(substrTargets, t)
		}
	}

	// 3) Main loop: color substrings first, else default, else plain
	out := make([]string, blockLines)
	i := 0
	for i < len(line) {
		matched := false

		for _, t := range substrTargets {
			if strings.HasPrefix(line[i:], t.Substring) {
				segRows, err := buildAsciiRows(t.Substring, banner)
				if err != nil {
					return nil, err
				}
				cc := parseColorCode(t.ColorCode)
				for r := range segRows {
					out[r] += cc + segRows[r] + "\033[0m"
				}
				i += len(t.Substring)
				matched = true
				break
			}
		}
		if matched {
			continue
		}

		ch := string(line[i])
		singleRows, err := buildAsciiRows(ch, banner)
		if err != nil {
			return nil, err
		}
		if defaultCC != "" {
			for r := range singleRows {
				out[r] += defaultCC + singleRows[r] + "\033[0m"
			}
		} else {
			for r := range singleRows {
				out[r] += singleRows[r]
			}
		}
		i++
	}
	return out, nil
}

func visibleLength(s string) int {
	return len(StripANSI(s))
}
