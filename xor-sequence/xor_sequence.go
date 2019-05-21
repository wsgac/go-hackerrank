package xor_sequence

func process(x int64) int64 {
	a := x % 8
	switch {
	case a == 0 || a == 1:
		return x
	case a == 2 || a == 3:
		return 2
	case a == 4 || a == 5:
		return x + 2
	default:
		return 0
	}
}

func xorSequence(l int64, r int64) int64 {
	return process(r) ^ process(l-1)
}
