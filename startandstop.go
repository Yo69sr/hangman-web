package hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// ================================================================================================== //
// Ce fichier contient toutes les fonctions utilisés pour les sauvegardes et chargements d'une partie //
// ================================================================================================== //

func Save(data *DataHang, filename string) error { // Save enregistre les données de type DataHang dans un fichier au format JSON
	SLData, err := json.Marshal(data) // Encode les données en format JSON
	if err != nil {                   // Retourne une erreur si l'encodage JSON échoue
		return fmt.Errorf("error while encoding JSON: %v", err)
	}
	err = ioutil.WriteFile(filename, SLData, 0644) // Écrit les données JSON dans le fichier spécifié avec les permissions 0644
	if err != nil {                                // Retourne une erreur si l'écriture du fichier échoue
		return fmt.Errorf("error while writing the file: %v", err)
	}
	ClearT()
	return nil // Retourne nil si la sauvegarde a réussi sans erreurs
}

func Load(filename string) (*DataHang, error) { // Load charge les données du jeu à partir d'un fichier spécifié. Elle retourne un pointeur vers une structure DataHang et une erreur éventuelle.
	SLData, err := ioutil.ReadFile(filename) // Lit le contenu du fichier spécifié par filename
	if err != nil {
		return nil, fmt.Errorf("error while reading the file: %v", err) // Retourne une erreur si la lecture échoue
	}
	var data DataHang
	err = json.Unmarshal(SLData, &data) // Décode le contenu JSON dans la structure data crée la ligne juste avant
	if err != nil {
		return nil, fmt.Errorf("error while decoding JSON: %v", err) // Retourne une erreur si le décodage échoue
	}
	for i := 5; i > 0; i-- {
		ClearT()
		fmt.Println("Game loaded from", filename)
		fmt.Printf("The game will start in %d seconds.", i)
		time.Sleep(1 * time.Second)
	}
	return &data, nil // Retourne un pointeur vers la structure data et aucune erreur
}
