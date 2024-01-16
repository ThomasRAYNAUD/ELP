package main

import (
	"net/http"
	"html/template"
)

type Page struct {
	Valeur string
}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/home", homeHandler)
	http.ListenAndServe(":9999", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Valeur: "mon premier essai"}
	err := templates.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
// http://decouvric.cluster013.ovh.net/generalites/une-premiere-application-web.html

























































// Creer les workers
    // Premiere routine qui remplit le channel de taches à effectuer
    // Le channel est rempli en continu des photos à analyser ( dans le dossier testdata/images )

    // Deuxième routine qui fait l'analyse de visage en se basant sur le channel des taches à effectuer et qui remplit le channel des taches terminees
    // Analyse les visages de chaque photo du channel des taches à effectuer et remplit le channel des taches terminees avec toutes les photos à copier/ajouter le rectangle

    // Troisieme routine qui prend le channel des taches terminées et qui enregistre les visages
    // Prend le channel des taches terminees et crée le jpg avec rectangle de chaque photo