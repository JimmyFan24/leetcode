package level2

import (
	"math"
	"strconv"
)

func Flatten(root *TreeNode) {
	list := []*TreeNode{}
	if root == nil {
		return
	}
	//1,null,2,3
	/*
			1
		      2
		     3
	*/
	list = append(list, root)
	newpreorder(root.Left, &list)
	newpreorder(root.Right, &list)
	//  2-->3-->4 2
	// 3 4          3

	for i := 0; i < len(list)-1; i++ {
		list[i].Left = nil
		list[i].Right = list[i+1]
	}
}
func newpreorder(root *TreeNode, list *[]*TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil && root.Right != nil {
		(*list) = append((*list), root)
		newpreorder(root.Left, list)
		newpreorder(root.Right, list)
	} else if root.Left == nil && root.Right != nil {
		(*list) = append((*list), root)
		newpreorder(root.Right, list)
	} else if root.Left != nil && root.Right == nil {
		(*list) = append((*list), root)
		newpreorder(root.Left, list)
	} else {
		(*list) = append((*list), root)
	}
	return
}

/*
	1
	11
	21
	1211
	f(n) = f(n-1)<count + element>
*/
func CountAndSay(n int) string {

	//n=4

	if n == 1 {
		return "1"
	}
	res := "1"

	return  countAndSay(n, res, 1)
}
func countAndSay(n int, res string, index int) string {

	if index == n {
		return res
	}
	strMap := make(map[string]int, len(res))
	for _, v := range res {
		strMap[string(v)]++
	}
	tmp := res
	res = ""


	//1211 211
	//111221
	for i:=0;i<len(tmp);i++{
		if i+1 <len(tmp){
			if tmp[i] != tmp[i+1] && strMap[string(tmp[i])] >0 {
				res += "1"
				strMap[string(tmp[i])]--
				res += string(tmp[i])
			} else {
				j:=i+1

				for ;j<len(tmp);{
					if tmp[i] == tmp[j]{
						j++
					}else{
						break
					}
				}
				if j-i == strMap[string(tmp[i])]{
					res += strconv.Itoa(strMap[string(tmp[i])])
					res += string(tmp[i])
				}else{
					strMap[string(tmp[i])] =  strMap[string(tmp[i])] - j+i
					res += strconv.Itoa(j-i)
					res += string(tmp[i])
				}
				i = j-1

			}
		}else{
			res += strconv.Itoa(strMap[string(tmp[i])])
			res += string(tmp[i])
		}



	}

	return countAndSay(n, res,index+1)

}

func IsValidBST(root *TreeNode) bool {


	return validBST(root,math.Inf(-1),math.Inf(1))


}
func validBST(root *TreeNode,min,max float64) bool{

	if root ==nil {
		return true
	}

	v := float64(root.Val)

	return v > min && v<max && validBST(root.Left,min,v) && validBST(root.Right,v,max)
}

func CanCompleteCircuit(gas []int, cost []int) int {
	//gas = [1,2,3,4,5], cost = [3,4,5,1,2]
	gasSum := 0
	costSum := 0
	for _,v := range gas{
		gasSum+=v
	}
	for _,v := range cost{
		costSum+=v
	}
	if gasSum < costSum{
		return -1
	}
	res := 0
	total := 0
	for i:=0;i<len(gas);i++{

		total += gas[i]-cost[i]
		if total <0{
			res = i+1
			total = 0
		}
	}
	return res
}





