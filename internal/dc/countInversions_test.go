package dc

import "testing"

type testData struct {
	s []int
	result int
}

func TestCountInversions(t *testing.T){

	sTest := []testData{
		{[]int{1, 8, 2, 4, 3, 7, 6, 5}, 10},
		{[]int{13, 2, 10, 3, 1, 12, 8, 4, 5, 9, 6, 15, 14, 11, 7}, 41},
		{[]int{6, 13, 7, 10, 8, 9, 11, 15, 2, 12, 5, 14, 3, 1, 4}, 62},
		{[]int{3, 9, 1, 15, 14, 11, 4, 6, 13, 10, 12, 2, 7, 8, 5}, 56},
	}

	for _, datum := range sTest {
		_, result := CountSliceInversions(datum.s)

		if result != datum.result {
			t.Errorf("FAILED. Expect %d got %d\n", datum.result, result)
		} else {
			t.Log("PASSED")
		}
	}
}
