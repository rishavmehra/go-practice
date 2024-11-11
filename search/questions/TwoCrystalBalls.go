package questions

import "math"

func TwoCrystalBalls(breaks []bool) bool {
	jumpAmount := math.Floor(math.Sqrt(float64((len(breaks)))))

	i := int(jumpAmount)
	for i < len(breaks) {
		if breaks[i] {
			break
		}
		i += int(jumpAmount)
	}

	i -= int(jumpAmount)

	for j := 0; j < int(jumpAmount) && i < len(breaks); j, i = j+1, i+1 {
		if breaks[i] {
			return true
		}
	}
	return false

}
