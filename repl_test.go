package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "     hello      world     ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "       Motherrrr      i killed a man",
			expected: []string{"motherrrr", "i", "killed", "a", "man"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("wrong amount of strings")
		}
		for i, _ := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				fmt.Println(expectedWord)
				fmt.Println(word)
				t.Errorf("Slices starting number %v doesn't match what was expected", i)
			}
		}
	}
}
