package logic

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"groupie/structs"
)

// function GetJson receives an API url and unparsing it into given struct. Also saving uparsed
// JSON data to a file.
func GetJson(url string, target interface{}, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error with receiving data: %v\n", err)
		return err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return err
	}
	file, _ := json.MarshalIndent(&target, "", "\t")
	if err = ioutil.WriteFile(filename, file, 0o644); err != nil {
		log.Printf("Error saving json file: %v\n", err)
		return err
	}
	return nil
}

// GetArtist , GetMembers, Relation functions  generates data from an API url,
// saves it to struct and returns it to handlers.
func GetArtist() []structs.Artists {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var user []structs.Artists
	file := "json/mainArtistList.json"
	err := GetJson(url, &user, file)
	if err != nil {
		log.Printf("Error getting json in GetArtist function: %s\n", err.Error())
	}

	return user
}

func GetMembers(id string) structs.Artists {
	url := "https://groupietrackers.herokuapp.com/api/artists/" + id
	var members structs.Artists
	file := "json/artistInfo.json"
	err := GetJson(url, &members, file)
	if err != nil {
		log.Printf("Error getting json in Members function: %s\n", err.Error())
	}
	return members
}

func Relation(id string) structs.Relations {
	url := "https://groupietrackers.herokuapp.com/api/relation/" + id
	var relation structs.Relations
	file := "json/concertDatesandLocation.json"
	err := GetJson(url, &relation, file)
	if err != nil {
		log.Printf("Error getting json Relation function: %s\n", err.Error())
	}
	return relation
}

// func SearchBar() {
// 	// var result []structs.Artists
// 	// var search []structs.Search

// 	// for i := range search {
// 	// 	if Compare(search[i], query) {
// 	// 		result = append(result, search[i].Name)
// 	// 	}
// 	// }
// 	url := "https://groupietrackers.herokuapp.com/api/artists"
// 	var all structs.Search
// 	file := "json/search1.json"
// 	if err := GetJson(url, &all, file); err != nil {
// 		log.Printf("Error getting json Relation function: %s\n", err.Error())
// 	}
// 	file = "json/search2.json"
// 	url = "https://groupietrackers.herokuapp.com/api/relation/"
// 	if err := GetJson(url, &all, file); err != nil {
// 		log.Printf("Error getting json Relation function: %s\n", err.Error())
// 	}
// 	file = "json/search3.json"
// 	url = "https://groupietrackers.herokuapp.com/api/dates"
// 	if err := GetJson(url, &all, file); err != nil {
// 		log.Printf("Error getting json Relation function: %s\n", err.Error())
// 	}
// 	file = "json/search4.json"
// 	url = "https://groupietrackers.herokuapp.com/api/locations"
// 	if err := GetJson(url, &all, file); err != nil {
// 		log.Printf("Error getting json Relation function: %s\n", err.Error())
// 	}
// 	// fmt.Println(all)
// }

// func Compare(query, compare string) bool {
// 	return strings.Contains(
// 		strings.ToLower(compare),
// 		strings.ToLower(query),
// 	)
// }
