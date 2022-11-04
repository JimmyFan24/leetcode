package level8

/*
Given a 0-indexed integer array nums,
return true if it can be made strictly increasing after removing exactly one element, or false otherwise.
If the array is already strictly increasing, return true.
The array nums is strictly increasing if nums[i - 1] < nums[i] for each index (1 <= i < nums.length).
*/
func CanBeIncreasing(nums []int) bool {

	//1.算出来多少个point是不对的
	count := 0
	index := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			count++
			index = i
		}
	}
	if count == 0 {
		return true
	}
	if count == 1 {
		if index == 0 || index == len(nums)-2 {
			return true
		}
		//需要比较index前后的值
		//1,2,1,3,4,index == 2 ,index+2 > index remove index+1
		//1,3,2,4,5 index ==2 index+1 ==4 index -1 == 3 remove index
		if nums[index-1] < nums[index+1] || (index+2 < len(nums) && nums[index+2] > nums[index]) {
			return true
		}

	}
	return false
}
func ValidMountainArray(arr []int) bool {

	i := 0
	for i < len(arr)-1 && arr[i] < arr[i+1] {
		i++
	}
	if i < 1 || i == len(arr)-1 {
		return false
	}
	for i < len(arr)-1 && arr[i] > arr[i+1] {
		i++
	}
	return i >= len(arr)-1
}
