package internal

func Map[T, U any](collection []T, fn func(T) U) []U {
	var out []U
	for i := range collection {
		out = append(out, fn(collection[i]))
	}
	return out
}

func Filter[T any](collection []T, fn func(T) bool) []T {
	var out []T
	for i := range collection {
		if fn(collection[i]) {
			out = append(out, collection[i])
		}
	}
	return out
}
