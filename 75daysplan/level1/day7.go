package level1

import "strconv"

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {

	oldcolor  := image[sr][sc]
	dfs(sr,sc,color,oldcolor,image)
	return image

}

func dfs(sr int,sc int,newcolor int,oldcolor int,image [][]int){


	if sr >=len(image)||sc >=len(image[0]) ||sc<0||sr<0{

		return
	}

	if image[sr][sc] == oldcolor && image[sr][sc] != newcolor{
		image[sr][sc] = newcolor
		dfs(sr,sc+1,newcolor,oldcolor,image)
		dfs(sr,sc-1,newcolor,oldcolor,image)
		dfs(sr+1,sc,newcolor,oldcolor,image)
		dfs(sr-1,sc,newcolor,oldcolor,image)
	}
	return

}

func DecodeString(s string) string {

	stack := make([]byte,0,len(s))

	for i:=0;i<len(s);i++{
		by := s[i]

		switch by {
		case ']':
			cnt := make([]byte, 0, len(s))

			// pop string
			for stack[len(stack)-1] != '[' {
				cnt = append(cnt,stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// pop '['
			stack = stack[:len(stack)-1]

			//pop count
			count := ""
			for len(stack) !=0 && stack[len(stack)-1]>= '0' && stack[len(stack)-1] <= '9'{
				count = count + string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]

			}
			num := 0
			for i := len(count) - 1; i >= 0; i-- {
				v, _ := strconv.Atoi(string(count[i]))
				num = num*10 + v
			}
			for i := 0; i < num; i++ {
				for j := len(cnt) - 1; j >= 0; j-- {
					stack = append(stack, cnt[j])
				}
			}

		default:
			stack = append(stack, by)
		}

	}
	return string(stack)
}