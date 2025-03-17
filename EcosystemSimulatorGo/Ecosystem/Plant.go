package Ecosystem

import "fmt"

type Plant struct {
	Coll int
	Name string
}

func (plant Plant) getColl() int {
	return plant.Coll
}
func (plant Plant) getName() string {
	return plant.Name
}
func (plant Plant) setColl(coll int) {
	plant.Coll = coll
}
func (plant Plant) setName(name string) {
	plant.Name = name
}
func newPlant() Plant {
	var name string
	var coll int
	fmt.Println("Enter plant name: ")
	fmt.Scan(&name)
	fmt.Println("Enter plant coll: ")
	fmt.Scan(&coll)
	return Plant{Name: name, Coll: coll}
}
