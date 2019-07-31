package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID            string `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	MiddleName    string `json:"middleName"`
	Age           int    `json:"age"`
	EmailAddress  string `json:"emailAddress"`
	MaritalStatus string `json:"maritalStatus"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "posttestgres"
	password = "sagarhello"
	dbname   = "GO"
)

func main() {
	r := gin.Default()
	//defining routes
	r.GET("/", GetUser)

	fmt.Println("Server is up ")
	// Start server
	r.Run(":6000")
}
