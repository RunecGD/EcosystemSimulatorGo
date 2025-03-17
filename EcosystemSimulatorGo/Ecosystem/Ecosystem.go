package Ecosystem

import "fmt"

var plants = []Plant{}
var animals = []Animal{}

func addPlant(p Plant) {
	plants = append(plants, p)
}
func addAnimal(a Animal) {
	animals = append(animals, a)
}
func getAllPlants() []Plant {
	return plants
}
func getAllAnimals() []Animal {
	return animals
}
func updatePlants(id int, p Plant) {
	if id >= 0 && id < len(plants) {
		fmt.Println(plants[id])
		plants[id] = p
	}
}
func updateAnimals(id int, p Animal) {
	if id > 0 && id < len(animals) {
		animals[id] = p
	}
}
func removePlant(id int) {
	if id >= 0 && id < len(plants) {
		plants = append(plants[:id], plants[id+1:]...)
	}
}
func removeAnimal(id int) {
	if id > 0 && id < len(animals) {
		animals = append(animals[:id], animals[id+1:]...)
	}
}
