package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1
	//var (
	//	a int = 1
	//	b int = 1
	//)
	a := 1
	b := 1
	fmt.Println(a)
	for i := 0; i < 5; i++ {
		fmt.Println(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 12
	//fmt.Println("a: ",a)
	//fmt.Println("b: ",b)
	//temp:=a
	//a=b
	//b=temp
	//fmt.Println("a: ",a)
	//fmt.Println("b: ",b)
	a, b = b, a
	t.Log(a, b)
}

// 常量初始化
const (
	Readable = 1 << iota
	lihaoyu
)
