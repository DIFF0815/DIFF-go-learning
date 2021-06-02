package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){

	//数据类型转换
	v1 := 19
	result := strconv.Itoa(v1)
	fmt.Println(result,reflect.TypeOf(result))
	var v2 int8 = 17
	data := strconv.Itoa(int(v2))
	fmt.Println(data,reflect.TypeOf(data))

	v3 := "666"
	ret,err := strconv.Atoi(v3)
	if err == nil{
		fmt.Println("转换成功",reflect.TypeOf(ret))
	}else {
		fmt.Println("转换失败")
	}

	//进制转换
	v4 := 99
	r4 := strconv.FormatInt(int64(v4),16)
	fmt.Println(r4,reflect.TypeOf(r4))

	v5 := "1001000101"
	r5,err5 := strconv.ParseInt(v5,2,0)
	fmt.Println(r5,err5,reflect.TypeOf(r5))


}
