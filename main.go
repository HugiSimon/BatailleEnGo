package main

import (
	"GoTpBataille/pile"
	"fmt"
	"math/rand"
	"os"
	"time"

	term "github.com/nsf/termbox-go"
)

func generer(p *pile.Pile) {
	couleurs := []string{"Coeur", "Carreau", "Pique", "Trèfle"}
	for i := 0; i < 4; i++ {
		for j := 1; j < 14; j++ {
			c := pile.Carte{}
			c.InitCarte(j, couleurs[i])
			p.Empiler(c)
		}
	}
}

func melanger(p *pile.Pile) {
	rand.Seed(time.Now().UnixNano())
	alea := rand.Intn(22) + 5
	temp := [26]pile.Pile{}
	for i := 0; i < alea; i++ {
		temp[i].Init(12)
	}
	for !p.EstVide() {
		for i := 0; i < alea; i++ {
			if !p.EstVide() {
				temp[i].Empiler(p.Depiler())
			}
		}
	}

	for i := 0; i < alea; i++ {
		for !temp[i].EstVide() {
			p.Empiler(temp[i].Depiler())
		}
	}
}

func nombreMelange(p *pile.Pile) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rand.Intn(5)+5; i++ {
		melanger(p)
	}
}

func distribuer(p *pile.Pile, j1 *pile.Pile, j2 *pile.Pile) {
	for !p.EstVide() {
		j1.Empiler(p.Depiler())
		j2.Empiler(p.Depiler())
	}
}

func afficheCalcul(j1 *pile.Pile, j2 *pile.Pile, t1 *pile.Pile, t2 *pile.Pile, robot bool) {
	j1c1 := j1.Depiler()
	j1c2 := j1.Depiler()
	j2c1 := j2.Depiler()
	j2c2 := j2.Depiler()

	score1 := 0

	fmt.Print("\033[H\033[2J")

	fmt.Println("Joueur 1 : ", j1c1.Valeur(), j1c1.Couleur(), " et ", j1c2.Valeur(), j1c2.Couleur(), "| il reste ", j1.Taille()+2+t1.Taille(), " cartes")

	if robot {
		fmt.Println("\n                     ", calculScore(&j1c1, &j1c2, robot), " VS ", calculScore(&j2c1, &j2c2, robot), "\n")
		score1 = calculScore(&j1c1, &j1c2, robot)

	} else {
		fmt.Println("Quelle carte voulez-vous jouer en premier ? (1 ou 2)")
		choix := 0

	keyPressListenerLoop:
		for {
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				if ev.Ch == '1' {
					choix = 1
					break keyPressListenerLoop
				}
				if ev.Ch == '2' {
					choix = 2
					break keyPressListenerLoop
				}
			}
		}

		if choix == 1 {
			score1 = calculScore(&j1c1, &j1c2, robot)
		} else {
			score1 = calculScore(&j1c2, &j1c1, robot)
		}

		fmt.Println("\n                     ", score1, " VS ", calculScore(&j2c1, &j2c2, true), "\n")
	}
	fmt.Println("Joueur 2 : ", j2c1.Valeur(), j2c1.Couleur(), " et ", j2c2.Valeur(), j2c2.Couleur(), "| il reste ", j2.Taille()+2+t2.Taille(), " cartes")

	if robot {
		Attendre(robot)
	}

	if score1 > calculScore(&j2c1, &j2c2, true) {
		fmt.Println("\nJoueur 1 gagne !")
		t1.Empiler(j1c1)
		t1.Empiler(j1c2)
		t1.Empiler(j2c1)
		t1.Empiler(j2c2)
	} else if score1 < calculScore(&j2c1, &j2c2, true) {
		fmt.Println("\nJoueur 2 gagne !")
		t2.Empiler(j1c1)
		t2.Empiler(j1c2)
		t2.Empiler(j2c1)
		t2.Empiler(j2c2)
	} else {
		fmt.Println("\nEgalité !")
		t1.Empiler(j1c1)
		t1.Empiler(j1c2)
		t2.Empiler(j2c1)
		t2.Empiler(j2c2)
	}

	Attendre(robot)
}

func calculScore(c1 *pile.Carte, c2 *pile.Carte, robot bool) int {
	temp := 0
	if robot {
		if c1.Valeur() > 9 {
			if c2.Valeur() < 9 {
				temp = c2.Valeur() * 10
			}
		} else if c2.Valeur() > 9 {
			if c1.Valeur() < 9 {
				temp = c1.Valeur() * 10
			}
		} else if c1.Valeur() > c2.Valeur() {
			temp = c1.Valeur()*10 + c2.Valeur()
		} else {
			temp = c2.Valeur()*10 + c1.Valeur()
		}
	} else {
		if c1.Valeur() < 10 {
			temp += c1.Valeur() * 10
		}
		if c2.Valeur() < 10 {
			temp += c2.Valeur()
		}
	}
	return temp
}

func Attendre(robot bool) {
	if robot {
		time.Sleep(1 * time.Second)
	} else {
	keyPressListenerLoop:
		for {
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				if ev.Key == term.KeyEnter {
					break keyPressListenerLoop
				}
			}
		}
	}
}

func reRemplir(p *pile.Pile, t *pile.Pile) {
	for !t.EstVide() {
		p.Empiler(t.Depiler())
	}
	nombreMelange(p)
}

func main() {
	robot := false
	if len(os.Args) > 1 && os.Args[1] == "robot" {
		robot = true
	}

	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

	deck := pile.Pile{}
	deck.Init(52)

	generer(&deck)
	nombreMelange(&deck)

	joueur1 := pile.Pile{}
	joueur2 := pile.Pile{}
	joueur1.Init(52)
	joueur2.Init(52)

	distribuer(&deck, &joueur1, &joueur2)

	joueur1temp := pile.Pile{}
	joueur2temp := pile.Pile{}
	joueur1temp.Init(52)
	joueur2temp.Init(52)

	for !joueur1.EstVide() && !joueur2.EstVide() {
		afficheCalcul(&joueur1, &joueur2, &joueur1temp, &joueur2temp, robot)
		if joueur1.EstVide() && !joueur1temp.EstVide() {
			reRemplir(&joueur1, &joueur1temp)
			fmt.Println("Re-remplissage de la pile du joueur 1")
		}
		if joueur2.EstVide() && !joueur2temp.EstVide() {
			reRemplir(&joueur2, &joueur2temp)
			fmt.Println("Re-remplissage de la pile du joueur 2")
		}
	}
	Attendre(robot)

	fmt.Print("\033[H\033[2J")
	fmt.Println("\n\nFin de la partie !")

	if joueur1.Taille() > joueur2.Taille() {
		fmt.Println("--- Joueur 1 gagne ! Bravo ! ---")
	} else {
		fmt.Println("--- Joueur 2 gagne ! Bravo ! ---")
	}

	Attendre(robot)

}

// TODO
// -L'égalité doit être mieux gérée
