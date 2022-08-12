package mongo

import (
	"encoding/json"
	"fmt"
	"strings"
)

type OpsData struct {
	OpsJson  []*OpsJson
}
/*{
        "op" : "n",
        "ns" : "",
        "o" : {
                "msg" : "periodic noop"
        }
}*/

//var teststr ="\n{\n        \"ts\" : Timestamp(1641466786, 1),\n        \"t\" : NumberLong(10),\n        \"h\" : NumberLong(0),\n        \"v\" : 2,\n        \"op\" : \"n\",\n        \"ns\" : \"\",\n        \"wall\" : ISODate(\"2022-01-06T10:59:46.890Z\"),\n        \"o\" : {\n                \"msg\" : \"periodic noop\"\n        }\n}\n{\n        \"ts\" : Timestamp(1641466796, 1),\n        \"t\" : NumberLong(10),\n        \"h\" : NumberLong(0),\n        \"v\" : 2,\n        \"op\" : \"n\",\n        \"ns\" : \"\",\n        \"wall\" : ISODate(\"2022-01-06T10:59:56.892Z\"),\n        \"o\" : {\n                \"msg\" : \"periodic noop\"\n        }\n}"
//var teststr ="{\n        \"op\" : \"n\",\n        \"ns\" : \"\",\n        \"o\" : {\n                \"msg\" : \"periodic noop\"\n        }\n},{\n        \"op\" : \"n\",\n        \"ns\" : \"2222\",\n        \"o\" : {\n                \"msg\" : \"peri2222odic noop\"\n        }\n}"
func OpsInit(fs string) *OpsData{
	//1.定义字符数据接受输入的json字符串
	//var opsStrArray []string
	opsStrArray := strings.Split(fs,";")

	for _,v :=range opsStrArray{
		ops := OpsJson{}
		err := json.Unmarshal([]byte(v),&ops)
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println(ops)
	}
	return nil
}