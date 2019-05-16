package lonely_integer

import "testing"

func TestSampleInput0(t *testing.T) {
	out := lonelyinteger([]int32{1})
	if out != 1 {
		t.Fatalf("Expected: %d Got %d", 1, out)
	}
}

func TestSampleInput1(t *testing.T) {
	out := lonelyinteger([]int32{1, 1, 2})
	if out != 2 {
		t.Fatalf("Expected: %d Got %d", 2, out)
	}
}

func TestSampleInput2(t *testing.T) {
	out := lonelyinteger([]int32{0, 0, 1, 2, 1})
	if out != 2 {
		t.Fatalf("Expected: %d Got %d", 2, out)
	}
}
