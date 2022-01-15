package main

import (
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// nbFloors: number of floors
	// width: width of the area
	// nbRounds: maximum number of rounds
	// exitFloor: floor on which the exit is found
	// exitPos: position of the exit on its floor
	// nbTotalClones: number of generated clones
	// nbAdditionalElevators: ignore (always zero)
	// nbElevators: number of elevators
	var nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators int
	fmt.Scan(&nbFloors, &width, &nbRounds, &exitFloor, &exitPos, &nbTotalClones, &nbAdditionalElevators, &nbElevators)

	var elevatorPositions = make([]int, nbElevators+1)
	for i := 0; i < nbElevators; i++ {
		// elevatorFloor: floor on which this elevator is found
		// elevatorPos: position of the elevator on its floor
		var elevatorFloor, elevatorPos int
		fmt.Scan(&elevatorFloor, &elevatorPos)
		elevatorPositions[elevatorFloor] = elevatorPos
	}
	elevatorPositions[exitFloor] = exitPos

	for {
		// cloneFloor: floor of the leading clone
		// clonePos: position of the leading clone on its floor
		// direction: direction of the leading clone: LEFT or RIGHT
		var cloneFloor, clonePos int
		var direction string
		fmt.Scan(&cloneFloor, &clonePos, &direction)
		if cloneFloor == -1 {
			fmt.Println("WAIT")
			continue
		}
		fmt.Fprintln(os.Stderr, "cloneFloor : ", cloneFloor, "position : ", clonePos)
		fmt.Fprintln(os.Stderr, " elevator pos = ", elevatorPositions[cloneFloor])
		fmt.Fprintln(os.Stderr, " direction : ", direction)

		if elevatorPositions[cloneFloor] > clonePos && direction == "LEFT" {
			fmt.Fprintln(os.Stderr, "doit aller à Droite")
			fmt.Println("BLOCK")
			continue
		} else if elevatorPositions[cloneFloor] < clonePos && direction == "RIGHT" {
			fmt.Fprintln(os.Stderr, "doit aller à Gauche")
			fmt.Println("BLOCK")
			continue
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("WAIT") // action: WAIT or BLOCK
	}
}
