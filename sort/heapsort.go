package main

func heapSort(arr []int) {

	//1，构建大项堆
	heapify(arr, len(arr))
	//2.调整堆顶元素
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i)
	}

}

func heapify(arr []int, unSortCapacity int) {

	//从最后一个非叶子节点开始构建非叶子节点的范围:[0,unSortCapacity/2-1]

	for i := unSortCapacity/2 - 1; i >= 0; i-- {
		//调整左子树
		leftIndex := 2*i + 1
		if leftIndex < unSortCapacity && arr[leftIndex] > arr[i] {
			//左子树大于父节点，交换
			arr[leftIndex], arr[i] = arr[i], arr[leftIndex]
		}
		rightIndex := 2*i + 2
		if rightIndex < unSortCapacity && arr[rightIndex] > arr[i] {
			//左子树大于父节点，交换
			arr[rightIndex], arr[i] = arr[i], arr[rightIndex]
		}
	}
}
