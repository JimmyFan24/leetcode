package level6

func MinNumberOperations(target []int) int {
	amount := target[0]

	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			amount += target[i] - target[i-1]
		}
	}
	return amount
}

func NumSubmatrixSumTarget(matrix [][]int, target int) int {
	arr := []int{}
	row := len(matrix)
	rol := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < rol; j++ {
			arr = append(arr, matrix[i][j])
		}
	}
	res := 0
	for i := 1; i <= len(arr); i++ {
		dfsSumTarget(arr, &res, 0, i, target, 0)
	}

	return res
}
func dfsSumTarget(arr []int, sum *int, start int, end int, target int, count int) {

	if count >= end {
		if target == 0 {
			*sum++
		}
	}

	for j := start; j < len(arr); j++ {

		if target-arr[j] >= 0 {
			target -= arr[j]
			dfsSumTarget(arr, sum, j+1, end, target, count+1)
			target += arr[j]
		}
	}
}
