func swap(data []int, i int, j int) {
	data[i], data[j] = data[j], data[i]
}

/*
 * Selection Sort
 */
func SelectionSort(data []int) {
	length := len(data)

	for i:=0; i<length-1; i++ {
		smallest := i
		for j:=i+1; j<length; j++ {
			if data[j] < data[smallest] {
				smallest = j
			}
		}
		if i != smallest {
			swap(data, i, smallest)
		}
	}
}

/*
 * Quick Sort
 */
func QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, left int, right int) {
	if left >= right {
		return
	}

	std := data[left]
	l, r := left, right

	for l < r {
		for l < r && data[r] >= std {
			r--
		}
		for l < r && data[l] <= std {
			l++
		}
		if l < r {
			swap(data, l, r)
		}
	}
	if left != l {
		swap(data, left, l)
	}

	quickSort(data, left, l-1)
	quickSort(data, l+1, right)
}

/*
 * Heap Sort
 */
func HeapSort(data []int) {
	length := len(data)
	for i:=length/2-1; i>=0; i-- {
		shift(data, i, length)
	}
	var cnt = length
	for cnt > 0 {
		swap(data, 0, cnt-1)
		cnt--
		shift(data, 0, cnt)
	}
}

func shift(data []int, s, l int) {
	left, right := 2*s+1, 2*s+2
	pos := s
	if left<l && data[left] > data[pos] {
		pos = left
	} 
	if right<l && data[right] > data[pos] {
		pos = right
	}
	if (pos != s) {
		swap(data, s, pos)
		shift(data, pos, l)
	}
}
