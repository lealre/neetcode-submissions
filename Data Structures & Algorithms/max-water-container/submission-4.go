func maxArea(heights []int) int {
	imax := len(heights)-1
	imin := 0
	var maxArea int

	for imin < imax {
		area := min(heights[imin], heights[imax]) * (imax-imin)
		if area > maxArea {
			maxArea = area
		}

		if heights[imin] < heights[imax] {
			imin++
			continue
		}

		imax--

	}

	return maxArea
}