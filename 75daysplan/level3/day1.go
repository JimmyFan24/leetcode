package level3

import "strconv"

func ConvertToTitle(columnNumber int) string {

	res := ""

	for  columnNumber >0 {
		tmp := columnNumber % 2
		//div := columnNumber / 26
		if tmp == 0{
			tmp =2
		}

		res = string('1'+tmp-1) + res

		columnNumber = columnNumber/26

	}

	return res
}

func ContainsNearbyDuplicate(nums []int, k int) bool {

	nMap := make(map[int]int,len(nums))
	res := false
	for i:=0;i<len(nums);i++{

		if _,ok := nMap[nums[i]];!ok{
			nMap[nums[i]] = i
		}else{
			tmp := nMap[nums[i]]
			if i- tmp >0{
				if i-tmp<=k{
					res = true
				}else {
					nMap[nums[i]] = i
					res =false
				}


			}else{

				if i-tmp<=k{
					res = true
				}else {
					nMap[nums[i]] = i
					res =false
				}
			}
		}

	}
	return res
}

func SummaryRanges(nums []int) []string {

	res := []string{}

	for i:=0;i<len(nums);i++{
		left := i
		right:= i
		for j:=left;j<len(nums)-1;j++{

			if nums[j]+1 == nums[j+1]{
				right++
			}else{
				break
			}

		}

		l := strconv.Itoa(nums[left])
		r := strconv.Itoa(nums[right])
		if right>left{
			res = append(res,l+"->"+r)
		}else{
			res = append(res,strconv.Itoa(nums[right]))
		}

		i = right
	}
	return res
}

func AddDigits(num int) int {


	for num /10>0{
		sum := num
		tmp := 0
		for sum > 0{

			tmp = tmp + sum %10
			sum = sum/10

		}
		num = tmp
	}
	return num
}