package level2


func OddEvenList(head *ListNode) *ListNode {

	oddList := []*ListNode{}
	notoddList := []*ListNode{}
	flag := 0
	for head != nil{
		if flag % 2 ==0{
			oddList = append(oddList,head)

		} else {
			notoddList = append(notoddList,head)

		}
		flag++
		head = head.Next

	}
	for i:=0;i<len(oddList)+len(notoddList);i++{
		if i+1<len(oddList){
			oddList[i].Next = oddList[i+1]
		}else if i+1 == len(oddList){
			oddList[i].Next = notoddList[0]
		}else if i >=len(oddList) && i+1<(len(oddList)+len(notoddList)) {
			notoddList[i-len(oddList)].Next = notoddList[i-len(oddList)+1]
		}else{
			notoddList[i-len(oddList)].Next = nil
		}

	}
	return head

}
func GetSum(a int, b int) int {
	sum :=0
	for i:=0;i<a;i++{
		sum = 1 << sum
	}

	for i:=0;i<b;i++{
		sum = 1 << sum
	}

	return sum
}
func LengthOfLIS(nums []int) int {

	res := 0
	dp := make([]int,len(nums))
	dp[0] = 0
	// 1,2,7,3,4,2,10
	// 0,1,2,2,3,1,4
	for i:=1;i<len(nums);i++{

		for j:=0;j<i;j++{
			if nums[j] < nums[i]{
				dp[i] = max(dp[i],1+dp[j])
			}

		}
	}
	return res
}

func max(a ,b int)int{

	if a>b{
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {

	res := [][]int{}
	for i:=0;i<len(intervals);i++{
		v := intervals[i]
		if i+1<len(intervals){


			v1 := intervals[i+1]

			if v[1] >=v1[0] && v[1] <= v1[1] || v[0]>=v1[0] && v[0]<= v1[1]{

				left := min(v[0],v1[0])
				right := max(v[1],v1[1])
				res = append(res,[]int{left,right})
				i++
			}else{
				res = append(res,v)
			}

		}else{
			res = append(res,v)
		}
	}
	return res
}
func min(a,b int) int{
	if a >b{
		return b
	}
	return a
}

/*func merge(nums1 []int, m int, nums2 []int, n int)  {


	m1 := len(nums1)
	for p := m1 ;m>0 && n>0;p--{

		if nums1[m-1] >= nums2[n-1]{
			nums1[p] = nums1[m-1]
			m--
		}else {
			nums1[p] = nums2[n-1]
			n-
		}

	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}*/


