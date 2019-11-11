package mypackage

import "fmt"

type person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *person {

	p := person{name: name}
	p.age = age
	return &p
}

func SecondExportFunc(val string) string {
	return val
}

func ThirdExportFunc() void {
	fmt.Println("THIRD")
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(NewPerson("Jon", 1))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
