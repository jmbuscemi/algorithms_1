package mergesort

func Merge(left, right []int) []int {
	arr := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(arr, right...)
		}
		if len(right) == 0 {
			return append(arr, left...)
		}
		if left[0] <= right[0] {
			arr = append(arr, left[0])
			left = left[1:]
		} else {
			arr = append(arr, right[0])
			right = right[1:]
		}
	}
	return arr
}

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	left := MergeSort(s[:n])
	right := MergeSort(s[n:])
	return Merge(left, right)
}
