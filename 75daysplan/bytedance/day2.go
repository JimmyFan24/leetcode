package main

type MyQueue struct {
	Val []int

}


func Constructor() MyQueue {

	m := &MyQueue{
		Val :[]int{},
	}

	return *m
}


func (this *MyQueue) Push(x int)  {

	this.Val = append(this.Val,x)
}


func (this *MyQueue) Pop() int {


		tmp := this.Val[0]
		this.Val = this.Val[1:]
		return tmp


}


func (this *MyQueue) Peek() int {

		return this.Val[0]

}


func (this *MyQueue) Empty() bool {

	return len(this.Val) == 0
}
