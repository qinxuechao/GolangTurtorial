package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	ff, _ := ioutil.ReadFile("/Users/qinxuechao/Desktop/ICBC_beijing_cwl_ICBC_beijing_cwl_map.svg")
	fmt.Println(ff)

	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(ff)
	fmt.Println(sourcestring)


}