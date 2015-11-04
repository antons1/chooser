package main

import(
	
)

/**
Structs to store a dilemma, with requirements and alternatives.
*/
type Alternative struct {
	var name string
	var weight int
}

type Requirement struct {
	var name string
	var description string
	var importance int
	var alternatives Alternative[]
}

type Dilemma struct {
	var title string
	var description string
	var requirements Requirement[]
}