package main

import (
	"fmt"
	"net/http"

	"twilioIvr/handlers"
)

func main() {
	//prints to the terminal that the server is running. 
	fmt.Println("server running. . . .. . ... .")

	//route that handles the main_menu function
	http.HandleFunc("/main_menu", handlers.MainMenuHandler)

	//route that handles the choice function
	http.HandleFunc("/handle_choice", handlers.HandleChoiceHandler)

	//we create a port for our application to run @ 8083
	http.ListenAndServe(":8083", nil)
}
