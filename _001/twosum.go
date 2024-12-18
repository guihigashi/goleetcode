package _001

func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)

	for i, num := range nums {
		if j, ok := numToIndex[target-num]; ok {
			return []int{j, i}
		}

		numToIndex[num] = i
	}

	panic("invalid arguments")
}
