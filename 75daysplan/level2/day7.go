package level2

import (
	"sort"
	"strconv"
)

func AddBinary(a string, b string) string {

	l1 := len(a) - 1
	l2 := len(b) - 1
	// 10 11
	res := ""
	carry := 0
	for l1 >= 0 && l2 >= 0 {
		a1, _ := strconv.Atoi(string(a[l1]))
		b1, _ := strconv.Atoi(string(b[l1]))
		tmp := a1 + b1 + carry
		res = res + strconv.Itoa(tmp%2)
		carry = 0
		if tmp/2 >= 1 {
			carry++
		}
		l1--
		l2--
	}
	for l1 >= 0 {
		a1, _ := strconv.Atoi(string(a[l1]))
		tmp := a1 + carry
		res = res + strconv.Itoa(tmp%2)
		carry = 0
		if tmp/2 == 1 {
			carry++
		}
		l1--
	}
	for l2 >= 0 {
		b1, _ := strconv.Atoi(string(b[l2]))
		tmp := b1 + carry
		res = res + strconv.Itoa(tmp%2)
		carry = 0
		if tmp/2 == 1 {
			carry++
		}
		l2--
	}
	if carry > 0 {

		res = res + strconv.Itoa(carry)

	}
	r := ""
	for i := len(res) - 1; i >= 0; i-- {
		r = r + string(res[i])
	}
	return r
}

func MinimumSum(num int) int {

	nList := []int{}
	//12219 % 10
	//12219 /10 1221 ==> /10 122 ==> 12 ==> 1 ==> 0
	for num > 0 {
		tmp := num % 10
		nList = append(nList, tmp)
		num = num / 10
	}

	//nList [1,2,2,1,9] len == 5 ==>(1,4) 排序然后遍历
	sort.Ints(nList)
	// nList ==> [1,1,2,2,9]
	res := 0
	for i := 1; i < len(nList); i++ {

		if nList[i] == 0 {
			nList = nList[i+1:]
		}
	}
	//2239
	for i := 0; i < len(nList); i++ {

		if len(nList) == 4 {
			res = nList[0]*10 + nList[3] + nList[1]*10 + nList[2]
		} else if len(nList) == 3 {
			res = nList[0]*10 + nList[1] + nList[2]
		} else if len(nList) == 2 {
			res = nList[0] + nList[1]
		}
	}
	return res

}

func SubtractProductAndSum(n int) int {
	nList := []int{}

	for n > 0 {
		tmp := n % 10
		nList = append(nList, tmp)
		n = n / 10
	}
	product := 1
	sum := 0

	for _, v := range nList {

		product = product * v
		sum = sum + v
	}
	return product - sum
}
func RangeSumBST(root *TreeNode, low int, high int) int {



	if root == nil{
		return 0
	}
	sum := 0
	if root.Val >=low && root.Val <=high{
		sum = root.Val
	}


	left := RangeSumBST(root.Left,low,high)
	right := RangeSumBST(root.Right,low,high)

	return left + right + sum

}
func bst(node *TreeNode, sum, low, high int) int {


	if node.Val < low || node.Val > high{
		node.Val = 0
	}
	if node.Left != nil && node.Right != nil {
		left := bst(node.Left, sum+node.Val, low, high)
		right := bst(node.Right, sum+node.Val+left, low, high)
		return sum + left + right
	} else if node.Left == nil && node.Right != nil {
		right := bst(node.Right, sum+node.Val, low, high)
		return sum + right
	} else if node.Left != nil && node.Right == nil {
		left := bst(node.Left, sum+node.Val, low, high)
		return sum + left
	} else {
		return sum + node.Val
	}

}

func CreateTargetArray(nums []int, index []int) []int {
	result := make([]int, len(nums))
	for i, pos := range index {
		copy(result[pos+1:i+1], result[pos:i])
		result[pos] = nums[i]
	}
	return result
}

