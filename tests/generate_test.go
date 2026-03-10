package integration_test

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/nicola-strappazzon/password-manager/cli/generate"
	"github.com/stretchr/testify/assert"
)

func testGenerate(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		stdout, _, err := run(generate.NewCommand(), []string{})
		assert.NoError(t, err)
		assert.Equal(t, 16, len(strings.TrimSpace(stdout)))
		assert.True(t, strings.ContainsAny(stdout, generate.SPECIAL))
	})

	t.Run("copy", func(t *testing.T) {
		stdout, stderr, err := run(generate.NewCommand(), []string{"-c"})

		if runtime.GOOS == "linux" && os.Getenv("DISPLAY") == "" {
			assert.ErrorContains(t, err, "Clipboard disabled: no X11 display detected")
			assert.Equal(t, "Error: Clipboard disabled: no X11 display detected\n", stderr)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, "Generated password copied to clipboard.\n", stdout)
		}
	})

	t.Run("length-32", func(t *testing.T) {
		stdout, _, err := run(generate.NewCommand(), []string{"-m", "32"})
		assert.NoError(t, err)
		assert.Equal(t, 32, len(strings.TrimSpace(stdout)))
	})

	t.Run("no-symbols", func(t *testing.T) {
		stdout, _, err := run(generate.NewCommand(), []string{"-s"})
		assert.NoError(t, err)
		assert.False(t, strings.ContainsAny(stdout, generate.SPECIAL))
	})

	t.Run("no-numbers", func(t *testing.T) {
		stdout, _, err := run(generate.NewCommand(), []string{"-n"})
		assert.NoError(t, err)
		assert.False(t, strings.ContainsAny(stdout, generate.NUMBERS))
	})

	t.Run("no-lowercase", func(t *testing.T) {
		stdout, _, err := run(generate.NewCommand(), []string{"-l"})
		assert.NoError(t, err)
		assert.False(t, strings.ContainsAny(stdout, generate.LOWERCASE))
	})

	t.Run("no-uppercase", func(t *testing.T) {
		stdout, _, err := run(generate.NewCommand(), []string{"-u"})
		assert.NoError(t, err)
		assert.False(t, strings.ContainsAny(stdout, generate.UPPERCASE))
	})

	t.Run("no-symbols-no-numbers", func(t *testing.T) {
		stdout, _, err := run(generate.NewCommand(), []string{"-s", "-n"})
		assert.NoError(t, err)
		assert.False(t, strings.ContainsAny(stdout, generate.SPECIAL))
		assert.False(t, strings.ContainsAny(stdout, generate.NUMBERS))
	})
}
