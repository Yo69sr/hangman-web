# Hangman Game

[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/1YLV-els)

## Description

This project is a classic implementation of the Hangman game in Go (Golang). The game randomly selects a word from a provided file, and the player has a limited number of attempts to guess the word by suggesting letters. The game includes ASCII art of a hanging man, enhancing the visual experience.

## Table of Contents

- [Features](#features)
- [Game Rules](#game-rules)
- [Packages Used](#packages-used)
- [Installation](#installation)
- []()
- [Usage](#usage)


## Features

- **Difficulty Selection**: Players can choose from three difficulty levels: easy, medium, and hard.
- **Save and Load**: Games can be saved to a file and loaded later.
- **ASCII Art**: Displays ASCII illustrations for the Hangman and messages for welcome, victory, or defeat.
- **Letter Revelation**: A portion of the letters in the word can be randomly revealed at the start of the game.
- **Debug Mode**: Option to display the answer during the game for development purposes.

## Game Rules

1. **Objective**: Guess the hidden word before running out of allowed attempts.
2. **Attempts**: Players have 10 attempts to guess the letters of the word.
3. **Proposals**: Players can propose either a letter or a complete word.
4. **Already Used Letter**: If a letter or a word has already been proposed, the player is warned and does not lose an attempt.
5. **Incorrect Word**: If the proposed word is incorrect, the player loses 2 attempts.
6. **Ending condition**: The game ends when the word is correctly guessed or when the number of attempts reaches zero.

## Packages Used

- `bufio`: For reading user input from the console.
- `encoding/json`: For serializing and deserializing game data during save and load.
- `flag`: For managing command-line arguments, like the save file.
- `fmt`: For formatted text output in the console.
- `io/ioutil`: For reading files containing words and save data.
- `math/rand`: For randomly selecting words and shuffling indices.
- `os`: For managing system operations, like clearing the screen and exiting the program.
- `os/exec`: For executing system commands, like clearing the screen.
- `strings`: For manipulating strings.
- `time`: For managing timing and delays.
- `unicode`: For validating that inputs are letters.

## Installation

To run this project, ensure you have Go installed on your system. Clone this repository and navigate to the project directory:

```bash
git clone https://github.com/noshagit/hangman-classic.git
cd hangman-classic
```

## Structuration du fichier

### Fichier Golang 

- **main.go**: main.go contient Le main de mon jeu, la fonction qui sera lu par mon ordinateur.
- **annexe.go**: annexe.go contient toutes les fonctions annexe que j'ai crée sans qu'elles soient forcément demandé. Je les ai crée pour mon confort personnel et surtout pour un confort visuel.
- **asciiart.go**: asciiart.go contient toutes les fonctions permettant l'affichage de José ( le Pendu ) ou d'un message en ASCII art.
- **hangman.go**: hangman.go contient La fonction principal de mon jeu ainsi que la création du type de structure que j'ai utilisé.
- **init.go**: init.go permet de remplir toutes les variables que j'ai créer dans mon hangman.go.
- **input.go**: input.go contient toutes les fonctions gérant les entrées utilisateurs.
- **startandstop.go**: startandstop.go contient toutes les fonctions utilisés pour les sauvegardes et chargements d'une partie.
- **verif.go**: verif.go a comme objectif de contenir toutes mes fonctions de vérification et de renvoi de booléen.

### Fichier .txt

- **words.txt**: words.txt regroupe l'ensemble des mots pour la difficulté "easy".
- **words2.txt**: words2.txt regroupe l'ensemble des mots pour la difficulté "medium".
- **words3.txt**: words3.txt regroupe l'ensemble des mots pour la difficulté "hard".
- **hangman.txt**: hangman.txt contient tous les dessins/position en ASCII art de José, le pendu.
- **standard.txt**: standard.txt lui contiendra tous les dessins en ASCII art excepté José, il y aura quelques ponctuations, les lettres en majuscule et minuscule ainsi que les chiffres.
- **save.txt**: Fichier qui sera crée lors d'une sauvegarde de jeu, et supprimé à la fin d'une partie contenant toutes les informations stockées dans une structure.

### Les autres fichiers

- **go.mod**: go.mod définis un nom de module "hangman" permettant d'appeler les fonctions contenu dans le ce module dans d'autre fichier ne començant pas par "package hangman".

## Usage

Now, if you're ready to play:

```bash
go run main/main.go
```

if you want to load from a previous game use this command:

```bash
go run main/main.go --startWith save.txt
```

Good luck and have fun playing Hangman!
