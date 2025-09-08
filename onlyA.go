package student

func OnlyA(s string) bool {
	for _, r := range s {
		if r != 'a' {
			return false
		}
	}
	return true
}
