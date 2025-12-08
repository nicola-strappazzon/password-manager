package card

import (
	"strings"
)

const FieldSeparator = ":"

type Item struct {
	Key   string
	Value string
}

func (i *Item) Parser(in string) *Item {
	tmp := strings.Split(in, FieldSeparator)

	if len(tmp) != 2 {
		return &Item{}
	}

	i.Key = strings.TrimSpace(tmp[0])
	i.Value = strings.TrimSpace(tmp[1])

	return i
}
