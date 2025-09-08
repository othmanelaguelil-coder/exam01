package student

func OnlyZ(s string) bool {
	for _, r := range s {
		if r != 'z' {
			return false
		}
	}
	return true
}
