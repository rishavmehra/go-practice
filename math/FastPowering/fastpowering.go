package fastpowering

import (
	"math"
)

func FastPowering(base float64, power int) float64 {

	if power == 0 {
		return 1
	}

	if power%2 == 0 {
		multiplier := FastPowering(base, power/2)
		return multiplier * multiplier
	}

	multiplier := FastPowering(base, int(math.Floor(float64(power/2))))
	return multiplier * multiplier * base
}
