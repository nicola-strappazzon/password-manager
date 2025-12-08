package card

import (
	"strings"
)

type Body string

func (b Body) ToString() string {
	return string(b)
}

func (b Body) ReadByLine(yield func(string) bool) {
	for line := range strings.Lines(b.ToString()) {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if yield(line) {
			return
		}
	}
}
