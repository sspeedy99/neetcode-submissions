func hasDuplicate(nums []int) bool {
    nums_set := make(map[int]struct{}, len(nums))
	for _, item := range nums {
		if _, exists := nums_set[item]; exists{
			return true;
		}
		nums_set[item] = struct{}{}
	}
	return false;
}
