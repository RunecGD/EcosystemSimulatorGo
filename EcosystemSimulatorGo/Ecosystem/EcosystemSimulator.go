package Ecosystem

import "fmt"

func Menu() {
	createFolderSaves()
	fmt.Println("1. Add Plant")
	fmt.Println("2. View Plants")
	fmt.Println("3. Update Plant")
	fmt.Println("4. Remove Plant")
	fmt.Println("5. Add Animal")
	fmt.Println("6. View Animals")
	fmt.Println("7. Update Animal")
	fmt.Println("8. Remove Animal")
	fmt.Println("9. Save Ecosystem")
	fmt.Println("10. Load Ecosystem")
	fmt.Println("0. Exit")
	for {
		var choice int
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			addPlant(newPlant())
		case 2:
			fmt.Println("Plants:", getAllPlants())
		case 3:
			var id int
			fmt.Print("Id Plants:")
			fmt.Scanln(&id)
			updatePlants(id-1, newPlant())
		case 4:
			var id int
			fmt.Print("Id Plants: ")
			fmt.Scanln(&id)
			removePlant(id - 1)
		case 5:
			addAnimal(newAnimal())
		case 6:
			fmt.Println("Animal:", getAllAnimals())
		case 7:
			var id int
			fmt.Print("Id Plants:")
			fmt.Scanln(&id)
			updateAnimals(id-1, newAnimal())
		case 8:
			var id int
			fmt.Print("Id Animal: ")
			fmt.Scanln(&id)
			removeAnimal(id - 1)
		case 9:
			var save string
			fmt.Print("Saves: ")
			fmt.Scanln(&save)
			err := saveEcosystem(save)
			if err != nil {
				fmt.Println("Error saving ecosystem:", err)
			} else {
				fmt.Println("Ecosystem saved successfully!")
			}
		case 10:

		case 0:
			fmt.Println("Exiting...")
			return
		}
	}
}
