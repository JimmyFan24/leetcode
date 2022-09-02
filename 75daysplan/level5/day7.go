package level5

import "sort"

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func IntToRoman(num int) string {

	res := ""
	var m = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	for num > 0 {
		n := 0
		if num >= 1000 {
			n, res = getRoman(num, 1000, res, m)
		} else if num >= 900 {
			n, res = getRoman(num, 900, res, m)
		} else if num >= 500 {
			n, res = getRoman(num, 500, res, m)
		} else if num >= 400 {
			n, res = getRoman(num, 400, res, m)
		} else if num >= 100 {
			n, res = getRoman(num, 100, res, m)
		} else if num >= 90 {
			n, res = getRoman(num, 90, res, m)
		} else if num >= 50 {
			n, res = getRoman(num, 50, res, m)
		} else if num >= 40 {
			n, res = getRoman(num, 40, res, m)
		} else if num >= 10 {
			n, res = getRoman(num, 10, res, m)
		} else if num >= 9 {
			n, res = getRoman(num, 9, res, m)
		} else if num >= 5 {
			n, res = getRoman(num, 5, res, m)
		} else if num >= 4 {
			n, res = getRoman(num, 4, res, m)
		} else {
			n, res = getRoman(num, 1, res, m)
		}
		num = n
	}
	return res
}
func getRoman(num, k int, res string, m map[int]string) (int, string) {

	tmp := num / k
	for i := 0; i < tmp; i++ {
		res = res + m[k]
	}
	num = num % k
	return num, res
}

func NextPermutation(nums []int) {

	//1.在 nums[i] 中找到 i 使得 nums[i] < nums[i+1]，此时较小数为 nums[i]，并且 [i+1, n) 一定为下降区间
	//2.步：如果找到了这样的 i ，则在下降区间 [i+1, n) 中从后往前找到第一个 j ，使得 nums[i] < nums[j] ，此时较大数为 nums[j]
	//3.第三步，交换 nums[i] 和 nums[j]，此时区间 [i+1, n) 一定为降序区间。最后原地交换 [i+1, n) 区间内的元素，使其变为升序，无需对该区间进行排序
	i, j := 0, 0
	//1,3,2
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(&nums, i+1, len(nums)-1)

}
func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}
func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

/*
Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
*/
func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	tmp := make([]int, len(num1)+len(num2))
	b1, b2 := []byte(num1), []byte(num2)

	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			tmp[i+j+1] += int(b1[i]-'0') * int(b2[j]-'0')
		}
	}
	for i := len(tmp) - 1; i > 0; i-- {
		tmp[i-1] += tmp[i] / 10
		tmp[i] = tmp[i] % 10
	}
	if tmp[0] == 0 {
		tmp = tmp[1:]
	}
	res := make([]byte, len(tmp))
	for i := 0; i < len(tmp); i++ {
		res[i] = '0' + byte(tmp[i])
	}
	return string(res)
}

/*
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
*/
func CombinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	c := []int{}
	res := [][]int{}
	dfsCombination(candidates, &res, target, c, 0)
	return res
}
func dfsCombination(candidates []int, res *[][]int, target int, tmparr []int, start int) {

	if target == 0 {
		tt := make([]int, len(tmparr))
		copy(tt, tmparr)
		*res = append(*res, tt)
		return
	}

	for j := start; j < len(candidates); j++ {
		//排序之后,重复的元素会连着,这里如果是不跳过的话会选到重复的元素
		if j > start && candidates[j] == candidates[j-1] {
			continue
		}
		if target >= candidates[j] {
			tmparr = append(tmparr, candidates[j])
			dfsCombination(candidates, res, target-candidates[j], tmparr, j+1)
			tmparr = tmparr[:len(tmparr)-1]
		}

	}
	return
}
func GenerateMatrix(n int) [][]int {

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	generate(&res, 0, n-1, 0, n-1, n, 1)
	return res
}
func generate(res *[][]int, left, right, top, button int, n int, value int) {

	if left > right || top > button {
		return
	}
	//第一行
	for j := left; j <= right; j++ {
		(*res)[top][j] = value
		value++
	}
	top++
	//最后一列
	for j := top; j <= button; j++ {
		(*res)[j][right] = value
		value++
	}
	right--
	//最后一行
	for j := right; j >= left; j-- {
		(*res)[button][j] = value
		value++
	}
	//第一列
	button--
	for j := button; j >= top; j-- {
		(*res)[j][left] = value
		value++
	}
	left++
	generate(res, left, right, top, button, n, value)

}
