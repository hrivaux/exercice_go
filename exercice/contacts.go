package main

import "fmt"

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans, email : %s", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"Employe : %s\nPoste : %s\nAdresse : %s\nSalaire : %.2f euros",
		e.Presentation(),
		e.Poste,
		e.Adresse.Format(),
		e.Salaire,
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire + e.Salaire*pct/100
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "TB"
	case e.Moyenne >= 14:
		return "B"
	case e.Moyenne >= 12:
		return "AB"
	case e.Moyenne >= 10:
		return "P"
	default:
		return "Aucune mention"
	}
}

func (e Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"Etudiant : %s\nPromo : %s\nMoyenne : %.2f\nMention : %s",
		e.Presentation(),
		e.Promo,
		e.Moyenne,
		e.MentionObtenue(),
	)
}

func main() {
	employe1 := Employe{
		Personne: Personne{
			Prenom: "Alice",
			Nom:    "Martin",
			Age:    32,
			Email:  "alice.martin@entreprise.fr",
		},
		Adresse: Adresse{
			Rue:        "12 rue des Lilas",
			Ville:      "Lyon",
			CodePostal: "69000",
		},
		Poste:   "Developpeuse Go",
		Salaire: 3200,
	}

	employe2 := Employe{
		Personne: Personne{
			Prenom: "Karim",
			Nom:    "Benali",
			Age:    41,
			Email:  "karim.benali@entreprise.fr",
		},
		Adresse: Adresse{
			Rue:        "8 avenue Victor Hugo",
			Ville:      "Paris",
			CodePostal: "75016",
		},
		Poste:   "Chef de projet",
		Salaire: 4100,
	}

	etudiant1 := Etudiant{
		Personne: Personne{
			Prenom: "Emma",
			Nom:    "Durand",
			Age:    21,
			Email:  "emma.durand@ecole.fr",
		},
		Promo:   "M2",
		Moyenne: 16.5,
	}

	etudiant2 := Etudiant{
		Personne: Personne{
			Prenom: "Lucas",
			Nom:    "Petit",
			Age:    22,
			Email:  "lucas.petit@ecole.fr",
		},
		Promo:   "M2",
		Moyenne: 13.2,
	}

	employe1.AugmenterSalaire(10)
	employe2.AugmenterSalaire(5)

	fmt.Println("===== FICHES EMPLOYES =====")
	fmt.Println(employe1.FicheEmploye())
	fmt.Println()
	fmt.Println(employe2.FicheEmploye())

	fmt.Println()
	fmt.Println("===== FICHES ETUDIANTS =====")
	fmt.Println(etudiant1.FicheEtudiant())
	fmt.Println()
	fmt.Println(etudiant2.FicheEtudiant())
}
