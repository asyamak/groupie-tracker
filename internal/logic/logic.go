package logic

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-tracker/models"
)

type Account map[string]interface{}

var (
	ArrayRelations []models.Relations
	Array          []models.ArtistResult
)

var (
	MapArtists   = make(map[int]models.ArtistResult)
	MapRelations = make(map[int]models.Relations)
)

func GetRelations() {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	// x, _ := http.Get(url)
	// fmt.Println(x)
	group := models.GroupId{}
	err := GetJson(url, &group)
	// fmt.Println(group)
	if err != nil {
		fmt.Printf("error getting json: %s\n", err.Error())
	}
	for _, rel := range group.Index {
		ArrayRelations = append(ArrayRelations, rel)
		MapRelations[rel.Id] = rel
		// fmt.Println(rel)
	}
}

func GetArtistss() []models.ArtistResult {
	// fmt.Println(1)
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var user []models.ArtistResult
	err := GetJson(url, &user)
	if err != nil {
		fmt.Printf("error getting json: %s\n", err.Error())
	}
	for _, us := range user {
		Array = append(Array, us)
		MapArtists[us.ID] = us
	}
	return user
}

func GetJson(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&target)
	return nil
}
