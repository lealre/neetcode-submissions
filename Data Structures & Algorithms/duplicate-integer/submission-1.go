func hasDuplicate(nums []int) bool {
    
    record := make(map[int]bool)
    for _,v := range nums {
        if _,ok := record[v]; ok {
            return true
        }

        record[v] = true

    }

    return false
}
