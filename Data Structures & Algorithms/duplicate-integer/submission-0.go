func hasDuplicate(nums []int) bool {
    
    record := make(map[int]int)
    for _,v := range nums {
        record[v]++
        if j,ok := record[v]; ok && j>1 {
            return true
        }

    }

    return false
}
