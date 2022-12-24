package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func GetAtividades() []Atividade {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM atividade")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	atividades := []Atividade{}
	for results.Next() {
		var atividade Atividade
        // for each row, scan into the Atividade struct
		err = results.Scan(&atividade.Id, &atividade.ProfessorId, &atividade.Titulo, &atividade.Descricao, &atividade.DataEntrega)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
        // append the atividade into atividades array
		atividades = append(atividades, atividade)
	}

	return atividades

}

func GetAtividade(id int) *Atividade {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	atividade := &Atividade{}
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}
    
	defer db.Close()

	results, err := db.Query("SELECT * FROM atividade where id=?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&atividade.Id, &atividade.ProfessorId, &atividade.Titulo, &atividade.Descricao, &atividade.DataEntrega)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return atividade
}

func AddAtividade(atividade Atividade) {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO atividade (id, professor_id, titulo, descricao, data_entrega) VALUES (?,?,?,?, now())",
		atividade.Id, atividade.ProfessorId, atividade.Titulo, atividade.Descricao)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}