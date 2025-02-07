package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
func DrawSolidTriangle(length int) {
	// TODO
	h := 1
	for i := 0; i < length; i++ {

		for j := 0; j < h; j++ {
			fmt.Print("#")
		}
		h++
		fmt.Println()
	}

}
