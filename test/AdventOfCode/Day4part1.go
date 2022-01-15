package main

import (
	"fmt"
	"strings"
)

func main() {

	boards := initSliceBoard()
	inputData := initsliceInput()
	//fmt.Println("elem :", boards)
	fmt.Println("len : ", len(boards))
	fmt.Println("input : ", inputData)
	// just need to mark some as already tested. then you tests if they are all marked.
	// sinon on vérifie si chaque element de la ligne est dans le slice inputData déjà récupéré, si oui ont arrète la ?

	for i := 0; i <= len(inputData); i++ {
		for j := 0; j <= len(boards); j++ {
			res := checkBoardAck(boards[j], inputData[:i])
			fmt.Println("res : ", res)
		}
	}
}

func checkBoardAck(board [][]string, input []string) []string {
	if Find(input, board[0][0]) {
		fmt.Println("ont a trouve l'input.")
	}
	return board[0]

}
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func initsliceInput() []string {
	input := "17,11,37,7,89,48,99,28,56,55,57,27,83,59,53,72,6,87,33,82,13,23,35,40,71,47,78,2,39,4,51,1,67,31,79,69,15,73,80,22,92,95,91,43,26,97,36,34,12,96,86,52,66,94,61,76,64,77,85,98,42,68,84,63,60,30,65,19,54,58,24,20,25,75,93,16,18,44,14,88,45,10,9,3,70,74,81,90,46,38,21,49,29,50,0,5,8,32,62,41"

	var inputData []string
	for _, elem := range strings.Split(input, ",") {
		inputData = append(inputData, elem)
	}
	return inputData
}

func initSliceBoard() [][][]string {
	array := "57 80 91 40 12\n62 36 72  0 20\n55 60 25 92 96\n14  2 17 18 86\n 1  4 90 66 38\n\n1 25 81 16 24\n33 40 86 28 96\n 4 97 90 32 13\n50 38 35 14 56\n73 42  9 36 67\n\n46 25 28 59 88\n33 48 15  0 95\n16 30 24  2 71\n11 81 58 70 65\n37  1 26 94 75\n\n53 56  5 72 99\n46 27 23 49  4\n 3 93 21 65 37\n35  7 50 32 24\n81 78 52 54 88\n\n76 14 58 84 17\n 5 89 97 57 15\n19 77 35 56 93\n31 71 92  2 39\n10 27 91 67 47\n\n27 95 74 40 50\n49 66 37 80 75\n54 70  1 16 41\n11 77 61  6 89\n71 45 73 23 29\n\n87 72 85 89 18\n91 56 54 14 88\n84 46 35  9 69\n36 66 53 34 17\n71 57  4 16  7\n\n35 85 76  2 59\n47 51 36 29 22\n25 62 86 78 95\n61 13 97 99 65\n92 52 39  9 73\n\n51 39 78 88 62\n79 91 58 50 15\n25 82 89 20  1\n12 73  7 95 32\n54 67 98 77  4\n\n49 38 14 68 13\n33 42 44 58 65\n15 47 29 84 82\n34 94 18 99 11\n89 43 23  0 78\n\n90 66 44 50 75\n71 57 87 40 98\n73 49 13 38 21\n35  6 11 91 84\n26 14 52 82 56\n\n81 38 54 73 42\n16 12 94 36 60\n45 98 10 70 46\n 4 58 96 53 64\n48 84 30  0  2\n\n37 16 48 41 97\n85 43 78 35 42\n39 18 30 44 82\n93 74 14 19 56\n87 10 80 77 17\n\n27 38 36 91 45\n71 33 72  9 89\n26 39 25 74 88\n 1 81 35 86 59\n10 63 14 61 53\n\n96 44 76 21 91\n82 31 90  4 48\n25 85 20 18 75\n64 71 98 47 49\n11 70  3  8 99\n\n71 73 69 20 39\n63 25 28 32  1\n 8 29 41 36 74\n24 59 95 16 76\n70  4 68 60 46\n\n35 73 66 20 44\n91 87 24 72 90\n64 81 84 83 30\n93 14 80 49 56\n53 94 10 92 34\n\n 6 79 37 43 78\n81 64 27 26 86\n11 80 21 52 93\n30 72 23 51 40\n12 97 71 38 41\n\n30 74 71 69 78\n85 56 23 58 24\n32 87 60 29 47\n67 19 57 35 14\n84 55 52 54 10\n\n81 23 12 54 20\n27 67 33 17 15\n16 11 40 73 32\n72 30 35  7 71\n94 68  3 61 84\n\n 9 12 73 84 90\n99 36 93 62 47\n28 54 59 17 74\n87 32 15 52 85\n31 11 82  1 43\n\n 7 92 57  6 71\n28 90 78 66 48\n25 35 72 58 65\n64 68 49 83  2\n36  0 52 88  3\n\n31 21 25 76 55\n32 53 92 38  6\n36 43 64 10 95\n96 79 90  8 54\n40 18 30 72 71\n\n61  1 52 70 60\n47 69 57 74 35\n 3 34 85 66 14\n92 33  7 31 20\n29 54 23 38  2\n\n39 49 87 43 11\n74 16 42 69 29\n17 60  0 38 66\n93 31 50 28 40\n75 76 12 79 88\n\n93 44 15 97  1\n47 14 75  0 28\n27 49 45 59 91\n 7 42 72 86 74\n57  3 16 36 79\n\n11 78  4  0 21\n71  9 18 88 80\n57 16  1  7 75\n25 50 70  3 22\n23 68 89 19 65\n\n99 17 60 70 66\n88 24 30  8 62\n49 59  4 29  7\n14 81 65 26 36\n46 15 80 95  6\n\n70 79 91 85 26\n89 63 86 73 82\n41 29  6 27 66\n59 39 37 17 51\n98 11  7  5 15\n\n47 13 36 27 83\n55 64  8 87 76\n 4  2 75 71 28\n79 12 85 94 54\n45 17 14 30 62\n\n21 20 46 80 25\n71 33 36 54 81\n76 34 63 40 35\n16 24  8 89 82\n 1 87 64 84 14\n\n43  3 63 50 25\n65 14  4 92 60\n26 45 49 51 96\n74 11 40 46  7\n90 67 77 56 68\n\n 4 24 92 95 44\n68 78 42  6 52\n17 23 89 56 38\n35 30 55 22 53\n73 64 61 34 19\n\n77 50 94 27 31\n16 39 12 14 95\n24 48 26 71 67\n85  3 43 66 82\n92 56  8 44  0\n\n37 46 31 67  9\n30 89 90  0 75\n99 45 73 93 52\n24 64 70 56 98\n53 48 10 51 14\n\n84 28 33 23 27\n53 13 87  1 24\n86 16 10 90 62\n88 12 99 43 22\n40 97 78 77  9\n\n97 43 61 53 90\n89 19  7 60 66\n 5 58 18 16 10\n14 13 47 39 20\n67 72 51 85 25\n\n54  2 70 69 91\n14 23 49 59  7\n64 88 33 65 56\n31 12 36 73 46\n82 87 66 71 95\n\n78 37 84 17 34\n64  6 68 74 46\n28 39 24 20 55\n12 61 97 60 98\n56 22  7  8 26\n\n97 46 67 51 92\n79 82 63 69 80\n83 84 26 10 65\n90 35 49 95 14\n89 36 17 39 45\n\n56 85 31 43 97\n48 84 83 12 28\n39 42 14 64 82\n98 17 72 29 91\n32  6 75 15 37\n\n35 75 70 16 55\n60 47 96 51 21\n65 52 54  6  3\n24 12 72 64 78\n 8 17 34 83 71\n\n62 34 27 90 22\n53 79 40 19 78\n52 77 23 60 54\n68 47 59 85 31\n 1 26 45  6 29\n\n64 74 89  6 97\n99 94 36 72 32\n59 13 42 69 16\n66 37 41 55 83\n35 75 54 95 57\n\n61 52  9 18 70\n41 53 90 94 87\n80 20  0 11 42\n 8 46 63 29  5\n34 15  7 91 10\n\n 4 90 30 28 64\n38 35 42 79 21\n37 65 88 83 56\n78 11 12 22 50\n74 49 81 82 71\n\n99 89 13 71 39\n96 78 67 82 59\n65 75 55 69 17\n16 15 86 27 85\n58 46 11 32  9\n\n76 42 24 97 47\n72 55 28 32 18\n 8 65 50 84 71\n52 56 36 87 59\n20 33 89 46 41\n\n59 84 46 15 11\n52 60 31 97 16\n67  3 53 10 61\n36 80  6 43  2\n95 91 14 76 77\n\n89 27 59  4 99\n71 56 74 19  5\n11 84 61 90 16\n29 93 97 24 52\n64 88 18 34 50\n\n63 18 38 79 34\n17 67 96 87 70\n60  2 54 29 81\n24 72 12 39  8\n59 36 44 97 69\n\n47 87 74 21 73\n64 67 89 92  6\n95 26 77 65 45\n22 69 31 13 61\n29 43 68 81 28\n\n12 30 18 91 27\n37  4 75 67 72\n99 33 19 90 25\n86 24 22 52 80\n38 62 94 11 44\n\n53 47 36 57  9\n14 92 27 29 67\n21 87 76  5 60\n88 10 54 95 48\n32 18 64  8 37\n\n83 53 88 51 46\n36 85 31  0 95\n58 52 57  2 75\n78 74  3 27 20\n86 99  7 29 96\n\n48 32 52 95 21\n80 20 75 68 29\n38 56 94 16 50\n99 91 26  6 82\n 8 93 25 43 61\n\n92  4 79 73 82\n68 20 44 15 70\n37 69 96 58 53\n25 76 86 54 47\n38  3  1 33 46\n\n62 84 30 15 35\n65 90 46  0 38\n87 93 70 26 44\n41  1 11 66 14\n45 22 55 96 57\n\n31 52 45 82 53\n61 15 40 83 68\n87 29  8 36 85\n26 47 12 37 32\n67  4 58 17 65\n\n31 14  4 22 93\n47 50 41 29  6\n69 63  3 38  7\n82 23 75 48 90\n73 66 83 87  5\n\n23 71 21 24 29\n52 94 84 14 58\n70 49 27 96 64\n48 66 38 65 89\n67 36 19 26  9\n\n 2 22 40 34 12\n13 78 75 68 36\n19 50 79  9  5\n95 10 54 71 37\n20 63  7 97 31\n\n 8 31 88 80 63\n33 25 77 37 89\n22 49 66 20 43\n50 21 30 57 99\n 9 47 48 26 44\n\n22 28 39 26 57\n49 43 37 76 95\n83 30 41 50 27\n46 51 55 96 47\n87 62 24 19 60\n\n48 68 34 73 84\n21 28 90 38  4\n13 19 27 45 36\n32 44 69 37 75\n 8 60 92 85 58\n\n53 77 25 45  8\n34 41 80 79 81\n32 65 98 86 91\n74 38 50 58 44\n56 75 52 33 97\n\n85 12 64  3 58\n90 36 27 71 78\n44 84 54 87 31\n16 50 14 17 93\n72 55 97  4 38\n\n99 41 38 68 51\n73 25  5 59  9\n48 83 36 58 55\n31 53 35 64 98\n66 80 30 13 92\n\n20  7 90  3 59\n83 49 42 26 22\n53 56 23 13 19\n63  5 74 97 79\n 1 92 32 73  4\n\n52 79 29 27 14\n 8 39  0 71 82\n18 85 88 81 10\n15  2 63  9 23\n26 84 12 46 40\n\n56 63 16  1 40\n44 18  5 50 80\n19 88 25 79 36\n17 89 78 76 46\n99 29 35  2 10\n\n62 96 27 68 31\n34 30  7  8 21\n70 22 80 40 42\n61 18 98 23  5\n38 69 24 78 65\n\n75 12  1 56 73\n25 90  9 10 53\n18 30 59 76 77\n39 40 28 68 46\n82 83 42 15 95\n\n47 83 65 24 35\n94 30 12 17 63\n66 75 56 40 41\n67 53 26 27 64\n55 22 31 74 11\n\n14 21  4 79 38\n52 70 50 44 29\n93 67 65 54 73\n83 34 97 30 60\n27 76 33  6 39\n\n16 62  3 75 84\n54 10  1  2 44\n21 42  0 99 64\n91 45 67 34 82\n83 78 17 53  6\n\n83  0 44 31 72\n73 19 18 85 45\n59 51 11 75 64\n90 56 97 46 53\n39 32 17 60 10\n\n50 57 93 43  9\n63 20 15  5 17\n35 48  2 52 60\n34 33 90 85 16\n13 53 47 59 62\n\n25 75  7 58 98\n28 66 20 21 84\n16 19 85 89  0\n 4 63 70 48 61\n13  8 23 15 37\n\n19 15  8 88 97\n81 70  4 59 72\n34 24 39 47 58\n11 10 50 18 99\n40 16 78 45 74\n\n82 98 67 94 22\n43 81 96 38 77\n15  6  0 40 62\n52 11 83 54 69\n12  7 70  5 31\n\n30 57 73 38 85\n35 84 25 19 65\n17  6 33 54 13\n10 24 31 50 88\n 2  8  1 59 43\n\n34 86 21 13 33\n37 20 39 57 66\n10 72 98 40  3\n83 23 87 11 70\n71 24 38  2 58\n\n47 99 60 43  0\n66 84 40 29 79\n52 58 25 91 55\n37 69 22 82  4\n65 62  2 31 28\n\n42 10  1 23  3\n27 53 46 89 31\n55  9 36 49 85\n87 73 92  5 61\n98 72 86 65 94\n\n 6 58 78 75  3\n72 84 79 17 71\n66  0 90 81 28\n94  7 63 76 40\n85 82 49 74 65\n\n46 27 38 87 55\n93 21 40 61 82\n43 75 54 32 77\n 6 23 84 10  4\n92 89 51  0 86\n\n71  6 28 94 26\n12 84 39 65 22\n73 23 78  2 59\n40 66 83 91 92\n67 25 64 42 41\n\n43 92 81 83 99\n59 18 23 78 94\n89 40 47 56  4\n24 73  9 29 60\n 5 26 51 68  6\n\n59  7 83 32 72\n46 82 58 78 44\n70 45 64 84 38\n34  8 79  1 81\n41 98 48 33 18\n\n68 30 32 16 88\n33 76 41 99 60\n11 38 14 49 65\n89 58 24 26 34\n70 92 29 64 94\n\n78 87 15 29  5\n75 54 63 20 33\n84 88 25 80 11\n93 73 92 39 22\n 2 10  4 82 77\n\n93 33 71 39  3\n27 42 76 12 13\n20 64 52 37 72\n94 21 23 99 91\n95 97 10 63 92\n\n30 65 97 73 85\n90 60 34  1 37\n79 47 40 31  9\n62 51 15 19 76\n96 77 64 72 71\n\n69 88 51 14 64\n77 12 96 46 58\n89  2 49 80 37\n32 62 44 23 84\n78 97 87 41 85\n\n17 40 61 68 81\n52 24 77 30 80\n75 88 27 78 87\n94 16 36 67 19\n95 43 97 46 93\n\n32 60  9 75 98\n74 82 94 20 81\n47  7 10 76 48\n88 11 86 69 67\n72 36 83 62 34\n\n26  8 64 47 76\n27 15 82 65 80\n54 79 84 87 37\n97 18  3 93 20\n88  7 40 99 35\n\n40 94  6 53 58\n45 60 81 64 30\n66 55 74 37  7\n63 95 23 72 17\n80 18 69 89 28\n\n48 83 84 85  3\n98 80 50 96 91\n18 69 44 62 15\n20 88 12 45 28\n 8 29  0 37 27\n"

	boards := make([][][]string, 0)
	boardNumber := 0
	lineNumber := 0

	for _, board := range strings.Split(array, "\n\n") {
		board = strings.Trim(board, " ")
		if board == "" {
			continue
		}
		boards = append(boards, make([][]string, 5))

		boards[boardNumber] = make([][]string, 5)

		for _, line := range strings.Split(board, "\n") {
			line = strings.Trim(line, " ")
			if line == "" || line == "  " {
				continue
			}
			if len(boards[boardNumber][lineNumber]) == 0 {
				boards[boardNumber][lineNumber] = make([]string, 0)
			}
			for _, elem := range strings.Split(line, " ") {
				elem = strings.Trim(elem, " ")
				elem = strings.Trim(elem, "  \n")
				if elem == "" || elem == " " || elem == "  " {
					//	fmt.Println("continuing")
					continue
				}
				boards[boardNumber][lineNumber] = append(boards[boardNumber][lineNumber], elem)
				//fmt.Println("board : ", boards)
			}
			lineNumber++
		}
		lineNumber = 0
		boardNumber++
	}
	return boards
}
