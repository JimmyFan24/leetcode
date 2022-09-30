package weeklycontest

func HardestWorker(n int, logs [][]int) int {

	hardest := logs[0][0]
	tmp := logs[0][1]
	for i := 1; i < len(logs); i++ {

		if logs[i][1]-logs[i-1][1] > tmp {
			tmp = logs[i][1] - logs[i-1][1]
			hardest = logs[i][0]
		} else if logs[i][1]-logs[i-1][1] == tmp {
			if hardest > logs[i][0] {
				hardest = logs[i][0]
			}
			tmp = logs[i][1] - logs[i-1][1]
		}
	}

	return hardest
}

func FindArray(pref []int) []int {

	dp := make([]int, len(pref))

	//dp[i] 表示 原数组的第i个数字
	//perf[i] = dp[0] ^ dp[1] ^ .... dp[i]

	dp[0] = pref[0]
	for i := len(pref) - 1; i >= 1; i-- {
		dp[i] = pref[i] ^ pref[i-1]
	}
	return dp
}
func RobotWithString(s string) string {

	robot := []rune{}

	paper := ""
	for len(s) > 0 {
		//如果当前robot为空,也入栈
		if len(robot) < 1 {
			robot = append(robot, rune(s[0]))
			s = s[1:]
			continue
		}
		if len(s) > 0 && rune(s[0]) <= robot[len(robot)-1] {
			//如果现在第一个字符比robot最后一个字符小,就入栈
			robot = append(robot, rune(s[0]))
			s = s[1:]
			continue
		}
		if len(s) < 0 && len(robot) > 0 {

			paper += string(robot[len(robot)-1])
			robot = robot[:len(robot)-1]
		}
		//如果s无法出栈,robot无法入栈,就先把robot出栈,直到s可以出栈
		for len(s) > 0 && len(robot) > 0 && rune(s[0]) > robot[len(robot)-1] {
			paper += string(robot[len(robot)-1])
			robot = robot[:len(robot)-1]
		}

	}

	for len(robot) > 0 {
		if len(paper) > 0 {
			if robot[len(robot)-1] < rune(paper[0]) {
				p := ""
				for len(robot) > 0 {
					p += string(robot[len(robot)-1])
					robot = robot[:len(robot)-1]
				}
				paper = p + paper
			} else {
				paper += string(robot[len(robot)-1])
				robot = robot[:len(robot)-1]
			}

		} else {
			paper += string(robot[len(robot)-1])
			robot = robot[:len(robot)-1]
		}

	}
	return paper

}
