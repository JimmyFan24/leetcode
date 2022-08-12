package level2

import (
	"fmt"
	"sort"
	"strconv"
)

func SetZeroes(matrix [][]int) {

	row := []int{}
	col := []int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}

		}
	}
	//row set zero
	for i := 0; i < len(row); i++ {
		matrix[row[i]] = make([]int, len(matrix[0]))
	}

	//col set zero
	for _, v := range matrix {
		for j := 0; j < len(col); j++ {
			v[col[j]] = 0
		}
	}

}

/*
	1-->2-->3-->4
    |           |
   10-->11-->12-->5
	|           |
    9<--8<--7<--6

*/
func SpiralOrder(matrix [][]int) []int {
	res := []int{}

	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}

	count, sum := 0, n*m

	top, left, buttom, right := 0, 0, m-1, n-1

	for count < sum {
		i, j := top, left

		for j <= right && count < sum {
			res = append(res, matrix[i][j])
			j++
			count++
		}
		i, j = top+1, right
		for i <= buttom && count < sum {
			res = append(res, matrix[i][j])
			i++
			count++
		}
		i, j = buttom, right-1
		for j >= left && count < sum {
			res = append(res, matrix[i][j])
			j--
			count++
		}

		i, j = buttom-1, left
		for i > top && count < sum {
			res = append(res, matrix[i][j])
			i--
			count++
		}

		top, left, buttom, right = top+1, left+1, buttom-1, right-1

	}
	return res
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {

	if root == nil {

		return nil
	}
	queue := []*Node{root}

	for len(queue) != 0 {

		l := len(queue)
		//var p []*Node
		for i, v := range queue {
			if i+1 < l {
				v.Next = queue[i+1]
			} else {
				v.Next = nil
			}

			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[l:]
	}
	return root

}

func GameOfLife(board [][]int) {
	offset := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	zeroList := [][]int{}
	oneList := [][]int{}
	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board[0]); j++ {

			count := 0

			for k := 0; k < len(offset); k++ {
				i0 := i + offset[k][0]
				j0 := j + offset[k][1]
				if inBoard(i0, j0, board) {
					if board[i0][j0] == 1 {
						//fmt.Printf("this is i:%d, this is j:%d=======",i0,j0)
						//fmt.Println(count)
						count++
						if i == 1 && j == 0 {
							fmt.Printf("i = %d,j = %d", i0, j0)
						}
					}
				}
			}
			//fmt.Printf("print count:%d -----",count)
			//fmt.Println(i,j)
			if count > 3 || (count < 2 && board[i][j] == 1) {
				zeroList = append(zeroList, []int{i, j})
			} else if (count == 2 || count == 3) && board[i][j] == 1 {
				continue
			} else if count == 3 && board[i][j] == 0 {
				fmt.Println("turn to  alive")
				oneList = append(oneList, []int{i, j})
			}

		}
	}

	for _, v := range zeroList {
		board[v[0]][v[1]] = 0
	}
	for _, v := range oneList {
		board[v[0]][v[1]] = 1
	}
}
func inBoard(x, y int, board [][]int) bool {

	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func AddStrings(num1 string, num2 string) string {
	if num1 == "" || num2 == "" {
		return ""
	}
	l1 := len(num1) - 1
	l2 := len(num2) - 1

	res := ""
	carry := 0
	for l1 >= 0 && l2 >= 0 {

		b1 := num1[l1]
		b2 := num2[l2]
		n1, _ := strconv.Atoi(string(b1))
		n2, _ := strconv.Atoi(string(b2))

		//93
		//22
		//5 + 0*10
		//4*10 +
		tmp := carry
		res = res + strconv.Itoa((n1+n2+tmp)%10)
		carry = 0
		if (n1+n2+tmp)/10 > 0 {
			carry++
		}
		l1--
		l2--
	}

	for l1 >= 0 {
		n1, _ := strconv.Atoi(string(num1[l1]))
		n1 = n1 + carry
		res = res + strconv.Itoa((n1 % 10))
		carry = 0
		if (n1)/10 > 0 {
			carry++
		}
		l1--
	}
	for l2 >= 0 {
		n2, _ := strconv.Atoi(string(num2[l2]))
		n2 = n2 + carry
		res = res + strconv.Itoa((n2 % 10))
		carry = 0
		if (n2)/10 > 0 {
			carry++
		}
		l2--
	}
	if carry > 0 {

		res = res + strconv.Itoa(carry)
	}
	tmp := ""
	for i := len(res) - 1; i >= 0; i-- {
		tmp = tmp + string(res[i])
	}

	return tmp
}

func SwapPairs(head *ListNode) *ListNode {
	cur := head
	if head.Next == nil {
		return head
	}
	for cur != nil && cur.Next != nil {
		cur.Val, cur.Next.Val = cur.Next.Val, cur.Val
		cur = cur.Next.Next
	}
	return head

}
func CombinationSum(candidates []int, target int) [][]int {

	if len(candidates) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	tmp := []int{}
	sort.Ints(candidates)
	comDfs(candidates, &res,  0, target, tmp)

	return res

}
func comDfs(candi []int, res *[][]int,  index int, target int, tmp []int) {

	if target <=0 {
		if target == 0 {
			t := make([]int,len(tmp))
			copy(t,tmp)
			*res = append(*res, t)
		}
		return
	}
	for i := index; i < len(candi); i++ {
		if candi[i] > target{
			break
		}

		tmp = append(tmp, candi[i])
		comDfs(candi, res,  i, target-candi[i], tmp)
		tmp = tmp[:len(tmp)-1]
	}
	return

}
