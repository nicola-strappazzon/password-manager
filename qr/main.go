package qr

import (
	"os"

	"github.com/mdp/qrterminal/v3"
)

func Generate(in string) {
	config := qrterminal.Config{
		HalfBlocks: true,
		Level:      qrterminal.M,
		Writer:     os.Stdout,
	}

	qrterminal.GenerateWithConfig(in, config)
}
