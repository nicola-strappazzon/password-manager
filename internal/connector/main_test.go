package connector_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/connector"

	"github.com/stretchr/testify/assert"
)

func TestBuildWithoutPasswordHasNoEnv(t *testing.T) {
	c := card.Card{}
	c.Database.Engine = "mysql"

	spec, err := connector.Build(c)

	assert.NoError(t, err)
	assert.Empty(t, spec.Env)
}

func TestBuildMissingEngine(t *testing.T) {
	spec, err := connector.Build(card.Card{})

	assert.Error(t, err)
	assert.Empty(t, spec.Bin)
}

func TestBuildUnsupportedEngine(t *testing.T) {
	c := card.Card{}
	c.Database.Engine = "oracle"

	_, err := connector.Build(c)

	assert.Error(t, err)
}

func TestSpecString(t *testing.T) {
	spec := connector.Spec{
		Bin: "mysql",
		Args: []string{
			"--host=127.0.0.1",
			"--user=app user",
			"--database=team's db",
		},
	}

	assert.Equal(t, "mysql --host=127.0.0.1 '--user=app user' '--database=team'\"'\"'s db'", spec.String())
}

func TestStringForPrintIncludesMySQLPassword(t *testing.T) {
	c := card.Card{}
	c.Database.Engine = "mysql"
	c.Host = "127.0.0.1"
	c.Username = "root"
	c.Password = "secret value"

	out, err := connector.StringForPrint(c)

	assert.NoError(t, err)
	assert.Equal(t, "mysql --host=127.0.0.1 --user=root '-psecret value'", out)
}
