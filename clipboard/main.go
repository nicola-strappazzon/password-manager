package clipboard

import (
	"log"

	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	check(err)
}

func Write(in string) {
	clipboard.Write(clipboard.FmtText, []byte(in))
}

func check(in error) {
	if in != nil {
		log.Fatalf(in.Error())
	}
}
