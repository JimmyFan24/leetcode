package util

import "io/ioutil"

func Read(path string) (string,error) {
	fs,err := ioutil.ReadFile(path)
	if err != nil {
		return "",err
	}
	return string(fs),err
}
