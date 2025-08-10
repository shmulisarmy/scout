package kanban

func list_contains[T comparable](list []T, item T) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func list_find_index[T any](list []T, compare_func func(T) bool) int {
	for i, j := range list {
		if compare_func(j) {
			return i
		}
	}
	return -1
}
