package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawRectangle(height, width int, inner, outer string) {
	// TODO
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i < 1 || i == height-1 {
				fmt.Print(outer)
			} else {
				if j < 1 || j == width-1 {
					fmt.Print(outer)
				} else {
					fmt.Print(inner)
				}
			}

		}
		fmt.Println()
	}

}

// REMARKS
// - Wenn Sie diese Aufgabe gelöst haben, können Sie die Aufgaben
//   für das Zeichnen von leeren und gefüllten Rechtecken sehr viel einfacher lösen.
