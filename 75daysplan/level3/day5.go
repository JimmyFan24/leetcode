package level3

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/
func ThreeSumClosest(nums []int, target int) int {
	//-4,-1,1,2
	/*sort.Ints(nums)
	min := int(math.Abs(float64(nums[0]+nums[1]+nums[2]-target)))
	res := nums[0]+nums[1]+nums[2]
	for i:=0;i<len(nums);i++{
		sumDFS(nums,target,&min,i+1,nums[i],1,&res)
	}
	return res*/

	n, res, diff := len(nums), 0, math.MaxInt32

	if n > 2 {

		for i := 0; i < n-2; i++ {

			if i > 0 && nums[i] == nums[i-1] {
				continue
			}

			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if int(math.Abs(float64(sum-target))) < diff {
					res, diff = sum, int(math.Abs(float64(sum-target)))
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}
func sumDFS(nums []int, target int, min *int, index int, sum int, count int, res *int) {

	if count == 3 {
		if *min >= int(math.Abs(float64(sum-target))) {
			*min = int(math.Abs(float64(sum - target)))
			*res = sum
		}
		if *min == 0 {
			return
		}
		return
	}
	if nums[index] > 0 && *min <= int(math.Abs(float64(sum-target))) {
		return
	}
	for j := index; j < len(nums); j++ {
		count++
		sumDFS(nums, target, min, j+1, sum+nums[j], count, res)

		count--
	}
	return

}
func getmin(a, b int) int {

	if a > b {
		return a
	}
	return b
}

/*
Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
*/
//-2,-1,0,0,1,2
func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	nMap := make(map[int]int, n)
	for i, v := range nums {
		nMap[v] = i
	}
	res := [][]int{}
	i, k := 0, n-1
	for i < n-2 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if nMap[target-sum] > 0 {

			}
			if sum == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++

			} else if sum < target {
				j++
			} else {

			}

		}
		k--

	}
	return res
}

func smallestRangeI(nums []int, k int) int {

	max := nums[0]
	min := nums[0]
	for _, v := range nums {
		if v >= max {
			max = v
		}
		if v <= min {
			min = v
		}
	}
	if max-min <= 2*k {
		return 0
	}
	return max - min - 2*k
}

/*
Example 1:

Input: s1 = "this apple is sweet", s2 = "this apple is sour"
Output: ["sweet","sour"]
*/
func UncommonFromSentences(s1 string, s2 string) []string {
	s1List := strings.Split(s1, " ")
	s2List := strings.Split(s2, " ")
	s1Map := make(map[string]int, len(s1List))
	s2Map := make(map[string]int, len(s2List))
	for _, v := range s1List {
		s1Map[v]++
	}
	for _, v := range s2List {
		s2Map[v]++
	}
	res := []string{}
	for _, v := range s1List {
		if s1Map[v] == 1 && s2Map[v] < 1 {
			res = append(res, v)
		}
	}
	for _, v := range s2List {
		if s2Map[v] == 1 && s1Map[v] < 1 {
			res = append(res, v)
		}
	}
	return res
}

/*
Example 1:

Input: deck = [1,2,3,4,4,3,2,1]
Output: true
Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].
*/
func HasGroupsSizeX(deck []int) bool {
	dMap := make(map[int]int, len(deck))

	for _, v := range deck {
		dMap[v]++
	}
	tmp := []int{}
	for _, v := range deck {
		if dMap[v] < 2 {
			return false
		}
		tmp = append(tmp, dMap[v])
	}

	//2,4,6,8
	//3,6,3 ==> 12
	//1,2,1 ==> 4
	//key : ==> 3
	//2,2,2,4,4,4,4,4,4,6,6,6
	//key :2,4,6
	//1,2,3,4
	// 4      2    2   2 ==> 10
	//0,0,0,0,1,1,2,2,3,3
	//key : 4
	min := dMap[deck[0]]
	for _, v := range dMap {
		if min > v {
			min = v
		}
	}
	//1,1,1,1,2,2,2,2,2,2
	g := -1
	for _, v := range dMap {
		if g == -1 {
			g = v
		} else {
			g = Gcd(g, v)
		}
	}
	return g >= 2
}
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

/*
Example 1:

Input: num = 5
Output: 2
Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.
*/
func FindComplement(num int) int {

	bNum := ""
	for num > 0 {

		bNum = strconv.Itoa(num%2) + bNum
		num = num / 2
	}
	res := 0
	for _, v := range bNum {

		t, _ := strconv.Atoi(string(v))
		if t == 1 {
			res = res*2 + 0
		} else {
			res = res*2 + 1
		}

	}
	return res
}

/*
Example 1:

Input: s = "5F3Z-2e-9-w", k = 4
Output: "5F3Z-2E9W"
Explanation: The string s has been split into two parts, each part has 4 characters.
Note that the two extra dashes are not needed and can be removed.
*/
func LicenseKeyFormatting(s string, k int) string {

	sList := strings.Split(s, "-")
	res := ""
	count := 0
	//"2-5g-3-J" "2-5G-3J" k==2
	//"2-4A0r7-4k" "24A0-R74K" k==4
	for i:=len(sList)-1;i>=0;i-- {

		if len(sList[i]) < k {

			for j:= len(sList[i])-1;j>=0 ;j--{
				if count==k{
					res =  "-" +res
					count = 0
				}
				if sList[i][j] >= 'a' && sList[i][j] <= 'z' {
					count++
					res =  strings.ToUpper(string(sList[i][j])) + res
				} else if sList[i][j] >='0' && sList[i][j]<='9' ||  (sList[i][j] >='A' && sList[i][j]<='Z'){
					count++
					res = string(sList[i][j]) + res
				}
			}
		} else {

			for j:= len(sList[i])-1;j>=0 ;j--{
				if count==k{
					res =  "-" +res
					count = 0
				}
				if sList[i][j] >= 'a' && sList[i][j] <= 'z' {
					count++
					res =  strings.ToUpper(string(sList[i][j]))+res
				} else if sList[i][j] >='0' && sList[i][j]<='9' ||  (sList[i][j] >='A' && sList[i][j]<='Z') {
					count++
					res = string(sList[i][j]) +res
				}
			}

		}
		if count == k   {
			res =  "-" +res
			count = 0
		}

	}
	if res[0] =='-'{
		res = res[1:]
	}
	return res
}
func FindMaxConsecutiveOnes(nums []int) int {

	max :=0
	res :=0
	//1,1,0,1,1,1

	for i:=0;i<len(nums);i++{
		if nums[i] ==1 {
			max++
		}else  {
			max = 0
		}
		res = getMax(res,max)

	}
	return res
}
func getMax(a,b int)int{
	if a>b{
		return a
	}
	return b

}
/*
   1
  2 3
 */
func FindTilt(root *TreeNode) int {

	if root ==nil{
		return 0
	}
	left := getSum(root.Left)
	right := getSum(root.Right)

	return FindTilt(root.Left) + FindTilt(root.Right) + int(math.Abs(float64(left-right)))


}
func getSum(root *TreeNode)int{

	if root ==nil{
		return 0
	}

	return getSum(root.Left) + getSum(root.Right) + root.Val
}
 type NNode struct {
	   Val int
	     Children []*NNode
	 }
func MaxDepth(root *NNode) int {

	if root ==nil{
		return 0
	}
	m := 0
	for _,v:= range root.Children{
		m = maxx(m,MaxDepth(v))
	}
	return m+1

}
func maxx(a,b int)int{
	if a>b{
		return a

	}
	return b
}