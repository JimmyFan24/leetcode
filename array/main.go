package main

import (
	"fmt"
	"strings"
)

func main()  {
	//var nums = []int{2,7,11,15}
	//fmt.Println(twoSum(nums,9))
	/*s := "abcbc"
	lengthOfLongestSubstring(s)*/
	str := "122"
	list :=strings.Split(str,",")
	fmt.Println(list)

	/*
	  704.
	 */
	nums :=[]int{5}
	search(nums,5)

}