package hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ================================================================================================== //
// Ce fichier contient toutes les fonctions utilisés pour les sauvegardes et chargements d'une partie //
// ================================================================================================== //

func GetOneF(filename string) string { // Fonction pour obtenir un nom de fichier unique si le fichier existe déjà
	ext := filepath.Ext(filename)             // Extrait l'extension du fichier (par exemple, ".txt")
	name := strings.TrimSuffix(filename, ext) // Supprime l'extension du nom de fichier
	uniqueFilename := filename                // Initialise le nom de fichier unique avec le nom fourni
	i := 1                                    // Compteur pour le suffixe
	for {                                     // Incrémente le suffixe jusqu'à ce que le fichier n'existe pas
		if _, err := os.Stat(uniqueFilename); os.IsNotExist(err) {
			break // Si le fichier n'existe pas, sort de la boucle
		}
		uniqueFilename = fmt.Sprintf("%s%d%s", name, i, ext) // Génère un nouveau nom de fichier avec un suffixe
		i++                                                  // Incrémente le suffixe
	}
	return uniqueFilename // Retourne le nom de fichier unique
}

func Save(data *DataHang) error { // Fonction Save qui gère la sauvegarde des données du jeu
	reader := bufio.NewReader(os.Stdin) // Crée un lecteur pour les entrées utilisateur
	ClearT()
	var response string
	for { // Demande à l'utilisateur s'il souhaite sauvegarder la partie
		fmt.Print("Do you want to save the game? (y/n): ")
		response, _ = reader.ReadString('\n')   // Lit la réponse de l'utilisateur
		response = strings.TrimSpace(response)  // Supprime les espaces inutiles autour de la réponse
		if response == "y" || response == "n" { // Vérifie si la réponse est valide
			break // Sort de la boucle si la réponse est "y" ou "n"
		}
		ClearT()
		fmt.Println("Please enter 'y' for yes or 'n' for no.")
	}
	if response == "n" { // Si l'utilisateur choisit de ne pas sauvegarder, affiche un message d'annulation
		fmt.Println("Save canceled.")
		time.Sleep(1 * time.Second)
		ClearT()
		return nil // Quitte la fonction sans effectuer de sauvegarde
	}
	fmt.Print("Enter the name of the save file (without extension): ")
	filename, _ := reader.ReadString('\n')          // Lit le nom du fichier
	filename = strings.TrimSpace(filename) + ".txt" // Supprime les espaces et ajoute l'extension .txt
	filename = GetOneF(filename)                    // Vérifie si le fichier existe déjà et obtient un nom unique si nécessaire
	SLData, err := json.Marshal(data)               // Convertit les données en format JSON
	if err != nil {
		return fmt.Errorf("error while encoding JSON: %v", err) // Retourne une erreur si l'encodage échoue
	}
	err = ioutil.WriteFile(filename, SLData, 0644) // Écrit les données JSON dans le fichier avec les permissions 0644
	if err != nil {
		return fmt.Errorf("error while writing the file: %v", err) // Retourne une erreur si l'écriture échoue
	}
	for i := 2; i > 0; i-- { // Confirmation que la sauvegarde a réussi
		fmt.Printf("Game successfully saved as '%s'\n", filename)
		time.Sleep(1 * time.Second)
		ClearT()
	}
	return nil // Retourne nil pour indiquer que la fonction s'est terminée sans erreur
}

func Load(filename string) (*DataHang, string, error) { // Load charge les données du jeu à partir d'un fichier spécifié. Elle retourne un pointeur vers une structure DataHang et une erreur éventuelle.
	SLData, err := ioutil.ReadFile(filename) // Lit le contenu du fichier spécifié par filename
	if err != nil {
		return nil, filename, fmt.Errorf("error while reading the file: %v", err) // Retourne une erreur si la lecture échoue
	}
	var data DataHang
	err = json.Unmarshal(SLData, &data) // Décode le contenu JSON dans la structure data crée la ligne juste avant
	if err != nil {
		return nil, filename, fmt.Errorf("error while decoding JSON: %v", err) // Retourne une erreur si le décodage échoue
	}

	return &data, filename, nil // Retourne un pointeur vers la structure data et aucune erreur
}
