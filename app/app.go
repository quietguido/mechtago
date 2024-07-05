package app

import (
	"fmt"
	"log"

	"github.com/quietguido/mechtago/calculator"
	"github.com/quietguido/mechtago/generator"
)

func Run() {
	generateAndEvaluate()
}

func generateAndEvaluate() {
	var inputFileName string
	var itemCount int
	var workerCount int

	fmt.Println("Enter inputfile name to generate: ")
	fmt.Scanln(&inputFileName)

	fmt.Println("Enter item count:")
	fmt.Scanln(&itemCount)

	fmt.Println("Enter number of concurrent workers: ")
	fmt.Scanln(&workerCount)

	groundTruth := generator.GenerateJson(inputFileName, itemCount)

	answ := calculator.Calculate(inputFileName, workerCount)

	log.Println("groundTruth: ", groundTruth)
	log.Println("calculated value: ", answ)
	if groundTruth != answ {
		log.Fatalln("answer is not same as groundtruth")
	}
}
