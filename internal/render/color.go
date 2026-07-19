// Copyright (C) 2026 4rat
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package render

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// ANSI escape code constants.
const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Orange  = "\033[38;5;208m"

	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
)

// NoColorMode disables all ANSI output when true.
var NoColorMode bool

// C returns the ANSI escape sequence for the named color, or "" if NoColorMode.
func C(name string) string {
	if NoColorMode {
		return ""
	}
	switch strings.ToLower(name) {
	case "red":
		return BrightRed
	case "green":
		return BrightGreen
	case "yellow":
		return BrightYellow
	case "blue":
		return BrightBlue
	case "magenta":
		return BrightMagenta
	case "cyan":
		return BrightCyan
	case "white":
		return BrightWhite
	case "orange":
		return Orange
	case "bold":
		return Bold
	case "reset":
		return Reset
	}
	return ""
}

// R returns the reset sequence, or "" if NoColorMode.
func R() string {
	if NoColorMode {
		return ""
	}
	return Reset
}

// Wrap wraps text in a named color and resets afterward.
func Wrap(color, text string) string {
	if NoColorMode {
		return text
	}
	return C(color) + text + R()
}

// BoldWrap wraps text in bold + named color.
func BoldWrap(color, text string) string {
	if NoColorMode {
		return text
	}
	return Bold + C(color) + text + R()
}

// ansiRegexp matches ANSI escape sequences for stripping.
var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*[mK]`)

// StripANSI removes all ANSI escape sequences from s.
func StripANSI(s string) string {
	return ansiRegexp.ReplaceAllString(s, "")
}

// VisibleLen returns the number of visible (non-ANSI) rune columns in s.
func VisibleLen(s string) int {
	return utf8.RuneCountInString(StripANSI(s))
}

// PadRight pads s with spaces on the right so its visible width equals width.
func PadRight(s string, width int) string {
	n := VisibleLen(s)
	if n >= width {
		return s
	}
	return s + strings.Repeat(" ", width-n)
}
