package main

import(
	"fmt"
)

const(
	BASE float64 = 10	// Base points

	RM1 float64 = 10	// (Required)		Multiplier for importance 1
	RM2 float64 = 2		// (Presedence)		...2
	RM3 float64 = 1		// (Important)		...3
	RM4 float64 = 0.5	// (Wanted)			...4
	RM5 float64 = 0.1	// (Unimportant)	...5
	
	AM1 float64 = 1.1	// (Wanted)			Multiplier for weight 1
	AM2 float64 = 1		// (Equal)			...2
	AM3 float64 = 0.9	// (Unwanted)		...3
	AM4 float64 = 0.1	// (Deal breaker)	...4
)

type Score struct {
	score float64
	dilemma Dilemma
}

func (sc *Score) initScore(dil Dilemma) {
	sc.dilemma = dil;
}

func getAnswer() {
	
}

func askQuestions() {
	
}

func (sc *Score) FindScore() {
	dil := sc.dilemma
	req := dil.Requirements
	
	fmt.Println();
	fmt.Printf("Help me choose!\n");
	fmt.Printf("----------------------\n");
	fmt.Printf("Case: %s\n", dil.Title);
	fmt.Printf("%s\n", dil.Description);
	fmt.Printf("----------------------\n");
	fmt.Println();
	
	for i := 0; i < len(req); i++ {
		fmt.Printf("%s\n---\n%s\n\n", req[i].Name, req[i].Description);
	}
}

func GetScore() (int) {
	return 0
}