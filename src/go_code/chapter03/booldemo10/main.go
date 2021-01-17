package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func pasestudent() {
	m := make(map[string]*student)

	fmt.Println(m)

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	fmt.Println(stus)
	fmt.Println(m)
}

func main() {

	pasestudent()
}
