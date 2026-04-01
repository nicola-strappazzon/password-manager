package cli

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestPersistentPreRunE(t *testing.T) {
	t.Setenv("PM_RECIPIENT", "")

	t.Run("allows version without setup", func(t *testing.T) {
		cmd := &cobra.Command{Use: "version"}
		err := PersistentPreRunE(cmd, []string{})
		assert.NoError(t, err)
	})

	t.Run("requires setup for other commands", func(t *testing.T) {
		cmd := &cobra.Command{Use: "ls"}
		err := PersistentPreRunE(cmd, []string{})
		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "Run 'pm setup'"))
	})
}
