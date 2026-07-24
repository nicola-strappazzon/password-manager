package connector

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/nicola-strappazzon/password-manager/internal/card"
)

// Spec describes the client invocation used to connect to a database: the
// binary to run, its arguments and any extra environment variables (used to
// pass the password without exposing it in the process list).
type Spec struct {
	Bin  string
	Args []string
	Env  []string
}

// Build returns the connection Spec for the given card based on its configured
// database engine. It does not check whether the client is installed; use
// Command for that. It fails when the engine is missing or unsupported.
func Build(c card.Card) (Spec, error) {
	engine := strings.ToLower(strings.TrimSpace(c.Database.Engine))

	switch engine {
	case "mysql", "mariadb":
		return mysql(c), nil
	case "postgres", "postgresql", "pgsql", "psql":
		return postgres(c), nil
	case "":
		return Spec{}, errors.New("The database engine is not set.")
	default:
		return Spec{}, fmt.Errorf("Unsupported database engine: %s", c.Database.Engine)
	}
}

// Command builds the client command for the card, validating that the required
// client is installed before returning it ready to run.
func Command(c card.Card) (*exec.Cmd, error) {
	spec, err := Build(c)
	if err != nil {
		return nil, err
	}

	bin, err := exec.LookPath(spec.Bin)
	if err != nil {
		return nil, fmt.Errorf("The %s client is not installed.", spec.Bin)
	}

	cmd := exec.Command(bin, spec.Args...)
	if len(spec.Env) > 0 {
		cmd.Env = append(os.Environ(), spec.Env...)
	}

	return cmd, nil
}

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
