package GridCreator

import (
	"net/http"
	"strconv"
)

func AddX(grille *Grille) {
	grille.XTaille++
	for i := 0; i <= grille.YTaille; i++ {
		grille.Grid[i] = append(grille.Grid[i], Case{grille.XTaille + grille.Start, i, "VIDE", "VIDE", "VIDE", false, false, 0, 0})
	}
}

func AddY(grille *Grille) {
	grille.YTaille++
	slices := func() []Case {
		var slice []Case
		for i := 0; i <= grille.XTaille; i++ {
			slice = append(slice, Case{i + grille.Start, grille.YTaille, "VIDE", "VIDE", "VIDE", false, false, 0, 0})
		}
		return slice
	}
	grille.Grid = append(grille.Grid, slices())
}

func TailleGrilleParNumber(w http.ResponseWriter, r *http.Request, grille *Grille) {
	change := r.FormValue("choice")
	number, _ := strconv.Atoi(r.FormValue("number"))
	switch change {
	case "y":
		for i := 0; i < number; i++ {
			AddY(grille)
		}
	case "x":
		for i := 0; i < number; i++ {
			AddX(grille)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func TailleGrille(w http.ResponseWriter, r *http.Request, grille *Grille) {
	change := r.FormValue("Change")
	switch change {
	case "y+":
		AddY(grille)
	case "x+":
		AddX(grille)
	case "x-":
		RmX(grille)
	case "y-":
		RmY(grille)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func RmX(grille *Grille) {
	grille.XTaille--
	if len(grille.Grid) == 0 || len(grille.Grid[0]) == 0 {
		return
	}

	for i := range grille.Grid {
		grille.Grid[i] = grille.Grid[i][:len(grille.Grid[i])-1]
	}
}

func RmY(grille *Grille) {
	grille.YTaille--
	if len(grille.Grid) == 0 || len(grille.Grid[0]) == 0 {
		return
	}

	grille.Grid = grille.Grid[:len(grille.Grid)-1]
}
