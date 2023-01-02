package main

import "fmt"

func main() {

	arr := []int{5, 6, 4, 2, 9, 7, 8}
	//1.冒泡算法
	//bubbleArr := BubbleSort(arr)
	//2.选择算法
	//selectArr := SelectSort(arr)
	//fmt.Println(selectArr)
	//2.插入算法
	insertArr := InsertSort(arr)
	fmt.Println(insertArr)

}
