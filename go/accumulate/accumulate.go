// Package accumulate is a solution to exercism.io exercise titled Accumulate.
package accumulate

// Accumulate returns a new collection given a collection and an operation to perform
// on each element of the collection.
func Accumulate(collection []string, operation func(string) string) []string {
	result := []string{}

	for _, item := range collection {
		result = append(result, operation(item))
	}

	return result
}
