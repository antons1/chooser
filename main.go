package main

import(
	"fmt"
)

func main() {
	dil := Dilemma{};
	dil.createDilemmaFromFile("./pc.json");
	
	if dil.Title == "" {
		fmt.Printf("There was an error reading the file. See above for more info\n");
		return;
	}
	
	sc := Score{};
	sc.initScore(dil);
	sc.FindScore();
}