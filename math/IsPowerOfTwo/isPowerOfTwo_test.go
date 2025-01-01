package ispoweroftwo

import (
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	t.Run("check Power of 64(bitwise)", func(t *testing.T) {
		got := BitwiseIsPower(64)
		want := true

		if got != want {
			t.Errorf("error while check the power of 64")
		}

	})
	t.Run("check Power of 64", func(t *testing.T) {
		got := IsPowerOfTwo(64)
		want := true

		if got != want {
			t.Errorf("error while check the power of 64")
		}

	})
}
