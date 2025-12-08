package card_test

import (
	"testing"

	"github.com/nicola-strappazzon/pm/card"
	"github.com/stretchr/testify/assert"
)

func TestItemParser(t *testing.T) {
	assert.Equal(t, "", (&card.Item{}).Parser("").Key)
	assert.Equal(t, "", (&card.Item{}).Parser("").Value)

	assert.Equal(t, "key", (&card.Item{}).Parser("key: value").Key)
	assert.Equal(t, "value", (&card.Item{}).Parser("key: value").Value)

	assert.Equal(t, "key", (&card.Item{}).Parser(" key: value").Key)
	assert.Equal(t, "value", (&card.Item{}).Parser("key:value").Value)

	assert.Equal(t, "", (&card.Item{}).Parser(" : ").Key)
	assert.Equal(t, "", (&card.Item{}).Parser(" : ").Value)
}
