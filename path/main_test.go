package path_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nicola-strappazzon/password-manager/path"

	"github.com/stretchr/testify/assert"
)

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
	var d string = t.TempDir()
	var p path.Path = "test/foo"

	t.Setenv("PM_PATH", d)

	assert.Equal(t, filepath.Join(d, "test"), p.Absolute())
}

func TestFull(t *testing.T) {
	var d string = t.TempDir()
	var p path.Path = "test/foo"
	var s string = filepath.Join(d, "test")

	t.Setenv("PM_PATH", d)

	assert.NoError(t, os.Mkdir(s, 0o755))
	assert.Equal(t, filepath.Join(d, "test/foo.gpg"), p.Full())
}

func TestIsDirectory(t *testing.T) {
	var d string = t.TempDir()
	var s string = filepath.Join(d, "test")
	var p path.Path = path.Path("test/foo")

	t.Setenv("PM_PATH", d)

	assert.NoError(t, os.Mkdir(s, 0o755))
	assert.True(t, p.IsDirectory())

	f, err := os.CreateTemp(t.TempDir(), "*-foo.gpg")
	assert.NoError(t, err)
	// t.Log(f.Name())

	var tf path.Path = path.Path(f.Name())
	// t.Log(tf.Directory())
	assert.False(t, tf.IsDirectory())
}

func TestIsFile(t *testing.T) {
	// var d string = t.TempDir()
	// var s string = filepath.Join(d, "test")
	// var p path.Path = path.Path("foo")

	// t.Setenv("PM_PATH", d)
	// t.Log(d)
	// t.Log(p.Directory())
	// assert.NoError(t, os.Mkdir(s, 0o755))
	// assert.False(t, p.IsDirectory())

	// f, err := os.CreateTemp(t.TempDir(), "foo.gpg")
	// assert.NoError(t, err)
	// t.Log(f.Name())

	// var tf path.Path = path.Path(f.Name())
	// t.Log(tf.Full())
	// assert.True(t, tf.IsFile())
}
