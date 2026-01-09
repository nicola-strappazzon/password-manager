package path_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nicola-strappazzon/password-manager/config"
	"github.com/nicola-strappazzon/password-manager/path"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	var err error

	testDir, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}

	config.DataDir = ""

	orig := config.UserHomeDir
	defer func() { config.UserHomeDir = orig }()

	config.UserHomeDir = func() (string, error) {
		return testDir, nil
	}

	code := m.Run()

	os.RemoveAll(testDir)
	os.Exit(code)
}

func TestPath(t *testing.T) {
	var p path.Path = "test/foo"
	assert.Equal(t, "test/foo", p.Path())
}

func TestDirectory(t *testing.T) {
	var p path.Path = "test/foo"
	assert.Equal(t, "test", p.Directory())
	p = "test"
	assert.Equal(t, ".", p.Directory())
	p = "test/foo/bas"
	assert.Equal(t, "test/foo", p.Directory())
	p = "test/foo/bas.gpg"
	assert.Equal(t, "test/foo", p.Directory())
}

func TestName(t *testing.T) {
	var p path.Path = "test/foo"
	assert.Equal(t, "foo.gpg", p.Name())
}

func TestBaseName(t *testing.T) {
	var p path.Path = "test/foo"
	assert.Equal(t, "foo", p.BaseName())
}

func TestAbsolute(t *testing.T) {
	var d string = UserHomeDir()
	var p path.Path = "test/foo"

	assert.Equal(t, filepath.Join(d, "test"), p.Absolute())
}

func TestFull(t *testing.T) {
	var d string = UserHomeDir()
	var p path.Path = "test/foo"

	assert.Equal(t, filepath.Join(d, "test/foo.gpg"), p.Full())
}

func TestIsDirectory(t *testing.T) {
	var d string = UserHomeDir()
	var p path.Path = path.Path("test/foo")
	var s string = filepath.Join(d, "test")

	assert.NoError(t, os.Mkdir(s, 0o755))
	assert.True(t, p.IsDirectory())

	f, err := os.CreateTemp(t.TempDir(), "*-foo.gpg")
	assert.NoError(t, err)

	var tf path.Path = path.Path(f.Name())
	assert.False(t, tf.IsDirectory())
}

func TestIsFile(t *testing.T) {
	var p path.Path = "foo"
	assert.False(t, p.IsInvalid())

	p = "test/foo"
	assert.False(t, p.IsInvalid())

	p = "test/foo/"
	assert.False(t, p.IsInvalid())

	p = "test/foo-bar"
	assert.False(t, p.IsInvalid())

	p = "test/foo_bar"
	assert.False(t, p.IsInvalid())

	p = "test/foo.txt"
	assert.True(t, p.IsInvalid())

	p = "test-foo.txt"
	assert.True(t, p.IsInvalid())
}

func UserHomeDir() (out string) {
	out, _ = config.UserHomeDir()
	return out
}
