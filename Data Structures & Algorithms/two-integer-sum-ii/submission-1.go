func twoSum(numbers []int, target int) []int {
        imax := len(numbers)-1
        imin := 0

        for imin < imax {

                total := numbers[imin] + numbers[imax]
                if total == target {
                        return []int{imin+1, imax+1}
                }

                if total > target {
                        imax--
                        continue
                }

                imin++
        }

        return []int{}
}
