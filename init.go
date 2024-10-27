package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ============================================================================== //
// Ce fichier contient toutes les fonctions permettant d'initialiser une variable //
// ============================================================================== //

func ChMots(filename string) ([]string, error) { // ChMots lit les mots à partir d'un fichier spécifié et retourne une slice de mots.
	fichier, err := os.Open(filename) // Ouvre le fichier spécifié par filename
	if err != nil {
		return nil, err // Retourne une erreur si l'ouverture du fichier échoue
	}
	defer fichier.Close() // S'assure que le fichier sera fermé à la fin de la fonction
	var mots []string
	scanner := bufio.NewScanner(fichier) // Crée un scanner pour lire le fichier ligne par ligne
	for scanner.Scan() {                 // Ajoute chaque ligne lue à la slice mots
		mots = append(mots, scanner.Text())
	}
	return mots, scanner.Err() // Retourne la slice de mots et l'erreur éventuelle du scanner
}

func RevL(mot string, n int, data *DataHang) string { // RevL révèle un certain nombre de lettres dans un mot donné, en fonction de la limite spécifiée.
	revele := []rune(strings.Repeat("_", len(mot)))
	lettresDejaRevelees := make(map[rune]bool) // Utilisé pour suivre les lettres déjà révélées afin d'éviter les doublons.
	indices := rand.Perm(len(mot))             // Génère une permutation aléatoire des indices du mot pour révéler des lettres dans un ordre aléatoire.
	for _, i := range indices {                // Parcourt les indices mélangés pour révéler les lettres.
		char := rune(mot[i])            // Obtient le caractère à la position indexée.
		if !lettresDejaRevelees[char] { // Vérifie si la lettre n'a pas déjà été révélée.
			occurrences := 0        // Compteur pour les occurrences de la lettre révélée.
			for j, c := range mot { // Parcourt le mot pour révéler toutes les occurrences de la lettre.
				if c == char && n > 0 { // Si le caractère correspond à la lettre et qu'il reste des lettres à révéler.
					revele[j] = c // Révèle la lettre dans le mot.
					occurrences++ // Incrémente le compteur d'occurrences.
					n--           // Décrémente le nombre de lettres restantes à révéler.
				}
			}
			lettresDejaRevelees[char] = true              // Marque la lettre comme révélée.
			data.UsedL = append(data.UsedL, string(char)) // Ajoute la lettre révélée à la liste des lettres utilisées.
			if n <= 0 {                                   // Si le nombre de lettres à révéler atteint zéro, on sort de la boucle.
				break
			}
		}
	}

	return string(revele) // Renvoie la version partiellement révélée du mot sous forme de chaîne.
}

func MotR(mots []string) string { // MotR retourne un mot aléatoire à partir d'une liste donnée de mots
	rand.Seed(time.Now().UnixNano())  // Initialise le générateur de nombres aléatoires avec l'heure actuelle pour garantir l'unicité des résultats
	return mots[rand.Intn(len(mots))] // `rand.Intn(len(mots))` génère un index aléatoire valide pour le tableau
}
