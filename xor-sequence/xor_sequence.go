package xor_sequence

func xorSequence(l int64, r int64) int64 {
	result := int64(0)
	a := int64(0)
	for i := int64(0); i <= r; i++ {
		a ^= i
		if i >= l {
			result ^= a
		}
	}
	return result
}
