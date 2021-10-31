package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name = %v, Age = %v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func main() {
	a
}
