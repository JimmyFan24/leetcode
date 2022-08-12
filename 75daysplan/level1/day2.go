package level1

func IsIsomorphic(s string, t string) bool {
	strList := []byte(t)
	patternByte := []byte(s)
	if (s == "" && t != "") || (len(patternByte) != len(strList)) {
		return false
	}
	pMap := map[byte]byte{}
	sMap := map[byte]byte{}

	for index,b := range patternByte{
		if _,ok := pMap[b];!ok{
			if _,ok := sMap[strList[index]];!ok{
				pMap[b] = strList[index]
				sMap[strList[index]] = b
			}
		}else {
			if pMap[b] != strList[index]{
				return false
			}
		}

	}
	return true
}


func IsSubsequence(s string, t string) bool {

	tList := []byte(t)
	sList := []byte(s)
	if s == ""{
		return true
	}
	if  len(tList)<len(sList) {
		return false
	}

	tMap := map[byte]byte{}


	for _,b := range tList{
		if _,ok := tMap[b];!ok{
			tMap[b] = 1
		}else{
			tMap[b]++
		}
	}

	for _,b := range sList{
		if _,ok := tMap[b];!ok{
			return false
		}else{

			tMap[b]--
		}
	}
	for _,v := range tMap{
		if v<0{
			return false
		}
	}

	return true

}
type ListNode struct {
	Val int
	Next *ListNode
}
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = MergeTwoLists(l1, l2.Next)
	return l2

}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}