// Package utilities, this file is for testing, any file ending with the word
// _test is considered a test file
package utilities

import "testing"

// The function should always begin with the word "Test"
// with the following signature
/*
func (t *testing.T)
*/
func TestReverse(t *testing.T) {
	// Specify the test cases, the input is in and want is the second value in each set
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
