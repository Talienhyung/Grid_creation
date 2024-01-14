package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Case struct {
	X    int
	Y    int
	Type string
}

type Grille struct {
	XTaille int
	YTaille int
	Grid    [][]Case
}

func main() {
	var grille Grille
	grille.Grid = append(grille.Grid, []Case{{0, 0, "VIDE"}})
	printGrille(grille.Grid)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Home(w, r, grille)
	})
	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		TailleGrille(w, r, &grille)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}

// Home handles HTTP requests for the home page and renders the appropriate HTML templates
func Home(w http.ResponseWriter, r *http.Request, infos Grille) {
	template, err := template.ParseFiles(
		"index.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

func TailleGrille(w http.ResponseWriter, r *http.Request, grille *Grille) {
	change := r.FormValue("Change")
	switch change {
	case "y+":
		grille.YTaille++
		slices := func() []Case {
			var slice []Case
			for i := 0; i <= grille.XTaille; i++ {
				slice = append(slice, Case{i, grille.YTaille, "VIDE"})
			}
			return slice
		}
		grille.Grid = append(grille.Grid, slices())
	case "x+":
		grille.XTaille++
		for i := 0; i <= grille.YTaille; i++ {
			grille.Grid[i] = append(grille.Grid[i], Case{grille.XTaille, i, "VIDE"})
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func printGrille(grille [][]Case) {
	for _, item := range grille {
		for _, item2 := range item {
			fmt.Print(item2)
		}
		fmt.Println("")
	}
}