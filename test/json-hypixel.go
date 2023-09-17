package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting the script....")

	argLimit := 2000
	cmdArg := os.Args[1]
	if cmdArg != "" && cmdArg != "0" {
		argLimit, _ = strconv.Atoi(cmdArg)
	}

	fmt.Printf("limit of price difference to show item : %d\n", argLimit)

	jsonFile, err := os.Open("hypixel-list-craft.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	listCraft, err := gabs.ParseJSON(byteValue)
	if err != nil {
		fmt.Println(err)
	}

	// -------------------

	priceList := getPriceFiles()
	//var resultElem = make(map[string]map[string]string)
	for itemName := range priceList.S("products").ChildrenMap() {
		//itemName := "ENCHANTED_RABBIT_HIDE"
		nbOrder, _ := priceList.Path("products." + itemName + ".quick_status.buyOrders").Data().(float64)

		if nbOrder != 0 {

			//fmt.Println(priceList.Path("products." + itemName))
			sellPrice := priceList.Path("products." + itemName + ".quick_status.sellPrice").Data().(float64)
			if sellPrice < 5 {
				continue
			}

			listRessources := getListRessources(listCraft, itemName)
			if len(listRessources) == 0 {
				continue
			}
			var fullPrice float64 = 0
			for requirementName := range listRessources {
				requirementPrice := getPrice(priceList, requirementName)
				if requirementPrice == 0 {
					fullPrice = 0

					break
				}
				fullPrice += listRessources[requirementName] * requirementPrice

			}
			if fullPrice != 0 && sellPrice-fullPrice > float64(argLimit) {
				fmt.Println("need to check the item   " + itemName)
				fmt.Println("selling price of the item :")
				fmt.Println(sellPrice)
				fmt.Println("price of material : ")
				fmt.Println(fullPrice)
			}
		}

	}
}
func getPriceFiles() *gabs.Container {

	resp, err := http.Get("https://api.hypixel.net/skyblock/bazaar")
	if err != nil {
		fmt.Println("error sending the http request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	json, _ := gabs.ParseJSON(body)
	//fmt.Println(json)
	return json
}
func getPrice(priceList *gabs.Container, itemName string) float64 {

	price := priceList.Path("products." + itemName + ".quick_status.buyPrice")
	//fmt.Println(price)
	if price == nil {
		return 0
	}
	realPrice := price.Data().(float64)
	return realPrice
}
func getListRessources(listCraft *gabs.Container, itemName string) map[string]float64 {
	var res = make(map[string]float64)

	craftable := listCraft.Path(itemName + ".recipe")

	if craftable == nil {
		//fmt.Println("can't craft this Item, go to next")
		return res
	}

	for key, child := range listCraft.Path(itemName + ".recipe").ChildrenMap() {
		if key == "count" {
			continue
		}
		val := child.Data().(string)
		if val == "" {
			continue
		}
		sli := strings.Split(val, ":")
		nbElem, err := strconv.Atoi(sli[1])
		if err != nil {
			fmt.Println("err on the element ")
			fmt.Println(sli)
		}
		t := float64(nbElem)

		res[sli[0]] += t
	}

	return res
}
