package repl

import "testing"

func TestCleanInput(t *testing.T) {
	t.Log("Init test")
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "bulbasaur charmander pikachu",
			expected: []string{"bulbasaur", "charmander", "pikachu"},
		},
		{
			input: "PIKACHU",
			expected: []string{"pikachu"},
		},
		{
			input: " bulbasaur  charmander  squirtle",
			expected: []string{"bulbasaur", "charmander", "squirtle"},
		},
	}

	for _, c := range cases {
		t.Log("start testing")
		actual := CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected: %v, actual: %v", c.expected, actual)
		}


		for i, expectedWord := range c.expected {
			actualWord := actual[i]
			if actualWord != expectedWord {
				t.Errorf("expected: %v, actual: %v", expectedWord, actualWord)
			}
		}
	}
}
