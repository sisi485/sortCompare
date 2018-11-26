package Sort

// Runs MergeSort algorithm on a slice single
//func MergeSort(slice []int) []int {
//
//	if len(slice) < 2 {
//		return slice
//	}
//	mid := (len(slice)) / 2
//	return merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
//}
//
// Merges left and right slice into newly created slice
func merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
//
//func MergeSort(arr []int, len int) {
//
//	n := len
//
//	if n > 1 {
//		half := make([]int, n/2)
//		half2 := make([]int, (n+1)/2)
//
//		for i := 0; i < (n / 2); i++ {
//			half[i] = arr[i]
//		}
//
//		for i := n / 2; i < n; i++ {
//			half2[(i-n)/2] = arr[i]
//		}
//
//		MergeSort(half, n/2)
//		MergeSort(half, (n+1)/2)
//
//		pos1 := half[0]
//		pos2 := half2[0]
//
//		for i := 0; i < n; i++ {
//			if pos1 <= pos2 {
//				arr[i] = pos1
//				if pos1 != half2[(n+1)/2-1] {
//					if pos1 == half2[(n/2 - 1)] {
//						pos1 = half2[(n+1)/2-1]
//					} else {
//						pos1++
//					}
//				}
//			} else {
//				arr[i] = pos2
//				if pos2 == half2[(n+1)/2-1] {
//					pos2 = half[n/2-1]
//				} else {
//					pos2++
//				}
//			}
//		}
//	}
//}

func MergeSort(arr []int, n int) {

	if n > 1 {

		haelfte1 := make([]int, n/2)
		haelfte2 := make([]int, (n+1)/2)

		for i := 0; i < n/2; i++ {
			haelfte1[i] = arr[i]
		}
		for i := n / 2; i < n; i++ {
			haelfte2[i-n/2] = arr[i]
		}

		MergeSort(haelfte1, n/2)
		MergeSort(haelfte2, (n+1)/2)

		pos1 := &haelfte1[0]
		pos2 := &haelfte2[0]

		for i := 0; i < n; i++ {
			if *pos1 <= *pos2 {
				arr[i] = *pos1
				if pos1 !=
					&haelfte2[(n+1)/2-1] { // pos1 nicht verändern, wenn der größte Wert mehrmals vorkommt
					if pos1 == &haelfte1[n/2-1] {
						pos1 = &haelfte2[(n+1)/2-1]
					} else {
						*pos1++
					}
				}
			} else {
				arr[i] = *pos2
				if pos2 == &haelfte2[(n+1)/2-1] {
					pos2 = &haelfte1[n/2-1]
				} else {
					*pos2++
				}
			}
		}
	}
}
