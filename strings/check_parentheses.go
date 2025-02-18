package strings

// Erwartet einen String und prüft, ob er korrekte Klammerpaare enthält.
// D.h. der Eingabestring enthält Klammern '(' und ')'.
// Die Funktion soll nun prüfen, ob der String für jede öffnende Klammer auch eine
// schließende Klammer enthält.
// Außerdem darf es keine schließenden Klammern geben, für die es nicht vorher eine
// passende öffnende Klammer gegeben hat.
// Die Funktion soll true liefern, falls der String korrekt geklammert ist.
func CheckParentheses(s string) bool {
	// TODO
	count := 0
	for _, i := range s {

		if i == '(' {
			count++
		}
		if i == ')' {
			count--
		}
		if count < 0 {
			return false
		}

	}
	if count == 0 {
		return true
	}
	return false

}
