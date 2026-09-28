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

// String returns a shell-readable representation of the client invocation.
func (s Spec) String() string {
	parts := append([]string{s.Bin}, s.Args...)
	for i, part := range parts {
		parts[i] = shellQuote(part)
	}

	return strings.Join(parts, " ")
}

// StringForPrint returns the command shown to users for copy/paste.
func StringForPrint(c card.Card) (string, error) {
	spec, err := Build(c)
	if err != nil {
		return "", err
	}

	engine := strings.ToLower(strings.TrimSpace(c.Database.Engine))
	if c.Password != "" && (engine == "mysql" || engine == "mariadb") {
		spec.Args = append(spec.Args, "-p"+c.Password)
	}

	return spec.String(), nil
}

// Build returns the connection Spec for the given card based on its configured
// database engine. It does not check whether the client is installed; use
// Command for that. It fails when the engine is missing or unsupported.
func Build(c card.Card) (Spec, error) {
	engine := strings.ToLower(strings.TrimSpace(c.Database.Engine))

	switch engine {
	case "mysql", "mariadb":
		return mysql(c), nil
	case "mongodb-atlas", "mongo-atlas":
		return mongoAtlas(c), nil
	case "mongodb", "mongo", "docdb", "documentdb":
		return mongo(c), nil
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

func shellQuote(s string) string {
	if s == "" {
		return "''"
	}

	if strings.IndexFunc(s, func(r rune) bool {
		return !(r >= 'a' && r <= 'z') &&
			!(r >= 'A' && r <= 'Z') &&
			!(r >= '0' && r <= '9') &&
			!strings.ContainsRune("@%_+=:,./-", r)
	}) == -1 {
		return s
	}

	return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'"
}
