package handlers

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"twilioIvr/models"
	"twilioIvr/utils"
)

// MainMenuHandler handles the initial request and presents the main menu to the caller
func MainMenuHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new Gather object to collect a single digit input from the caller
	gather := &models.Gather{
		Action:    "/handle_choice", // The URL to send the gathered input to
		Method:    "POST",           // Use POST method to send the input
		NumDigits: 1,                // Collect only one digit
		Say: &models.Say{ // Say object to provide instructions to the caller
			Text: "Welcome to the California Weather Tower Report. Press a digit to get the weather report for a specific city: " +
				"1 for Los Angeles, 2 for San Diego, 3 for San Jose, 4 for San Francisco, 5 for Fresno, 6 for Sacramento, " +
				"7 for Long Beach, 8 for Oakland, 9 for Bakersfield, 0 for Anaheim.",
		},
	}

	// Create a new TwimlResponse object containing the Gather object
	response := &models.TwimlResponse{
		Gather: gather,
	}

	// Marshal the TwimlResponse object to XML format
	res, err := xml.Marshal(response)
	if err != nil {
		// If there is an error during marshaling, return an internal server error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type of the response to XML
	w.Header().Set("Content-Type", "application/xml")
	// Write the XML response to the ResponseWriter
	w.Write(res)
}

// HandleChoiceHandler handles the input choice from the user and returns the appropriate weather report
func HandleChoiceHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form values from the request
	err := r.ParseForm()
	if err != nil {
		// If there is an error during parsing, return an internal server error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the digit pressed by the user from the form values
	digit := r.FormValue("Digits")
	var response *models.TwimlResponse

	// Switch case to handle different digit inputs
	switch digit {
	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
		// Convert the digit character to an integer index
		digitIndex := int(digit[0] - '0')
		if digitIndex == 0 {
			digitIndex = 10 // Map '0' to 10 for Anaheim
		}

		// Retrieve the city name from a predefined list using the digit index
		city := utils.Cities[digitIndex-1]
		// Get a random weather report
		weatherReport := utils.GetRandomWeatherReport()
		// Format the full weather report message
		fullReport := fmt.Sprintf("The weather in %s: %s", city, weatherReport)

		// Create a TwimlResponse object with the weather report message
		response = &models.TwimlResponse{
			Say: &models.Say{Text: fullReport, Voice: "woman", Language: "en-US"},
		}
	default:
		// If the input is not a valid digit, create a TwimlResponse object with an error message
		response = &models.TwimlResponse{
			Say: &models.Say{Text: "Sorry, I didn't catch that. Please try again.", Voice: "woman", Language: "en-US"},
		}
	}

	// Marshal the TwimlResponse object to XML format
	res, err := xml.Marshal(response)
	if err != nil {
		// If there is an error during marshaling, return an internal server error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type of the response to XML
	w.Header().Set("Content-Type", "application/xml")
	// Write the XML response to the ResponseWriter
	w.Write(res)
}
