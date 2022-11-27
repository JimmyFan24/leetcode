package level10

import (
	"fmt"
	"reflect"
)

func ForTest() {

	arr := []int{1, 2, 3, 4}
	res := []*int{}
	for _, v := range arr {
		fmt.Println(reflect.TypeOf(v))
		res = append(res, &v)
	}
	fmt.Println(res)
}
