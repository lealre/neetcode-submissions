func maxProfit(prices []int) int {
        if len(prices) <= 1 {
                return 0
        }

        minValue := prices[0]
        bestResult := 0
        for _,v := range prices {
                result := v - minValue
                if result > bestResult {
                        bestResult = result
                }

                if minValue > v {
                        minValue = v
                }
        }

        return bestResult
}
