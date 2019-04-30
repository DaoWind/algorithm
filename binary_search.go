func binarySearch(data []int, val int) int  {
	end := len(data) - 1
	start := 0

	for start < end {
		mid := (start + end) / 2
		if val == data[mid] {
			return mid
		} else if val > data[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}	
	return -1
}
