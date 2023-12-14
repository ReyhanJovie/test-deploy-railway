package utility

func Sum(nums ...int) (total int) {
	if len(nums) == 0 {
		return -1
	}

	for _, num := range nums {
		if num <= 0 {
			return -1
		}
		total += num
	}
	return
}
