package level9

import (
	"math"
	"sort"
)

func FourSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	res := [][]int{}
	cur := []int{}

	kSum(nums, 0, len(nums)-1, target, 4, cur, &res)
	return res
}
func kSum(nums []int, left, right int, target int, k int, cur []int, res *[][]int) {
	//边界条件
	//1.总个数小于k
	//2.k<2
	//3. target 小于最小的值的k倍
	//4.target大于最大值的k倍
	if right-left+1 < k || k < 2 || target < nums[left]*k || target > nums[right]*k {
		return
	}
	if k == 2 {
		twoSum(nums, left, right, target, cur, res)
	} else {
		//left 为左边的起点
		for i := left; i < len(nums); i++ {
			//如果i是left，说明是第一个元素，不需要跳过
			//如果i 等于left，但是i-1 ！= i，说明这个i可以取，不会重复，否则这个i需要继续往下找
			if i == left || (i > left && nums[i-1] != nums[i]) {
				next := make([]int, len(cur))
				copy(next, cur)
				//将当前的nums[i]加到数组传下去，比如说现在k是4，应该next只有一个元素，然后传给k=3 的情况
				next = append(next, nums[i])
				//这里可以看到，right的边界一直都是最右，left在增加
				kSum(nums, i+1, len(nums)-1, target-nums[i], k-1, next, res)
			}
		}
	}
}
func twoSum(nums []int, left, right int, target int, cur []int, res *[][]int) {

	//双指针算two sum

	//1.死循环，left<right
	for left < right {

		//2.算下左边界和右边界现在的和，进行判断
		sum := nums[left] + nums[right]
		if sum == target {

			//如果相等，说明找到一对，先加到cur
			cur = append(cur, nums[left], nums[right])
			temp := make([]int, len(cur))
			copy(temp, cur)
			//加到res
			*res = append(*res, temp)
			//cur 还原
			cur = cur[:len(cur)-2]
			left++
			right--
			//由于left 和right已经进行更新，所以下面的left应该和left-1相比，right应该和right+1相比，后者才是之前的值
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}

		} else if sum < target {
			//如果小
			left++
		} else {
			right++
		}
	}
}
func Two(nums []int, target int) []int {
	nMap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, ok := nMap[nums[i]]; ok {
			if nMap[nums[i]] != i {
				return []int{i, nMap[nums[i]]}
			}

		}
		nMap[target-nums[i]] = i

	}
	return nil
}

/*
[[0,0,0,0,0,0,0,0,0,0],
[0,1,1,1,1,1,1,1,1,0],
[0,1,0,0,0,0,0,0,0,0],
[0,1,0,1,1,1,1,1,1,1],
[0,1,0,0,0,0,0,0,0,0],
[0,1,1,1,1,1,1,1,1,0],
[0,1,0,0,0,0,0,0,0,0],
[0,1,0,1,1,1,1,1,1,1],
[0,1,0,1,1,1,1,0,0,0],
[0,1,0,0,0,0,0,0,1,0],
[0,1,1,1,1,1,1,0,1,0],
[0,0,0,0,0,0,0,0,1,0]]
1
*/
func ShortestPath(grid [][]int, k int) int {

	res := math.MaxInt32
	visited := [][]bool{}
	for i := 0; i < len(grid); i++ {
		visited = append(visited, make([]bool, len(grid[0])))
	}
	dfsPath(grid, 0, 0, &res, k, 0, &visited)
	if res == math.MaxInt32 {
		res = -1
	}
	return res
}

var position = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func dfsPath(grid [][]int, x, y int, res *int, k int, tmp int, visited *[][]bool) {

	if x == len(grid)-1 && y == len(grid[0])-1 {
		*res = getTheMinOne(*res, tmp)
		return
	}
	for _, v := range position {
		x0 := x + v[0]
		y0 := y + v[1]
		if inBoard(grid, x0, y0) {

			if !(*visited)[x0][y0] {
				(*visited)[x0][y0] = true
				if grid[x0][y0] == 1 {
					if k > 0 {

						dfsPath(grid, x0, y0, res, k-1, tmp+1, visited)

					}
					(*visited)[x0][y0] = false
				} else {
					dfsPath(grid, x0, y0, res, k, tmp+1, visited)
					(*visited)[x0][y0] = false
				}
			}
		}

	}

}
func inBoard(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])

}
func getTheMinOne(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func CommonFactors(a int, b int) int {

	res := 0
	factor := a
	if a > b {
		factor = b
	}
	for i := 1; i*i < factor; i++ {
		if a%i == 0 && b%i == 0 {
			res++
		}
	}
	return res * 2
}
