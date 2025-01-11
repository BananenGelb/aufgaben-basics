package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.
func MinListRecursive(nums []int) int {
	// TODO
	if len(nums) == 0 { //es wird getestet, ob die Liste Leer ist
		return 0 //wenn die Liste leer ist wird 0 ausgegeben
	}
	if len(nums) == 1 { //es wird getestet, ob die Lioste nur 1 Element hat
		return nums[0] //wenn die Liste nur ein element hat wird dieses ausgegeben
	}
	minOfRest := MinListRecursive(nums[1:])
	if nums[0] < minOfRest {
		return nums[0]
	}
	return minOfRest
}

// REMARKS
// - Diese Funktion nennt man "rekursiv".
// - Rekursion ist ein größeres Thema, das in der Vorlesung separat behandelt wird.
//   Um die Denkweise frühzeitig zu üben, gibt es aber gelegentlich auch vorab Aufgaben dazu
