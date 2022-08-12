package main

import (
	"sort"
)
type ListNode struct {
	Val int
	Next *ListNode
}
func main()  {

	//MySqrt(8)
	//NextGreatestLetter([]byte{'c','f','j'},'j')

	//3 -- 1 -- 3 -- 4 -- 2
	//FindDuplicate([]int{3,3,4,1,2})
	//min := MinConstructor()

	/*min := MinStack{
		Val:    []int{},
		Lenght: 0,
	}
	min.Push(-2)
	min.Push(0)
	min.Push(-3)
	m := min.GetMin()
	fmt.Println(m)
	min.Pop()
	min.Top()
	min.GetMin()*/

	/*_ = &ListNode{
		Val:  3,
		Next: nil,
	}
	node3 := &ListNode{
		Val:  2,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  1,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}
	DeleteDuplicates(node1)*/
}


func DeleteDuplicates(head *ListNode) *ListNode {


	pList := []*ListNode{}
	for head != nil && head.Next!=nil{

		for  head.Val == head.Next.Val{
			if head.Next.Next != nil{
				head.Next = head.Next.Next
				head = head.Next
			}else{
				head.Next = nil
				break

			}

		}
		//1 2 3 3
		tmp := head.Next
		head.Next = nil
		pList = append(pList,head)
		head = tmp

	}

	for i:=0;i<len(pList)-1;i++{
		pList[i].Next = pList[i+1]
	}
	return pList[0]
}


type MinStack struct {
	Val []int
	Lenght int
}

func MinConstructor() MinStack {

	return MinStack{
		Val:[]int{},
	}
}


func (this *MinStack) Push(val int)  {

	this.Val = append(this.Val,val)
	this.Lenght++
}


func (this *MinStack) Pop()  {

	this.Val = this.Val[:this.Lenght-1]
	this.Lenght--
}


func (this *MinStack) Top() int {
	tmp := this.Val[this.Lenght-1]
	return tmp
}


func (this *MinStack) GetMin() int {

	sort.Ints(this.Val)
	return this.Val[0]
}


