package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	// n: the number of temperatures to analyse
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)
	var min int64 = 5526
	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	negative := false
	for i := 0; i < n; i++ {
		negative = false
		// t: a temperature expressed as an integer ranging from -273 to 5526
		temperature, _ := strconv.ParseInt(inputs[i], 10, 64)
		if temperature < 0 {
			negative = true
			temperature = temperature * -1
		}
		if min > temperature {
			min = temperature
		}

	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	//fmt.Println(min) // Write answer to stdout

	if min == 5526 {
		fmt.Println("0")
	} else if negative {
		fmt.Println(min * -1)
	} else {
		fmt.Println(min)
	}
}
