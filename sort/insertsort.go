package main

func insertSort(arr []int) {

	for unSortedBoundry := 1; unSortedBoundry < len(arr); unSortedBoundry++ {

		sortedBoundry := unSortedBoundry - 1
		current := arr[unSortedBoundry]
		for sortedBoundry >= 0 && arr[sortedBoundry] > current {
			arr[sortedBoundry+1] = arr[sortedBoundry]
			sortedBoundry--
		}
		arr[sortedBoundry+1] = current

	}
	return
}
