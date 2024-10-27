package hangman

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ================================================================================================================= //
// Ce fichier contient toutes les fonctions permettant l'affichage de José ( le Pendu ) ou d'un message en ASCII art //
// ================================================================================================================= //

func ChHangman(filename string) ([10]string, error) { // ChHangman lit les données à partir d'un fichier spécifié et retourne un tableau de chaînes représentant les positions d'un pendu.
	data, err := ioutil.ReadFile(filename) // Lit le contenu du fichier spécifié par filename
	if err != nil {
		return [10]string{}, err // Retourne un tableau vide et l'erreur si la lecture échoue
	}
	positions := strings.Split(string(data), "\n\n") // Divise le contenu en segments basés sur les double sauts de ligne
	var positionsPendu [10]string
	copy(positionsPendu[:], positions) // Copie les positions lues dans le tableau positionsPendu
	return positionsPendu, nil         // Retourne le tableau de positions et nil pour l'erreur
}

func AffP(data DataHang) { // AffP affiche la représentation graphique du pendu en fonction du nombre de tentatives restantes (Pv).
	if data.Pv > 0 && data.Pv <= len(data.PosHang) { // Vérifie si le nombre de tentatives restantes est positif et ne dépasse pas le nombre de positions disponibles.
		fmt.Println(data.PosHang[10-data.Pv]) // Affiche la position correspondante du pendu en utilisant l'index basé sur les tentatives restantes.
	}
}

func AffPF(data *DataHang) { // AffPF affiche la représentation finale du pendu lorsque le joueur a perdu.
	fmt.Println(data.PosHang[9]) // Affiche la position du pendu correspondant à une perte (décès) du joueur.
}

func ChASCIIArt(filePath string, chars []string) (map[string]string, error) { //ChASCIIArt lit un fichier contenant de l'art ASCII et l'associe à des caractères spécifiés. Elle retourne une carte associant chaque caractère à sa représentation ASCII correspondante.
	file, err := os.Open(filePath) // Ouvre le fichier spécifié par filePath
	if err != nil {
		return nil, err // Retourne une erreur si l'ouverture échoue
	}
	defer file.Close()                  // Assure la fermeture du fichier à la fin de la fonction
	asciiMap := make(map[string]string) // Initialise une map pour stocker les caractères et leurs représentations ASCII
	scanner := bufio.NewScanner(file)   // Crée un scanner pour lire le fichier
	var asciiArt []string
	currentIndex := 0
	for scanner.Scan() { // Parcourt chaque ligne du fichier
		line := scanner.Text()             // Lit la ligne actuelle
		if strings.TrimSpace(line) == "" { // Vérifie si la ligne est vide
			if len(asciiArt) > 0 { // Si on a déjà collecté des lignes d'art ASCII
				if currentIndex < len(chars) {
					asciiMap[chars[currentIndex]] = strings.Join(asciiArt, "\n") // Ajoute l'art ASCII à la map pour le caractère actuel
					currentIndex++                                               // Passe au caractère suivant
				}
				asciiArt = []string{} // Réinitialise le slice pour la prochaine entrée
			}
		} else {
			asciiArt = append(asciiArt, line) // Ajoute la ligne d'art ASCII au slice
		}
	}
	if len(asciiArt) > 0 && currentIndex < len(chars) { // Ajoute le dernier art ASCII si le fichier ne se termine pas par une ligne vide
		asciiMap[chars[currentIndex]] = strings.Join(asciiArt, "\n")
	}
	if err := scanner.Err(); err != nil { // Vérifie s'il y a eu une erreur lors de la lecture du fichier
		return nil, err // Retourne une erreur si une erreur de lecture a eu lieu
	}
	return asciiMap, nil // Retourne la map d'art ASCII et aucune erreur
}

func PrintASCIIM(asciiMap map[string]string, message string) { // PrintASCIIM affiche un message sous forme d'art ASCII en utilisant un mappage
	var lines = make([]string, 7)
	for _, char := range message { // Parcourt chaque caractère du message
		asciiChar, found := asciiMap[string(char)] // Récupère la représentation ASCII du caractère dans le mappage
		if found {                                 // Si le caractère est trouvé
			asciiLines := strings.Split(asciiChar, "\n") // divise la représentation ASCII en lignes
			for i := 0; i < len(asciiLines); i++ {       // Ajoute chaque ligne de l'ASCII à la ligne correspondante dans le tableau lines et ajoute un espace après chaque ligne pour la séparation
				lines[i] += asciiLines[i] + " " //
			}
		} else { // Si le caractère n'est pas trouvé, ajoute des espaces vides pour maintenir l'alignement
			for i := 0; i < 7; i++ {
				lines[i] += "        "
			}
		}
	}
	for _, line := range lines { // Affiche chaque ligne de l'art ASCII résultant
		fmt.Println(line)
	}
}
