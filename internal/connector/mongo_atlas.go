package connector

import "github.com/nicola-strappazzon/password-manager/internal/card"

func mongoAtlas(c card.Card) Spec {
	args := []string{}

	if c.Host != "" {
		args = append(args, "mongodb+srv://"+c.Host+"/")
	}

	args = append(args, "--apiVersion=1")

	if c.Username != "" {
		args = append(args, "--username="+c.Username)
	}

	if c.Password != "" {
		args = append(args, "--password="+c.Password)
	}

	return Spec{Bin: "mongosh", Args: args}
}
