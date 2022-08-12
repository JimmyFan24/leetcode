package level3

import (
	"leetcode/75daysplan/level2"
	"strings"
)

func CanWinNim(n int) bool {

	if n<=3{
		return true
	}
	//1 :win
	//2 :win
	//3 :win
	//4 :lose
	//5 :5-1 ==4 win
	//6 :6-2 ==4 win
	//7 :7-3 ==4 win
	//8 :8-3 ==5 lose

	if n%4==0{
		return false
	}
	return true
}
func CountSegments(s string) int {

	if s=="" {
		return 0
	}
	res := 0

	sList := strings.Split(s," ")
	for _,v:= range sList{
		if v==""{
			continue
		}else{
			res++
		}
	}
	return  res
}
func SumOfLeftLeaves(root *level2.TreeNode) int {

	if root == nil{
		return 0
	}

	if root.Left !=nil && root.Left.Left==nil && root.Left.Right==nil{
		return root.Left.Val + SumOfLeftLeaves(root.Right)
	}
	return SumOfLeftLeaves(root.Left) + SumOfLeftLeaves(root.Right)

}
/*
Example 1:

Input: s = "abcd", t = "abcde"
Output: "e"
Explanation: 'e' is the letter that was added.
 */
func FindTheDifference(s string, t string) byte {
	sMap := make([]rune,26)

	for _,v := range s{
		sMap[v]++
	}
	var res byte
	for _,v := range t{
		sMap[v]--
		if sMap[v]<0{
			res =  byte(v)
		}
	}
	return res
}
func HammingDistance(x int, y int) int {
	dis := 0

	//0111先异或,把相同位置零,不同位置1,然后在循环去低位1,累加次数
	for xor := x^y;xor!=0;xor&=(xor-1){
		dis++
	}
	return dis
}



func LeafSimilar(root1 *level2.TreeNode, root2 *level2.TreeNode) bool {

	list1 :=[]int{}
	list2 := []int{}
	preOrder(root1,&list1)
	preOrder(root2,&list2)
	/*if len(list2) != len(list1){
		return false
	}*/

	return true
}
func preOrder(root *level2.TreeNode ,res *[]int){

	if root == nil{
		return
	}
	preOrder(root.Left,res)
	if root.Left == nil && root.Right==nil{
		*res = append(*res,root.Val)
	}
	preOrder(root.Right,res)

}
func level(root *level2.TreeNode)[][]int{

	res := [][]int{}
	queue := []*level2.TreeNode{root}

	for len(queue)>0{
		l:=len(queue)
		tmp := make([]int,0,l)
		for _,v := range queue{
			if v.Left !=nil{
				queue = append(queue,v.Left)
			}
			if v.Right !=nil{
				queue = append(queue,v.Right)
			}
			if v.Left == nil && v.Right == nil{
				tmp = append(tmp,v.Val)
			}
		}
		res = append(res,tmp)
		queue = queue[l:]
	}
	return res
}
/*
Example 1:

Input: s = "ab", goal = "ba"
Output: true
Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.
 */
func BuddyStrings(s string, goal string) bool {

	if len(s) != len(goal){
		return false
	}
	flag := false
	sMap := make(map[rune]int,26)
	for _,v := range s{
		sMap[v-'a']++
		if sMap[v-'a']>1{
			flag = true
			break
		}
	}
	if s == goal{
		if !flag{
			return false
		}
		return true
	}else {
		p:=-1
		q:=-1

		sByte := []byte(s)
		for i,v := range sByte{
			if p ==-1 && v != goal[i]{
				p = i
				continue
			}
			if q ==-1 && v!= goal[i]{
				q = i
				continue
			}
		}
		if p != -1 && q!= -1{
			sByte[p] ,sByte[q] = sByte[q],sByte[p]
			if string(sByte) == goal{
				return true
			}
			return false
		}
		return false


	}
}

func IncreasingBST(root *level2.TreeNode) *level2.TreeNode {

	head := &level2.TreeNode{}
	tail := head
	preOrder1(root,tail)
	return head.Right
}
func preOrder1(root *level2.TreeNode,newhead *level2.TreeNode) *level2.TreeNode{
	if root == nil{
		return newhead
	}

	newhead = preOrder1(root.Left,newhead)
	root.Left = nil
	newhead.Right,newhead = root,root

	newhead = preOrder1(root.Right,newhead)
	return newhead
}
func SortArrayByParity(nums []int) []int {

	res := []int{}
	for _,v := range nums{
		if v%2==0{
			res =append(res,v)
		}

	}
	for _,v := range nums{
		if v%2!=0{
			res =append(res,v)
		}

	}
	return res
}
func SortArrayByParityII(nums []int) []int {
	res := make([]int,len(nums))
	odd :=0
	even :=1
	for _,v:= range nums{

		if v%2==0{
			res[odd] = v
			odd+=2
		}else{
			res[even] = v
			even+=2
		}
	}
	return res
}
/*
如果单词以元音开头（'a', 'e', 'i', 'o', 'u'），在单词后添加"ma"。
例如，单词 "apple" 变为 "applema" 。
如果单词以辅音字母开头（即，非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。
例如，单词 "goat" 变为 "oatgma" 。
根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从 1 开始。
例如，在第一个单词后添加 "a" ，在第二个单词后添加 "aa" ，以此类推。

 */
func ToGoatLatin(sentence string) string {

	sentenceList := strings.Split(sentence," ")

	res := ""
	for i,v := range sentenceList{
		tmp := ""
		if isVowel(v){
			tmp = v+"ma"
		}else{
			tmp = v[1:] + string(v[0])+ "ma"
		}
		for j:=0;j<i+1;j++{
			tmp = tmp +"a"
		}
		res = res + " " + tmp
	}
	return res
}
func isVowel(s string)bool{
	if s[0] == 'a'|| s[0] == 'e'||s[0] == 'i'||s[0] == 'o'|| s[0] == 'u' || s[0]=='A'|| s[0]=='E' ||
		s[0]=='I' || s[0]=='O' || s[0]=='U'{
		return true
	}

	return false
}

func Transpose(matrix [][]int) [][]int {

	m:=len(matrix)
	n:=len(matrix[0])
	res := make([][]int,n)
	for  i:=0;i<n;i++{
		res[i] = make([]int,m)
	}

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			res[j][i] = matrix[i][j]
		}
	}
	return res

}
func IsMonotonic(nums []int) bool {

	if len(nums)==1{
		return true
	}

	return increase(nums)||discrease(nums)
}
func increase(nums []int)bool{

	for i:=0;i<len(nums)-1;i++{
		if nums[i+1]>=nums[i]{
			continue
		}
		return false
	}
	return true
}
func discrease(nums []int)bool{

	for i:=0;i<len(nums)-1;i++{
		if nums[i+1]<=nums[i]{
			continue
		}
		return false
	}
	return true
}