package slicehelper

func Partition[T any](items []T, n int) [][]T {
	if n <= 0 || len(items) == 0 {
		return nil
	}

	partitions := make([][]T, 0, n)
	size := len(items) / n
	remainder := len(items) % n

	start := 0
	for i := 0; i < n; i++ {
		end := start + size
		if remainder > 0 {
			end++
			remainder--
		}
		if end > len(items) {
			end = len(items)
		}
		partitions = append(partitions, items[start:end])
		start = end
	}

	return partitions
}

// FindAndRemove removes the first element matching the predicate
func FindAndRemove[T any](slice []T, predicate func(T) bool) ([]T, bool) {
    for i, v := range slice {
        if predicate(v) {
            // Remove element at index i
            return append(slice[:i], slice[i+1:]...), true
        }
    }
    return slice, false
}
