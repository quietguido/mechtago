package generator

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func GenerateJson(filename string, itemCount int) int {
	var groundTruthValue int

	output, err := os.Create(filename + ".json")
	if err != nil {
		log.Fatalf("file creating failed")
	}

	defer output.Close()

	_, err = output.Write([]byte("["))
	if err != nil {
		log.Fatalf("failed to write opening braket")
	}

	for i := 0; i < itemCount; i++ {
		aVal := rand.Intn(21) - 10
		bVal := rand.Intn(21) - 10

		groundTruthValue += aVal
		groundTruthValue += bVal

		jsonData, err := json.Marshal(
			Data{
				A: aVal,
				B: bVal,
			},
		)
		if err != nil {
			log.Fatalf("encoding problem")
		}

		_, err = output.Write(jsonData)
		if err != nil {
			log.Fatalf("failed to write")
		}

		if i != itemCount-1 {
			_, err = output.Write([]byte(",\n"))
			if err != nil {
				log.Fatalf("failed to write new line")
			}
		}
	}

	_, err = output.Write([]byte("]"))
	if err != nil {
		log.Fatalf("failed to write closing braket")
	}

	log.Println("Finished generating successfully")

	return groundTruthValue
}
