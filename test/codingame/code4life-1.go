package main

import (
	"fmt"
	"os"
)

/**
 * Bring data on patient samples from the diagnosis machine to the laboratory with enough molecules to produce medicine!
 **/

func main() {
	var projectCount int
	fmt.Scan(&projectCount)

	for i := 0; i < projectCount; i++ {
		var a, b, c, d, e int
		fmt.Scan(&a, &b, &c, &d, &e)
		fmt.Fprintf(os.Stderr, "A %d B %d C %d D %d E %d \n", a, b, c, d, e)
	}
	var storage = make([][]int, 2)
	storagePlayer := make([]int, 5)
	storageEnemy := make([]int, 5)
	storage[0] = storagePlayer
	storage[1] = storageEnemy

	for {
		for i := 0; i < 2; i++ {

			var location string

			var eta, score, expertiseA, expertiseB, expertiseC, expertiseD, expertiseE int
			fmt.Scan(&location, &eta, &score, &storage[i][0], &storage[i][1], &storage[i][2], &storage[i][3], &storage[i][4], &expertiseA, &expertiseB, &expertiseC, &expertiseD, &expertiseE)

			fmt.Fprintf(os.Stderr, "location %s score %d storage : \n %v\n",
				location, score, storage[i])

		}
		getAvailableElem()

		samples := getSamples()

		fmt.Printf("GOTO DIAGNOSIS\n")
	}
}

func getSamples() []map[string]int {

	var sampleCount int
	fmt.Scan(&sampleCount)
	fmt.Fprintf(os.Stderr, "nb fichier : %d\n", sampleCount)

	samples := make([]map[string]int, sampleCount)

	for i := 0; i < sampleCount; i++ {
		sampleTemp := make(map[string]int, 10)

		var sampleId, carriedBy, rank int
		var expertiseGain string
		var health, costA, costB, costC, costD, costE int
		fmt.Scan(&sampleId, &carriedBy, &rank, &expertiseGain, &health, &costA, &costB, &costC, &costD, &costE)
		sampleTemp["id"] = sampleId
		sampleTemp["carriedBy"] = carriedBy
		sampleTemp["health"] = health
		sampleTemp["costA"] = costA
		sampleTemp["costB"] = costB
		sampleTemp["costC"] = costC
		sampleTemp["costD"] = costD
		sampleTemp["costE"] = costE
		samples[i] = sampleTemp

		fmt.Fprintf(os.Stderr, "id fichier %d   récompense %d \n A %d B %d C %d D %d E %d", sampleId, health, costA, costB, costC, costD, costE)

	}
	return samples

}

func getAvailableElem() {
	var availableA, availableB, availableC, availableD, availableE int
	fmt.Scan(&availableA, &availableB, &availableC, &availableD, &availableE)
}
