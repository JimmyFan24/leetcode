package main

func MergerSort(arr []int) []int {

	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) >> 1
	return mergesort(arr[:mid], arr[mid:])

}
func mergesort(arr1 []int, arr2 []int) []int {

	res := []int{}

	if len(arr1) >= 2 {
		arr1 = mergesort(arr1[:len(arr1)/2], arr1[len(arr1)/2:])
	}
	if len(arr2) >= 2 {
		arr2 = mergesort(arr2[:len(arr2)/2], arr2[len(arr2)/2:])
	}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			res = append(res, arr1[i])
			i++
		} else if arr1[i] > arr2[j] {
			res = append(res, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		res = append(res, arr1[i])
		i++

	}
	for j < len(arr2) {
		res = append(res, arr2[j])
		j++

	}
	return res
}
