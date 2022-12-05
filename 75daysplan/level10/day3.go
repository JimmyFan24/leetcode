package level10

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetMinimumDifference(root *TreeNode) int {

	arr := []int{}
	inOrder(root, &arr)
	res := 0
	for i := 1; i < len(arr); i++ {
		res = min(res, arr[i]-arr[i-1])
	}

	return res
}
func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MatrixReshape(mat [][]int, r int, c int) [][]int {
	row := len(mat)
	col := len(mat[0])
	if row*col != r*c {
		return mat
	}
	res := make([][]int, r)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, c)
	}
	oriCol, oriRow := 0, 0
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			if oriRow < len(mat) && oriCol < len(mat[0]) {
				res[i][j] = mat[oriRow][oriCol]
				oriCol++
			} else if oriRow >= len(mat) {
				//这个时候列已经遍历完了，需要下一列
				oriCol++
				oriRow = 0
				res[i][j] = mat[oriRow][oriCol]
				oriRow++
			} else if oriCol >= len(mat[0]) {
				//这个时候行已经遍历完了，需要跳到下一行
				oriRow++
				oriCol = 0
				res[i][j] = mat[oriRow][oriCol]
				oriCol++
			}
		}
	}

	return res
}
func FindLHS(nums []int) int {

	res := 0
	for before := 0; before < len(nums)-1; before++ {
		count := 1
		max, min := nums[before], nums[before]
		for after := before + 1; after < len(nums); after++ {

			if nums[after] >= max {
				if nums[after]-min <= 1 && nums[after]-min > 0 {
					count++
					max = nums[after]
				}

			} else if nums[after] <= min {
				if max-nums[after] <= 1 && max-nums[after] > 0 {
					count++
					min = nums[after]
				}
			}

		}
		res = getMax(res, count)
	}

	return res
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
