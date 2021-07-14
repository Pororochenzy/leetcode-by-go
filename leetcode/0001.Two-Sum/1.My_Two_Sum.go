package leetcode

//用map来保存 需要找的数字
func MytwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for key, val := range nums {
		if idx, ok := numMap[target-val]; ok {
			return []int{idx, key}
		}
		numMap[val] = key //记录 当前值:索引
	}
	return nil
}
