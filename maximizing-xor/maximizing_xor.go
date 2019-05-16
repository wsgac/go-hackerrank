package maximizing_xor

import "math"

func maximizingXor(l int32, r int32) int32 {
	max := int32(math.MinInt32)
	for a := l; a <= r; a++ {
		for b := a; b <= r; b++ {
			if a != b {
				if c := a ^ b; c > max {
					max = c
				}
			}
		}
	}
	return max
}
