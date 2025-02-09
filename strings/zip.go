package strings

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	// TODO
	result := ""
	l := len(s1)
	if len(s1) > len(s2) {
		l = len(s2)
	}

	for i := 0; i < l; i++ {
		result = result + string(s1[i]) + string(s2[i])
	}
	result += s1[l:]
	result += s2[l:]

	return result
}
