package card

type Files []File

func (f *Files) Add(in File) {
	*f = append(*f, in)
}

func (f *Files) Delete(in File) {
	for i := (*f).Count() - 1; i >= 0; i-- {
		if (*f)[i].Name == in.Name {
			(*f) = append((*f)[:i], (*f)[i+1:]...)
			break
		}
	}
}

func (f Files) Exist(in File) bool {
	for i := f.Count() - 1; i >= 0; i-- {
		if f[i].Name == in.Name {
			return true
		}
	}
	return false
}

func (f *Files) Count() int {
	return len(*f)
}

func (f Files) Get(in File) File {
	for i := f.Count() - 1; i >= 0; i-- {
		if f[i].Name == in.Name {
			return f[i]
		}
	}
	return File{}
}
