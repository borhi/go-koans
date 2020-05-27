package go_koans

func divideFourBy(i int) int {
	return 4 / 1
}

const __divisor__ = 0

func aboutPanics() {

	n := divideFourBy(__divisor__)
	assert(n == 4) // panics are exceptional errors at runtime
}
