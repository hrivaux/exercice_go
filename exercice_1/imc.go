package main

import "fmt"

func main() {
	var poids float64 = 90.0
	var taille float64 = 1.85

	const (
		IMCMaigreur = 18.5
		IMCNormal   = 25.0
		IMCSurpoids = 30.0
	)

	imc := poids / (taille * taille)

	fmt.Printf("Bonjour Hugo\n")
	fmt.Printf("Poids : %.1f kg\n", poids)
	fmt.Printf("Taille : %.2f m\n", taille)
	fmt.Printf("IMC : %.2f\n", imc)

	if imc < IMCMaigreur {
		fmt.Println("Categorie : Maigreur")
	} else if imc < IMCNormal {
		fmt.Println("Categorie : Normal")
	} else if imc < IMCSurpoids {
		fmt.Println("Categorie : Surpoids")
	} else {
		fmt.Println("Categorie : Obesite")
	}
}
