package main

import (
	"blueking/pkg/mongo"
	_ "blueking/pkg/util"
	_ "log"
)


type OpJson struct {
	Op string
	Ns string
	O string
}
var str = `{"Id":11,"Name":"王老师","Age":30,"Tno":"t201205"}`
var fs = `{"Op":"n","Ns":"-","O":"====="};{"Op":"0000","Ns":"122","O":"=2="}`
func main() {


	mongo.OpsInit(fs)
	//mongo.RunAnalyse(od)
}
