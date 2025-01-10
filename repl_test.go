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
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch")
		}

		for index := range actual {
			word := actual[index]
			expectedWord := c.expected[index]

			if word != expectedWord {
				t.Errorf("word mismatch")
			}
		}
	}

}
