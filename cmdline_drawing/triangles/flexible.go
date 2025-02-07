package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawTriangle(length int, inner, outer string) {
	// TODO
	h := 1
	for i := 0; i < length; i++ {

		for j := 0; j < h; j++ {
			if i < 1 || i == length-1 {
				fmt.Print(outer)
			} else {
				if j < 1 || j == h-1 {
					fmt.Print(outer)
				} else {
					fmt.Print(inner)
				}
			}
		}
		h++
		fmt.Println()
	}
}
