package main

import (
	"log"

	server "groupie-tracker/internal/app"
)

// type Client struct {
// 	Cl *http.Client
// }

// func getClient() *Client {
// 	return &Client{
// 		cl & http.Client,
// 	}
// }

func main() {
	// client := &Cl.Client{Timeout: 10 * time.Second}
	if err := server.App(); err != nil {
		log.Println(err.Error())
		return
	}
}
