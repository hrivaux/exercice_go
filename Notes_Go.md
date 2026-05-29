# Notes de cours — Langage Go

## Les entiers

Les entiers en Go servent à stocker des nombres sans partie décimale.

Exemples : `int`, `int8`, `int16`, `int32`, `int64`, `uint`.

Le type `int` est le plus couramment utilisé pour représenter des nombres entiers.

---

## Variables et constantes

Une variable peut voir sa valeur changer au cours de l’exécution du programme.

```go
var age int = 20
age = 21
```

Une constante garde toujours la même valeur après sa déclaration.

```go
const PI = 3.14
```

---

## Différence entre `:=` et `=`

L’opérateur `:=` permet de déclarer une variable tout en lui affectant une valeur.

```go
nom := "Youssef"
```

L’opérateur `=` permet de modifier la valeur d’une variable déjà déclarée.

```go
nom = "Ali"
```

---

## Les slices

Un slice est une structure de données dynamique.

Sa taille peut évoluer contrairement à celle d’un tableau.

```go
notes := []int{12, 15, 18}
notes = append(notes, 20)
```

---

## `len` et `cap`

`len` retourne le nombre d’éléments contenus dans un slice.

```go
len(notes)
```

`cap` retourne sa capacité, c’est-à-dire l’espace qui lui est alloué en mémoire.

```go
cap(notes)
```

---

## `copy` des slices

La fonction `copy` sert à copier les éléments d’un slice vers un autre.

```go
source := []int{1, 2, 3}
destination := make([]int, len(source))

copy(destination, source)
```

---

## `iota`

`iota` permet de générer automatiquement des valeurs constantes.

```go
const (
    Faible = iota
    Moyen
    Eleve
)
```

Dans cet exemple : `Faible = 0`, `Moyen = 1` et `Eleve = 2`.

---

## Maps

Une map stocke des données sous forme de couples clé/valeur.

```go
ages := map[string]int{
    "Youssef": 24,
    "Ali": 22,
}
```

Une valeur est récupérée à partir de sa clé.

```go
fmt.Println(ages["Youssef"])
```

---

## Boucle `for`

Go ne possède qu’un seul type de boucle : `for`.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Elle peut également être utilisée comme une boucle `while`.

```go
for condition {
    // code
}
```

---

## Fonctions variadiques

Une fonction variadique accepte un nombre variable d’arguments d’un même type.

```go
func afficherNoms(noms ...string) {
    for _, nom := range noms {
        fmt.Println(nom)
    }
}
```

---

## Closure

Une closure est une fonction capable de conserver une variable en mémoire ou de retourner une autre fonction.

```go
func compteur() func() int {
    n := 0

    return func() int {
        n++
        return n
    }
}
```

---

## `fallthrough`

Le mot-clé `fallthrough` permet de poursuivre l’exécution dans le cas suivant d’un `switch`.

```go
switch niveau {
case 1:
    fmt.Println("Niveau 1")
    fallthrough
case 2:
    fmt.Println("Niveau 2")
}
```

---

## Structures

Une structure permet de regrouper plusieurs données dans un même type.

```go
type Personne struct {
    Nom string
    Age int
}
```

---

## Notion de classe en Go

Go ne possède pas de classes comme Java.

À la place, il utilise des `structs` associées à des méthodes.

---

## Méthodes

Une méthode ajoute un comportement à une struct.

```go
func (p Personne) SePresenter() {
    fmt.Println("Bonjour, je suis", p.Nom)
}
```

---

## Pointeurs

Un pointeur permet de manipuler directement l’adresse mémoire d’une variable.

```go
func modifierAge(p *Personne) {
    p.Age = 25
}
```

---

## Objets en Go

Les objets en Go sont construits à partir de structs et de méthodes.

```go
type Produit struct {
    Nom  string
    Prix float64
}
```

---

## Embedding

Go utilise la composition via l’embedding.

```go
type Personne struct {
    Nom string
}

type Employe struct {
    Personne
    Poste string
}
```

---

## Struct tags

Les struct tags sont des informations complémentaires associées aux champs d’une struct.

```go
type Personne struct {
    Nom string `json:"nom,omitempty"`
}
```

```go
MotDePasse string `json:"-"`
```

---

## Visibilité en Go

La visibilité dépend de la première lettre du nom.

```go
type Personne struct {}
```

```go
type personne struct {}
```

---

## Camel Case et Pascal Case

```go
nomProduit
calculerPrix
```

```go
NomProduit
CalculerPrix
```

---

## Gestion des erreurs : Go vs Java

```java
try {
    // code
} catch (Exception e) {
    // erreur
}
```

```go
resultat, err := calculer()

if err != nil {
    fmt.Println("Erreur :", err)
}
```

---

## Mot-clé `defer`

```go
defer fmt.Println("Fin du programme")
```

```go
defer fichier.Close()
```
