package main

import "fmt"

func main() {

	//fmt.scan 一直等待直到输入个数等于要捕获的个数
	//var name string
	//var age int
	//fmt.Println("请输入用户名")
	//count,err := fmt.Scan(&name,age)
	//fmt.Println(name)
	//fmt.Println(age)
	//fmt.Println(count,err)

	//fmt.Scanln  等回车就结束
	//var name string
	//var age int
	//fmt.Println("请输入用户名")
	//count,err := fmt.Scanln(&name,age)
	//fmt.Println(name)
	//fmt.Println(age)
	//fmt.Println(count,err)

	//fmt.Scanf
	var name string
	var age int
	fmt.Println("请输入用户名")
	count,err := fmt.Scanf("我叫%s 今年%d岁",&name,&age)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(count,err)

}
