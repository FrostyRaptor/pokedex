package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "THIS SHOULD     FIX     THIS",
			expected: []string{"this", "should", "fix", "this"},
		},
		{
			input:    "CHarmanDER      BULBasauR     pikACHU        ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of lists doesn't match")
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Mismatched words")
			}
		}
	}
}
