package main

//  go run main.go  -name wangk -age 100
// 参考 ： https://www.cnblogs.com/chenqionghe/p/8295807.html
import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "默认A", "姓名")
	age := flag.Int("age", 20, "年龄")

	flag.Parse()
	fmt.Printf("您的姓名是： %s ,年龄: %s ", *name, *age)

}
