package main

import "fmt"

/*
1.时间复杂度o(N)
2.空间复杂度o(1)
要求:找出一组数组里面的出现奇数次的数字
 */
func findOddNum(arr[]int)(int){
	var eor = 0
	for _,a := range arr{
		eor ^= a
	}
	return eor
}
func findOddNumTwo(arr[]int)(int,int){
	var eor = 0
	for _,a := range arr{
		eor ^= a

	}

	rightone := eor & ( (^eor) + 1 )
	fmt.Printf("--%b--",rightone)
	fmt.Printf("--%b--",eor)
	var onlyone = 0
	for _,b := range arr{
		if (b & rightone) != 0{
			onlyone ^= b
		}
	}
	return onlyone,onlyone ^ eor
}
func main(){
	arr :=[]int {1,1,2,3,3,3,3,9}
	num1,num2 := findOddNumTwo(arr)
	fmt.Println(num1,num2)




}