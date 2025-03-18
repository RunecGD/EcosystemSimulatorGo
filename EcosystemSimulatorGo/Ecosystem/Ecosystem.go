package Ecosystem

import (
	"fmt"
	"os"
	"path/filepath"
)

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
		plants[id].setName(p.getName())
		plants[id].setColl(p.getColl())
	}
}
func updateAnimals(id int, a Animal) {
	if id > 0 && id < len(animals) {
		animals[id].setName(a.getName())
		animals[id].setColl(a.getColl())
		animals[id].setType(a.getType())
	}
}
func removePlant(id int) {
	if id >= 0 && id < len(plants) {
		plants = append(plants[:id],
			plants[id+1:]...)
	}
}
func removeAnimal(id int) {
	if id > 0 && id < len(animals) {
		animals = append(animals[:id],
			animals[id+1:]...)
	}
}
func saveEcosystem(filename string) error {
	var data string

	data += "Plants:\n"
	for _, plant := range plants {
		data += fmt.Sprintf("Name: %s, Coll: %d\n", plant.Name, plant.Coll)
	}

	data += "\nAnimals:\n"
	for _, animal := range animals {
		data += fmt.Sprintf("Name: %s, Coll: %d, Type: %s\n", animal.Name, animal.Coll, animal.Type)
	}

	path := filepath.Join("Saves", filename)
	return os.WriteFile(path,
		[]byte(data),
		0644)
}
