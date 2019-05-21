package sum_vs_xor

import "math"

// XOR does not diverge from addition only when the 1 bits of the addends do not overlap, otherwise there's the carry bit, which addition considers while XOR discards
func countZeros(n int64) int {
	zeros := 0
	for {
		if n&1 == 0 {
			zeros++
		}
		n >>= 1
		if n == 0 {
			break
		}
	}
	return zeros
}

func sumXor(n int64) int64 {
	if n == 0 {
		return 0
	}
	return int64(math.Pow(2.0, float64(countZeros(n))))
}
