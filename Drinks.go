package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/tidwall/gjson"
)

type drinks struct {
	strDrink      string
	strCategory   string
	strGlass      string
	strDrinkThumb string
}

var newDrink string

func randomDrink() string {

	//Query DB and get response
	url := "https://the-cocktail-db.p.rapidapi.com/random.php"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "the-cocktail-db.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "62a789f53bmsh39781af3958139dp14b2d8jsn9041f89ac768")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	//Convert response to []byte
	body, _ := ioutil.ReadAll(res.Body)

	//Convert response to string
	bodyJSON := string(body)

	//Iterate through JSON and pull out applicable info
	drinkName := gjson.Get(bodyJSON, "drinks.#.strDrink")
	drinkPic := gjson.Get(bodyJSON, "drinks.#.strDrinkThumb")
	drinkRecipe := gjson.Get(bodyJSON, "drinks.#.")
	drinkNameDirty := drinkName.String()
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	newDrink := reg.ReplaceAllString(drinkNameDirty, "")

	return newDrink
}

//
