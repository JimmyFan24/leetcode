package level6

import (
	"sort"
	"strconv"
	"time"
)

type Twitter struct {
	Users map[int]*User
}
type User struct {
	Id       int
	Flow     []*User
	Fans     []*User
	Contents Contents
}
type Content struct {
	ContentID int
	Time      int64
}
type Contents []Content

func (C Contents) Len() int {
	return len(C)
}

func (C Contents) Less(i, j int) bool {

	return C[i].Time > C[j].Time
}

func (C Contents) Swap(i, j int) {
	C[i], C[j] = C[j], C[i]
}
func Constructor355() Twitter {

	return Twitter{Users: make(map[int]*User)}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {

	if _, ok := this.Users[userId]; !ok {
		user := &User{
			Id:   userId,
			Flow: nil,
			Fans: nil,
		}

		user.Contents = append(user.Contents, Content{
			ContentID: tweetId,
			Time:      time.Now().UnixNano(),
		})
		this.Users[userId] = user
	} else {
		for _, c := range this.Users[userId].Contents {
			if c.ContentID == tweetId {
				return
			}
		}
		this.Users[userId].Contents = append(this.Users[userId].Contents, Content{
			ContentID: tweetId,
			Time:      time.Now().UnixNano(),
		})
	}

	return
}

func (this *Twitter) GetNewsFeed(userId int) []int {

	if _, ok := this.Users[userId]; !ok {
		this.Users[userId] = &User{Id: userId}
		return nil
	}
	user := this.Users[userId]
	flow := user.Flow
	var contentsList Contents

	for _, f := range flow {
		if f != nil {
			for _, c := range f.Contents {
				contentsList = append(contentsList, c)
			}
		}

	}
	for _, v := range user.Contents {
		contentsList = append(contentsList, v)
	}
	sort.Sort(contentsList)

	res := []int{}
	for i := 0; i < 10 && i < len(contentsList); i++ {
		res = append(res, contentsList[i].ContentID)
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {

	if _, ok := this.Users[followeeId]; !ok {
		this.Users[followeeId] = &User{
			Id:       followeeId,
			Flow:     nil,
			Fans:     nil,
			Contents: nil,
		}

	}

	if _, ok := this.Users[followerId]; !ok {
		this.Users[followerId] = &User{
			Id:       followerId,
			Flow:     nil,
			Fans:     nil,
			Contents: nil,
		}
	}

	for _, v := range this.Users[followerId].Flow {
		if v.Id == followeeId {
			return
		}
	}
	this.Users[followerId].Flow = append(this.Users[followerId].Flow, this.Users[followeeId])
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	index := -1
	for i, v := range this.Users[followerId].Flow {
		if v.Id == followeeId {
			index = i
			break
		}
	}
	this.Users[followerId].Flow[index] = nil
}
func EvalRPN(tokens []string) int {

	stack := []string{}
	for _, v := range tokens {

		if v == "+" || v == "-" || v == "/" || v == "*" {

			if len(stack) > 1 {
				num1 := stack[len(stack)-2]
				num2 := stack[len(stack)-1]
				n1, _ := strconv.Atoi(num1)
				n2, _ := strconv.Atoi(num2)
				stack = stack[:len(stack)-2]
				newNumber := caculateNumber(n1, n2, v)
				stack = append(stack, strconv.Itoa(newNumber))

			}
		} else {
			stack = append(stack, v)
		}
	}
	res, _ := strconv.Atoi(stack[0])
	return res
}
func caculateNumber(a, b int, sign string) int {

	res := 0
	switch sign {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}
	return res
}
func TrailingZeroes(n int) int {

	if n/5 == 0 {
		return 0
	}
	return n/5 + TrailingZeroes(n/5)
}
