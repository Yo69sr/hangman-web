package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

// ======================================================================== //
// Ce fichier contient toutes les fonctions gérant les entrées utilisateurs //
// ======================================================================== //

func ChooseD() string { // ChooseD demande à l'utilisateur de sélectionner une difficulté pour le jeu du Pendu
	ClearT()
	for {
		fmt.Println("Choose a difficulty: easy, medium, hard.")     // Affiche les options de difficulté à l'utilisateur
		scanner := bufio.NewScanner(os.Stdin)                       // Crée un scanner pour capturer l'entrée de l'utilisateur depuis la console
		scanner.Scan()                                              // Lit l'entrée de l'utilisateur
		choix := strings.TrimSpace(strings.ToLower(scanner.Text())) // Normalise l'entrée : enlève les espaces et met en minuscules
		mots := strings.Fields(choix)                               // Divise la chaîne d'entrée en mots basés sur les espaces
		if len(mots) == 1 {                                         // Vérifie qu'il n'y a qu'un seul mot
			switch mots[0] { // Compare le mot choisi à des options valides
			case "easy", "medium", "hard": // Si le choix est valide
				return mots[0]
			default:
				ClearT()
				fmt.Println("Invalid choice, please choose between easy, medium, or hard.")
			}
		} else {
			ClearT()
			fmt.Println("Invalid choice, please enter a single word.")
		}
	}
}

func ValL(data *DataHang) string { // ValL demande à l'utilisateur de proposer une lettre ou un mot et gère les entrées
	for {
		fmt.Print("Propose a letter or a word (or type STOP to quit):")
		var proposition string
		fmt.Scan(&proposition)                     // Lit l'entrée de l'utilisateur
		proposition = strings.ToLower(proposition) // Convertit la proposition en minuscules

		if strings.ToUpper(proposition) == "STOP" { // Vérifie si l'utilisateur souhaite quitter le jeu

			err := Save(data, "save.txt") // Sauvegarde l'état actuel du jeu dans un fichier
			if err != nil {               // Affiche une erreur si la sauvegarde échoue
				for i := 5; i > 0; i-- {
					fmt.Println("Error while saving the game:", err)
					time.Sleep(1 * time.Second)
					fmt.Printf("Closing the program in %d seconds.", i)
					ClearT()
				}
			} else {
				for i := 5; i > 0; i-- { // Confirme que la partie a été sauvegardée avec succès
					fmt.Println("Game saved in save.txt. See you soon!")
					fmt.Printf("Closing the program in %d seconds.", i)
					time.Sleep(1 * time.Second)
					ClearT()
				}
			}
			os.Exit(0) // Quitte le programme
		}
		if len(proposition) == 1 && unicode.IsLetter(rune(proposition[0])) { // Vérifie si la proposition est une seule lettre
			if !Cont(data.UsedL, proposition) { // Vérifie si la lettre a déjà été utilisée
				data.UsedL = append(data.UsedL, proposition) // Ajoute la lettre aux lettres utilisées
				return proposition                           // Retourne la lettre proposée
			}

			if data.Pv <= 0 { // Vérifie si le nombre de tentatives restantes est épuisé
				ClearT()
				AffPF(data)
				fmt.Println("You lost. The word was:", data.SearchedW)
				break
			}
			ClearT()
			fmt.Printf("You have already used this letter. %d attempts remaining\n", data.Pv)
			AffP(*data)
			fmt.Println("Word:", data.Word)
			fmt.Printf("Letters already used: %v\n", data.UsedL)
			fmt.Printf("Words already used: %v\n", data.UsedW)
		} else if len(proposition) > 1 && IsMVal(proposition) { // Vérifie si la proposition est un mot valide
			if !Cont(data.UsedW, proposition) { // Vérifie si le mot a déjà été utilisé
				data.UsedW = append(data.UsedW, proposition) // Ajoute le mot aux mots utilisés
				return proposition                           // Retourne le mot proposé
			}
			if data.Pv <= 0 { // Vérifie si le nombre de tentatives restantes est épuisé
				ClearT()
				AffPF(data)
				fmt.Println("You lost. The word was:", data.SearchedW)
				break
			}
			ClearT()
			fmt.Printf("You have already used this word. %d attempts remaining.", data.Pv) // Indique que le mot a déjà été utilisé et affiche les tentatives restantes
			AffP(*data)
			fmt.Println("Word:", data.Word)
			fmt.Printf("Letters already used: %v\n", data.UsedL)
			fmt.Printf("Words already used: %v\n", data.UsedW)
		}
	}
	return "error" // Retourne "error" si aucune condition n'est remplie
}
