func maxProfit(prices []int) int {
	// init 2 ptr -- 0th and 1st
	// find difference between both ptr
	// if right < left, update left to right, and right to right+1
	// store largest difference
	l := 0
	maxProfit := 0
	for r := 1; r < len(prices); r++ {
		if prices[l] > prices[r] {
			// update ptr
			l = r
		} else if (prices[r] - prices[l]) > maxProfit {
			maxProfit = prices[r] - prices[l]
		}
	}
	return maxProfit
}