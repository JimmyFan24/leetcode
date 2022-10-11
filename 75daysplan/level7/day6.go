package level7

import "math"

func MatrixScore(grid [][]int) int {

	//1.保证第一列全为1
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		if grid[i][0] == 0 {
			oneflipzero(&grid[i])
		}
	}
	//2.其他列的1尽量多
	for i := 1; i < col; i++ {
		oneCount := 0
		zeroCount := 0
		for j := 0; j < row; j++ {
			if grid[j][i] == 1 {
				oneCount++
			}
			if grid[j][i] == 0 {
				zeroCount++
			}
		}
		//如果1的数量小于列总数的一半,那么需要翻转
		if oneCount < zeroCount {
			for j := 0; j < row; j++ {
				if grid[j][i] == 1 {
					grid[j][i] = 0
				} else if grid[j][i] == 0 {
					grid[j][i] = 1
				}
			}
		}
	}
	//3.算总和
	//1,0
	//1*2 + 0
	score := 0
	for i := 0; i < row; i++ {
		sum := 0
		for j := 0; j < col; j++ {
			sum = sum*2 + grid[i][j]
		}
		score += sum
	}
	return score
}
func oneflipzero(arr *[]int) {

	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == 0 {
			(*arr)[i] = 1
		} else if (*arr)[i] == 1 {
			(*arr)[i] = 0
		}
	}
}

func PathInZigZagTree(label int) []int {

	ans := []int{}
	pathInZigZagTree(label, &ans)

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}

	return ans
}
func findHight(label int) int {
	i := 1
	level := 1
	for level < label {
		i++
		level = int(math.Pow(2, float64(i))) - 1
	}
	return i
}
func pathInZigZagTree(label int, ans *[]int) {

	if label <= 1 {
		*ans = append(*ans, label)
		return
	}
	*ans = append(*ans, label)
	//1.找到第几层
	height := findHight(label)
	//2.找到label是第几个

	index := label - int(math.Pow(2, float64(height-1))) + 1
	//15,14,13,12
	if height%2 == 0 {
		index = (int(math.Pow(2, float64(height))) - 1) - label + 1
	}
	//3.找到父节点,father是第i-1层,第j/2
	if index == 0 {
		//说明是最左边的那个
		father := int(math.Pow(2, float64(height-2)))
		pathInZigZagTree(father, ans)
		return
	}
	fatherIndex := index / 2
	if index%2 != 0 {
		fatherIndex = index/2 + 1
	}
	//4.找到第i-1层第father个
	min := int(math.Pow(2, float64(height-2)))
	max := int(math.Pow(2, float64(height-1))) - 1
	father := min + fatherIndex - 1
	if (height-1)%2 == 0 {
		father = max - fatherIndex + 1
	}
	pathInZigZagTree(father, ans)

}
