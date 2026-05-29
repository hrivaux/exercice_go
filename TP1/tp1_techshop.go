package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, produit := range c.Produits {
		if produit.ID == p.ID {
			return errors.New("un produit avec cet ID existe deja")
		}
	}

	c.Produits = append(c.Produits, p)
	return nil
}

func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, produit := range c.Produits {
		if produit.ID == id {
			return produit, nil
		}
	}

	return Produit{}, errors.New("produit introuvable")
}

func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultats []Produit

	for _, produit := range c.Produits {
		if produit.Actif && memeTexte(produit.Categorie, cat) {
			resultats = append(resultats, produit)
		}
	}

	return resultats
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nbModifies := 0

	for i, produit := range c.Produits {
		if produit.Actif && memeTexte(produit.Categorie, categorie) {
			c.Produits[i].Prix = produit.Prix - produit.Prix*pct/100
			nbModifies++
		}
	}

	return nbModifies
}

func (c *Catalogue) Vendre(id int, qte int) error {
	if qte <= 0 {
		return errors.New("la quantite doit etre positive")
	}

	for i, produit := range c.Produits {
		if produit.ID == id {
			if !produit.Actif {
				return errors.New("ce produit est inactif")
			}
			if produit.Stock < qte {
				return errors.New("stock insuffisant")
			}

			c.Produits[i].Stock -= qte
			if c.Produits[i].Stock == 0 {
				c.Produits[i].Actif = false
			}
			return nil
		}
	}

	return errors.New("produit introuvable")
}

func (c Catalogue) Rapport() string {
	nbProduits := 0
	valeurTotale := 0.0

	for _, produit := range c.Produits {
		if produit.Actif {
			nbProduits++
			valeurTotale += produit.Prix * float64(produit.Stock)
		}
	}

	return fmt.Sprintf("Produits actifs : %d\nValeur totale du stock : %.2f euros", nbProduits, valeurTotale)
}

func (p Produit) String() string {
	statut := "actif"
	if !p.Actif {
		statut = "inactif"
	}

	return fmt.Sprintf(
		"ID %d | %s %s | %.2f euros | stock : %d | categorie : %s | %s",
		p.ID,
		p.Marque,
		p.Nom,
		p.Prix,
		p.Stock,
		p.Categorie,
		statut,
	)
}

func minusculeSimple(texte string) string {
	resultat := []byte(texte)

	for i, lettre := range resultat {
		if lettre >= 'A' && lettre <= 'Z' {
			resultat[i] = lettre + 32
		}
	}

	return string(resultat)
}

func memeTexte(a string, b string) bool {
	return minusculeSimple(a) == minusculeSimple(b)
}

func lireTexte(message string) string {
	var texte string

	fmt.Print(message)
	fmt.Scan(&texte)
	return texte
}

func lireInt(message string) int {
	for {
		var texte string

		fmt.Print(message)
		fmt.Scan(&texte)

		nombre, err := strconv.Atoi(texte)
		if err == nil {
			return nombre
		}
		fmt.Println("Erreur : veuillez entrer un nombre entier.")
	}
}

func lireFloat(message string) float64 {
	for {
		var texte string

		fmt.Print(message)
		fmt.Scan(&texte)

		nombre, err := strconv.ParseFloat(texte, 64)
		if err == nil {
			return nombre
		}
		fmt.Println("Erreur : veuillez entrer un nombre.")
	}
}

func ajouterProduitCLI(catalogue *Catalogue) {
	produit := Produit{
		ID:        lireInt("ID : "),
		Nom:       lireTexte("Nom : "),
		Marque:    lireTexte("Marque : "),
		Prix:      lireFloat("Prix : "),
		Stock:     lireInt("Stock : "),
		Categorie: lireTexte("Categorie : "),
		Actif:     true,
	}

	if err := catalogue.AjouterProduit(produit); err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("Produit ajoute.")
}

func chercherProduitCLI(catalogue Catalogue) {
	id := lireInt("ID du produit : ")
	produit, err := catalogue.TrouverParID(id)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println(produit)
}

func soldesCLI(catalogue *Catalogue) {
	categorie := lireTexte("Categorie : ")
	pct := lireFloat("Reduction en % : ")

	if pct <= 0 || pct > 100 {
		fmt.Println("Erreur : la reduction doit etre entre 1 et 100.")
		return
	}

	nbModifies := catalogue.AppliquerReduction(categorie, pct)
	fmt.Println(nbModifies, "produit(s) modifie(s).")
}

func vendreCLI(catalogue *Catalogue) {
	id := lireInt("ID du produit : ")
	qte := lireInt("Quantite : ")

	if err := catalogue.Vendre(id, qte); err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("Vente effectuee.")
}

func afficherCategorieCLI(catalogue Catalogue) {
	categorie := lireTexte("Categorie : ")
	produits := catalogue.TrouverParCategorie(categorie)

	if len(produits) == 0 {
		fmt.Println("Aucun produit actif dans cette categorie.")
		return
	}

	for _, produit := range produits {
		fmt.Println(produit)
	}
}

func afficherMenu() {
	fmt.Println()
	fmt.Println("===== TechShop =====")
	fmt.Println("[1] Ajouter")
	fmt.Println("[2] Chercher par ID")
	fmt.Println("[3] Soldes")
	fmt.Println("[4] Vendre")
	fmt.Println("[5] Rapport")
	fmt.Println("[6] Chercher par categorie")
	fmt.Println("[0] Quitter")
}

func catalogueDeDepart() Catalogue {
	return Catalogue{
		Produits: []Produit{
			{ID: 1, Nom: "iPhone 15", Marque: "Apple", Prix: 899.99, Stock: 8, Categorie: "smartphone", Actif: true},
			{ID: 2, Nom: "MacBook Air", Marque: "Apple", Prix: 1299.99, Stock: 5, Categorie: "ordinateur", Actif: true},
			{ID: 3, Nom: "Galaxy S24", Marque: "Samsung", Prix: 799.99, Stock: 10, Categorie: "smartphone", Actif: true},
			{ID: 4, Nom: "ThinkPad X1", Marque: "Lenovo", Prix: 1499.99, Stock: 3, Categorie: "ordinateur", Actif: true},
			{ID: 5, Nom: "AirPods Pro", Marque: "Apple", Prix: 249.99, Stock: 15, Categorie: "audio", Actif: true},
		},
	}
}

func main() {
	catalogue := catalogueDeDepart()

	for {
		afficherMenu()
		choix := lireInt("Votre choix : ")

		switch choix {
		case 1:
			ajouterProduitCLI(&catalogue)
		case 2:
			chercherProduitCLI(catalogue)
		case 3:
			soldesCLI(&catalogue)
		case 4:
			vendreCLI(&catalogue)
		case 5:
			fmt.Println(catalogue.Rapport())
		case 6:
			afficherCategorieCLI(catalogue)
		case 0:
			fmt.Println("Au revoir.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
