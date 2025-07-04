package utils

import (
	"strings"
)

// alignRows adjusts the position of ASCII art based on alignment type.
// It no longer returns an error if width < content; it simply uses zero padding.
func alignRows(rows []string, align string, width int) ([]string, error) {
	// Left alignment → no padding
	if align == "left" {
		return rows, nil
	}

	// Measure the visible length (strip ANSI escapes)
	lineLen := visibleLength(rows[0])

	// Calculate pad, but never negative
	pad := 0
	switch align {
	case "right":
		if width > lineLen {
			pad = width - lineLen
		}
	default:
		// Unknown alignment: treat as left
		return rows, nil
	}

	// Prepend spaces to each row
	for i := range rows {
		if pad > 0 {
			rows[i] = strings.Repeat(" ", pad) + rows[i]
		}
	}
	return rows, nil
}
