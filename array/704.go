package main
func search(nums []int, target int) int {

	index := findTarget(nums,0,len(nums)-1,target)
	return index
}
func findTarget(nums[]int,L int,R int,target int) int{
	middle := L + (R-L)/2
	result := 0
	if R - L < 0{
		   return -1
	}
	if nums[middle] > target{
		result =  findTarget(nums,0,middle-1,target)
	}else if nums[middle] < target{
		result =  findTarget(nums,middle+1,R,target)
	}else{
		return middle
	}


	return result
}

