// Package reverse is an adapted version from package morestrings in
// found https://golang.org/doc/code.html.
//
// It serves the sole purpose of exploring how to add a remote module via
// github, nothing more, nothing less.
package morestringtools

// ReverseRunes returns the given string s reverse rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
