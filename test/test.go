package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func printBoard(list []string) {
	for i, elem := range list {
		if i%3 == 0 {
			fmt.Println()
		}
		fmt.Print(elem, " ")
	}
	fmt.Println()
}

type type_name struct{}

func (m []type_name) function_name() int {
	//code
}

func checkForOneLine(line string) {
	var stack [][]rune

	stack = append(stack, make([]rune, 0))

	for _, char := range line {
		if stack[0][len(stack[0])-1] > char {
			stack[len(stack)-1] = append(stack[len(stack)-1], char)
		} else {
			stack = append(stack, make([]rune, 0))
			stack[len(stack)-1] = append(stack[len(stack)-1], char)
		}

		fmt.Fprintf(os.Stderr, "heu c'était quoi ? %+v ", string(stack[0][len(stack[0])-1]))
	}
	fmt.Fprintf(os.Stderr, "   %+v ", string(stack[0]))
	fmt.Fprintf(os.Stderr, "\nEND OF LINE\n")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var board = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	condition := true
	printBoard(board)
	for condition {
		fmt.Print("saisissez un nombre entre 1 et 9  ")
		scanner.Scan()
		nbr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("c pas un nombre. ")
			continue
		} else if nbr > 9 || nbr < 1 {
			fmt.Println("c pas entre 0 et 9")
			continue
		}
		board[nbr-1] = "X"
		printBoard(board)
		fmt.Println()
		separator := ""
		string := strings.Join(board, separator)
		regex, _ := regexp.MatchString("[XO]{9}", string)
		if regex {
			fmt.Println("FIN du game")
			condition = false
		}
	}
}
