package GridCreator

import (
	"log"
	"net/http"
	"text/template"
)

func Root() {
	var grille Grille
	grille.Grid = append(grille.Grid, []Case{{0, 0, "VIDE", "VIDE", "VIDE", false, false, 0, 0}})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Home(w, r, grille)
	})
	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		TailleGrille(w, r, &grille)
	})
	http.HandleFunc("/affiche", func(w http.ResponseWriter, r *http.Request) {
		Affiche(w, r, &grille)
	})
	http.HandleFunc("/couche", func(w http.ResponseWriter, r *http.Request) {
		Couche(w, r, &grille)
	})
	http.HandleFunc("/changeNbr", func(w http.ResponseWriter, r *http.Request) {
		TailleGrilleParNumber(w, r, &grille)
	})
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		Reset(w, r, &grille)
	})
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		StartChange(w, r, &grille)
	})
}

// Home handles HTTP requests for the home page and renders the appropriate HTML templates
func Home(w http.ResponseWriter, r *http.Request, infos Grille) {
	template, err := template.ParseFiles(
		"./index.html",
		"./template/case.html",
		"./template/footer.html",
		"./template/text.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}
