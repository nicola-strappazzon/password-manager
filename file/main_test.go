package file_test

import (
	"path/filepath"
	"testing"

	"github.com/nicola-strappazzon/pm/file"

	"github.com/stretchr/testify/assert"
)

func TestIsDir(t *testing.T) {
	dirTest := t.TempDir()
	fileTest := filepath.Join(dirTest, "test.txt")

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "existing directory",
			path: dirTest,
			want: true,
		},
		{
			name: "existing file",
			path: fileTest,
			want: false,
		},
		{
			name: "non existing path",
			path: filepath.Join(dirTest, "does-not-exist"),
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, file.IsDir(test.path), test.want)
		})
	}
}
