package card_test

import (
	"testing"

	"github.com/nicola-strappazzon/pm/card"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	var f card.Files
	f.Add(card.File{
		Name: "foo.txt",
	})
	assert.Equal(t, len(f), 1)
}

func TestCount(t *testing.T) {
	var f card.Files
	assert.Equal(t, f.Count(), 0)
	f.Add(card.File{
		Name: "foo.txt",
	})
	assert.Equal(t, f.Count(), 1)
}

func TestDelete(t *testing.T) {
	var f card.Files
	f.Add(card.File{
		Name: "foo.txt",
	})
	assert.Equal(t, f.Count(), 1)
	f.Delete(card.File{
		Name: "bar.txt",
	})
	assert.Equal(t, f.Count(), 1)
	f.Delete(card.File{
		Name: "foo.txt",
	})
	assert.Equal(t, f.Count(), 0)
}

func TestExist(t *testing.T) {
	var f card.Files
	f.Add(card.File{
		Name: "foo.txt",
	})
	assert.Equal(t, f.Count(), 1)
	assert.False(t, f.Exist(card.File{Name: "bar.txt"}))
	assert.True(t, f.Exist(card.File{Name: "foo.txt"}))
}
