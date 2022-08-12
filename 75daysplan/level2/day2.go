package level2

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res   = []string{}
	final = 0
)

func LetterCombinations(digits string) []string {

	if len(digits) == 0{
		return []string{}
	}
	//res   = []string{}
	dfs(&digits,0,"")
	return res

}
func dfs(digits *string,index int,s string){

	if index == len(*digits){
		res = append(res,s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num -'0']
	for i:=0;i<len(letter);i++{

		dfs(digits,index+1,s + string(letter[i]))
	}
	return

}
