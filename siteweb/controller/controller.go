package controller

import (
	"bytes"
	"html/template"
	"net/http"
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
func Track(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	RenderTemplate(w, "track.html", nil)
}
func Album(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	RenderTemplate(w, "album.html", nil)
}
func AlbumDamso(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	albumDamso := []album{
		{Name: "Titre1", Cover: "", Date: "Unknow", Number: "0"},
		{Name: "Titre2", Cover: "", Date: "Unknow2", Number: "00"},
		{Name: "Titre3", Cover: "", Date: "Unknow3", Number: "000"},
	}
	pageData := Data{
		Artist: "Damso",
		Album:  albumDamso,
		Track:  nil,
	}
	RenderTemplate(w, "albumdamso.html", pageData)

}
func TrackLaylow(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	trackLaylow := []infoTrack{
		{Title: "Maladresse", CoverAlbum: "[empty]", NameAlbum: ".Raw-Z", NameArtist: "Laylow", Date: "2018", Link: "https://open.spotify.com/intl-fr/track/0nAHBAlzkyaQXUp7qTULqv"},
	}
	pageData := Data{
		Artist: "Laylow",
		Album:  nil,
		Track:  trackLaylow,
	}
	RenderTemplate(w, "tracklaylow.html", pageData)
}
