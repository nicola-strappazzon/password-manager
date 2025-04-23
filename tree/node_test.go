package tree

import (
	"testing"

	"github.com/nicola-strappazzon/pm/tree"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	node := tree.Node{
		Name: "Password Store",
	}
	node.Add("foo")
	node.Get("foo").Add("foo01")
	node.Get("foo").Add("foo02")
	node.Get("foo").Add("foo03")
	node.Get("foo").Get("foo03").Add("foo0301")
	node.Get("foo").Get("foo03").Add("foo0302")
	node.Add("bar")
	node.Add("baz")
	node.Get("baz").Add("baz01")
	node.Get("baz").Get("baz01").Add("baz0101")
	node.Get("baz").Get("baz01").Add("baz0102")
	node.Print()

	assert.True(t, true)
}
