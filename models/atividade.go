package models

type Atividade struct {
	Id int `json:"id"`
	ProfessorId int `json:"professor_id"`
	Titulo string `json:"titulo"`
	Descricao string `json:"descricao"`
	DataEntrega string `json:"data_entrega"`
}