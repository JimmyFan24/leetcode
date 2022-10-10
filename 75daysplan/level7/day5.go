package level7

import (
	"sort"
)

func FrequencySort(s string) string {
	m := make(map[rune]int, 26)
	for _, v := range s {
		m[v]++
	}

	mm := make(map[int][]rune, len(s))
	for _, v := range s {
		if m[v] > 0 {
			mm[m[v]] = append(mm[m[v]], v)
			delete(m, v)
		}

	}
	freq := []int{}
	for k, _ := range mm {
		freq = append(freq, k)
	}
	res := ""
	sort.Ints(freq)
	for i := len(freq) - 1; i >= 0; i-- {
		f := freq[i]
		for _, v := range mm[freq[i]] {
			for j := 0; j < f; j++ {
				res += string(v)
			}
		}
	}
	return res
}

/*
Input: s = "][]["
Output: 1
Explanation: You can make the string balanced by swapping index 0 with index 3.
The resulting string is "[[]]".
*/
func MinSwaps(s string) int {

	leftstack := []byte{}
	rightstack := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			leftstack = append(leftstack, s[i])
		}
		if s[i] == ']' {
			if len(leftstack) < 1 {
				rightstack = append(rightstack, s[i])
			} else {
				leftstack = leftstack[:len(leftstack)-1]
			}
		}
	}

	lenght := len(leftstack)
	if lenght%2 == 0 {
		return lenght / 2
	} else {
		return lenght/2 + 1
	}
}
func CountSubIslands(grid1 [][]int, grid2 [][]int) int {

	//两者相加,然后1变成2,有2就是岛屿,算出岛屿个数即可

	count := 0
	visited := make([][]bool, len(grid2))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid2[0]))
	}
	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfsIsland(&grid2, &visited, i, j)
			}
		}
	}

	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			if grid2[i][j] == 1 {
				dfsIsland(&grid2, &visited, i, j)
				count++
			}
		}
	}

	return count
}

var Position = [][]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func dfsIsland(grid2 *[][]int, visited *[][]bool, x, y int) {

	if !isInBoard(*grid2, x, y) {
		return
	}
	if (*visited)[x][y] || (*grid2)[x][y] == 0 {
		return
	}
	(*visited)[x][y] = true
	(*grid2)[x][y] = 0
	for _, v := range Position {
		x0 := x + v[0]
		y0 := y + v[1]
		dfsIsland(grid2, visited, x0, y0)
	}

}
func isInBoard(grid3 [][]int, x, y int) bool {

	return x >= 0 && x < len(grid3) && y >= 0 && y < len(grid3[0])
}

func NumOfSubarrays(arr []int, k int, threshold int) int {

	//k个数的平均数大于等于threshold,也就是和要大于等于threhold*k

	//sort.Ints(arr)
	//dp:
	//dp[i] : 表示前i个元素中,满足条件的子数组的个数
	//dp[i] = if arr[i-2]+arr[i-1] + arr[i] >= sum :dp[i] = dp[i]+1 ,else dp[i] = dp[i-1]

	dp := make([]int, len(arr)+1)
	dp[0] = 0

	for i := k; i <= len(arr); i++ {
		sum := 0
		for j := i - 1; j >= i-k; j-- {
			sum += arr[j]
		}
		if sum >= threshold*k {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}
	//dfsNumofSubarrays(arr, k, &count, threshold*k, 1, 1, arr[0])

	return dp[len(arr)]
}
func dfsNumofSubarrays(arr []int, k int, count *int, sum int, start int, end int, tmp int) {

	if end >= k {
		if tmp >= sum {
			*count++
		}
		return
	}

	//2,2, 2,2,5,5,5,8
	//2.2.8
	//2.
	for j := start; j < len(arr); j++ {
		for j+1 < len(arr) && arr[j] == arr[j+1] {
			j++
		}
		tmp += arr[j]
		dfsNumofSubarrays(arr, k, count, sum, j+1, end+1, tmp)
		tmp -= arr[j]
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapNodes(head *ListNode, k int) *ListNode {
	nodes := []*ListNode{}
	tmp := head
	for tmp != nil {
		nodes = append(nodes, tmp)
		tmp = tmp.Next
	}
	nodes[k-1].Val, nodes[len(nodes)-k].Val = nodes[len(nodes)-k-1].Val, nodes[k-1].Val
	return head

}
func MatrixBlockSum(mat [][]int, k int) [][]int {

	ans := make([][]int, len(mat))
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]int, len(mat[0]))
	}

	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans[0]); j++ {

			ans[i][j] = countSum(mat, i, j, k)
		}
	}
	return ans
}
func countSum(mat [][]int, i, j int, k int) int {
	//i - k <= r <= i + k,
	sum := 0
	for r := i - k; r <= i+k && i < len(mat); r++ {
		for c := j - k; c <= j+k && c < len(mat[0]); c++ {
			if isInBoard(mat, r, c) {
				sum += mat[r][c]
			}

		}
	}
	return sum
}

/*
i - k <= r <= i + k,
j - k <= c <= j + k,
*/

func CountTriplets(arr []int) int {

	//1.a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
	//2.b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
	//3. a== b
	//4.0 <= i < j <= k < arr.length
	//1.翻译一下,a==b ,也就是 从i^i+1^...^k == 0
	i, k := 0, len(arr)-1
	count := 0

	for ; i < len(arr)-1; i++ {
		tmpk := k
		for ; tmpk >= i+1; tmpk-- {
			j := i + 1
			for ; j <= tmpk; j++ {

				if countA(arr, i, j) == countB(arr, j, tmpk) {
					count++
					//fmt.Println(i,j,tmpk)
				}

			}
		}

	}

	return count

}
func countA(arr []int, i, j int) int {
	a := 0
	for q := i; q < j; q++ {
		a ^= arr[q]
	}
	return a
}
func countB(arr []int, j, k int) int {
	b := 0
	for q := j; q <= k; q++ {
		b ^= arr[q]
	}
	return b
}
func CountMaxOrSubsets(nums []int) (ans int) {
	//1.位图法,用一个长度为len(nums)的数字的二进制的0,1来表示子集
	// [2,3,4,5,6] ==> 1-1-1-0-0 ==> [2,3,4]
	//2.一共有2^n-1个非空子集
	maxOr := 0
	//这里是全部的子集的二进制表示,一个i表示一个子集
	for i := 1; i < 1<<len(nums); i++ {

		or := 0
		for j, num := range nums {
			//第j个数是否是1,可以用i右移j位再与1 &得到
			if i>>j&1 == 1 {
				or |= num
			}
		}
		if or > maxOr {
			maxOr = or
			ans = 1
		} else if or == maxOr {
			ans++
		}

	}

	return ans

}
