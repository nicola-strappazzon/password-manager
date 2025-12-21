package base64

import (
	"encoding/base64"

	"github.com/nicola-strappazzon/pm/check"
)

func Decode(in string) (out []byte) {
	out, e := base64.StdEncoding.DecodeString(in)
	check.Check(e)
	return
}

func Encode(in []byte) string {
	return base64.StdEncoding.EncodeToString(in)
}
