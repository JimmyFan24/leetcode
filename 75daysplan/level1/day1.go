package level1

func RunningSum(nums []int) []int {
	sumArr := make([]int,len(nums))
	temp :=0
	for  i :=0;i<len(nums);i++{
		temp += nums[i]
		sumArr[i] = temp
	}
	return sumArr
}

func PivotIndex(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	var sum, left ,right int
	for _,num := range nums{
		sum += num
	}
	for i,num := range nums{
		if left == right{
			return i
		}else if i+1<len(nums){
			left += num
			right -= nums[i+1]
		}
	}
	return -1
}