package dc

func CountSliceInversions(s []int) ([]int, int) {
	totalSlice, totalInv := sortAndCountInversions(s, 0)
	return totalSlice, totalInv
}

func sortAndCountInversions(s []int, inv int) ([]int, int) {
	if len(s) == 1 {
		return s, inv
	} else {
		s1, s2 := splitSlice(s)

		leftS, leftInv := sortAndCountInversions(s1, inv)
		rightS, rightInv := sortAndCountInversions(s2, inv)
		mergedS, mergedInv := mergeAndCountInversions(leftS, rightS)

		totalInv := leftInv + rightInv + mergedInv

		return mergedS, totalInv
	}
}

func mergeAndCountInversions(leftS, rightS []int) ([]int, int) {
	var mergedS []int
	mergedInv := 0
	leftI, rightI := 0, 0

	for leftI < len(leftS) && rightI < len(rightS) {
		if leftS[leftI] <= rightS[rightI] {
			mergedS = append(mergedS, leftS[leftI])
			leftI++
		} else {
			mergedS = append(mergedS, rightS[rightI])
			mergedInv += len(leftS)-leftI
			rightI++
		}
	}
	if leftI < len(leftS) {
		mergedS = append(mergedS, leftS[leftI:]...)
	} else {
		mergedS = append(mergedS, rightS[rightI:]...)
	}

	return mergedS, mergedInv
}

func splitSlice (s []int) (s1, s2 []int) {
	if len(s)%2 == 0 {
		s1 = s[:len(s)/2]
		s2 = s[len(s)/2:]
	} else {
		s1 = s[:(len(s)+1)/2]
		s2 = s[(len(s)+1)/2:]
	}
	return
}
