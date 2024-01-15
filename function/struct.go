package GridCreator

type Case struct {
	X      int
	Y      int
	Type   string
	Type2  string
	Type3  string
	Item   bool
	Flower bool
	ItemX  int
	ItemY  int
}

type Grille struct {
	XTaille int
	YTaille int
	Grid    [][]Case
	Couche  int
	Start   int
	NbrItem int
}
