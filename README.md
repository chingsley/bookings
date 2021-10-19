# BOOKINGS AND RESERVATIONS

A Web Application for making Room reservations in a Bed and Breakfast lodging establishment that offers overnight accommodation and breadkfast

## A Fullstack Application build with Go

This is a fullstack application. I am working on this project with the purpose of gaining in-depth understanding of the Go Programing language

## About the Stack

The application is fullstack built with Golang

## Topics Covered in this project

- The syntax of Go language
- Writing tests in Go
- Creating HTML pages using HTML5
- Writing Javascript for the browser without JQuery
- Cascading Style Sheets
- Building modern, secure web applications
- How to correctly organize the codebase in a Go application
- How to connect to a database in Go
- Managing sessions in Go
- Go middlewares
- 3rd party packages with Go Modules
- How to implement secure user authentication in Go
- Best practices for security in web applications
- How to deploy a Go web application to a live server

## How to test locally

- Clone the repo
- Cd into the project root directro
  > cd bookings
- Instal the required packages with
  > go mod tidy
- Ensure you've created a db named **bookings**
- Use the following command to create and seed the tables
  > ./scripts/db-reset.sh
- To start the app, run the command
  > ./scripts/start.sh
- To access the app, visit the browser on localhost:8080
- Use the following cmmands to run the autmated tests
  > go test (Runs the test)<br/>
  > go test -v (Runs the test in verbose mode)<br />
  > go test -cover (Run test, show coverage in terminal)<br />
  > go test -coverprofile=coverage.out && go tool cover -html=coverage.out (Run test, show coverage in browser)

## Deployment

N/A yet

## Authors

[chingsley](https://github.com/chingsley)

## Acknowledgements

- [Trevor Sawler](https://www.udemy.com/course/building-modern-web-applications-with-go/?referralCode=0415FB906223F10C6800#instructor-1)
- [Udemy](https://www.udemy.com)

## Resources and Packages

- Built in Go version 1.15
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)
