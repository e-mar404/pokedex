package main 

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "  ",
			expected: []string{},
		},
		{
			input: "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "   pokemon    pika",
			expected: []string{"pokemon", "pika"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range(cases) {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths dont match: len(actual): %v, len(expected): %v", actual, c.expected)
		}
		
		for i := range(c.expected) {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("actual: %s and expected: %s do not match", word, expectedWord)
			}
		}
	}
}
