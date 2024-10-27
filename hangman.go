package hangman

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// =========================================================================================== //
// Ce fichier contient La fonction principal de mon jeu ainsi que le type de structure utilisé //
// =========================================================================================== //

type DataHang struct { // DataHang contient toutes les informations nécessaires pour suivre une partie
	Word      string     // Le mot choisi pour la partie actuelle
	SearchedW string     // Le mot affiché au joueur
	Pv        int        // Le nombre de tentatives restantes
	PosHang   [10]string // Les différentes étapes du pendu
	UsedL     []string   // liste contenant les lettres que le joueur a déjà essayées
	UsedW     []string   // liste des mots déjà utilisés
}

func Play() { // Play gère les appels de fonctions pour le jeu du Pendu.
	var debug bool
	var filename string
	var load bool
	var err error
	var message_bjr string = "HELLO!"
	var message_gg string = "VICTORY!"
	var message_wp string = "NICE TRY"
	chars := []string{ // Tableau des caractères utilisés pour l'art ASCII.
		"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		":", ";", "<", "=", ">", "?", "@",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z", "[", "\\", "]", "^", "_", "`",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z", "{", "|", "}", "²",
	}
	asciiMap, err := ChASCIIArt("standard.txt", chars) // Charge les caractères ASCII à partir du fichier "standard.txt".
	if err != nil {
		fmt.Println("Error loading the file:", err)
		return
	}
	for i := 5; i > 0; i-- { // Affiche un compte à rebours de 5 secondes avant de commencer la partie.
		ClearT()
		PrintASCIIM(asciiMap, message_bjr)
		fmt.Printf("The game will start in %d seconds.", i)
		time.Sleep(1 * time.Second) // Attente d'une seconde.
	}
	ClearT()
	PrintASCIIM(asciiMap, message_bjr)
	var saveFile string
	flag.StringVar(&saveFile, "startWith", "", "Fichier de sauvegarde à charger") // Récupère un fichier de sauvegarde potentiel via un flag "startWith".
	flag.Parse()
	var data *DataHang
	if saveFile != "" { // Si un fichier de sauvegarde est fourni, on charge la partie sauvegardée.
		data, filename, err = Load(saveFile)
		load = true
		if err != nil {
			fmt.Println("Error loading the save:", err) // Message d'erreur en cas d'échec de chargement de la sauvegarde.
			return
		}
	} else { // Si aucun fichier de sauvegarde n'est fourni, on commence une nouvelle partie.
		difficulte := ChooseD() // Demande la difficulté au joueur.
		var fichierMots string
		switch difficulte { // Selon la difficulté choisie, on charge un fichier.
		case "easy":
			fichierMots = "words.txt"
		case "medium":
			fichierMots = "words2.txt"
		case "hard":
			fichierMots = "words3.txt"
		}
		mots, err := ChMots(fichierMots) // Charge les mots correspondant à la difficulté choisie.
		if err != nil {
			fmt.Println("Error loading the words:", err) // Message d'erreur en cas d'échec de chargement des mots.
			return
		}
		positionsPendu, err := ChHangman("hangman.txt") // Charge les différentes positions du pendu à afficher au fur et à mesure.
		if err != nil {
			fmt.Println("Error loading hangman positions:", err) // Message d'erreur en cas d'échec de chargement des positions.
			return
		}
		motATrouver := MotR(mots) // Choisit un mot aléatoire.
		data = &DataHang{
			SearchedW: motATrouver,
			Pv:        10,
			PosHang:   positionsPendu,
		}
		motRevele := RevL(motATrouver, len(motATrouver)/2, data) // Révèle certaines lettres du mot.
		data.Word = motRevele
		ClearT()
		fmt.Println("Would you like to start the game on debug mode (with the answer printed)? [y/n]")
		for {
			scanner := bufio.NewScanner(os.Stdin) // Utilisation d'un scanner pour capturer l'entrée de l'utilisateur.
			scanner.Scan()
			res := strings.TrimSpace(strings.ToLower(scanner.Text()))
			switch res {
			case "y":
				debug = true
				ClearT()
				fmt.Println("Debug mode is ON.")
			case "n":
				debug = false
				ClearT()
				fmt.Println("Debug mode is OFF.")
			default:
				ClearT()
				fmt.Println("Invalid input. Please enter 'y' for yes or 'n' for no.")
				fmt.Println("Would you like to start the game on debug mode (with the answer printed)? [y/n]")
			}

			if res == "y" || res == "n" {
				break // On sort explicitement de la boucle après une réponse valide.
			}
		}
		fmt.Println("Good luck, you have 10 attempts.")
		fmt.Println("Press Enter to start the game.")
		fmt.Scanln() // Attend une entrée de l'utilisateur.
		print("\n")
	}
	ClearT()
	for data.Pv > 0 && data.Word != data.SearchedW { // Boucle principale du jeu : continue tant que le joueur a des tentatives et n'a pas deviné le mot.
		if debug {
			fmt.Println("Answer:", data.SearchedW) // Affiche la réponse si le mode débogage est activé.
		}
		fmt.Println("Word:", data.Word)
		fmt.Printf("Letters already used: %v\n", data.UsedL)
		fmt.Printf("Words already used: %v\n", data.UsedW)
		proposition := ValL(data) // Le joueur propose une lettre ou un mot.
		if len(proposition) > 1 { // Si la proposition est un mot complet.
			if proposition == data.SearchedW { // Si le mot est correct, le joueur gagne.
				ClearT()
				for i := 5; i > 0; i-- {
					PrintASCIIM(asciiMap, message_gg) // Affiche un message de victoire en ASCII art.
					fmt.Println("You found the word:", data.SearchedW)
					fmt.Printf("The program will close in %d seconds.\n", i)
					time.Sleep(1 * time.Second)
					ClearT()
				}
				PrintASCIIM(asciiMap, message_gg)
				fmt.Println("You found the word:", data.SearchedW)
				fmt.Println("Press Enter to close the program.")
				fmt.Scanln()
				break
			} else { // Si le mot est incorrect, deux tentatives sont retirées.
				ClearT()
				data.Pv -= 2
				if data.Pv > 0 {
					fmt.Printf("Incorrect word, %d attempts remaining\n", data.Pv)
					AffP(*data)
				}
			}
		} else {
			if strings.Contains(data.SearchedW, proposition) { // Si la proposition est une lettre.
				if LREver(data.Word, data.SearchedW, proposition) { // Si la lettre est dans le mot.
					fmt.Printf("Letter already revealed everywhere, %d attempts remaining\n", data.Pv)
				} else { // Révèle la lettre dans le mot.
					for i, c := range data.SearchedW {
						if string(c) == proposition {
							data.Word = data.Word[:i] + proposition + data.Word[i+1:] // Met à jour le mot avec la lettre révélée.
						}
					}
				}
				ClearT()
				AffP(*data)
			} else { // Si la lettre n'est pas dans le mot, une tentative est retirée.
				ClearT()
				data.Pv--
				fmt.Printf("Letter not found, %d attempts remaining\n", data.Pv) // Message indiquant que la lettre n'a pas été trouvée.
				AffP(*data)
			}
		}
		if data.Pv <= 0 { // Si le joueur n'a plus de tentatives, il perd.
			ClearT()
			for i := 5; i > 0; i-- {
				PrintASCIIM(asciiMap, message_wp) // Affiche un message de défaite.
				fmt.Println("You lost. The word was:", data.SearchedW)
				fmt.Printf("The program will close in %d seconds.\n", i)
				time.Sleep(1 * time.Second)
				ClearT()
			}
			PrintASCIIM(asciiMap, message_wp)
			fmt.Println("You lost. The word was:", data.SearchedW)
			fmt.Println("Press Enter to close the program.")
			fmt.Scanln()
		}
	}
	if data.Word == data.SearchedW { // Si le joueur devine le mot avant d'épuiser ses tentatives, il gagne.
		ClearT()
		for i := 5; i > 0; i-- {
			PrintASCIIM(asciiMap, message_gg)
			fmt.Println("You found the word:", data.SearchedW)
			fmt.Printf("The program will close in %d seconds.\n", i)
			time.Sleep(1 * time.Second)
			ClearT()
		}
		PrintASCIIM(asciiMap, message_gg)
		fmt.Println("You found the word:", data.SearchedW)
		fmt.Println("Press Enter to close the program.")
		fmt.Scanln()
	}
	ClearT()
	for i := 2; i < 0; i-- {
		time.Sleep(1 * time.Second)
	}
	ClearT()
	if load {
		RmFile(filename) // Supprime le fichier de sauvegarde à la fin, si il y en a un.
	}
}
