package level4

import (
	"math"
	"sort"
)

func MemLeak(memory1 int, memory2 int) []int {

	res := []int{}

	start := 1
	for memory1 >= 0 && memory2 >= 0 {

		if memory1 >= memory2 {
			if start <= memory1 {
				memory1 -= start
				start++
			} else {
				break
			}
		} else {
			if start <= memory2 {
				memory2 -= start
				start++
			} else {
				break
			}
		}

	}
	res = []int{start, memory1, memory2}
	return res
}
func QueensAttacktheKing(queens [][]int, king []int) [][]int {

	canKill := [][]int{}

	canKillRow(queens, king, &canKill)
	canKillRal(queens, king, &canKill)
	otherCanKill(queens, king, &canKill)
	return canKill

}
func canKillRow(queen [][]int, king []int, cankill *[][]int) {

	//1.先扫描行
	rowRight := 8
	rowLeft := -1
	for _, v := range queen {
		if v[0] == king[0] && v[1] > king[1] {

			if rowRight > v[1] {
				rowRight = v[1]
			}
		} else if v[0] == king[0] && v[1] < king[1] {
			if rowLeft < v[1] {
				rowLeft = v[1]
			}
		}
	}
	if rowRight != 8 {
		*cankill = append(*cankill, []int{king[0], rowRight})
	}
	if rowLeft != -1 {
		*cankill = append(*cankill, []int{king[0], rowLeft})
	}

}

func canKillRal(queen [][]int, king []int, cankill *[][]int) {

	//扫描列
	ralTop := -1
	ralButtom := 8
	for _, v := range queen {
		if v[1] == king[1] && v[0] > king[0] {

			if ralButtom > v[0] {
				ralButtom = v[0]
			}
		} else if v[1] == king[1] && v[0] < king[0] {
			if ralTop < v[0] {
				ralTop = v[0]
			}
		}
	}
	if ralButtom != 8 {
		*cankill = append(*cankill, []int{ralButtom, king[1]})
	}
	if ralTop != -1 {
		*cankill = append(*cankill, []int{ralTop, king[1]})
	}

}
func otherCanKill(queen [][]int, king []int, cankill *[][]int) {
	//扫描列
	leftTop := []int{-1, -1}
	leftButtom := []int{8, -1}
	rightTop := []int{-1, 8}
	rightButtom := []int{8, 8}
	for _, v := range queen {

		//右下角
		if v[0] > king[0] && v[1] > king[1] {
			if canGetKing(v, king) {
				if rightButtom[1] > v[1] {
					rightButtom[0] = v[0]
					rightButtom[1] = v[1]
				}

			}

			//左上角
		} else if v[0] < king[0] && v[1] < king[1] {
			if canGetKing(v, king) {
				if leftTop[1] < v[1] {
					leftTop[0] = v[0]
					leftTop[1] = v[1]
				}

			}
			//右上角
		} else if v[0] < king[0] && v[1] > king[1] {
			if canGetKing(v, king) {
				if rightTop[1] > v[1] {
					rightTop[0] = v[0]
					rightTop[1] = v[1]
				}

			}
			//左下角
		} else if v[0] > king[0] && v[1] < king[1] {
			if canGetKing(v, king) {
				if leftButtom[1] < v[1] {
					leftButtom[0] = v[0]
					leftButtom[1] = v[1]
				}

			}
		}

	}
	if rightButtom[0] != 8 {
		*cankill = append(*cankill, rightButtom)
	}
	if leftTop[1] != -1 {
		*cankill = append(*cankill, leftTop)
	}
	if leftButtom[1] != -1 {
		*cankill = append(*cankill, leftButtom)
	}
	if rightTop[0] != -1 {
		*cankill = append(*cankill, rightTop)
	}

	*cankill = append(*cankill, rightButtom)

}
func canGetKing(queen []int, king []int) bool {
	return int(math.Abs(float64(king[0])-float64(queen[0]))) == int(math.Abs(float64(king[1])-float64(queen[1])))
}

type SmallestInfiniteSet struct {
	InfiniteSet []int
}

func NewConstructor() SmallestInfiniteSet {
	arr := []int{}
	count := 1
	for i := 0; i < 1001; i++ {
		arr = append(arr, count)
		count++
	}
	return SmallestInfiniteSet{InfiniteSet: arr}
	//["SmallestInfiniteSet","addBack","popSmallest","popSmallest","popSmallest","addBack","popSmallest","popSmallest","popSmallest"]
	//[[],[2],[],[],[],[1],[],[],[]]
}
func (this *SmallestInfiniteSet) PopSmallest() int {

	tmp := this.InfiniteSet[0]
	this.InfiniteSet = this.InfiniteSet[1:]
	return tmp

}
func (this *SmallestInfiniteSet) AddBack(num int) {

	find := binarySearch(this.InfiniteSet, 0, len(this.InfiniteSet)-1, num)
	if find == -1 {

		this.InfiniteSet = append(this.InfiniteSet, num)
		sort.Ints(this.InfiniteSet)

	}

}
func binarySearch(arr []int, l, r int, num int) int {

	for l < r {
		mid := l + ((r - l) >> 2)
		if arr[mid] > num {
			r = mid - 1
		} else if arr[mid] < num {
			l = mid + 1
		}
		if arr[mid] == num {
			return mid
		}
	}

	if l >= r && arr[l] != num {
		return -1
	}

	return l
}

/*func pop(root *SmallestInfiniteSet)int{
	if root.Left != nil{
		tmp := root.Val
		root = nil
		return tmp
	}
	return pop(root.Left)
}
func (this *SmallestInfiniteSet) AddBack(num int) {

	if !getNum(this,num){
		insertNum(this,num)
	}
}
func insertNum(root *SmallestInfiniteSet,num int){
	if root == nil{
		return
	}
	if root.Val > num{

		if root.Right == nil{
			root.Right = &SmallestInfiniteSet{
				Val:   num,
				Left:  nil,
				Right: nil,
			}
		}else {
			insertNum(root.Right,num)
		}
	}else{
		if root.Left == nil{
			root.Right = &SmallestInfiniteSet{
				Val:   num,
				Left:  nil,
				Right: nil,
			}
		}else {
			insertNum(root.Left,num)
		}
	}

}
func build(arr []int,l,r int) *SmallestInfiniteSet{

	if l >= r{
		return nil
	}
	mid := l + ((r-l)>>2)
	root := &SmallestInfiniteSet{
		Val: arr[mid],
	}
	left := build(arr,l,mid-1)
	right := build(arr,mid+1,r)
	root.Left  = left
	root.Right = right
	return root

}
func getNum(root *SmallestInfiniteSet, num int) bool {

	if root == nil {
		return false
	}
	if root.Val < num {
		return getNum(root.Right, num)
	} else if root.Val > num {
		return getNum(root.Left, num)
	}

	return true

}*/
