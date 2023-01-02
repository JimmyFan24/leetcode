package main

func SelectSort(arr []int) []int {

	//[1,5,4,3,8,2,6]
	//[1,4,5,3,8,2,6]
	//外循环代表排好序的位置,比如说,一开始i==0,意思是在1-n这个区间内找到比i小的从而确定i这个位置的值
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		//此时,min已经是最小元素的索引,直接和i交换
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}

func swap(arr []int, num1 int, num2 int) {
	temp := 0
	temp = arr[num1]
	arr[num1] = arr[num2]
	arr[num2] = temp

}

func main() {
	arr := []int{7, 6, 5, 4}
	//	selectSort(arr)
	//sum := rescursionSort(arr,0,len(arr)-1)
	//leftZoom := -1

	//quickSort(arr,0,len(arr)-1)
	//heapSort(arr)
	//sortColors(arr)
	//containsDuplicate(arr)
	//result := majorityElement(arr)

	selectSort(arr)
	fmt.Println(arr)
}
