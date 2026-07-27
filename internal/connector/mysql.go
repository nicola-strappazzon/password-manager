package connector

import "github.com/nicola-strappazzon/password-manager/internal/card"

func mysql(c card.Card) Spec {
	args := []string{}

	if c.Host != "" {
		args = append(args, "--host="+c.Host)
	}

	if c.Port != "" {
		args = append(args, "--port="+c.Port)
	}

	if c.Username != "" {
		args = append(args, "--user="+c.Username)
	}

	if c.Database.Schema != "" {
		args = append(args, "--database="+c.Database.Schema)
	}

	spec := Spec{Bin: "mysql", Args: args}

	if c.Password != "" {
		spec.Env = []string{"MYSQL_PWD=" + c.Password}
	}

	return spec
}
