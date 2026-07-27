package connector

import "github.com/nicola-strappazzon/password-manager/internal/card"

func postgres(c card.Card) Spec {
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

	if c.Database.Schema != "" {
		args = append(args, "--dbname="+c.Database.Schema)
	}

	spec := Spec{Bin: "psql", Args: args}

	if c.Password != "" {
		spec.Env = []string{"PGPASSWORD=" + c.Password}
	}

	return spec
}
