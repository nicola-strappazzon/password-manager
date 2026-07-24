package term

// ANSI color helpers. Each function wraps the given text with an SGR escape
// code and a reset, so callers never deal with raw escape sequences. Colors use
// the 256-color palette (Nord-inspired: cohesive cool tones with muted labels).

const colorReset = "\033[0m"

func wrap(code, text string) string {
	return "\033[" + code + "m" + text + colorReset
}

// Warning returns the text in bold muted gold, used for warnings.
func Warning(text string) string {
	return wrap("1;38;5;179", text)
}

// Notice returns the text in steel blue, used for informational notices.
func Notice(text string) string {
	return wrap("38;5;110", text)
}

// Label returns the text in muted gray-cyan, used for field labels so the value
// stays the visual focus.
func Label(text string) string {
	return wrap("38;5;109", text)
}
