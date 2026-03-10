package arguments

func First(in []string) string {
	if len(in) >= 1 {
		return in[0]
	}

	return ""
}
