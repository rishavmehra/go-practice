package euclideanalgorithm_test

import (
	"testing"

	euclideanalgorithm "github.com/rishavmehra/golangPractice/math/EuclideanAlgorithm"
)

func TestGCD(t *testing.T) {
	t.Run("test for 252 & 105 ", func(t *testing.T) {
		got := euclideanalgorithm.GCD(252, 105)
		want := 21

		if want != got {
			t.Errorf("result is incorrect")
		}

	})
	// data := []struct {
	// 	n    int
	// 	k    int
	// 	want int
	// }{
	// 	{0, 0, 0}, {2, 0, 2}, {0, 2, 2}, {12, 4, 4}, {252, 105, 21}, {-462, -1071, -21},
	// }

	// for _, d := range data {
	// 	if got := euclideanalgorithm.GCD(d.n, d.k); got != d.want {
	// 		t.Errorf("Invalid value for N: %d, got: %d, want: %d", d.n, got, d.want)
	// 	}
	// }

}
