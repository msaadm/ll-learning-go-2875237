package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

type Tour struct {
	Name  string
	Price string
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()

	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}

		tours = append(tours, tour)
	}

	return tours
}

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(nil)
	}

	content := string(bytes)

	// fmt.Print(content)

	tours := toursFromJson(content)

	for _, tour := range tours {
		priceInFloat, _ := strconv.ParseFloat(tour.Price, 64)
		fmt.Printf("Tour \"%s\" in price of %.2f$\n", tour.Name, priceInFloat)
	}
}
