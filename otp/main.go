package otp

import (
	"log"
	"time"

	"github.com/pquerna/otp/totp"
)

func Get(in string) string {
	code, err := totp.GenerateCode(in, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	return code
}
