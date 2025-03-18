package Ecosystem

import "fmt"

type Animal struct {
	Name string
	Coll int
	Type string
}

func (animal Animal) getName() string {
	return animal.Name
}
func (animal Animal) getColl() int {
	return animal.Coll
}
func (animal Animal) getType() string {
	return animal.Type
}
func (animal Animal) setName(name string) {
	animal.Name = name
}
func (animal Animal) setColl(coll int) {
	animal.Coll = coll
}
func (animal Animal) setType(t string) {
	animal.Type = t
}
func newAnimal() Animal {
	var name, t string
	var coll int
	fmt.Println("Enter animal name: ")
	fmt.Scan(&name)
	fmt.Println("Enter animal coll: ")
	fmt.Scan(&coll)
	fmt.Println("Enter type: ")
	fmt.Scan(&t)
	return Animal{Name: name, Coll: coll, Type: t}
}
