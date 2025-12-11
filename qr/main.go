package qr

import (
	"os"

	"github.com/mdp/qrterminal/v3"
)

func Generate(in string) {
	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}

	qrterminal.GenerateWithConfig(in, config)
}
