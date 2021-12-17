package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello lihaoyu")
	//os.Exit(0)
	if len(os.Args) > 1 {
		fmt.Println("hello", os.Args[1])
	}
	os.Exit(0)
}
