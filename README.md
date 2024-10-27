# Hangman Game

[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/1YLV-els)

## Description

This project is a classic implementation of the Hangman game in Go (Golang). The game randomly selects a word from a provided file, and the player has a limited number of attempts to guess the word by suggesting letters. The game includes ASCII art of a hanging man, enhancing the visual experience.

## Table of Contents

- [Features](#features)
- [Game Rules](#game-rules)
- [Packages Used](#packages-used)
- [Installation](#installation)
- [Project Structure](#project-structure)
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

## Project Structure

### Go Files

- **main.go**: Contains the main function of the game, which will be executed by the computer.
- **annexe.go**: Includes auxiliary functions created for personal convenience and code readability, even if they are not strictly necessary.
- **asciiart.go**: Contains all functions for displaying José (the Hangman) or other messages in ASCII art.
- **hangman.go**: Houses the main logic of the game as well as the definition of the data structures used.
- **init.go**: Used to initialize all the variables defined in `hangman.go`.
- **input.go**: Groups all functions that handle user inputs.
- **startandstop.go**: Contains functions used for saving and loading a game.
- **verif.go**: Intended to hold all verification functions and return booleans.

### .txt Files

- **words.txt**: Contains the words for the "easy" difficulty level.
- **words2.txt**: Contains the words for the "medium" difficulty level.
- **words3.txt**: Contains the words for the "hard" difficulty level.
- **hangman.txt**: Contains ASCII art drawings/positions of José, the Hangman.
- **standard.txt**: Includes various ASCII art drawings, featuring some punctuation, uppercase and lowercase letters, as well as numbers.
- **save.txt**: A file created during game saving, which will be deleted at the end of the game. It contains all information stored in a structure.

### Other Files

- **go.mod**: Defines the module name as "hangman," allowing the functions within this module to be called from other files not starting with "package hangman."

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
