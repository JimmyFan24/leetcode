package main

import (
	"fmt"
)

func selectSort(arr []int) {

	//[1,5,4,3,8,2,6]
	//[1,4,5,3,8,2,6]
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}

}

func swap(arr []int, num1 int, num2 int) {
	temp := 0
	temp = arr[num1]
	arr[num1] = arr[num2]
	arr[num2] = temp

}

func main() {
	arr := []int{7, 9, 6, 5, 4}
	//	selectSort(arr)
	//sum := rescursionSort(arr,0,len(arr)-1)
	//leftZoom := -1

	//quickSort(arr,0,len(arr)-1)
	//heapSort(arr)
	//sortColors(arr)
	//containsDuplicate(arr)
	//result := majorityElement(arr)

	insertSort(arr)
	fmt.Println(arr)
}
