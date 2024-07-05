package calculator

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/quietguido/mechtago/generator"
)

type SpecificCounter struct {
	counter int
	_       [128]byte
}

func Calculate(inputfilePath string, workerCount int) int {
	inputFile, err := os.Open(fmt.Sprintf("%v.json", inputfilePath))
	if err != nil {
		log.Fatalf("failed to open file inputfile")
	}

	defer inputFile.Close()

	var wg sync.WaitGroup
	specificCounter := make([]SpecificCounter, workerCount)

	jobs := make(chan int, 1000)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(jobs, &specificCounter[i].counter, &wg)
	}

	reader := bufio.NewReaderSize(inputFile, 8*1024)
	decoder := json.NewDecoder(reader)

	_, err = decoder.Token()
	if err != nil {
		log.Fatalf("Failed to read opening token")
	}

	var data generator.Data

	for decoder.More() {
		err := decoder.Decode(&data)
		if err != nil {
			log.Fatalf("failed sequnetion decode")
		}

		jobs <- data.A
		jobs <- data.B
	}

	close(jobs)
	wg.Wait()

	var answ int
	for i := 0; i < workerCount; i++ {
		answ += specificCounter[i].counter
	}

	return answ
}

func worker(jobs <-chan int, specificCounter *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range jobs {
		*specificCounter += num
	}
}
