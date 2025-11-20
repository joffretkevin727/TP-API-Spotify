package controller

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	template := template.Must(template.ParseFiles("template/" + filename))
	if err := template.Execute(w, data); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html", nil)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func Track(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "track.html", nil)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func Album(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "album.html", nil)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func AlbumDamso(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "albumdamso.html", nil)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func TrackLaylow(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "tracklaylow.html", nil)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
