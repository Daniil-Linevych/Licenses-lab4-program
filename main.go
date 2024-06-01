package main

import (
	//"context"
	"bufio"
	"fmt"
	ai "lab4-program-api/cohere_package"
	"os"
	"strings"

	cohere "github.com/cohere-ai/cohere-go/v2"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var chat_history []*cohere.Message
	var response_text string
	firstmessage := "До теми дипломної роботи визначити мету, завдання, предмет, обʼєкт, актуальність, практичне значення, проаналізувати " +
		"аналоги (2-3) та визначити потенційний стек проєкту з аргументацією вибору, відповідність 121 спеціальності." +
		"Документ повинен розкривати в цілому, ідею проекту. Тема: "
	firstIteration := true

	fmt.Print("Enter topic of your nightmare(diplom):0 ")

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if firstIteration {
			message = firstmessage + message
		}

		response_text, chat_history = ai.GetAiResponse(message, chat_history)
		fmt.Println(response_text)

		continue_answer := ContinueWish()
		if !continue_answer {
			break
		}

		fmt.Print("k, what?: ")
	}

}

func ContinueWish() bool {

	reader := bufio.NewReader(os.Stdin)
	var continue_answer string

	for {

		fmt.Print("Wanna squeeze more out of me?(ofc/enough): ")

		continue_answer, _ = reader.ReadString('\n')
		continue_answer = strings.TrimSpace(continue_answer)

		if continue_answer == "ofc" {
			return true
		} else if continue_answer == "enough" {
			return false
		}
	}

}
