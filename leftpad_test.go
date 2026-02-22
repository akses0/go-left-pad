package leftpad

import "testing"

func TestLeftPad(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		length   int
		padStr   string
		expected string
	}{
		{"Basic padding", "foo", 5, " ", "  foo"},
		{"Already long enough", "foobar", 5, " ", "foobar"},
		{"Empty pad string", "foo", 5, "", "foo"},
		{"Padding with multi-char string", "foo", 7, "ab", "ababfoo"},
		{"Padding with multi-char string (partial)", "foo", 6, "ab", "abafoo"},
		{"Exact length", "foo", 3, " ", "foo"},
		{"Length less than current", "foo", 2, " ", "foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LeftPad(tt.s, tt.length, tt.padStr)
			if result != tt.expected {
				t.Errorf("LeftPad(%q, %d, %q) = %q; want %q", tt.s, tt.length, tt.padStr, result, tt.expected)
			}
		})
	}
}
