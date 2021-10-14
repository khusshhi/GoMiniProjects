package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		fmt.Println(">ME : ")
		var chatbot string
		fmt.Scanln(&chatbot)
		//chatbot = strings.ToLower(chatbot)
		chatbot = strings.TrimSpace(chatbot)
		chatMap := map[string]string{
			"Hello":             "Hi, how are you?",
			"What's your name?": "I usually go bye Happybot101",
			"I'm fine":          "That's nice",
			"How are you":       "I'm Good!",
			"Thanks":            "Your welcome",
			"Dou you like it?":  "Yes i loved it",
			"Sorry":             "No problem",
			"Bye":               "Take care",
			"All the best":      "Thankyou",
			"Please":            "Sure, I'll do it",
			"quit":              "See you soon!",
			"help":              "Allowed commands : \nHello, What's your name?, I'm fine, How are you, Thanks, bye",
		}
		fmt.Println("MYBOT : ")
		if chatMap[chatbot] != "" {
			fmt.Println(chatMap[chatbot])
		} else {
			fmt.Println("Use 'help' command for allowed keywords")
		}
		if strings.Contains(chatbot, "Bye") || strings.Contains(chatbot, "quit") {
			break
		}
	}
}
