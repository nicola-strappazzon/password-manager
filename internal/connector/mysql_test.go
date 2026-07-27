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

func TestBuildMySQLAliases(t *testing.T) {
	for _, engine := range []string{"mysql", "mariadb"} {
		t.Run(engine, func(t *testing.T) {
			c := card.Card{}
			c.Database.Engine = engine

			spec, err := connector.Build(c)

			assert.NoError(t, err)
			assert.Equal(t, "mysql", spec.Bin)
		})
	}
}
