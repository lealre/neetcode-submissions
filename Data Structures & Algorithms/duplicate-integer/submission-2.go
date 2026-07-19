func hasDuplicate(nums []int) bool {
    
    duplicate := map[int]int{} 
    for _,v := range nums {
        if _, ok := duplicate[v]; ok {
            return true
        }
        duplicate[v]++
    }

    return false
}
