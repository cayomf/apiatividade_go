package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root"
const dbpass = "root"
const dbname = "atividade_db"

func GetAlunos() []Aluno {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM aluno")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	alunos := []Aluno{}
	for results.Next() {
		var aluno Aluno
		// for each row, scan into the Product struct
		err = results.Scan(&aluno.Id, &aluno.Nome)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the product into products array
		alunos = append(alunos, aluno)
	}

	return alunos

}

func GetAluno(id int) *Aluno {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	aluno := &Aluno{}
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM aluno where id=?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&aluno.Id, &aluno.Nome)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return aluno
}

func AddAluno(aluno Aluno) {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO aluno (id,nome) VALUES (?,?)",
		aluno.Id, aluno.Nome)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
