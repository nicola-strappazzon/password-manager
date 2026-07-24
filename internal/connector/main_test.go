package connector_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/connector"

	"github.com/stretchr/testify/assert"
)

func TestBuildMySQL(t *testing.T) {
	c := card.Card{}
	c.Database.Engine = "mysql"
	c.Database.Schema = "shop"
	c.Host = "127.0.0.1"
	c.Port = "3306"
	c.Username = "root"
	c.Password = "secret"

	spec, err := connector.Build(c)

	assert.NoError(t, err)
	assert.Equal(t, "mysql", spec.Bin)
	assert.Equal(t, []string{
		"--host=127.0.0.1",
		"--port=3306",
		"--user=root",
		"--database=shop",
	}, spec.Args)
	assert.Equal(t, []string{"MYSQL_PWD=secret"}, spec.Env)
}

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

func TestBuildMariaDBUsesMySQL(t *testing.T) {
	c := card.Card{}
	c.Database.Engine = "MariaDB"

	spec, err := connector.Build(c)

	assert.NoError(t, err)
	assert.Equal(t, "mysql", spec.Bin)
}

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
