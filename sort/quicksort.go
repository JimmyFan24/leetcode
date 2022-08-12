package main

func quickSort(arr []int,L int,R int){

	if L < R{
		partitionLeft,partitionRight := partition(arr,L,R)
		quickSort(arr,L,partitionLeft-1)
		quickSort(arr,partitionRight+1,R)
	}


}

func containsDuplicate(nums []int) bool {

	numMap := make(map[int]int,len(nums))
	for _,num := range nums{
		numMap[num]++
	}
	for _,n := range numMap{
		if n == 0{
			return false
		}
	}
	return true

	return false
}

func majorityElement(nums []int) int {
	result := 0
	numMap := make(map[int]int,len(nums))
	majorityCount := 0
	if len(nums) % 2 == 0 {
		majorityCount = len(nums)/2
	}else {
		majorityCount = len(nums)/2+1
	}

	for _,i := range nums{
		numMap[i]++
	}
	for i,c := range numMap{
		if c >= majorityCount{
			result = i
		}
	}
	return result
}
func sortColors(nums []int)  {
	zero,one := 0,0
	for i,n := range nums{
		nums[i] = 2
		if n <= 1{
			nums[one]=1
			one++
		}
		if n == 0{
			nums[zero] = 0
			zero++
		}
	}
}

func partition(arr []int,L int,R int)(int,int){
	less := L - 1
	more := R
	//temp := arr[R]
	for{
		if L < more{
			if arr[L] < arr[R]{
				swap(arr,less+1,L)
				less++
				L++
			}else if arr[L] > arr[R]{

				swap(arr,more-1,L)
				more--

			}else {
				L++
			}
		}else {
			break
		}
	}
	swap(arr,more,R)
	return less+1, more
}