package main

import (
	"LolGo/apikey"
	"LolGo/riotapi"
	"fmt"
)

func main() {
	ap := riotapi.NewAPI("euw", apikey.Key)
	a, err := ap.GetSummonerByName("Spellmaker")
	if err != nil {
		fmt.Println("error: ")
		fmt.Println(err)
	}

	id := a.ID

	b, err := ap.GetRankedStats(id, "SEASON3")

	if err != nil {
		fmt.Println("error: ")
		fmt.Println(err)
	}

	for _, el := range b.Champions {
		fmt.Println(el.Name)
	}
}
