package main
func boundCheck(length int, lbound int, ubound int) bool {
	return (lbound < ubound && lbound>=0 && ubound<length)
}