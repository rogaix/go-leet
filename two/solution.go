package two

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	numMap := make(map[int]int)
	for idx, num := range nums {
		diff := target - num
		nidx, ok := numMap[diff]
		if ok {
			return []int{idx, nidx}
		}

		numMap[num] = idx
	}

	return []int{}
}
