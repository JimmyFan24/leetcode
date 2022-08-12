package main


func heapSort(arr []int){
	if arr == nil || len(arr) < 2{
		return
	}
	for i,_ := range arr{
		heapInsert(arr,i)
	}
	heapsize := len(arr)
	heapsize--
	swap(arr,0 ,heapsize)
	for{
		if heapsize > 0{

			heapify(arr,0,heapsize)
			swap(arr,0,heapsize-1)
			heapsize--
		}else {
			break
		}
	}
}
func heapInsert(arr []int,index int){
	for{
		if arr[index] > arr[(index-1)/2]{
			swap(arr,index,(index-1)/2)
			index = (index-1)/2
		}else {
			break
		}
	}
}
func heapify(arr []int,index int,heapsize int){

	left := index * 2 + 1
	largest := 0
	for{
		if left < heapsize{
			if left + 1 < heapsize && arr[left + 1]>arr[left]{
				largest = left + 1
			}else {
				largest = left
			}

			if arr[largest] > arr[index]{
				largest = largest
			}else {
				break
			}
			swap(arr,index,largest)
			index = largest
			left = 2 * index + 1

		}else {
			break
		}
	}

}