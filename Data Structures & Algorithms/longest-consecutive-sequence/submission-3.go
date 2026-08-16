func longestConsecutive(nums []int) int {
        if len(nums) == 0 {
                return 0
        }
        sorted := mySort(nums)
        
        maxInc := 0
        tempInc := 0

        fmt.Printf("Starting algo for %v\n", sorted)
        for i:=0;i<len(sorted)-1;i++ {
                if sorted[i+1] == sorted[i]+1 {
                        tempInc++
                        if i < len(sorted)-2 {
                                continue
                        }
                }

                if tempInc > maxInc  {
                        fmt.Printf("New longest sequence in indice %d: %d\n",i,tempInc )
                        maxInc = tempInc
                }

                tempInc = 0
        }

        return maxInc + 1
}

func mySort(nums []int) []int {
        sort.Ints(nums)
        set := []int(nil)
        for i:=0;i<len(nums);i++ {
                if len(set) == 0 || nums[i] != set[len(set)-1] {
                        set = append(set, nums[i])
                }
        }

        return set

}
