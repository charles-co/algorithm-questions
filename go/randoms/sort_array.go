package randoms

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func MergeSort(arr []int) []int {

	merge := func(left, right []int) []int {
		res := make([]int, 0)
		for len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				res = append(res, left[0])
				left = left[1:]
			} else {

				res = append(res, right[0])
				right = right[1:]
			}
		}
		if len(left) > 0 {
			res = append(res, left...)
		}
		if len(right) > 0 {
			res = append(res, right...)
		}
		return res
	}

	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	return append(append(left, pivot), right...)
}
