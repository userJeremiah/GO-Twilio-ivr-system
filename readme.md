# twilio ivr doc

this walks through the setup and running of this application.



## needs to have
- Golang installed on machine to compile and run the application
- ngrok installed on machine to expose the port
- a source of stable internet connection throughout the process
- also good to add that the internet source should be seperate from the device used for making calls
- n/b: the phone number used for calling the twilio number has to be the same as the number you used in registering the twilio account.

## clone the application or open the application

you're finding yourself in the root directory that has 3 layers or files: the app folder, the `quick.go` file and this `readme.md` file. 

- run quick.go
- run ngrok (`ngrok http 3000`)
- copy the ngrok url(eg `https://cd2b-105-112-125-164.ngrok-free.app`) and paste it into the twilio console and append with `/twiml`(eg. `https://cd2b-105-112-125-164.ngrok-free.app/twiml`)
- call your twilio number

this is the first demo to check if your twilio is configured properly and if the application runs fine.

## main application

- cd `app`
- run `go run main.go`
- run ngrok (`ngrok http 8083`)
- copy the ngrok url(eg `https://cd2b-105-112-125-164.ngrok-free.app`) and paste it into the twilio console and append with `/main_menu`(eg. `https://cd2b-105-112-125-164.ngrok-free.app/main_menu`)
- call your twilio number


## code 

i'll go ahead to explain or give a walkthrough of the codebase, i'll be focusing on what goes down in the `app` folder. it contains the `handlers`, `models`, and `utils` directories and the `main.go` file. 
no frameworks, just vanilla golang.

the `main.go` handles the routes(too few to have its own package and i didnt want to deal with libraries), which calls the handlers as well cause that's how the `http` package works.
- handled what's going on in `main.go`. check comments

let's head over to the `handlers` directory and explain that package which sheds more light on the rest of the app.

### Functionality:

- This function is designed to be a handler for an HTTP request, likely triggered by a user visiting a specific URL.
- It creates a TwiML (Twilio Markup Language) response that prompts the user to select a city for a weather report.

#### Breakdown:


- starting out with the `models` package/directory

This code defines three structs used for building TwiML (Twilio Markup Language) responses in your Go application. Here's a detailed explanation of each struct and its fields:

1. `TwimlResponse`:

This struct represents the main container for a TwiML response.
It has an XMLName field set to "Response" which is the root element of the TwiML document.
It includes two optional fields: Gather and Say.
Gather: Points to a Gather struct instance if the response involves collecting user input via a keypad.
Say: Points to a Say struct instance if the response involves playing a message to the user.

2. `Gather`:

This struct defines the properties for a TwiML element that collects user input via a keypad.
It has an XMLName field set to "Gather" which is the element name within the TwiML document.
It has several attributes to configure the user input collection:
- Action: Specifies the URL path where the user's digit selection will be submitted (mandatory).
- Method: Sets the HTTP method for submitting the data (typically "POST").
NumDigits: Limits the user input to a specific number of digits (optional).
- Timeout (optional): Sets the maximum wait time in seconds for user input before moving on.
- Say: Points to a Say struct instance defining the message played to the user before collecting input (optional).

3. `Say`:

This struct defines the properties for a TwiML element that plays a message to the user.
It has an XMLName field set to "Say" which is the element name within the TwiML document.
The message content is defined by the Text field.
It has optional attributes for customizing the voice and language of the message:
Voice: Specifies a voice name from Twilio's available options (e.g., "Polly.Amy").
Language: Sets the language for the message (e.g., "en-US" for English).

`models.Gather` struct:

This code assumes the existence of a models package containing helper structs.
It defines a models.Gather struct with properties relevant to collecting user input.
Action: Specifies the URL path (/handle_choice) where the user's digit selection will be submitted.
Method: Sets the HTTP method for submitting the data (POST).
NumDigits: Limits the user input to a single digit (1-9 or 0).
Say: Defines the message played to the user using a nested models.Say struct.
Text: Contains the actual message prompting the user to choose a city by pressing a corresponding digit (1-9 or 0). The message lists all available city options.

- moving onto the `handlers` package/directory

This code defines a handler function for handling incoming calls to the application. 

- It starts off with the `MainMenuHandler` function.
This code assumes the existence of a `models.TwimlResponse `struct likely used for building TwiML responses.
It creates an instance of this struct and populates the Gather field with the previously defined gather object.

XML Marshalling and Response:
The code marshals the models.TwimlResponse struct into XML format using xml.Marshal.
If an error occurs during marshalling, it writes an error message to the response with an internal server error status code (500).
If successful, it sets the response content type to "application/xml" and writes the marshaled XML data to the response body.

- next is the `HandleChoiceHandler` function.
Breakdown:

Parsing Form Values:
The function starts by parsing the form data submitted by the user in the HTTP request using `r.ParseForm()`.
If there's an error during parsing, it returns an internal server error response (http.StatusInternalServerError) with the error message.

Retrieving User Input:
It retrieves the digit pressed by the user using the `r.FormValue`("Digits") method.
The value is stored in the digit variable.

Handling User Choice:
A switch statement handles different user input digit values.
It checks if the digit is between "1" and "0" (inclusive).
If a valid digit is entered:
It converts the digit character to an integer index by subtracting '0' from its ASCII code.
It adjusts the index for '0' as it's mapped to position 10 (Anaheim) in the city list.
It retrieves the corresponding city name from a predefined list in utils.Cities using the adjusted index.
It gets a random weather report using `utils.GetRandomWeatherReport`.
It formats a complete weather report message combining the city and report.
Finally, it creates a `models.TwimlResponse` object with a Say element containing the formatted message. Voice and language are set to "woman" and "en-US" respectively.
If an invalid digit is entered (anything other than 1-0), it creates a `models.TwimlResponse` object with an error message using the same voice and language settings.

Marshalling and Response:
The code marshals the constructed `models.TwimlResponse` object into XML format using `xml.Marshal`.
If there's an error during marshalling, it returns an internal server error response with the error message.
On successful marshalling, it sets the response content type to "application/xml".
Finally, it writes the marshaled XML data to the response body using w.Write.


- Lastly, the `mock.db.go` in the `utils` package acts as a database of weather reports. Additionally, the `GetRandomWeatherReport()` function helps the application to feel more dynamic as the weather reports switches up with every request. #   G O - T w i l i o - i v r - s y s t e m 
 
 #   G O - T w i l i o - i v r - s y s t e m 
 
 #   G O - T w i l i o - i v r - s y s t e m 
 
 
