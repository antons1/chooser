package main

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/**
Structs to store a dilemma, with requirements and alternatives.
*/
type Alternative struct {
	Name string
	Weight float64
}

type Requirement struct {
	Name string
	Description string
	Importance float64
	Alternatives []Alternative
}

type Dilemma struct {
	Title string
	Description string
	Requirements []Requirement
}

// Reads a dilemma from the given file.
func (dil *Dilemma) createDilemmaFromFile(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print("Error not null!\n")
		fmt.Printf("%s\n", err)
		fmt.Println()
		dil.Title = "";
		return
	}
	
	err = json.Unmarshal(data, dil)	
	if err != nil {
		fmt.Print("Error not null!\n")
		fmt.Printf("%s\n", err)
		fmt.Println()
		dil.Title = "";
		return
	}
}