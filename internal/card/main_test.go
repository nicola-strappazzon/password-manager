package card_test

import (
	"testing"

	"github.com/nicola-strappazzon/password-manager/internal/card"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	yml := `---
password: "#O123456@bc"
`

	assert.Equal(t, card.New(yml).Password, "#O123456@bc")
}

func TestGetValue(t *testing.T) {
	c := card.Card{}
	c.Password = "#O123456@bc"
	c.Phone = "+34123456789"
	c.PUK = "12345678"
	c.Username = "foo"
	c.Notes = "This is a test note."
	c.AWS.Region = "baz-1"

	assert.Equal(t, c.GetValue("password"), "#O123456@bc")
	assert.Equal(t, c.GetValue("phone"), "+34123456789")
	assert.Equal(t, c.GetValue("puk"), "12345678")
	assert.Equal(t, c.GetValue("username"), "foo")
	assert.Equal(t, c.GetValue("notes"), "This is a test note.")
	assert.Equal(t, c.GetValue("aws.region"), "baz-1")
}

func TestSetValue(t *testing.T) {
	c := card.Card{}
	c.SetValue("phone", "+34123456789")
	c.SetValue("puk", "12345678")
	c.SetValue("username", "foo")

	assert.Equal(t, c.GetValue("phone"), "+34123456789")
	assert.Equal(t, c.GetValue("puk"), "12345678")
	assert.Equal(t, c.GetValue("username"), "foo")
}
