package main

import "fmt"

func main() {

	//索引切片和循环
	var name string = "中国人"
	v1 := name[0:1]
	fmt.Println(v1)
	v2 := name[0:3]
	fmt.Println(v2)

	for i:=0;i<len(name);i++{
		//fmt.Println(i,name[i])
	}

	for index,item := range name{
		fmt.Println(index,item,string(item))
	}

	dataList := []rune(name)
	fmt.Println(dataList[0],string(dataList[0]))


	//数组
	var numbers [3]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	fmt.Println(numbers)

	var names = [2]string{"中","国"}
	fmt.Println(names)

	var ages = [2]int{23,18}
	fmt.Println(ages)

	var names2 = []string{"aa","bb","cc"}
	fmt.Println(names2)

	var numberArr *[3]int
	fmt.Println(numberArr)

	numberArr2 := new([2]int)
	fmt.Println(numberArr2)



}
