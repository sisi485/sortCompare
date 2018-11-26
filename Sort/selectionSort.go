package Sort


func SelectionSort(arr []int) {
	var n = len(arr)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		//swap
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}