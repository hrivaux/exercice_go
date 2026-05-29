package main

import (
	"errors"
	"fmt"
	"strconv"
)

func operation(op string) (func(float64, float64) (float64, error), error) {
	switch op {
	case "+":
		return func(a, b float64) (float64, error) {
			return a + b, nil
		}, nil
	case "-":
		return func(a, b float64) (float64, error) {
			return a - b, nil
		}, nil
	case "*":
		return func(a, b float64) (float64, error) {
			return a * b, nil
		}, nil
	case "/":
		return func(a, b float64) (float64, error) {
			if b == 0 {
				return 0, errors.New("division par zero impossible")
			}

			return a / b, nil
		}, nil
	default:
		return nil, errors.New("operation inconnue")
	}
}

func main() {
	for {
		var premierNombre string
		var b float64
		var op string

		fmt.Print("Entrez deux nombres et une operation (+, -, *, /) ou quit : ")

		if _, err := fmt.Scan(&premierNombre); err != nil {
			fmt.Println("Erreur : saisie invalide")
			continue
		}

		if premierNombre == "quit" {
			break
		}

		a, err := strconv.ParseFloat(premierNombre, 64)
		if err != nil {
			fmt.Println("Erreur : saisie invalide")
			continue
		}

		if _, err := fmt.Scan(&b, &op); err != nil {
			fmt.Println("Erreur : saisie invalide")
			continue
		}

		if op == "quit" {
			break
		}

		calcul, err := operation(op)
		if err != nil {
			fmt.Println("Erreur :", err)
			continue
		}

		resultat, err := calcul(a, b)
		if err != nil {
			fmt.Println("Erreur :", err)
			continue
		}

		fmt.Printf("Resultat : %.2f\n", resultat)
	}
}
