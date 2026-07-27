package connector_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/connector"

	"github.com/stretchr/testify/assert"
)

func TestBuildPostgreSQL(t *testing.T) {
	c := card.Card{}
	c.Database.Engine = "postgresql"
	c.Database.Schema = "shop"
	c.Host = "127.0.0.1"
	c.Port = "5432"
	c.Username = "postgres"
	c.Password = "secret"

	spec, err := connector.Build(c)

	assert.NoError(t, err)
	assert.Equal(t, "psql", spec.Bin)
	assert.Equal(t, []string{
		"--host=127.0.0.1",
		"--port=5432",
		"--username=postgres",
		"--dbname=shop",
	}, spec.Args)
	assert.Equal(t, []string{"PGPASSWORD=secret"}, spec.Env)
}

func TestBuildPostgreSQLAliases(t *testing.T) {
	for _, engine := range []string{"postgres", "postgresql", "pgsql", "psql"} {
		t.Run(engine, func(t *testing.T) {
			c := card.Card{}
			c.Database.Engine = engine

			spec, err := connector.Build(c)

			assert.NoError(t, err)
			assert.Equal(t, "psql", spec.Bin)
		})
	}
}
