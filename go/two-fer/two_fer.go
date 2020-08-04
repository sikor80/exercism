// Package twofer is a solution to exercism.io exercise titled Two Fer.
package twofer

// ShareWith returns a different string depending on the input.
// If the given name is "Alice", the result should be "One for Alice, one for me."
// If no name is given, the result should be "One for you, one for me."
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
