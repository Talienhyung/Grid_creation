package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Case struct {
	X      int
	Y      int
	Type   string
	Type2  string
	Type3  string
	Item   bool
	Flower bool
}

type Grille struct {
	XTaille int
	YTaille int
	Grid    [][]Case
	Couche  int
	Start   int
	NbrItem int
}

func main() {
	var grille Grille
	grille.Grid = append(grille.Grid, []Case{{0, 0, "VIDE", "VIDE", "VIDE", false, false}})

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
				slice = append(slice, Case{i, grille.YTaille, "VIDE", "VIDE", "VIDE", false, false})
			}
			return slice
		}
		grille.Grid = append(grille.Grid, slices())
	case "x+":
		grille.XTaille++
		for i := 0; i <= grille.YTaille; i++ {
			grille.Grid[i] = append(grille.Grid[i], Case{grille.XTaille, i, "VIDE", "VIDE", "VIDE", false, false})
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Affiche(w http.ResponseWriter, r *http.Request, grille *Grille) {
	affiche := r.FormValue("affiche")
	tab := strings.Split(affiche, " ")
	x, _ := strconv.Atoi(tab[0])
	y, _ := strconv.Atoi(tab[1])
	ChangeType(&grille.Grid[y][x], grille.Couche)
	grille.NbrItem = CaculNombreItem(grille.Grid)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ChangeType(cases *Case, couche int) {
	switch couche {
	case 0:
		switch cases.Type {
		case "VIDE":
			cases.Type = "ONTOP"
		case "ONTOP":
			cases.Type = "ONBOTTOM"
		case "ONBOTTOM":
			cases.Type = "ONLEFT"
		case "ONLEFT":
			cases.Type = "ONRIGHT"
		case "ONRIGHT":
			cases.Type = "CENTER"
		case "CENTER":
			cases.Type = "BLOCK"
		case "BLOCK":
			cases.Type = "VIDE"
		}
	case 1:
		switch cases.Type2 {
		case "VIDE":
			cases.Type2 = "ONTOP"
		case "ONTOP":
			cases.Type2 = "ONBOTTOM"
		case "ONBOTTOM":
			cases.Type2 = "ONLEFT"
		case "ONLEFT":
			cases.Type2 = "ONRIGHT"
		case "ONRIGHT":
			cases.Type2 = "VIDE"
		}
	case 2:
		switch cases.Type3 {
		case "VIDE":
			cases.Type3 = "ONTOP"
		case "ONTOP":
			cases.Type3 = "ONBOTTOM"
		case "ONBOTTOM":
			cases.Type3 = "ONLEFT"
		case "ONLEFT":
			cases.Type3 = "ONRIGHT"
		case "ONRIGHT":
			cases.Type3 = "VIDE"
		}
	case 3:
		cases.Flower = !cases.Flower
	case 4:
		cases.Item = !cases.Item
	}
}

func Couche(w http.ResponseWriter, r *http.Request, grille *Grille) {
	grille.Couche, _ = strconv.Atoi(r.FormValue("couche"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CaculNombreItem(grid [][]Case) int {
	items := 0
	for _, item := range grid {
		for _, cases := range item {
			if cases.Item {
				items++
			}
		}
	}
	return items
}
