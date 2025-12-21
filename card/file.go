package card

import (
	"os"
	"path/filepath"

	"github.com/nicola-strappazzon/pm/base64"
	"github.com/nicola-strappazzon/pm/check"
	"github.com/nicola-strappazzon/pm/file"
)

type File struct {
	Name    string
	Content string
}

func (f File) Save(in string) {
	os.MkdirAll(filepath.Dir(in), 0700)
	check.Check(os.WriteFile(in, f.Decode(), 0600))
}

func (f File) Decode() (out []byte) {
	return base64.Decode(f.Content)
}

func (File) Load(in string) (out File) {
	out.Name = file.Name(in)
	out.Content = base64.Encode(file.ReadInBytes(in))
	return
}
