package main

import "fmt"

func main() {
	maps := make([]map[string]int, 3)

	map1 := make(map[string]int)
	map1["1"] = 45521

	map2 := make(map[string]int)
	map2["2"] = 45

	map3 := make(map[string]int)
	map3["3"] = 9
	map3["4"] = 5

	maps[0] = map1
	maps[1] = map2
	maps[2] = map3
	maps[0]["1"] = 98956
	fmt.Println(maps[0]["1"])
	for _, m := range maps {
		fmt.Println(m)
	}
}
