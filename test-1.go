package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// bioHandler handles the root route
func bioHandler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Bio</title>
				<style>
					body {
						font-family: Arial, sans-serif;
						font-size: 16px;
						line-height: 1.5;
						background-color: #f7f7f7;
						text-align: center;
						padding: 2em;
					}
					h1 {
						font-size: 2.5em;
						margin-bottom: 1em;
					}
				</style>
			</head>
			<body>
			<h1>Hello! My name is Alex Peraza, I am 20 years old. 
			I love to code, get hands on experience in hardware and spend time with friends. 
			My hobbies are: Watching series and playing sports.</h1>			
			</body>
		</html>
	`
	fmt.Fprint(w, html)
} 

// getGreeting returns a greeting based on the time of day
func getGreeting(t time.Time) string {
	switch t.Weekday() {
	case time.Monday:
		return "Monday!"
	case time.Tuesday:
		return "Tuesday!"
	case time.Wednesday:
		return "Wednesday!"
	case time.Thursday:
		return "Thursday!"
	case time.Friday:
		return "Friday!"
	case time.Saturday:
		return "Saturday!"
	case time.Sunday:
		return "Sunday!"
	default:
		return ""
	}
}

// greetingHandler handles the /greeting route
func greetingHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	greeting := getGreeting(now)
	html := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Greeting</title>
				<style>
					body {
						font-family: Arial, sans-serif;
						font-size: 16px;
						line-height: 1.5;
						background-color: #f7f7f7;
						text-align: center;
						padding: 2em;
					}
					h1 {
						font-size: 2.5em;
						margin-bottom: 1em;
					}
				</style>
			</head>
			<body>
				<h1>Right now is {{time}}</h1>
				<p>Enjoy the rest of your {{greeting}}.</p>
			</body>
		</html>
	`
	html = strings.ReplaceAll(html, "{{time}}", now.Format("3:04pm"))
	html = strings.ReplaceAll(html, "{{greeting}}", greeting)
	fmt.Fprint(w, html)
}
// randomHandler handles the /random route
func randomHandler(w http.ResponseWriter, r *http.Request) {
	quotes := map[int]string{
		1:  "Believe in yourself and all that you are. Know that there is something inside you that is greater than any obstacle.",
		2:  "Success is not final, failure is not fatal: it is the courage to continue that counts.",
		3:  "Your only limit is the amount of willingness you have to succeed.",
		4:  "Don't let yesterday take up too much of today.",
		5:  "The greatest glory in living lies not in never falling, but in rising every time we fall.",
		6:  "The best way to predict the future is to invent it. - Alan Kay",
		7:  "The technology you use impresses no one. The experience you create with it is everything. - Sean Gerety",
		8:  "The digital revolution is far more significant than the invention of writing or even of printing. - Douglas Engelbart",
		9:  "Technology is a tool, and like any tool, its value depends on how it is used. - David Crystal",
		10: "Innovation distinguishes between a leader and a follower. - Steve Jobs",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes)) + 1

	html := `
        <!DOCTYPE html>
        <html>
            <head>
                <title>Random Quote</title>
                <style>
                    body {
                        font-family: Arial, sans-serif;
                        font-size: 16px;
                        line-height: 1.5;
                        background-color: #f7f7f7;
                        text-align: center;
                    }
                    h1 {
                        margin-top: 50px;
                        font-size: 28px;
                    }
                    p {
                        margin-top: 20px;
                        font-size: 20px;
                    }
                </style>
            </head>
            <body>
                <h1>Random Quote</h1>
                <p>` + quotes[randomIndex] + `</p>
            </body>
        </html>
    `
	fmt.Fprintf(w, html)
}

func main() {
	// Create a new multiplexer (map)
	mux := http.NewServeMux()

	// Register handlers for each route
	mux.HandleFunc("/", bioHandler)
	mux.HandleFunc("/greeting", greetingHandler)
	mux.HandleFunc("/random", randomHandler)

	// Create a new server and listen on port 4000
	fmt.Println("Starting server on port 4000...")
	http.ListenAndServe(":4000", mux)
}

