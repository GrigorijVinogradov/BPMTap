package util

func Average(collection []int64) int64 {
	var total int64 = 0
	for _, v := range collection {
		total += v
	}

	return total / int64(len(collection))
}
