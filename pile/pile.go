package pile

type Carte struct {
	valeur  int
	couleur string
}

type Pile struct {
	cartes   []Carte
	nbCartes int
}

func (p *Pile) Init(nbr int) {
	p.cartes = make([]Carte, nbr)
	p.nbCartes = 0
}

func (c *Carte) InitCarte(valeur int, couleur string) {
	c.valeur = valeur
	c.couleur = couleur
}

func (p *Pile) Empiler(c Carte) {
	p.cartes[p.nbCartes] = c
	p.nbCartes++
}

func (p *Pile) Depiler() Carte {
	c := p.cartes[p.nbCartes-1]
	p.nbCartes--
	return c
}

func (p *Pile) Sommet() Carte {
	return p.cartes[p.nbCartes-1]
}

func (p *Pile) EstVide() bool {
	return p.nbCartes == 0
}

func (p *Pile) Taille() int {
	return p.nbCartes
}

func (c *Carte) Valeur() int {
	return c.valeur
}

func (c *Carte) Couleur() string {
	return c.couleur
}

func (p *Pile) DebugNombres(nbr int) {
	p.nbCartes = nbr
}
