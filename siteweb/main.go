package main

import (
	api "TP-API-Spotify/API"
	"TP-API-Spotify/router"
	"fmt"
	"net/http"
)

func main() {
	token := api.GetToken()
	albums, err := api.Api("https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums", token)
	if err != nil {
		fmt.Printf("Erreur lors de l'appel API: %v\n", err)
	}
	fmt.Println("Albums reçus de l'API Spotify:", albums)
	r := router.New()
	fmt.Println("http://localhost:8080/home")
	http.ListenAndServe(":8080", r)

}
