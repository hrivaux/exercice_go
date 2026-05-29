package main

import "fmt"

const (
	Petit = iota
	Moyen
	Grand
)

func main() {
	bonbons := []string{"fraise", "citron", "cola", "menthe"}
	taillePaquet := Grand

	fmt.Println("Bonbons disponibles :")

	for i, bonbon := range bonbons {
		fmt.Println(i+1, "-", bonbon)
	}

	fmt.Println()
	fmt.Println("Votre paquet contient :")

	switch taillePaquet {
	case Grand:
		fmt.Println("- bonus : bonbon cola")
		fallthrough
	case Moyen:
		fmt.Println("- bonus : bonbon citron")
		fallthrough
	case Petit:
		fmt.Println("- bonbon fraise")
	default:
		fmt.Println("Taille inconnue")
	}
}
