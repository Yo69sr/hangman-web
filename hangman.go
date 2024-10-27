package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Structure pour stocker les données du jeu
type DonneesPendu struct {
	Mot         string // Le mot actuel avec les lettres devinées
	MotATrouver string // Le mot à deviner, choisi au hasard
	Tentatives  int    // Nombre de tentatives restantes
}

// Charger les mots depuis le fichier fourni
func ChargerMots(nomFichier string) ([]string, error) {
	fichier, err := os.Open(nomFichier) // Ouvre le fichier contenant les mots
	if err != nil {
		return nil, err // Retourne une erreur si le fichier ne peut pas être ouvert
	}
	defer fichier.Close() // Assure la fermeture du fichier à la fin de la fonction

	var mots []string
	scanner := bufio.NewScanner(fichier) // Crée un scanner pour lire le fichier
	for scanner.Scan() {
		mots = append(mots, scanner.Text()) // Ajoute chaque ligne (mot) à la liste
	}
	return mots, scanner.Err() // Retourne la liste de mots et une éventuelle erreur de lecture
}

// Choisir un mot aléatoire parmi la liste de mots
func MotAleatoire(mots []string) string {
	rand.Seed(time.Now().UnixNano())  // Initialise le générateur de nombres aléatoires
	return mots[rand.Intn(len(mots))] // Retourne un mot aléatoire de la liste
}

// Révéler quelques lettres dans le mot à deviner
func RevelerLettres(mot string, n int) string {
	revele := []rune(strings.Repeat("_", len(mot))) // Crée une représentation du mot avec des underscores
	indices := rand.Perm(len(mot))[:n]              // Génère des indices aléatoires pour révéler des lettres

	for _, i := range indices {
		revele[i] = rune(mot[i]) // Remplace les underscores par les lettres révélées
	}
	return string(revele) // Retourne la chaîne avec les lettres révélées
}

// Fonction principale pour jouer au jeu de Pendu
func Jouer() {
	mots, err := ChargerMots("words.txt") // Charge les mots depuis le fichier
	if err != nil {
		fmt.Println("Error loading words :", err) // Affiche une erreur en cas d'échec
		return
	}

	motATrouver := MotAleatoire(mots)                              // Choisit un mot à deviner
	motRevele := RevelerLettres(motATrouver, len(motATrouver)/2-1) // Révèle quelques lettres

	donnees := DonneesPendu{
		Mot:         motRevele,   // Initialise le mot affiché avec les lettres révélées
		MotATrouver: motATrouver, // Stocke le mot à deviner
		Tentatives:  10,          // Définit le nombre de tentatives
	}

	fmt.Println("Good luck, you have 10 attempts.") // Message d'accueil

	// Boucle principale du jeu
	for donnees.Tentatives > 0 && donnees.Mot != donnees.MotATrouver {
		fmt.Println("Word:", donnees.Mot) // Affiche le mot avec les lettres devinées
		fmt.Print("Propose a letter: ")   // Demande à l'utilisateur de proposer une lettre
		var proposition string
		fmt.Scan(&proposition) // Lit la proposition de l'utilisateur

		if len(proposition) != 1 { // Vérifie si la proposition est une seule lettre
			fmt.Println("Please enter a single letter.")
			continue // Si non, redemande une lettre
		}

		// Vérifie si la lettre proposée est dans le mot à deviner
		if strings.Contains(donnees.MotATrouver, proposition) {
			for i, c := range donnees.MotATrouver {
				if string(c) == proposition {
					donnees.Mot = donnees.Mot[:i] + proposition + donnees.Mot[i+1:] // Remplace le underscore par la lettre proposée
				}
			}
		} else {
			donnees.Tentatives--                                                          // Décrémente le nombre de tentatives si la lettre n'est pas présente
			fmt.Printf("Letter not present, %d attempts remaining\n", donnees.Tentatives) // Affiche le nombre de tentatives restantes
		}
	}

	// Fin du jeu, affichage du résultat
	if donnees.Mot == donnees.MotATrouver {
		fmt.Println("Congratulations! You have found the word:", donnees.MotATrouver) // Félicite l'utilisateur s'il a gagné
	} else {
		fmt.Println("Too bad! The word was:", donnees.MotATrouver) // Informe l'utilisateur du mot correct s'il a perdu
	}
}
