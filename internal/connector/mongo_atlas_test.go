package connector_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/internal/card"
	"github.com/nicola-strappazzon/password-manager/internal/connector"

	"github.com/stretchr/testify/assert"
)

func TestBuildMongoDBAtlas(t *testing.T) {
	c := card.Card{}
	c.Database.Engine = "mongodb-atlas"
	c.Host = "test-mongo.example.net"
	c.Username = "test_user"
	c.Password = "secret"

	spec, err := connector.Build(c)

	assert.NoError(t, err)
	assert.Equal(t, "mongosh", spec.Bin)
	assert.Equal(t, []string{
		"mongodb+srv://test-mongo.example.net/",
		"--apiVersion=1",
		"--username=test_user",
		"--password=secret",
	}, spec.Args)
	assert.Empty(t, spec.Env)
}

func TestBuildMongoDBAtlasAliases(t *testing.T) {
	for _, engine := range []string{"mongodb-atlas", "mongo-atlas"} {
		t.Run(engine, func(t *testing.T) {
			c := card.Card{}
			c.Database.Engine = engine

			spec, err := connector.Build(c)

			assert.NoError(t, err)
			assert.Equal(t, "mongosh", spec.Bin)
			assert.Contains(t, spec.Args, "--apiVersion=1")
		})
	}
}
