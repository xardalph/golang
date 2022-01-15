package main

import (
	"bufio"
	"fmt"
	"strings"

	"os"
	"regexp"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// N: Number of elements which make up the association table.
	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	// Q: Number Q of file names to be analyzed.
	var Q int

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &Q)
	var extensions map[string]string
	extensions = make(map[string]string)
	for i := 0; i < N; i++ {
		// EXT: file extension
		// MT: MIME type.
		var EXT, MT string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &EXT, &MT)
		extensions[strings.ToLower(EXT)] = MT

	}
	for i := 0; i < Q; i++ {
		scanner.Scan()
		FNAME := scanner.Text()
		_ = FNAME // to avoid unused error // One file name per line.

		r := regexp.MustCompile(`(?i)(?:\.)([^.]+)$`)
		result := r.FindStringSubmatch(FNAME)

		if len(result) != 2 {
			fmt.Println("UNKNOWN")
			continue
		}
		ext := result[1]
		if val, ok := extensions[strings.ToLower(ext)]; ok {
			fmt.Println(val)
			continue
		}
		fmt.Println("UNKNOWN")
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")

	// For each of the Q filenames, display on a line the corresponding MIME type. If there is no corresponding type, then display UNKNOWN.

}
