package the_great_xor

import "testing"

func TestTheGreatXor(t *testing.T) {
	var tests = []struct {
		input    int64
		expected int64
	}{
		{input: 2, expected: 1},
		{input: 10, expected: 5},
		{input: 5, expected: 2},
		{input: 100, expected: 27},
	}
	for _, test := range tests {
		if output := theGreatXor(test.input); output != test.expected {
			t.Fatalf("Expected: %d Got: %d", test.expected, output)
		}
	}
}
