package main

import (
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseFiles("index.html")) 
//charge le fichier "index.html" comme modèle
// stocke le modèle prêt à être utilisé dans la variable templates

func main() {
	http.HandleFunc("/home", homeHandler) // correspondance entre l'URL "/home" et la fonction homeHandler -> homeHandler sera executé quand qlq accède à l'URL
	http.ListenAndServe(":9999", nil) // démarre un serveur web HTTP sur le port 9999
}

func homeHandler(w http.ResponseWriter, r *http.Request) { //w = définit comment écrire une réponse HTTP
    err := templates.ExecuteTemplate(w, "index.html", nil) //applique et transmet la template en http -> sans trasmettre de valeur
    if err != nil { // si erreur retourne l'erreur
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
