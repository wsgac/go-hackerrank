package maximizing_xor

import "testing"

type Arguments struct {
	L int32
	H int32
}

func TestMaximizingXOR(t *testing.T) {
	var tests = []struct {
		input    Arguments
		expected int32
	}{
		{input: Arguments{L: 10, H: 15}, expected: 7},
		{input: Arguments{L: 11, H: 100}, expected: 127},
	}
	for _, test := range tests {
		if output := maximizingXor(test.input.L, test.input.H); output != test.expected {
			t.Fatalf("Expected: %d Got: %d", test.expected, output)
		}
	}
}
