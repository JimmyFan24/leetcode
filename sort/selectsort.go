package main

import "fmt"

func selectSort(arr []int){

	//[1,5,4,3,8,2,6]
	//[1,4,5,3,8,2,6]
	for i := 1;i < len(arr);i++{
		for j := i - 1;j >= 0 && arr[j] > arr[j+1];j--{
			swap(arr,j,j+1)
		}

	}
	
}

func swap(arr []int, num1 int,num2 int){
	temp := 0
	temp = arr[num1]
	arr[num1] = arr[num2]
	arr[num2] = temp

}

func main(){
	arr := []int{3,2,3}
	//	selectSort(arr)
	//sum := rescursionSort(arr,0,len(arr)-1)
	//leftZoom := -1

	//quickSort(arr,0,len(arr)-1)
	//heapSort(arr)
	//sortColors(arr)
	//containsDuplicate(arr)
	result := majorityElement(arr)

	fmt.Println(result)
}