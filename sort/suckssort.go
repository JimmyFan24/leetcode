package main

func sucksSort(arr []int) {

	//5,4,3,7,8
	//4,3,5,7,8
	for j := len(arr) - 1; j >= 0; j-- {
		for i := 0; i < j; i++ {
			if arr[i] >= arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	return
}
