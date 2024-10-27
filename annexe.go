package hangman

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// =============================================================================================================================== //
// Ce fichier contient toutes les fonctions annexe que j'ai crée sans qu'elles soient forcément demandé pour mon confort personnel //
// =============================================================================================================================== //

func RmFile(filename string) { // RmFile supprime le fichier de sauvegarde "save.txt" s'il existe à la fin d'une partie.
	_, err := os.Stat(filename) // Vérifier si le fichier existe en utilisant os.Stat.
	if os.IsNotExist(err) {     // Si le fichier n'existe pas, on sort de la fonction.
		return
	} else if err != nil { // Gérer d'autres erreurs potentielles lors de la vérification.
		fmt.Println("Error while checking the file:", err) // Affiche un message d'erreur.
		return
	}
	err = os.Remove(filename) // Supprimer le fichier en utilisant os.Remove.
	if err != nil {           // Gérer les erreurs lors de la suppression du fichier.
		fmt.Println("Error while deleting the file:", err) // Affiche un message d'erreur.
		return
	}
}

func ClearT() { // ClearT efface le terminal en fonction du système d'exploitation.
	var cmd *exec.Cmd
	switch runtime.GOOS { // Vérifie le système d'exploitation en cours d'exécution.
	case "windows": // Si l'OS est Windows
		if os.Getenv("PSExecutionPolicy") != "" { // Vérifie si l'application est exécutée dans PowerShell
			cmd = exec.Command("powershell", "-Command", "Clear-Host") // Utilise PowerShell pour effacer le terminal.
		} else {
			cmd = exec.Command("cmd.exe", "/c", "cls") // Utilise l'invite de commandes classique pour effacer le terminal.
		}
	default: // Pour les autres systèmes d'exploitation (Linux et macOS)
		cmd = exec.Command("clear") // Utilise la commande "clear" pour effacer le terminal.
	}
	cmd.Stdout = os.Stdout // Redirige la sortie de la commande vers la sortie standard.
	err := cmd.Run()       // Exécute la commande.
	if err != nil {        // Si une erreur se produit lors de l'exécution, elle est affichée à l'utilisateur.
		fmt.Println("error while clearing the terminal:", err)
	}
}
