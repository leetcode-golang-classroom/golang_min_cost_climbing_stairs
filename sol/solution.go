package sol

func minCostClimbingStairs(cost []int) int {
	minCost := 0
	costLen := len(cost)
	lastOne := 0
	lastTwo := cost[costLen-1]
	for startStep := costLen - 2; startStep >= 0; startStep-- {
		minCost = cost[startStep] + min(lastOne, lastTwo)
		lastOne = lastTwo
		lastTwo = minCost
	}
	return min(lastOne, lastTwo)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
