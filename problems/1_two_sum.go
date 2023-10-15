package problems

// 给定一个整数数组nums和一个整数目标值target 在数组中找出两数和为目标值的数
// 假设只有一个答案满足

func twoSum(nums []int, target int) []int {
	temp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		y, ok := temp[diff]
		if ok {

			return []int{i, y}
		}

		temp[nums[i]] = i
	}

	return []int{}
}

func init() {
	twoSum([]int{1, 3}, 4)
}
