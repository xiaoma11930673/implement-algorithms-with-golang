package main

import "math"

func main() {
	rob([]int{2,1,1,2})
}

func rob(nums []int) int {
	pre2, pre1 := 0, 0
	for i := 0; i < len(nums); i++{
		cur := int(math.Max(float64(pre2 + nums[i]), float64(pre1)))
		pre2, pre1 = pre1, cur
	}
	return pre1
}