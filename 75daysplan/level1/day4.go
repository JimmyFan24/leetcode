package level1

func LongestConsecutive(nums []int) int {
	res, numMap := 0, map[int]int{}
	for _, num := range nums {
		if numMap[num] == 0 {
			left, right, sum := 0, 0, 0


			if numMap[num-1] > 0 {

				left = numMap[num-1]
			} else {
				left = 0
			}
			if numMap[num+1] > 0 {
				right = numMap[num+1]
			} else {
				right = 0
			}
			// sum: length of the sequence n is in
			sum = left + right + 1
			numMap[num] = sum
			// keep track of the max length
			res = max(res, sum)
			// extend the length to the boundary(s) of the sequence
			// will do nothing if n has no neighbors
			numMap[num-left] = sum
			numMap[num+right] = sum
		} else {
			continue
		}
	}

	return res
}
func max(a int,b int) int{
	if a>=b{
		return a
	}
	return b
}

func IsInterleave(s1 string, s2 string, s3 string) bool {


	if len(s1) + len(s2) != len(s3){
		return false
	}

	n := len(s1)
	m := len(s2)
	var dp [][]bool

	for i:=0;i<=n;i++{
		line := make([]bool,m+1)
		dp = append(dp,line)
	}
	dp[0][0] = true
	for i:=1;i<n+1;i++{
		dp[i][0] = dp[i-1][0] && s3[i-1] == s1[i-1]
	}

	for j:=1;j<m+1;j++{
		dp[0][j] = dp[0][j-1] && s3[j-1] == s1[j-1]
	}
	for i:=1;i<len(s1)+1;i++{
		for j:=1;j<len(s2)+1;j++{

			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] ==s3[i+j-1])

		}
	}

	return dp[n][m]
}
