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

// LES DONNEES DU TOKEN
type TokenData struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// ----------- LES DONNEES DES ALBUM--------------
type AlbumList struct {
	Total int     `json:"total"`
	Items []Album `json:"items"`
}

type Album struct {
	Artist []struct {
		Name string `json:"name"`
	} `json:"artists"`
	Name        string  `json:"name"`
	ReleaseDate string  `json:"release_date"`
	TotalTracks int     `json:"total_tracks"`
	Images      []Image `json:"images"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

//------------------------------------------------

// ------- LES INFO DE LA CHANSON MALADRESSE-------
type TrackInfo struct {
	NameTrack string `json:"name"`

	ExternalUrl struct {
		LinkSpotify string `json:"spotify"`
	} `json:"external_urls"`

	Album struct {
		NameAlbum   string `json:"name"`
		ReleaseDate string `json:"release_date"`

		Artists []struct {
			ArtistName string `json:"name"`
		} `json:"artists"`

		Images []struct {
			ImageURL string `json:"url"`
		} `json:"images"`
	} `json:"album"`
}

// ------------------------------------------------

// RECUPERE LE TOKEN NECESSAIRE A Apidamso() et ApiLaylow()
func GetToken() string {
	httpClient := http.Client{
		Timeout: time.Second * 15,
	}
	//CETTE PARTIE VA CONSTRUIRE LE LIEN DE LA REQUETE POST
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

// INTERROGE LA BDD SPOTIFY ET RENVOIE LA LISTE D'ALBUMS DE DAMSO
func ApiDamso(urlapi string, token string) (AlbumList, error) {

	//init client HTTP
	httpClient := http.Client{
		Timeout: time.Second * 15,
	}

	//la requetes GET HTTP
	req, errReq := http.NewRequest(http.MethodGet, urlapi, nil)

	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue:", errReq.Error())
	}

	// Ajout du token le header de la requête HTTP
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

// INTERROGE LA BDD SPOTIFY ET RENVOIE LES INFO DE LA CHANSON MALADRESSE
func ApiLaylow(urlapi string, token string) (TrackInfo, error) {

	//init client HTTP
	httpClient := http.Client{
		Timeout: time.Second * 15,
	}

	//la requete GET HTTP
	req, errReq := http.NewRequest(http.MethodGet, urlapi, nil)

	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue:", errReq.Error())
	}

	// Ajout du token dans le header de la requête HTTP
	req.Header.Add("Authorization", "Bearer "+token)

	// Execution de la requête HTTP

	res, errResp := httpClient.Do(req)

	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue lors de l'exécution: %w", errResp)
		return TrackInfo{}, errResp
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

	var Maladresse TrackInfo

	json.Unmarshal(body, &Maladresse)
	fmt.Println("Affichage des données de l'API Spotify")

	return Maladresse, nil
}
