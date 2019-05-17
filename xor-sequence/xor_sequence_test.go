package xor_sequence

import "testing"

type Arguments struct {
	L int64
	H int64
}

func TestXorSequence(t *testing.T) {
	var tests = []struct {
		input    Arguments
		expected int64
	}{
		{input: Arguments{L: 3, H: 5}, expected: 5},
		{input: Arguments{L: 4, H: 6}, expected: 2},
		{input: Arguments{L: 15, H: 20}, expected: 22},
	}
	for _, test := range tests {
		if output := xorSequence(test.input.L, test.input.H); output != test.expected {
			t.Fatalf("Expected: %d Got: %d", test.expected, output)
		}
	}
}
