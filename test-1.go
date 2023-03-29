package main

import (
	"fmt"
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