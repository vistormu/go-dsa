package ansi

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// Cursor movement
	Home  = "\x1b[H"
	Up    = "\x1b[A"
	Down  = "\x1b[B"
	Right = "\x1b[C"
	Left  = "\x1b[D"
	Start = "\r"

	// Screen clearing
	ScreenEnd   = "\x1b[0J"
	ScreenStart = "\x1b[1J"
	Screen      = "\x1b[2J"
	LineEnd     = "\x1b[0K"
	LineStart   = "\x1b[1K"
	Line        = "\x1b[2K"

	// Text styles
	Bold      = "\x1b[1m"
	Dim       = "\x1b[2m"
	Italic    = "\x1b[3m"
	Underline = "\x1b[4m"
	Blink     = "\x1b[5m"
	Reverse   = "\x1b[7m"
	Hidden    = "\x1b[8m"
	Strike    = "\x1b[9m"

	// Text colors
	Reset    = "\x1b[0m"
	Black    = "\x1b[30m"
	Red      = "\x1b[31m"
	Green    = "\x1b[32m"
	Yellow   = "\x1b[33m"
	Blue     = "\x1b[34m"
	Magenta  = "\x1b[35m"
	Cyan     = "\x1b[36m"
	White    = "\x1b[37m"
	Black2   = "\x1b[90m"
	Red2     = "\x1b[91m"
	Green2   = "\x1b[92m"
	Yellow2  = "\x1b[93m"
	Blue2    = "\x1b[94m"
	Magenta2 = "\x1b[95m"
	Cyan2    = "\x1b[96m"
	White2   = "\x1b[97m"

	// Background colors
	BgBlack    = "\x1b[40m"
	BgRed      = "\x1b[41m"
	BgGreen    = "\x1b[42m"
	BgYellow   = "\x1b[43m"
	BgBlue     = "\x1b[44m"
	BgMagenta  = "\x1b[45m"
	BgCyan     = "\x1b[46m"
	BgWhite    = "\x1b[47m"
	BgBlack2   = "\x1b[100m"
	BgRed2     = "\x1b[101m"
	BgGreen2   = "\x1b[102m"
	BgYellow2  = "\x1b[103m"
	BgBlue2    = "\x1b[104m"
	BgMagenta2 = "\x1b[105m"
	BgCyan2    = "\x1b[106m"
	BgWhite2   = "\x1b[107m"
)

func Hex(code string) string {
	// Remove leading "#" if present
	if strings.HasPrefix(code, "#") {
		code = code[1:]
	}
	if len(code) != 6 {
		// Fallback: return reset code for invalid hex
		return Reset
	}
	r, err1 := strconv.ParseUint(code[0:2], 16, 8)
	g, err2 := strconv.ParseUint(code[2:4], 16, 8)
	b, err3 := strconv.ParseUint(code[4:6], 16, 8)
	if err1 != nil || err2 != nil || err3 != nil {
		return Reset
	}
	return Rgb(int(r), int(g), int(b))
}

func Rgb(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func BgHex(code string) string {
	if strings.HasPrefix(code, "#") {
		code = code[1:]
	}
	if len(code) != 6 {
		return Reset
	}
	r, err1 := strconv.ParseUint(code[0:2], 16, 8)
	g, err2 := strconv.ParseUint(code[2:4], 16, 8)
	b, err3 := strconv.ParseUint(code[4:6], 16, 8)
	if err1 != nil || err2 != nil || err3 != nil {
		return Reset
	}
	return BgRgb(int(r), int(g), int(b))
}

func BgRgb(r, g, b int) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b)
}
