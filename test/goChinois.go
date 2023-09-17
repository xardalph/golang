package main

import "fmt"

type Goban [9][9]int

func (m *Goban) getvalue() int {
	return 5
}
func main() {
	isWhiteTurn := 1

	fmt.Println("ca compiiile")
	goban := Goban{}
	fmt.Println("valeur : ", goban[2][1])
	for i := 1; i < 5; i++ {
		newPawn := getNewValue(goban)
		fmt.Println(newPawn)
		goban[newPawn[0]][newPawn[1]] = 1
		isWhiteTurn = (isWhiteTurn + 1) % 2

	}
}

func getNewValue(goban Goban) [2]int {
	printGoban(goban)
	var ret [2]int
	fmt.Println(len(goban), ret[0])
	for ok := true; ok; ok = (len(goban) < ret[0] || len(goban) < ret[1]) && goban[ret[0]][ret[1]] != 0 {
		fmt.Println("what column do you want to play ?")
		_, _ = fmt.Scan(&ret[0])
		fmt.Println("what line do you want to play ?")
		_, _ = fmt.Scan(&ret[1])
		fmt.Println("condition : ", goban[ret[0]][ret[1]] != 0)

	}

	//check value is not already taken, and that it's not illegal (pas de recurse nécessaire ici je pense, juste -+1 check)
	return ret
}

func printGoban(goban Goban) {
	for line, _ := range goban {
		for column, _ := range goban[line] {
			fmt.Printf("%d ", goban[line][column])
		}
		fmt.Println()

	}
}
