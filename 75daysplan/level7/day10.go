package level7

/*
Given a 0-indexed n x n integer matrix grid,
return the number of pairs (Ri, Cj) such that row Ri and column Cj are equal.
A row and column pair is considered equal if
they contain the same elements in the same order (i.e. an equal array).
Explanation: There is 1 equal row and column pair:
- (Row 2, Column 1): [2,7,7]
*/
func EqualPairs(grid [][]int) int {

	count := 0

	rowArray := [][]int{}
	colArray := [][]int{}
	for i := 0; i < len(grid); i++ {
		rowArray = append(rowArray, grid[i])
	}

	for i := 0; i < len(grid[0]); i++ {
		tmp := []int{}
		for j := 0; j < len(grid); j++ {
			tmp = append(tmp, grid[j][i])
		}
		colArray = append(colArray, tmp)
	}
	for i := 0; i < len(rowArray); i++ {
		for j := 0; j < len(colArray); j++ {
			if isEqualArray(rowArray[i], colArray[j]) {
				count++
			}
		}
	}
	return count
}
func isEqualArray(arr1 []int, arr2 []int) bool {

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

/*
You are given an array of strings of the same length words.
In one move, you can swap any two even indexed characters or any two odd indexed characters of a string words[i].

Two strings words[i] and words[j] are special-equivalent if after any number of moves, words[i] == words[j].

For example, words[i] = "zzxy" and words[j] = "xyzz" are special-equivalent because we may make
the moves "zzxy" -> "xzzy" -> "xyzz".
A group of special-equivalent strings from words is a non-empty subset of words such that:

Every pair of strings in the group are special equivalent, and
The group is the largest size possible (i.e., there is not a string words[i]
not in the group such that words[i] is special-equivalent to every string in the group).
Return the number of groups of special-equivalent strings from words.
input: words = ["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
Output: 3
Explanation:
One group is ["abcd", "cdab", "cbad"], since they are all pairwise special equivalent,
and none of the other strings is all pairwise special equivalent to these.
The other two groups are ["xyzz", "zzxy"] and ["zzyx"].
Note that in particular, "zzxy" is not special equivalent to "zzyx".
*/
func NumSpecialEquivGroups(words []string) int {

	equalGroup := make(map[string]int, len(words))
	visited := make([]bool, len(words))

	for i := 0; i < len(words); i++ {
		if !visited[i] {
			equalGroup[words[i]]++
		}
		visited[i] = true
		j := i + 1

		for ; j < len(words); j++ {
			if !visited[j] && isSpecialEqualString(words[i], words[j]) {
				equalGroup[words[i]]++
				visited[j] = true
			}

		}

	}
	res := 0

	for _, v := range equalGroup {
		if v > 0 {
			res++
		}
	}
	return res

}
func isSpecialEqualString(str1 string, str2 string) bool {
	evenMap := make(map[byte]int, 26)
	oddMap := make(map[byte]int, 26)
	for i := 0; i < len(str1); i++ {
		if i%2 == 0 {
			evenMap[str1[i]]++
		} else {
			oddMap[str1[i]]++
		}
	}
	for i := 0; i < len(str2); i++ {
		if i%2 == 0 {
			evenMap[str2[i]]--
		} else {
			oddMap[str2[i]]--
		}
	}
	for _, v := range oddMap {
		if v != 0 {
			return false
		}
	}
	for _, v := range evenMap {
		if v != 0 {
			return false
		}
	}
	return true
}

/*
You have a browser of one tab where you start on the homepage and you can visit another url,
get back in the history number of steps or move forward in the history number of steps.
Implement the BrowserHistory class:

BrowserHistory(string homepage) Initializes the object with the homepage of the browser.
void visit(string url) Visits url from the current page. It clears up all the forward history.
string back(int steps) Move steps back in history.
If you can only return x steps in the history and steps > x,
you will return only x steps.
Return the current url after moving back in history at most steps.
string forward(int steps) Move steps forward in history.
If you can only forward x steps in the history and steps > x,
you will forward only x steps.
Return the current url after forwarding in history at most steps.
*/
type BrowserHistory struct {
	HistoryStack []string
	BackIndex    int
	CurrentPage  string
}

func Constructor1472(homepage string) BrowserHistory {

	return BrowserHistory{
		CurrentPage:  homepage,
		BackIndex:    0,
		HistoryStack: []string{homepage},
	}
}

func (this *BrowserHistory) Visit(url string) {

	if len(this.HistoryStack) == 0 {
		this.HistoryStack = append(this.HistoryStack, url)
		return
	}
	this.HistoryStack = this.HistoryStack[:this.BackIndex+1]
	this.HistoryStack = append(this.HistoryStack, url)
	this.BackIndex++

	return
}

/*
["BrowserHistory","visit","visit","back","visit","visit","back","forward","visit","visit","visit","visit","visit","visit","forward","forward","visit","back","visit","visit","visit","visit","forward","visit","visit","visit"]
[["momn.com"],["bx.com"],["bjyfmln.com"],[3],["ijtrqk.com"],["dft.com"],[10],[10],["yc.com"],["yhl.com"],["xynxvix.com"],["izfscdv.com"],["cdenhm.com"],["ocgcjz.com"],[5],[5],["gtd.com"],[9],["hfeour.com"],["ghmh.com"],["nnm.com"],["knm.com"],[4],["cbtg.com"],["acyvwod.com"],["mydr.com"]]
*/
func (this *BrowserHistory) Back(steps int) string {

	if steps > this.BackIndex {
		this.BackIndex = 0
		resultUrl := this.HistoryStack[0]
		//this.BackIndex -= len(this.HistoryStack)-1
		return resultUrl

	}
	this.BackIndex -= steps
	return this.HistoryStack[this.BackIndex]

}

func (this *BrowserHistory) Forward(steps int) string {

	if this.BackIndex+steps >= len(this.HistoryStack) {
		this.BackIndex = len(this.HistoryStack) - 1
		return this.HistoryStack[len(this.HistoryStack)-1]
	}
	this.BackIndex += steps
	return this.HistoryStack[this.BackIndex]

}
