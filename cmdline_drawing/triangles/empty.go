package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Der Rand des Dreiecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyTriangle(length int) {
	// TODO
	h := 1
	for i := 0; i < length; i++ {

		for j := 0; j < h; j++ {
			if i < 1 || i == length-1 {
				fmt.Print("#")
			} else {
				if j < 1 || j == h-1 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
		}
		h++
		fmt.Println()
	}
}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Dreieck" bzw. "Leeres Rechteck".
