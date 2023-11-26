package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Game struct {
	Letters      []string // nul
	FoundLetters []string
	UsedLetters  []string
	Word         string
	TurnsLeft    int
	IsGame       bool
	IsWon        bool
}

var Player Game

func New(turns int, word string) (*Game, error) {
	if len(word) < 2 {
		return nil, fmt.Errorf("Le mot '%s' doit faire minimum 2 charactères . got=%v", word, len(word))
	}

	letters := strings.Split(strings.ToLower(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		IsWon:        false,
		IsGame:       true,
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
		Word:         strings.ToLower(word),
	}

	return g, nil
}

func HasWon(hiddenWord []string, word string) bool {
	if len(hiddenWord) == len(word) {
		for i := range hiddenWord {
			if string(hiddenWord[i]) != string(word[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

/*func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func (g *Game) LoseTurn(guess string) {
	g.TurnsLeft--
	g.UsedLetters = append(g.UsedLetters, guess)
}

func LetterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
} */

var words = make([]string, 0, 50)

func PrepareFileName(level string) string {
	switch level {
	case "1":
		return "champions.txt"
	case "2":
		return "items.txt"
	case "3":
		return "spells.txt"
	default:
		return "champions.txt"
	}
}

func Load(filename string) error { //charge le .txt avec les noms de champions
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func PickWord() string { //prend un nom de champion aléatoire
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(words))
	return words[i]
}

/*var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("Quelle est votre lettre? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		guess = strings.TrimSpace(guess)

		if len(guess) != 1 {
			fmt.Println("Format invalide.", guess, len(guess))
			continue
		}
		valid = true
	}
	return
} */

func Start(level string) {
	err := Load(PrepareFileName(level))
	if err != nil {
		fmt.Printf("Could not load dictionary: %v\n", err)
		os.Exit(1)
	}

	g, err := New(10, PickWord())
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}
	Player = *g
	fmt.Println("game init word : ", g.Word)
}

func (g *Game) CheckInput(value string) string {
	value = strings.ToLower(value)
	if len(value) != 1 {
		if g.Word == value {
			for index := range g.Word {
				g.FoundLetters[index] = string(g.Word[index])
			}
			return "Vous avez trouvé le mot"
		} else {
			g.TurnsLeft -= 2
			return "Vous avez perdu deux vies"
		}
	} else {

		for _, letter := range g.UsedLetters {
			if letter == value {
				return "Vous avez deja indiqué cette lettre"
			}
		}
		g.UsedLetters = append(g.UsedLetters, value)
		IsFind := false
		for i, v := range g.Word {
			if value == string(v) {
				IsFind = true
				g.FoundLetters[i] = string(v)
			}
		}
		if !IsFind {
			g.TurnsLeft -= 1
			return "Vous avez perdu une vie"
		}
		return "Vous avez trouvé la lettre"
	}
}
