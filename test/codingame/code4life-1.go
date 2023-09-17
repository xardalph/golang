package main

import (
	"fmt"
	"os"
	"sort"
)

/**
 * Bring data on patient samples from the diagnosis machine to the laboratory with enough molecules to produce medicine!
 **/

type BagMolecules struct {
	A int
	B int
	C int
	D int
	E int
}

type Personnage struct {
	location      string
	eta           int
	score         int
	carrySampleId int
	storage       BagMolecules
	expertise     BagMolecules
}

func NewPersonnage() Personnage {
	perso := Personnage{}
	fmt.Scan(&perso.location, &perso.eta, &perso.score, &perso.storage.A, &perso.storage.B, &perso.storage.C, &perso.storage.D, &perso.storage.E, &perso.expertise.A, &perso.expertise.B, &perso.expertise.C, &perso.expertise.D, &perso.expertise.E)
	return perso
}

type sampleList struct {
	s []map[string]int
}

func (samples *sampleList) getById(id int) map[string]int {
	for _, v := range samples.s {
		if v["id"] == id {
			return v
		}

	}
	return nil
}
func (samples *sampleList) getFirstAvailable() map[string]int {

	i := 0
	for samples.s[i]["carriedBy"] != -1 {
		i++
	}
	return samples.s[i]

}

func main() {
	removeUsellessData()
	tookSampleId := -1

	for {
		perso := NewPersonnage()
		_ = NewPersonnage()

		getAvailableElem()
		samples := getSamples()

		if tookSampleId == -1 && perso.location != "DIAGNOSIS" {
			fmt.Printf("GOTO DIAGNOSIS\n")
			continue
		}

		if perso.location == "DIAGNOSIS" && tookSampleId == -1 {
			tookSampleId = samples.getFirstAvailable()["id"]
			fmt.Printf("CONNECT %d\n", tookSampleId)

			fmt.Fprintf(os.Stderr, "connecting to : %v\n", tookSampleId)
			continue
		}

		// got a sample, goto molecule and take them

		if perso.location == "DIAGNOSIS" {
			fmt.Printf("GOTO MOLECULES\n")
			continue
		}
		toSend := samples.getById(tookSampleId)
		fmt.Fprintf(os.Stderr, "v is : %v\n storage : %v", toSend, perso.storage.A)
		if getMolecules(toSend, perso.storage) {
			continue
		}

		if perso.location != "LABORATORY" {
			fmt.Printf("GOTO LABORATORY\n")
			continue
		}
		if tookSampleId != -1 {

			fmt.Printf("CONNECT %d\n", tookSampleId)
			tookSampleId = -1
			continue
		}

		fmt.Printf("GOTO DIGNOSIS\n")

	}
}

func removeUsellessData() {
	var projectCount int
	fmt.Scan(&projectCount)

	for i := 0; i < projectCount; i++ {
		var a, b, c, d, e int
		fmt.Scan(&a, &b, &c, &d, &e)
		fmt.Fprintf(os.Stderr, "A %d B %d C %d D %d E %d \n", a, b, c, d, e)
	}
}

func getMolecules(samples map[string]int, persoInventory BagMolecules) bool {
	fmt.Fprintf(os.Stderr, "filling this order  :  %v \n", samples)

	if persoInventory.A < samples["costA"] {
		fmt.Printf("CONNECT A\n")
		fmt.Fprintf(os.Stderr, "Getting mol A \n")
		return true
	}
	if persoInventory.B < samples["costB"] {
		fmt.Printf("CONNECT B\n")
		return true
	}
	if persoInventory.C < samples["costC"] {
		fmt.Printf("CONNECT C\n")
		return true
	}
	if persoInventory.D < samples["costD"] {
		fmt.Printf("CONNECT D\n")
		return true
	}
	if persoInventory.E < samples["costE"] {
		fmt.Printf("CONNECT E\n")
		return true
	}
	return false
}

func getSamples() sampleList {

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
		sampleTemp["ratio"] = (health * 100) / ((costA + costB + costC + costD + costE + 2) * 100)
		samples[i] = sampleTemp

		//fmt.Fprintf(os.Stderr, "id fichier %d   récompense %d \n A %d B %d C %d D %d E %d", sampleId, health, costA, costB, costC, costD, costE)

	}
	sort.SliceStable(samples, func(i, j int) bool {
		return samples[i]["ratio"] > samples[j]["ratio"]
	})
	fmt.Fprintf(os.Stderr, "full list : %v", samples)
	return sampleList{s: samples}

}

func getAvailableElem() {
	var availableA, availableB, availableC, availableD, availableE int
	fmt.Scan(&availableA, &availableB, &availableC, &availableD, &availableE)
}

//[carriedBy:-1 costA:4 costB:0 costC:0 costD:0 costE:0 health:10 id:6 ratio:1]
