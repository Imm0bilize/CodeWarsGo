package FindTheUniqueNumber

func FindUniq(arr []float32) float32 {
	fstNumber := arr[0]
	fstNumberCounter := 1
	var otherNumber float32
	otherNumberCounter := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] != fstNumber {
			otherNumber = arr[i]
			otherNumberCounter++
		} else {
			fstNumberCounter++
		}

		if ((otherNumberCounter > 1) && fstNumberCounter == 1) || ((fstNumberCounter > 1) && otherNumberCounter == 1) {
			break
		}
	}

	if otherNumberCounter > fstNumberCounter {
		return fstNumber
	} else {
		return otherNumber
	}

}
