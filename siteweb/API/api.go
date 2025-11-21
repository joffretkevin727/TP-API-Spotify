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

// AlbumListe
type AlbumList struct {
	Total int     `json:"total"`
	Items []Album `json:"items"`
}

// Album info
type Album struct {
	Name        string  `json:"name"`
	ReleaseDate string  `json:"release_date"`
	TotalTracks int     `json:"total_tracks"`
	Images      []Image `json:"images"`
}

// image info
type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
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
func ApiDamso(urlapi string, token string) (AlbumList, error) {

	//init client HTTP
	httpClient := http.Client{
		Timeout: time.Second * 15,
	}

	//la requetes HTTP
	req, errReq := http.NewRequest(http.MethodGet, urlapi, nil)

	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue:", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header de la requête HTTP
	req.Header.Add("Authorization", "Bearer "+token)

	// Execution de la requête HTTP

	res, errResp := httpClient.Do(req)

	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue lors de l'exécution: %w", errResp)
		return AlbumList{}, errResp
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	fmt.Println("Statut de la réponse de l'API Spotify :", res.Body)

	// Lecture et récupération du corps de la requête HTTP

	body, errBody := io.ReadAll(res.Body)
	fmt.Println(io.ReadAll(res.Body))

	if errBody != nil {

		fmt.Println("Oupss, une erreur est survenue", errBody.Error())
	}

	var albumData AlbumList

	json.Unmarshal(body, &albumData)
	fmt.Println("Affichage des données de l'API Spotify")

	return albumData, nil
}
