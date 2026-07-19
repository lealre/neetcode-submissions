func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums))
	fc := make(map[int]int)
	for _, v := range nums {
		fc[v]++
	}

	for key, value := range fc {
		bucket[value-1] = append(bucket[value-1], key)
	}

	var results []int
	for j := len(bucket) - 1; j >= 0; j-- {
		for _, value := range bucket[j] {
			results = append(results, value)
			if len(results) == k {
				return results
			}
		}
	}

	return []int{}
}