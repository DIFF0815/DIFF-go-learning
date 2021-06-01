package main

import "fmt"

func main()  {
	////输出
	//print("我是print输出")
	//print("我是print输出2")
	//print("\n")
	//println("我是println输出")
	//println("我是println输出2")

	//fmt.Print("我是fmt.Print输出")
	//fmt.Print("我是fmt.Print输出2")
	//fmt.Print("\n")
	//fmt.Println("我是fmt.Println输出")
	//fmt.Println("我是fmt.Println输出2")
	//fmt.Println("我是fmt","Println输出3")
	//fmt.Printf("我是%s,现在在%s,%d岁,身高比例%.2f,是一个100%%好人","diff","gz",18,20.234)


	//数据类型
	//fmt.Println(123)
	//////fmt.Println("1"+2) //不能这样加 会报错
	//fmt.Println("1" + "2")
	//fmt.Println(1+2)
	//
	//fmt.Println(1==1)
	//fmt.Println(1>2)
	//fmt.Println(1<2)
	//fmt.Println("Fa"=="fa")
	//fmt.Println("1"=="01")

	//变量
	//var st string = "中国人"
	//fmt.Println(st)
	//var sd int = 123
	//fmt.Println(sd)
	//var flag bool
	//flag = true
	//fmt.Println(flag)

	//var name string
	//fmt.Scanf("%s",&name)
	//fmt.Println(name)
	//if name == "中国人"{
	//	fmt.Println("输入正确")
	//}else {
	//	fmt.Println("输入错误")
	//}

	//变量简写
	//var name string = "lisi"
	//fmt.Println(name)
	//var name2 = "zhangsan"
	//fmt.Println(name2)
	//name3 := "wangwu"
	//fmt.Println(name3)
	//
	//var name4,name5,name6 string
	//name4 = "4"
	//name5 = "5"
	//name6 = "6"
	//fmt.Println(name4,name5,name6)
	//
	//var (
	//	name7 = "7"
	//	name8 bool
	//	name9 = 12345
	//	name10 string //默认 ""
	//	name11 int  //默认 0
	//	name12 bool //默认 false
	//)
	//name8 = true
	//fmt.Println(name7,name8,name9,name10,name11,name12)


	//var v1,v2,v3 int
	//fmt.Println(v1,v2,v3)
	//var (
	//	v6 = 123
	//	v7 string
	//)
	//fmt.Println(v6,v7)
	////var v4,v5 = 11,12
	////v4,v5 = 11,12 //错误
	//v4,v5 := 11,12
	//fmt.Println(v4,v5)

	//number :=88
	//if true{
	//	number := 99
	//	fmt.Println(number)
	//}
	//fmt.Println(number)

	//常量
	const age = 1
	fmt.Println(age)

	const (
		name = "lisi"
		scroe = 99
	)
	fmt.Println(name)
	fmt.Println(scroe)

	const (
		v1 = iota + 1
		v2
		v3
		v4
	)
	fmt.Println(v1,v2,v3,v4)


}
