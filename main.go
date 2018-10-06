package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/igorpo/confusius/models"
	"github.com/igorpo/confusius/simulation"
)

func main() {
	choices, err := readEvaluations("./evaluations/evaluation.json")
	if err != nil {
		fmt.Printf("error reading evaluations file: %v\n", err)
		os.Exit(1)
	}

	simulator := simulation.New(100, choices)
	simulator.Simulate()
	fmt.Println(simulator.Top(5))
}

func readEvaluations(filename string) (models.Choices, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	evalBytes, err := ioutil.ReadAll(jsonFile)

	var evalChoices models.Choices
	json.Unmarshal(evalBytes, &evalChoices)
	return evalChoices, nil
}
