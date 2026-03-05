package output

import "github.com/jedib0t/go-pretty/v6/text"

// StatusColor returns an ANSI color for a build/event status string.
func StatusColor(status string) text.Color {
	switch status {
	case "RUNNING":
		return text.FgBlue
	case "SUCCESS":
		return text.FgGreen
	case "FAILURE":
		return text.FgRed
	case "ABORTED":
		return text.FgYellow
	case "QUEUED":
		return text.FgCyan
	case "BLOCKED":
		return text.FgMagenta
	default:
		return text.Reset
	}
}

// ColorizeStatus wraps a status string in ANSI color codes.
func ColorizeStatus(status string, noColor bool) string {
	if noColor {
		return status
	}
	c := StatusColor(status)
	return text.Colors{c}.Sprint(status)
}
