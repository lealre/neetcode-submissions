func productExceptSelf(nums []int) []int {
	var mul int
	mul = 1
	zeroCount := 0
	for _,v := range nums {
		if v == 0 {
			zeroCount++
			if zeroCount > 1 {
				return make([]int, len(nums))
			}
			continue
		}
		mul = mul * v
	}

	result := make([]int, len(nums))
	for i,v := range nums {
		if zeroCount == 1 && v != 0 {
			result[i] = 0
			continue
		}

		if v == 0 {
			result[i] = mul
			continue
		}
		result[i] = mul / v
	}

	return result
}
