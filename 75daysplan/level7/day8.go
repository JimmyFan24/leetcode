package level7

import (
	"sort"
	"strconv"
)

/*
Given the array orders, which represents the orders that customers have done in a restaurant.
More specifically orders[i]=[customerNamei,tableNumberi,foodItemi]
where customerNamei is the name of the customer,
tableNumberi is the table customer sit at, and foodItemi is the item customer orders.
Return the restaurant's “display table”. The “display table” is a table
whose row entries denote how many of each food item each table ordered.
The first column is the table number and the
remaining columns correspond to each food item in alphabetical order.
The first row should be a header whose first column is “Table”,
followed by the names of the food items. Note that the customer names are not part of the table.
Additionally, the rows should be sorted in numerically increasing order.

Input: orders = [["David","3","Ceviche"],["Corina","10","Beef Burrito"],["David","3","Fried Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","Ceviche"]]
Output:
[["Table","Beef Burrito","Ceviche","Fried Chicken","Water"],
["3","0","2","1","0"],
["5","0","1","0","1"],
["10","1","0","0","0"]]
Explanation:
The displaying table looks like:
Table,Beef Burrito,Ceviche,Fried Chicken,Water
3    ,0           ,2      ,1            ,0
5    ,0           ,1      ,0            ,1
10   ,1           ,0      ,0            ,0
*/
type Table struct {
	Foods map[string]string
}

func DisplayTable(orders [][]string) [][]string {
	Tables := make(map[string][]string, len(orders))
	header := make(map[string]int, len(orders))
	for _, v := range orders {
		header[v[2]]++
		Tables[v[1]] = append(Tables[v[1]], v[2])
	}
	res := [][]string{}
	h := []string{"Table"}

	for k, v := range header {
		if v > 0 {
			h = append(h, k)
		}
	}
	sort.Strings(h[1:])
	res = append(res, h)
	t := []int{}
	for k, _ := range Tables {
		kk, _ := strconv.Atoi(k)
		t = append(t, kk)
	}
	sort.Ints(t)

	for i := 0; i < len(t); i++ {

		v := Tables[strconv.Itoa(t[i])]
		tmp := make(map[string]int, len(v))
		for _, vv := range v {
			tmp[vv]++

		}
		order := []string{strconv.Itoa(t[i])}
		for i := 1; i < len(h); i++ {
			if tmp[h[i]] > 0 {
				order = append(order, strconv.Itoa(tmp[h[i]]))
			} else {
				order = append(order, strconv.Itoa(0))
			}
		}
		res = append(res, order)
	}

	return res
}
func MaxSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	total := 0

	for i := 1; i < m-1; i++ {
		//每一行的中心的起始点
		for c := 1; c < n-1; c++ {
			centre := grid[i][c]
			for j := -1; j <= 1; j++ {
				centre += grid[i-1][c+j]
				centre += grid[i+1][c+j]
			}
			if centre > total {
				total = centre
			}
		}
	}
	return total

}

/*
You start at the cell (rStart, cStart) of an rows x cols grid facing east. The northwest corner is at the first row and column in the grid, and the southeast corner is at the last row and column.

You will walk in a clockwise spiral shape to visit every position in this grid. Whenever you move outside the grid's boundary, we continue our walk outside the grid (but may return to the grid boundary later.). Eventually, we reach all rows * cols spaces of the grid.

Return an array of coordinates representing the positions of the grid in the order you visited them.
*/
func SpiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {

	position := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	steps, step, count, k, delta := 1, 1, 1, 0, 0
	res := [][]int{}
	res = append(res, []int{rStart, cStart})
	for count < rows*cols {
		step = steps
		for step > 0 {
			rStart += position[k][0]
			cStart += position[k][1]
			if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
				res = append(res, []int{rStart, cStart})
				count++
			}
			step--
		}
		k = (k + 1) % 4
		delta++
		if delta%2 == 0 {
			delta = 0
			steps++
		}
	}
	return res

}

/*
Design the CombinationIterator class:

CombinationIterator(string characters, int combinationLength) Initializes the object with a string
characters of sorted distinct lowercase English letters and a number combinationLength as arguments.
next() Returns the next combination of length combinationLength in lexicographical order.
hasNext() Returns true if and only if there exists a next combination.
*/

type CombinationIterator struct {
	CombinationList []string
}

func Constructor(characters string, combinationLength int) CombinationIterator {

	combinationList := []string{}

	dfsCombination(0, characters, &combinationList, combinationLength, 0, "")

	return CombinationIterator{
		CombinationList: combinationList,
	}

}
func dfsCombination(start int, characters string, arr *[]string, end int, count int, tmp string) {

	if count == end {

		*arr = append(*arr, tmp)
	}
	for j := start; j < len(characters); j++ {
		tmp += string(characters[j])
		dfsCombination(j+1, characters, arr, end, count+1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}

func (this *CombinationIterator) Next() string {

	top := this.CombinationList[0]
	this.CombinationList = this.CombinationList[1:]
	return top
}

func (this *CombinationIterator) HasNext() bool {

	return len(this.CombinationList) > 0
}

/*
You are given an array of people, people,
which are the attributes of some people in a queue (not necessarily in order).
Each people[i] = [hi, ki] represents the ith person of height hi
with exactly ki other people in
front who have a height greater than or equal to hi.

Reconstruct and return the queue that is represented by the input array people.
The returned queue should be formatted as an array queue,
where queue[j] = [hj, kj] is the attributes of the jth person in the queue
(queue[0] is the person at the front of the queue).

Input: people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
Output: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
Explanation:
Person 0 has height 5 with no other people taller or the same height in front.
Person 1 has height 7 with no other people taller or the same height in front.
*/
func ReconstructQueue(people [][]int) [][]int {

	//贪心算法,有两个维度,先确定一个维度

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	return people

}
