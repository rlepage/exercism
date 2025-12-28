package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func not[T any](predicate func(T) bool) func(T) bool {
	return func(e T) bool {
		return !predicate(e)
	}
}

func Keep[T any](collection []T, predicate func(T) bool) []T {
	res := []T{}
	for _, e := range collection {
		if predicate(e) {
			res = append(res, e)
		}
	}
	return res
}

func Discard[T any](collection []T, predicate func(T) bool) []T {
	return Keep(collection, not(predicate))
}
