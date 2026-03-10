package check

import (
	"log"
)

func Check(in error) {
	if in != nil {
		log.Fatal(in.Error())
	}
}
