package level1



func ArrangeCoins(n int) int {


	dp := make([]int,n)
	if n<1 {
		return 0
	}


	dp[0] = 1
	res :=0
	for i:=1;i<=n;i++{
		if dp[i-1] +i+1 >n{
			res = i
			break
		}else if dp[i-1] +i+1== n{
			res = i+1
			break
		}
		dp[i] = dp[i-1] + i+1
	}

	return res

}

func ThirdMax(nums []int) int {

	first ,second,third := 0,0,0
	if len(nums) <3{
		return getMax(nums)
	}else{
		for _,v := range nums{

			if v > first{
				first = v
			} else if v > second{
				second =v
			}else if v > third{
				third = v
			}
		}

	}
	if third == 0{
		return getMax(nums)

	}
	return third
}
func getMax(nums []int)int{

	tmp := 0
	for i:=0;i<len(nums);i++{
		if tmp <nums[i]{
			tmp = nums[i]
		}
	}
	return tmp
}

func LongestPalindrome(s string) string {

	dp := make([][]bool,len(s))
	for i:=0;i<len(dp);i++{
		dp[i] = make([]bool,len(s))
	}


	res := ""
	for i:=0; i <len(s); i++ {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}

	return res




}