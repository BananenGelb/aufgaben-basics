package strings

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {
	// TODO
	result := ""
	for i := 1; i < n; i++ {
		result = result + s + sep

	}
	result = result + s
	return result
}
