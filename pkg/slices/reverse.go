package slices

func ReverseCopy[E any](slice []E) []E {
	l := len(slice)
	cp := make([]E, l)
	for i, v := range slice {
		cp[l-i-1] = v
	}
	return cp
}
