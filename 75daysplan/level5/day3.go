package level5

import "sort"

func RestoreMatrix(rowSum []int, colSum []int) [][]int {

	//1.先迭代行
	res := make([][]int, len(rowSum))
	for i, _ := range res {
		res[i] = make([]int, len(colSum))
	}
	for i, _ := range rowSum {
		//第一行第一列
		for j := 0; j < len(colSum); j++ {
			if rowSum[i] <= colSum[j] {
				res[i][j] = rowSum[i]
				colSum[j] -= rowSum[i]
				rowSum[i] -= rowSum[i]
				break

			} else {

				res[i][j] = colSum[j]
				rowSum[i] -= colSum[j]
				colSum[j] = 0
			}
		}

	}
	return res
}

type CustomStack struct {
	Stack []int
	Size  int
}

func Constructor1381(maxSize int) CustomStack {

	return CustomStack{Stack: []int{}, Size: maxSize}
}

func (this *CustomStack) Push(x int) {

	if len(this.Stack) < this.Size {
		this.Stack = append(this.Stack, x)
	}
	return

}

func (this *CustomStack) Pop() int {
	tmp := -1
	if len(this.Stack) > 0 {
		tmp = this.Stack[len(this.Stack)-1]
		this.Stack = this.Stack[:len(this.Stack)-1]
	}

	return tmp
}

func (this *CustomStack) Increment(k int, val int) {

	if k >= len(this.Stack) {
		for i, _ := range this.Stack {
			this.Stack[i] += val
		}
	} else {
		for i := 0; i < k; i++ {
			this.Stack[i] += val
		}
	}
	return
}

type CheckIn struct {
	Place string
	Time  int
}
type StationTime struct {
	sum, count float64
}
type UndergroundSystem struct {
	checkIns     map[int]*CheckIn
	stationTimes map[string]map[string]*StationTime
}

func Constructor1396() UndergroundSystem {

	return UndergroundSystem{
		make(map[int]*CheckIn),
		make(map[string]map[string]*StationTime),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {

	this.checkIns[id] = &CheckIn{
		stationName,
		t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {

	checkin := this.checkIns[id]
	destination := this.stationTimes[checkin.Place]
	if destination == nil {
		this.stationTimes[checkin.Place] = make(map[string]*StationTime)
	}
	st := this.stationTimes[checkin.Place][stationName]
	if st == nil {
		st = new(StationTime)
		this.stationTimes[checkin.Place][stationName] = st
	}
	st.sum += float64(t - checkin.Time)
	st.count++
	delete(this.checkIns, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {

	st := this.stationTimes[startStation][endStation]
	return st.sum / st.count
}

func NumOfPairs(nums []string, target string) int {

	count := 0
	for i := 0; i < len(nums)-1; i++ {
		num1 := nums[i]
		j := i + 1
		num2 := ""
		for ; j < len(nums); j++ {
			num2 = nums[j]
			if num1+num2 == target {
				count++
			}
			if num2+num1 == target {
				count++
			}
		}

	}
	return count
}

func CheckIfCanBreak(s1 string, s2 string) bool {

	b1 := []int{}
	b2 := []int{}

	for _, v := range s1 {
		b1 = append(b1, int(v-'a'))
	}
	for _, v := range s2 {
		b2 = append(b2, int(v-'a'))
	}
	sort.Ints(b1)
	sort.Ints(b2)
	i := 0
	for ; i < len(b1); i++ {
		if b1[i] > b2[i] {
			break
		}
	}
	j := 0
	for ; j < len(b1); j++ {
		if b1[j] < b2[j] {
			break
		}
	}
	return i == len(b1) || j == len(b1)
}

func NumTeams(rating []int) int {

	num := 0

	//2,5,3,4,1
	for count := 0; count < len(rating)-2; count++ {

		for j := count + 1; j < len(rating)-1; j++ {
			k := len(rating) - 1
			jj := j
			for ; k > jj; k-- {
				if rating[count] < rating[jj] && rating[jj] < rating[k] {
					num++
				}
				if rating[count] > rating[jj] && rating[jj] > rating[k] {
					num++
				}
			}

		}

	}
	return num
}
