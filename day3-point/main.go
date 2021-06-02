package main

import "fmt"

func main(){

	//指针
	var v1 *int
	v2 := new(int)
	var v3 *int
	fmt.Println(&v1,v1)
	fmt.Println(&v2,v2)
	fmt.Println(&v3,v3)



}
