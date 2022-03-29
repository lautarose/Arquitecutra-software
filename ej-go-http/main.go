package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Categories []Category

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	var tam int
	cats, err := GetCategories("https://api.mercadolibre.com/sites/MLA/categories")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Las categorias de MLA son:")
	fmt.Println("")
	fmt.Println("  id             name")
	fmt.Println("------------------------------------------")

	tam = len(cats)
	for i := 0; i < tam; i++ {
		fmt.Println(cats[i])
	}

}

func GetCategories(siteID string) (Categories, error) {

	var cats Categories
	response, err := http.Get(siteID)

	if err != nil {
		return cats, err
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return cats, err
	}

	err = json.Unmarshal(bytes, &cats)

	if err != nil {
		return cats, err
	}

	return cats, err
}
