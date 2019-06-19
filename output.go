package main

import "fmt"

type Output struct{}

func (c Output) print(data interface{}) {
	fmt.Println(data)
}
