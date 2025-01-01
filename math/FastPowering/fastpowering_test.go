package fastpowering

import "testing"

func TestFastPowering(t *testing.T) {
	t.Run("fast powering of r 2^8", func(t *testing.T) {
		got := FastPowering(2, 8)
		want := 256
		if want != int(got) {
			t.Errorf("Result is incorrect")
		}
	})
}
