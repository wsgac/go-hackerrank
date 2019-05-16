package lonely_integer

// func lonelyinteger(a []int32) int32 {
// 	n := len(a)
// 	for i, el := range a {
// 		repeated := false
// 		for j := 0; j < n; j++ {
// 			if i != j {
// 				if 0 == el^a[j] {
// 					repeated = true
// 					break
// 				}
// 			}
// 		}
// 		if !repeated {
// 			return el
// 		}
// 	}
// 	return 0
// }

func lonelyinteger(a []int32) int32 {
	m := make(map[int32]int32)
	for _, el := range a {
		m[el]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
