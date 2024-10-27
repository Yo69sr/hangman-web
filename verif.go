package hangman

import "unicode"

// ================================================================================//
// Ce fichier contient toutes mes fonctions de vérification et de renvoi de booléen //
// ================================================================================ //

func Cont(list []string, item string) bool { // Cont vérifie si un élément donné existe dans une liste de chaînes de caractères
	for _, l := range list { // Parcourt chaque élément de la liste
		if l == item { // Compare l'élément actuel avec l'item recherché
			return true // Retourne true si l'élément est trouvé
		}
	}
	return false // Retourne false si l'élément n'est pas trouvé dans la liste
}

func IsMVal(mot string) bool { // IsMVal vérifie si une chaîne de caractères ne contient que des lettres
	for _, char := range mot { // Parcourt chaque caractère de la chaîne
		if !unicode.IsLetter(char) { // Vérifie si le caractère n'est pas une lettre
			return false // Retourne false si un caractère non lettre est trouvé
		}
	}
	return true // Retourne true si tous les caractères sont des lettres
}

func LREver(motRevele, motATrouver, lettre string) bool { // LREver vérifie si une lettre peut être révélée dans le mot à trouver
	for i := range motATrouver { // Parcourt chaque index du mot à trouver
		if string(motATrouver[i]) == lettre && string(motRevele[i]) == "_" { // Vérifie si la lettre correspond à la lettre du mot à trouver et que la position correspondante dans le mot révélé est un underscore
			return false // Retourne false si la lettre peut être révélée
		}
	}
	return true // Retourne true si la lettre ne peut pas être révélée
}
