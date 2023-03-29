package main

import (
	"fmt"
	"net/http"
	
	
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