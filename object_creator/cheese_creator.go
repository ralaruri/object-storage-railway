package object_creator

import (
	"encoding/json"
	"log"
	"os"
)

// write the cheeseMap to a json file
func WriteCheeseMapToJSONFile() {
	cheeseMap := map[int]string{100: "chedder", 200: "swiss", 300: "gouda", 400: "mozzarella"}
	// create a file
	file, err := os.Create("cheese.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create a json encoder
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	// write the cheeseMap to the file
	err = encoder.Encode(cheeseMap)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteJsonFile() {
	err := os.Remove("cheese.json")
	if err != nil {
		log.Fatal(err)
	}
}
