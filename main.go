package main

import (
	"fmt"
	"log"
	"os"
//	"strings"
	"github.com/FaTerminator/chatbot"
	"github.com/FaTerminator/fitify"
)

// Autoload environment variables in .env
import _ "github.com/joho/godotenv/autoload"

func chatbotProcess(session chatbot.Session, message string) (string, error) {
	resp := fitify.CaseMatch(message)

	return fmt.Sprintf("%s", resp), nil
}

func main() {
	//Uncomment the following lines to customize the chatbot
	fitify.GetMuscles()
	fitify.GetEquipments()
	fitify.GetImages()
	fitify.GetExercises()

	chatbot.WelcomeMessage = "Welcome to Fitify.\nYou can ask me about:\n-Which exercises train a certain muscle\n-What exercises can be done using certain equipment"
	chatbot.ProcessFunc(chatbotProcess)

	//Use the PORT environment variable
	port := os.Getenv("PORT")
	// Default to 3000 if no PORT environment variable was defined
	if port == "" {
		port = "8080"
	}

	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(chatbot.Engage(":" + port))
}
