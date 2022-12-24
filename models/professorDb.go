package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func GetProfessores() []Professor {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM professor")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	professores := []Professor{}
	for results.Next() {
		var professor Professor
        // for each row, scan into the Product struct
		err = results.Scan(&professor.Id, &professor.Nome)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
        // append the product into products array
		professores = append(professores, professor)
	}

	return professores

}

func GetProfessor(id int) *Professor {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	professor := &Professor{}
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}
    
	defer db.Close()

	results, err := db.Query("SELECT * FROM professor where id=?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&professor.Id, &professor.Nome)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return professor
}

func AddProfessor(professor Professor) {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO professor (id,nome) VALUES (?,?)",
		professor.Id, professor.Nome)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}