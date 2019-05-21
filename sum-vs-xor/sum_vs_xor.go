package sum_vs_xor

func sumXor(n int64) int64 {
	ans := int64(0)
	for x := int64(0); x <= n; x++ {
		if n+x == n^x {
			ans++
		}
	}
	return ans
}
