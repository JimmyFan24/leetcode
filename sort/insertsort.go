package main


func InsertSort(arr []int) []int {

	for unSortedBoundary := 1; unSortedBoundary < len(arr); unSortedBoundary++ {
		//3,2,4,6,5
		//sortedBoundary 代表有序序列的最后一个数字,unSortedBoundary代表未排序的第一个数字
		sortedBoundary := unSortedBoundary - 1
		//当前需要判断位置的值是current
		current := arr[unSortedBoundary]
		//如果sortedBoundary的位置大于等于0,而且sortedBoundary位置的值都比现在需要插入的值小,则需要继续往前找位置插入
		for sortedBoundary >= 0 && arr[sortedBoundary] > current {
			//3,2,4,6,5
			//这一步是sorted的位置往后挪一步,把位置腾空出来插入
			arr[sortedBoundary+1] = arr[sortedBoundary]
			sortedBoundary--
		}
		//这个时候,sortedBoundary位置的值是小于current的,所以插入的位置就是sortedBoundary+1
		arr[sortedBoundary+1] = current
	}
	return arr

}
