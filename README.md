
Voici le README adapté avec vos modifications :

Hangman Game
Game made by Yousri Berriche, Bachelor 1 Informatique.

⚙️ ◄| Project instructions |► 🎓

Description
This project is a classic implementation of the Hangman game in Go (Golang). The game randomly selects a word from a provided file, and the player has a limited number of attempts to guess the word by suggesting letters or full words. It also includes HTML/CSS for a modernized interface.

Table of Contents
Features
Game Rules
Packages Used
Installation
Project Structure
Usage
Features
Difficulty Selection: Choose between easy, medium, and hard levels.
ASCII Art: Displays an ASCII Hangman for a retro touch.
Submit Words or Letters: Players can propose either letters or full words. Incorrect words cost 2 lives.
HTML Frontend: Includes a clean web interface with animations, colors, and a modern layout.
History Tracking: Displays a history of letters and words already guessed.
Game Rules
Objective: Guess the hidden word before running out of lives.
Attempts: Players start with 10 lives.
Proposals: Propose a single letter or a complete word.
Incorrect Word: Guessing the wrong word reduces 2 lives.
Win Condition: Reveal all letters of the word or guess the word correctly.
Lose Condition: Lives reach 0 before guessing the word.
Packages Used
html/template: For rendering the game's HTML interface.
fmt: For printing messages to the console.
strings: For managing user inputs and word matching.
net/http: For serving the web-based Hangman game.
time: For handling animations and delays.
math/rand: For randomizing word selection.
Installation
Make sure Go is installed on your system. Clone the repository and navigate to the directory:




cd hangman-classic
Project Structure
Go Files
main.go: Main file that initializes the HTTP server and handles requests.
hangman.go: Core game logic, including word selection and game state management.
HTML & CSS
index.html: Welcome page with difficulty selection buttons.
page-2.html: Game page where players can submit guesses and see their progress.
index.css: Styles the welcome page with vibrant animations and responsive layout.
page-2.css: Styles the game page with a clean interface and animations.
.txt Files
words.txt, words2.txt, words3.txt: Word lists for each difficulty level.
hangman.txt: ASCII art for the Hangman stages.
Usage
Run the project using:



go run main.go
Access the game in your browser at: http://localhost:8080.

Web Interface Details
Welcome Page (index.html)
Displays a title and buttons to select difficulty (Easy, Medium, Hard).
CSS Highlights:
Vibrant gradients and button hover animations.
Center-aligned content for clean design.
Game Page (page-2.html)
Shows the word to guess, remaining lives, and guess history.
Input field to submit letters or words.
CSS Highlights:
Smooth animations for buttons and backgrounds.
Clean typography and responsive layout.