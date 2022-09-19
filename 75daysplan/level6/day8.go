package level6

import (
	"math"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp, res := make([]int, len(nums)), []int{}
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 1, 1
	for i := 1; i < len(nums); i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}
	if maxSize == 1 {
		return []int{nums[0]}
	}
	for i := len(nums) - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return res
}
func LargestDivisibleSubset(nums []int) []int {
	//1.首先将nums数组升序排序
	sort.Ints(nums)
	//2.定义状态：dp[i]表示以nums[i]为最大整数的最大可整除子集的大小，即一定包含nums[i]
	dp := make([]int, len(nums))
	//3.初始化dp[i]=1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
	}
	maxSize, maxVal := 1, 1
	for i := 1; i < len(nums); i++ {
		//4.对0,...,i-1的数进行枚举，如果可整除nums[i]，说明可被添加到以nums[i]为最大整数的可整除子集中，dp[i]++
		for j, v := range nums[:i] {
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}

		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}
	if maxSize == 1 {
		return []int{nums[0]}
	}
	res := []int{}
	//倒序遍历dp数组，找到最大的可整除子集个数dp[i]，
	//此时nums[i]一定在结果数组中；再找dp[j]=dp[i]-1且dp[j]可以被dp[i]整除的项，将dp[j]加入结果集；
	//令dp[i]=dp[j]，继续寻找，直到dp[i]=0为止。
	for i := len(nums) - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return res
}
func dfsLaegest(res *[]int, start int, nums []int, end int, count int, tmp []int) {

	if start > len(nums) {
		return
	}
	if len(tmp) > len(*res) {
		c := make([]int, len(tmp))
		copy(c, tmp)
		*res = c
	}

	for j := start; j < len(nums); j++ {
		if len(tmp) == 0 {
			tmp = append(tmp, nums[j])
			dfsLaegest(res, j+1, nums, end, count+1, tmp)
			tmp = tmp[:len(tmp)-1]
		} else if isDivide(tmp[len(tmp)-1], nums[j]) {
			tmp = append(tmp, nums[j])
			dfsLaegest(res, j+1, nums, end, count+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}

	}
}
func isDivide(num1 int, num int) bool {

	return num%num1 == 0
}

/*func LargestDivisibleSubset2(nums []int) []int {

	sort.Ints(nums)
	//9,18,54,90,180,360,540,720
	//1, 2, 6,10,20,40, 60, 80
	dp := make([][][]int,len(nums)+1)
	dp[1] = append(dp[1],[]int{nums[0]})

	for i:=2;i<=len(nums);i++{

		maxSubset := []int{}
		maxLength := math.MinInt32

		for _,d := range dp[i-1]{

			if isDivide(d,nums[i-1]){
				if len(d)+1 >= maxLength{
					copy(maxSubset,d)
					maxSubset = append(maxSubset,nums[i-1])
					dp[i] = append(dp[i],maxSubset)
					maxLength = len(maxSubset)
				}

			}
		}


	}
	return dp[len(nums)][0]
}*/

func UniquePathsIII(grid [][]int) int {

	visited := make([][]bool, len(grid))
	for i, _ := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	startx, starty := 0, 0
	for i, v := range grid {
		for ii, vv := range v {
			if vv == 1 {
				startx = i
				starty = ii
			}
		}
	}
	//0,1
	//2,0
	visited[startx][starty] = true
	res := 0
	dfsPath(&visited, &res, grid, startx, starty)
	return res

}

var optionPath = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func dfsPath(visited *[][]bool, res *int, grid [][]int, startx, starty int) {
	if startx >= len(grid) || starty >= len(grid[0]) {
		return
	}
	if grid[startx][starty] == 2 {
		for i, v := range *visited {
			for ii, vv := range v {
				if !vv && grid[i][ii] == 0 {
					return
				}
			}
		}
		*res++
		return
	}
	for i := 0; i < len(optionPath); i++ {
		x0 := startx + optionPath[i][0]
		y0 := starty + optionPath[i][1]
		if inBoard(x0, y0, grid) && !(*visited)[x0][y0] {
			(*visited)[x0][y0] = true
			dfsPath(visited, res, grid, x0, y0)
			(*visited)[x0][y0] = false
		}
	}
}
func inBoard(x, y int, grid [][]int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] != -1
}

/*
Input: words = ["dog","cat","dad","good"],
letters = ["a","a","c","d","d","d","g","o","o"],
score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
Output: 23
Explanation:
Score  a=1, c=9, d=5, g=3, o=2
Given letters, we can form the words "dad" (5+1+5) and "good" (3+2+2+5) with a score of 23.
Words "dad" and "dog" only get a score of 21.
*/

func MaxScoreWords(words []string, letters []byte, score []int) int {

	scoreMap := make(map[int]int, 26)
	for i, v := range score {
		scoreMap['a'+i] = v
	}
	letterMap := make(map[byte]int, len(letters))
	for _, v := range letters {
		letterMap[v]++
	}
	res := math.MinInt32

	maxscoredfs(&res, letterMap, scoreMap, words, 0, 0)

	return res
}
func maxscoredfs(res *int, letterMap map[byte]int, scoreMap map[int]int, words []string, start int, score int) {
	//dog
	if start >= len(words) {
		if score > *res {
			*res = score
		}
		return
	}
	for j := start; j < len(words); j++ {
		i := 0
		wordsocre := 0
		for ; i < len(words[j]); i++ {
			if letterMap[words[j][i]] >= 1 {
				wordsocre += scoreMap[int(words[j][i])]
				letterMap[words[j][i]]--
			} else {
				break
			}
		}
		if i != len(words[j]) {
			for q := 0; q < i; q++ {
				letterMap[words[j][q]]++

			}
			wordsocre = 0
			maxscoredfs(res, letterMap, scoreMap, words, j+1, score+wordsocre)
		} else {
			maxscoredfs(res, letterMap, scoreMap, words, j+1, score+wordsocre)

			for q := 0; q < len(words[j]); q++ {
				letterMap[words[j][q]]++
			}
		}

	}

}
