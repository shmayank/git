package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
