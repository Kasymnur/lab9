package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" 
)

type Person struct {
	ID    int
	Name  string
	Age   int
	Email string
}

func main() {

	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		rows, err := db.Query("SELECT * FROM people")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		people := []Person{}

		
		for rows.Next() {
			p := Person{}
			err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Email)
			if err != nil {
				log.Fatal(err)
			}
			people = append(people, p)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		tmpl := template.Must(template.ParseFiles("index.html"))
		err = tmpl.Execute(w, people)
		if err != nil {
			log.Fatal(err)
		}
	})

	
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
