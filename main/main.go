package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"strings"
)

func main() {
	var v string
	v = "file,ffffff,aaa"
	b := strings.SplitN(v, ",", 3)
	fmt.Println(b)
	logs.Info(b)
}
