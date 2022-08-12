package level1

import (
	"fmt"
	"regexp"
	"strconv"
)

func MinPartitions(n string) int {
	bi := len(n)

	sum := 0
	tmp := 1
	nu :=0
	on :=0
	for tmp>=1{
		num := ""
		for i:=0;i<bi;i++{
			num += "1"
		}
		nu,_ = strconv.Atoi(num)
		on,_ = strconv.Atoi(n)
		tmp = on - nu
		n = strconv.Itoa(tmp)
		sum++
		bi = len(n)
	}
	reg := regexp.MustCompile(`\(\)`)
	st := "G()()()()(al)"
	reg.ReplaceAllString(st,"o")
	fmt.Println()
	return sum
}

func Makesquare(matchsticks []int) bool {

	sum :=0
	max :=0
	for _,v := range matchsticks{
		sum +=v
		if v>max{
			max = v
		}
	}
	tmp := sum % 4
	width := sum / 4
	if tmp !=0{
		return false
	}

	if width < max{
		return false
	}
	size := make([]int,4)

	return backtrack(matchsticks,size,0,width)
}

func backtrack(sticks[]int,size[]int,index int,width int)bool{
	if (index == len(sticks)){
		if (size[0] == size[1]&&size[1]==size[2]&&size[2]==size[3]){
			return true
		}else{
			return false
		}
	}

	for i:=0;i<len(size);i++{
		if size[i] + sticks[index] > width{
			continue
		}

		size[i] += sticks[index]
		next := backtrack(sticks,size,index+1,width)
		if next == true{
			return true
		}
		size[i] -= sticks[index]
	}
	return false
}

func SingleNumber(nums []int) int {
	nMap := make(map[int]int,len(nums))
	for _,v := range nums{
		if _,ok := nMap[v];!ok{
			nMap[v] = 1
		}else{
			nMap[v] +=1
		}

	}
	res := 0
	for i,v := range nMap{
		if v == 1{
			res = i
		}
	}
	return res
}

func Generate(numRows int) [][]int {
	res := [][]int{}
	res = append(res,[]int{1})

	for i:=1;i<numRows;i++{
		tmp := make([]int,i+1)
		for j:=0;j<len(tmp);j++{
			if j-1<0 {
				tmp[j] = res[i-1][j]
			} else if j>=len(res[i-1]){
				tmp[j] = res[i-1][j-1]
			} else {
				tmp[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res = append(res,tmp)


	}
	return res
}

/*func HammingWeight(num uint32) int {
	sn := strconv.FormatUint(num,2)
	res :=0
	for _,v := range sn {
		if v == '1'{
			res+=1
		}
	}
	return res
}*/
func MoveZeroes(nums []int)  {

	queue := []int{}
	for i:=0;i<len(nums);i++{

		if nums[i]!=0{
			queue = append(queue,nums[i])
		}

	}

	for i:=0;i<len(nums);i++{
		nums[i] = queue[i]
	}

}