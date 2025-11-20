package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// token info
type TokenData struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// image albums
type Image struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// info albums
type AlbumsData struct {
	Name         string  `json:"name"`
	Cover        []Image `json:"images"`
	Date         string  `json:"release_date"`
	NumberTracks int     `json:"total"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
}

// liste albums
type Albums struct {
	Items []AlbumsData `json:"items"`
}

// renvoie le token d'accès en string
func GetToken() string {
	httpClient := http.Client{
		Timeout: time.Second * 15,
	}
	//Partie faite a l'IA
	payload := strings.NewReader("grant_type=client_credentials")
	req, errReq := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token?grant_type=client_credentials&client_id=c2c125d4756c4d8692282454614ca245&client_secret=b6385721a0b54f268f37f939f13e442e", payload)
	authString := base64.StdEncoding.EncodeToString([]byte("c2c125d4756c4d8692282454614ca245" + ":" + "b6385721a0b54f268f37f939f13e442e"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+authString)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue:", errReq.Error())
	}
	req.Header.Add("User-Agent", "Ynov Campus Cours")

	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue:", errResp.Error())
		return ""
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue", errBody.Error())
		return ""
	}
	var tokenData TokenData
	json.Unmarshal(body, &tokenData)
	fmt.Println(tokenData.AccessToken)
	return tokenData.AccessToken
}

// interroge la BDD Spotify et renvoie les albums
func Api(urlapi string, token string) ([]Albums, error) {
	// URL
	//album Damso : https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums"
	fmt.Println("rentrer dans Api()")
	//init client HTTP
	httpClient := http.Client{
		Timeout: time.Second * 15,
	}

	// Création de la requête HTTP vers l'API avec initialisation de la methode HTTP, la route et le corps de la requête

	req, errReq := http.NewRequest(http.MethodGet, urlapi, nil)

	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue:", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header, User-Agent permet d'identifier l'application, système ....
	req.Header.Add("Authorization", "Bearer "+token)
	fmt.Println("req.Header.Add(Authorization, Bearer + token)")
	req.Header.Add("User-Agent", "Ynov Campus Cours")
	fmt.Println("req.Header.Add(User-Agent, Ynov Campus Cours)")

	// Execution de la requête HTTP vars l'API

	res, errResp := httpClient.Do(req)

	if errResp != nil {
		return nil, fmt.Errorf("Oupss, une erreur est survenue lors de l'exécution: %w", errResp)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP

	body, errBody := io.ReadAll(res.Body)
	fmt.Println("io.ReadAll(res.Body)")

	if errBody != nil {

		fmt.Println("Oupss, une erreur est survenue", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données

	var albumData []Albums

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &albumData)
	fmt.Println("Affichage des données de l'API Spotify")
	if albumData == nil {
		fmt.Println("decodeData est nil")
	}
	return albumData, nil
}
