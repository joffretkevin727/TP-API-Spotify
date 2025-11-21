package controller

import (
	api "TP-API-Spotify/API"
	"bytes"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	template := template.Must(template.ParseFiles("template/" + filename))

	buf := new(bytes.Buffer)
	if err := template.Execute(buf, data); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}

func Home(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	RenderTemplate(w, "home.html", nil)
}

func AlbumDamso(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	albumUrl := "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums"
	albumData, err := api.ApiDamso(albumUrl, api.GetToken())

	if err != nil {
		http.Error(w, "Erreur lors de l'appel API: %v", http.StatusInternalServerError)
		return
	}
	RenderTemplate(w, "albumdamso.html", albumData)

}
func TracksLaylow(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//RenderTemplate(w, "tracks/laylow.html", api.TrackList)
}
