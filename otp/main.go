package otp

import (
	"time"

	"github.com/nicola-strappazzon/password-manager/check"

	"github.com/pquerna/otp/totp"
)

func Get(in string) string {
	code, err := totp.GenerateCode(in, time.Now())
	check.Check(err)

	return code
}
