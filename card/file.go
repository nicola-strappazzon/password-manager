package card

import (
	"github.com/nicola-strappazzon/password-manager/base64"
	"github.com/nicola-strappazzon/password-manager/file"
)

type File struct {
	Name    string
	Content string
}

func (f File) Save(in string) {
	file.Save(in, f.Decode())
}

func (f File) Decode() (out []byte) {
	return base64.Decode(f.Content)
}

func (File) Load(in string) (out File) {
	out.Name = file.Name(in)
	out.Content = base64.Encode(file.ReadInBytes(in))
	return
}

func (f File) Size() (out uint64) {
	return uint64(len(f.Decode()))
}
