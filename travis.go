package travis

func Bool(value bool) *bool {
	b := new(bool)
	*b = value
	return b
}

func String(value string) *string {
	b := new(string)
	*b = value
	return b
}
