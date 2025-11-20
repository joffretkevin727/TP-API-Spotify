package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type album struct {
	Name   string
	Cover  string
	Date   string
	Number string
}
type infoTrack struct {
	Title      string
	CoverAlbum string
	NameAlbum  string
	NameArtist string
	Date       string
	Link       string
}

type Data struct {
	Artist string
	Album  []album
	Track  []infoTrack
}

func Api(urlApi string) (interface{}, error) {

	httpClient := &http.Client{
		Timeout: time.Second * 2, // Timeout après 2sec
	}

	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil) // Crée la requête HTTP
	if errReq != nil {
		return nil, fmt.Errorf("erreur de création de requête: %w", errReq)
	}

	req.Header.Add("User-Agent", "Ynov Campus Cours") // Ajoute le User-Agent

	res, errResp := httpClient.Do(req) // Exécute la requête
	if errResp != nil {
		return nil, fmt.Errorf("erreur lors de l'exécution de la requête: %w", errResp)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("requête API échouée avec le statut: %d", res.StatusCode)
	}

	defer res.Body.Close() // Ferme le corps de la réponse après exécution

	body, errBody := io.ReadAll(res.Body) // Lit et récupère le corps de la requête
	if errBody != nil {

		return nil, fmt.Errorf("erreur de lecture du corps de réponse: %w", errBody)
	}
	var decodedData interface{}
	// Déclare la variable qui va contenir les données

	if err := json.Unmarshal(body, &decodedData); err != nil {
		return nil, fmt.Errorf("erreur de decodage JSON: %w", err)

	}
	// donne les données JSON
	return decodedData, nil
}
