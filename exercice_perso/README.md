# Exercice perso - Machine a bonbons

## Niveau

Tres debutant.

## Objectif

Creer un petit programme Go qui affiche les bonbons disponibles dans une
machine, puis donne un paquet de bonbons selon une taille choisie.

Cet exercice sert a pratiquer :

- les slices ;
- la boucle `for` ;
- le `switch case` ;
- le `fallthrough` ;
- les constantes avec `iota`.

## Consigne

1. Cree une slice avec plusieurs noms de bonbons.
2. Affiche tous les bonbons avec une boucle `for`.
3. Cree trois tailles de paquet avec `const` et `iota` :
   - petit ;
   - moyen ;
   - grand.
4. Choisis une taille de paquet dans une variable.
5. Utilise un `switch` pour afficher ce que contient le paquet.
6. Utilise `fallthrough` pour que :
   - un grand paquet contienne aussi ce qu'il y a dans un moyen paquet ;
   - un moyen paquet contienne aussi ce qu'il y a dans un petit paquet.

## Pour tester

Depuis ce dossier :

```bash
go run main.go
```

