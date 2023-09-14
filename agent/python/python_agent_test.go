package python

import (
	"testing"
)

func TestExtractPythonCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name:     "Valid Python Code Extraction",
			input:    "Some random text\n```python\nprint('Hello World!')\n```\nSome more text.",
			expected: "print('Hello World!')",
			hasError: false,
		},
		{
			name:     "No Python Code Block",
			input:    "Some random text without any python code block",
			expected: "",
			hasError: true,
		},
		{
			name:     "Missing End Block",
			input:    "Some random text\n```python\nprint('Hello World!')Some more text.",
			expected: "",
			hasError: true,
		},
		{
			name:     "Multiple Python Code Blocks",
			input:    "Random text\n```python\nprint('First!')\n```\nAnother text\n```python\nprint('Second!')\n```",
			expected: "print('First!')",
			hasError: false,
		},
	}

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				output, err := extractPythonCode(test.input)
				if output != test.expected {
					t.Fatalf("expected %s but got %s", test.expected, output)
				}
				if test.hasError && err == nil {
					t.Fatalf("expected an error but got none")
				}
				if !test.hasError && err != nil {
					t.Fatalf("didn't expect an error but got: %v", err)
				}

			},
		)
	}
}
