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
