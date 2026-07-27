package connector

import "github.com/nicola-strappazzon/password-manager/internal/card"

func mongo(c card.Card) Spec {
	args := []string{}

	if c.Host != "" {
		args = append(args, "--host="+c.Host)
	}

	if c.Port != "" {
		args = append(args, "--port="+c.Port)
	}

	if c.Username != "" {
		args = append(args, "--username="+c.Username)
	}

	if c.Password != "" {
		args = append(args, "--password="+c.Password)
	}

	if c.Database.Schema != "" {
		args = append(args, c.Database.Schema)
	}

	return Spec{Bin: "mongosh", Args: args}
}
