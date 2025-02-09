package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	// TODO
	t := ""
	for _, i := range s {

		t = t + string(i) + string(i)
	}
	return t

}
