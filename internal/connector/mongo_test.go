package connector_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/connector"

	"github.com/stretchr/testify/assert"
)

func TestBuildMongoDB(t *testing.T) {
	c := card.Card{}
	c.Database.Engine = "mongodb"
	c.Database.Schema = "shop"
	c.Host = "127.0.0.1"
	c.Port = "27017"
	c.Username = "root"
	c.Password = "secret"

	spec, err := connector.Build(c)

	assert.NoError(t, err)
	assert.Equal(t, "mongosh", spec.Bin)
	assert.Equal(t, []string{
		"--host=127.0.0.1",
		"--port=27017",
		"--username=root",
		"--password=secret",
		"shop",
	}, spec.Args)
	assert.Empty(t, spec.Env)
}

func TestBuildMongoAliases(t *testing.T) {
	for _, engine := range []string{"mongodb", "mongo", "docdb", "documentdb"} {
		t.Run(engine, func(t *testing.T) {
			c := card.Card{}
			c.Database.Engine = engine

			spec, err := connector.Build(c)

			assert.NoError(t, err)
			assert.Equal(t, "mongosh", spec.Bin)
		})
	}
}
