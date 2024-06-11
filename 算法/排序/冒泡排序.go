package 排序

func BubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	l := len(arr)
	for i := 0; i < l; i++ {
		for j := 1; j < l; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return
}
