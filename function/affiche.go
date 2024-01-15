package GridCreator

import (
	"net/http"
	"strconv"
	"strings"
)

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
	CalculItemPos(cases)
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

func Reset(w http.ResponseWriter, r *http.Request, grille *Grille) {
	for y, item := range grille.Grid {
		for x := range item {
			grille.Grid[y][x] = Case{x, y, "VIDE", "VIDE", "VIDE", false, false, 0, 0}
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CalculItemPos(cases *Case) {
	if cases.Flower {
		cases.ItemY = cases.Y*160 + 20
		cases.ItemX = cases.X*160 + 80
	} else if cases.Item {
		cases.ItemY = cases.Y*160 + 50
		cases.ItemX = cases.X*160 + 80
	} else {
		cases.ItemX = 0
		cases.ItemY = 0
	}
}
