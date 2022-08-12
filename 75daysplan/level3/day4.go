package level3

import (
	"math"
	"strings"
)

/*
Example 1:

Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
*/
func LongestCommonSubsequence(text1 string, text2 string) int {

	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	//dp[i][j] : 长度为 i 的 text1[0:i-1] 和长度为 j 的 text2[0:j-1]
	//dp[i][j] : 当 text1[i−1] = text2[j−1] 时，将这两个相同的字符称为公共字符
	/*
		dp[i][j] = max(dp[i-1][j], dp[i][j-1]) :text1[i−1] != text2[j−1]
		dp[i][j]=dp[i−1][j−1]+1 ,:text1[i−1] = text2[j−1]
	*/

	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {

			if text1[i-1] != text2[j-1] {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	return dp[len(text1)][len(text2)]

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func PartitionLabels(s string) []int {
	var lastIndexOf [26]int
	for i, v := range s {
		lastIndexOf[v-'a'] = i
	}
	res := []int{}
	//"ababcbacadefegdehijhklij"
	for start, end := 0, 0; start < len(s); start = end + 1 {
		end = lastIndexOf[s[start]-'a']
		for i := start; i < end; i++ {
			if end < lastIndexOf[s[i]-'a'] {
				end = lastIndexOf[s[i]-'a']
			}
		}
		res = append(res, end-start+1)
	}
	return res
}

/*
Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
*/
func Jump(nums []int) int {

	if len(nums) == 1 {
		return 0
	}
	step, needChoose, canReach := 0, 0, 0
	//4,1,1,3,1,1,1
	for i, v := range nums {

		if i+v > canReach {
			canReach = i + v
			if canReach > len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return step
}
func NumTrees(n int) int {

	//F(i,n) 代表以 i 为根节点，1-n 个数组成的二叉排序树的不同的个数
	//dp[n] = F(1,n) + F(2,n) + F(3,n) + …… + F(n,n)
	//dp 和 F(i,n) 的关系又可以得到下面这个等式 F(i,n) = dp[i-1] * dp[n-i]
	//dp[i] = dp[0] * dp[n-1] + dp[1] * dp[n-2] + …… + dp[n-1] * dp[0]

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	//1,2,3
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
*/
func Rob(nums []int) int {
	//dp[i] = max(dp[i-2]+nums[i],dp[i-1])
	//dp[0] = nums[0]
	//dp[1] = max(dp[0],dp[1])
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)]
}

/*
Example 1:

Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
*/
func MaxProduct(nums []int) int {

	// [-3,4,-2]
	//  -3 -12
	if len(nums) == 1 {
		return nums[0]
	}

	res := math.MinInt32
	for i := 0; i < len(nums); i++ {
		j := i + 1
		if i == len(nums)-1 {
			res = max(res, nums[i])
			break
		}
		// [-3,4,-2,5,-4,7,-8]
		tmp := 1
		if nums[i] < 0 {
			for ; j < len(nums); j++ {

				tmp = tmp * nums[j]
				if tmp < 0 {
					res = max(nums[i]*tmp, res)
				} else {
					if tmp == 0 {
						res = max(res, nums[i])
						break
					}
					res = max(res, nums[i])
				}
			}

		} else if nums[i] >= 0 {
			if nums[i] == 0 {
				res = max(res, nums[i])
				continue
			}
			for ; j < len(nums); j++ {

				tmp = tmp * nums[j]

				if tmp < 0 {
					res = max(nums[i], res)
				} else {
					if tmp == 0 {
						res = max(res, nums[i])
						break

					}
					res = max(nums[i]*tmp, res)
				}
			}

		}

	}
	return res

}
func FindTargetSumWays(nums []int, target int) int {

	//1,2,3,4,5

	sum := 0
	min := 0
	for _, v := range nums {
		sum += v
		min -= v
	}
	if sum < target || min > target {
		return 0
	}
	if sum == target || min == target {
		return 1
	}

	//sum==15,target==5 10
	//1+2+3+4+5 == 5
	step := 0
	//maxTmp := sum-target

	for i := 1; i < len(nums)+1; i++ {
		FindDfs(nums, sum, target, i, &step)
	}
	return min
}
func FindDfs(nums []int, sum, target, index int, step *int) {

	//1+2+3+4+5 == 5

	for j := index; j < len(nums); j++ {

		sum = sum - nums[j]
		if sum == target {

			return
		}
		FindDfs(nums, sum, target, index+1, step)
		sum = sum + nums[j]
	}
	return
}

/*
Example 1:

Input: strs = ["cba","daf","ghi"]
Output: 1
Explanation: The grid looks as follows:
  cba
  daf
  ghi
Columns 0 and 2 are sorted, but column 1 is not, so you only need to delete 1 column.
*/
/*
"zyx",
"wvu",
"tsr"]
*/
func MinDeletionSize(strs []string) int {

	m := len(strs)
	n := len(strs[0])
	col := 0
	for j := 0; j < n; j++ {
		for i := 0; i < m-1; i++ {
			if strs[i][j] > strs[i+1][j] {
				col++
				break
			}
		}
	}

	return col
}
func DiStringMatch(s string) []int {
	n := len(s)
	minNum := 0
	maxNum := n
	res := make([]int, n+1)
	for i, v := range s {
		if v == 'I' {
			res[i] = minNum
			minNum++
		}
		if v == 'D' {
			res[i] = maxNum
			maxNum--
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindMode(root *TreeNode) []int {

	res := findmode(root)
	m := make(map[int]int, len(res))
	for _, v := range res {
		m[v]++
	}
	r := []int{}
	maxFlu := 0
	for _, v := range res {
		if m[v] >= maxFlu {
			maxFlu = m[v]
		}
	}
	for i := 0; i < len(res); i++ {
		if m[res[i]] == maxFlu {
			delete(m, res[i])
			r = append(r, res[i])
		}
	}
	return r
}

func findmode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}

	left := FindMode(root.Left)
	right := FindMode(root.Right)

	for _, v := range left {

		res = append(res, v)
	}
	res = append(res, root.Val)
	for _, v := range right {
		res = append(res, v)
	}

	return res
}

/*
Example 1:

Input: num = 28
Output: true
Explanation: 28 = 1 + 2 + 4 + 7 + 14
1, 2, 4, 7, and 14 are all divisors of 28.
*/
func CheckPerfectNumber(num int) bool {

	bound := int(math.Sqrt(float64(num))) + 1
	sum := 0
	for i := 2; i < bound; i++ {
		if num%i != 0 {
			continue
		}
		sum += num/i + i
	}
	return sum == num
}

/*
Example 1:

Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
*/
func ReverseWords(s string) string {

	sList := strings.Split(s, " ")
	res := ""
	for _, v := range sList {

		tmp := []byte{}
		for i := len(v) - 1; i >= 0; i-- {
			tmp = append(tmp, v[i])
		}
		res = res + string(tmp) + " "
	}
	return res[:len(res)-1]
}

/*
Example 1:

Input: s = "abcdefg", k = 2
Output: "bacdfeg"
*/
func ReverseStr(s string, k int) string {

	res := ""

	if len(s) <= k {
		res = reverse1(s)
	} else if k < len(s) && 2*k > len(s) {
		res = reverse1(s[:k]) + s[k:]
	} else {
		res = reverse1(s[:k]) + s[k:2*k] + ReverseStr(s[2*k:],k)
	}

	return res
}
func reverse1(s string) string {

	res := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return string(res)
}
/*
Example 1:

Input: nums = [1,2,1,3,2,5]
Output: [3,5]
Explanation:  [5, 3] is also a valid answer.
 */
func singleNumber(nums []int) []int {
	nMap := make(map[int]int,len(nums))
	for _,v:= range nums{
		nMap[v]++
	}
	res := []int{}
	for i:=0;i<len(nums);i++{
		if nMap[nums[i]]==1{
			res =append(res,nums[i])
		}
	}
	return res
}
func SingleNumberIII(nums []int) []int {
	diff := 0
	//0001
	//0010
	//0001
	//0010
	//0011
	//0100

	//0111

	//0001
	//0011
	//0100
	//0001

	for _, num := range nums {
		diff ^= num
	}
	// Get its last set bit (lsb)
	//0111
	//1001
	//0001
	diff = diff & (-diff)
	res := []int{0, 0} // this array stores the two numbers we will return
	for _, num := range nums {
		if (num & diff) == 0 { // the bit is not set
			res[0] ^= num
		} else { // the bit is set
			res[1] ^= num
		}
	}
	return res
}