package mathfunc

func IsOdd(nb int) bool {
	if nb%2 != 0 {
		return true
	}
	return false
}

func IsEven(nb int) bool {
	if nb%2 == 0 {
		return true
	}
	return false
}
