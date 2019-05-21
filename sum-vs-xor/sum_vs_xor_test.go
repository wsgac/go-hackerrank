package sum_vs_xor

import "testing"

func TestSumXor(t *testing.T) {
	var tests = []struct {
		input    int64
		expected int64
	}{
		{input: 5, expected: 2},
		{input: 10, expected: 4},
	}
	for _, test := range tests {
		if output := sumXor(test.input); output != test.expected {
			t.Fatalf("Expected: %d Got: %d", test.expected, output)
		}
	}
}
