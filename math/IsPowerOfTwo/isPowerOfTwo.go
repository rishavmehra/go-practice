package ispoweroftwo

func IsPowerOfTwo(num int) bool {
	if num == 0 {
		return false
	}
	for num != 1 {
		if num%2 != 0 {
			return false
		}
		num = num / 2
	}
	return true
}

func BitwiseIsPower(num int) bool {
	if num < 0 {
		return false
	}
	return (num & (num - 1)) == 0
}
