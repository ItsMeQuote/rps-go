package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func printDivider() {
	fmt.Println(strings.Repeat("-", 25))
}

func printResult(status string, botChoice string, playerChoice string, score map[string]int, bot string, player string) {
	fmt.Println(status+"\nБот выбрал:", botChoice, "| Вы выбрали:", playerChoice, "\nСчёт бота:", score[bot], "| Ваш счёт:", score[player])
}

func main() {
	fmt.Println("Добро пожаловать в игру «Камень, ножницы, бумага»!")

	beats := map[string]string{
		"камень":  "ножницы",
		"ножницы": "бумага",
		"бумага":  "камень",
	}
	choices := []string{
		"камень",
		"ножницы",
		"бумага",
	}
	shorts := map[string]string{
		"к": "камень",
		"н": "ножницы",
		"б": "бумага",
	}
	score := map[string]int{
		"player": 0,
		"bot":    0,
	}

	for {
		var playerChoice string
		fmt.Print("Выберите: камень, ножницы или бумага: ")
		fmt.Scanln(&playerChoice)

		printDivider()

		botChoice := choices[rand.Intn(len(choices))]

		if fullChoice, ok := shorts[playerChoice]; ok {
			playerChoice = fullChoice
		}

		if _, ok := beats[playerChoice]; !ok {
			fmt.Println("Нужно ввести: камень/ножницы/бумага или первые буквы в этих словах.")
			printDivider()
			continue
		}

		if playerChoice == botChoice {
			printResult("🤝 Ничья!", botChoice, playerChoice, score, "bot", "player")
			printDivider()
		} else if beats[playerChoice] == botChoice {
			printResult("🎉 Вы выиграли!", botChoice, playerChoice, score, "bot", "player")
			score["player"] += 1
			printDivider()
		} else {
			printResult("❌ Вы проиграли!", botChoice, playerChoice, score, "bot", "player")
			score["bot"] += 1
			printDivider()
		}
	}
}
