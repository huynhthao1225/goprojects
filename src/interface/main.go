package main

import (
	"fmt"
	"os"
	"text/template"
)

type botString interface {
	getGreeting() string
}

type botInt interface {
	getValue() int
}

type bot interface {
	botString
	botInt
}

type enHello struct{}
type vnHello struct{}
type intHello struct{}

func (vnHello) getGreeting() string {
	return "Chào bạn"
}
func (vnHello) getValue() int {
	return 80
}

func (enHello) getGreeting() string {
	return "Hello there"
}
func (enHello) getValue() int {
	return 90
}

func (intHello) getGreeting() string {
	return "Int Hello"
}
func (intHello) getValue() int {
	return 100
}

type Inventory struct {
	Material string
	Count    uint
}

func main() {

	en := enHello{}
	vn := vnHello{}
	val := intHello{}
	printGreeting(vn)
	printValue(vn)
	printGreeting(en)
	printValue(en)
	printGreeting(val)
	printValue(val)

	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printValue(b bot) {
	fmt.Println(b.getValue())
}
