package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	//opening db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query(`SELECT * from users`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	person := []Person{}
	// fmt.Println(rows)

	sql.ErNoRows()
	for rows.Next() {
		var p Person
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.MiddleName, &p.Age, &p.EmailAddress, &p.MaritalStatus); err != nil {
			panic(err)

		}
		person = append(person, p)
	}

	// Returns the result as json
	c.JSON(200, gin.H{
		data: person,
	})
}
