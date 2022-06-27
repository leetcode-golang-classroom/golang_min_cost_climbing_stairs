# golang_min_cost_climbing_stairs

You are given an integer array `cost` where `cost[i]` is the cost of `ith` step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index `0`, or the step with index `1`.

Return *the minimum cost to reach the top of the floor*.

## Examples

**Example 1:**

```
Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.

```

**Example 2:**

```
Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.

```

**Constraints:**

- `2 <= cost.length <= 1000`
- `0 <= cost[i] <= 999`

## 解析

給一個陣列 cost, 其中每一個值 cost[i] 代表從第 i 階出發的 cost

每次從第i 階出發可以選擇爬1 階 或是爬 2 階

而最初只能從第0階或是第1 階出發

題目要求爬到最高階的最小 cost

舉例來看 這個問題

假設 cost 為 [10, 15, 20]

從第 0 階出發可以花出以下 決策樹

![](https://i.imgur.com/R6cFCkj.png)

會發現包含的 1 階出發的決策樹

從第 0 階出發的最小 cost  =  

到 0 階的 cost + min(從第1 階出發的最小 cost, 從第 2 階出發的最cost)

所以問題結構如下  min_cost_from_i = cost[i] + min(min_cost_from_i+1, min_cost_from_i+2)

解遞迴子問題，為了避免重複解問題，所以倒著從最大階往下解過去比較方邊

為了方便初始話 從最後 底點是因為已經在目標所以 cost = 0

在 len(cost)-1 階的 cost 為 cost[len(cost)-1] 因為只要從 len(cost) - 1 爬 1 階就到終點

然後開始算 f(i) = cost(i) + min(f(i+1), f(i+2))

到最後 把 min(f(0), f(1)) 及為答案因為只能從 index 0 還有 index 1 出發

## 程式碼
```go
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
```
## 困難點

1. 找出 minCost 的遞迴子關系

## Solve Point

- [x]  設定最開始最頂點的 cost =0 , 然後逐步往回推算
- [x]  計算 f(i) = cost(i)+ min(f(i+1), f(i+2))