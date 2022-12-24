package main

import (
	"apiatividade_go/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "apiatividade_go/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gerenciador de Atividades API
// @version 1.0
// @description This is an example API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8083
func main() {
	router := gin.Default()

	router.GET("/alunos", getAlunos)
	router.GET("/aluno/:id", getAluno)
	router.POST("/addAluno", addAluno)

	router.GET("/atividades", getAtividades)
	router.GET("/atividade/:id", getAtividade)
	router.POST("/addAtividade", addAtividade)

	router.GET("/entregas", getEntregas)
	router.GET("/entrega/:id", getEntrega)
	router.POST("/addEntrega", addEntrega)
	router.PUT("/updateEntrega", updateEntrega)

	router.GET("/professores", getProfessores)
	router.GET("/professor/:id", getProfessor)
	router.POST("/addProfessor", addProfessor)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8083")
}

// getAlunos pega todos os alunos.
// getAlunos             godoc
// @Summary      Pega todos os alunos
// @Description  Pega todos os alunos
// @Tags         getAlunos
// @Produce      json
// @Success      200  {array}  models.Aluno
// @Router       /alunos [get]
func getAlunos(c *gin.Context) {
	alunos := models.GetAlunos()

	if alunos == nil || len(alunos) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, alunos)
	}
}

// getAluno pega um aluno.
// getAluno             godoc
// @Summary      Pega um aluno
// @Description  Pega um aluno
// @Tags         getAluno
// @Produce      json
// @Param        id     path     string     true        "id"
// @Success      200  {array}  models.Aluno
// @Router       /aluno/{id} [get]
func getAluno(c *gin.Context) {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)

	if err != nil {
		return
	}

	aluno := models.GetAluno(convId)

	if aluno == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, aluno)
	}
}

// addAluno adiciona um aluno.
// addAluno             godoc
// @Summary      Adiciona um aluno
// @Description  Adiciona um aluno
// @Tags         addAluno
// @Produce      json
// @Param        body     body     models.Aluno     true        "Aluno"
// @Success      200  {array}  models.Aluno
// @Router       /addAluno [post]
func addAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.BindJSON(&aluno); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddAluno(aluno)
		c.IndentedJSON(http.StatusCreated, aluno)
	}
}

// getAtividades pega todas as atividades.
// getAtividades             godoc
// @Summary      Pega tados as atividades
// @Description  Pega tados as atividades
// @Tags         getAtividades
// @Produce      json
// @Success      200  {array}  models.Atividade
// @Router       /atividades [get]
func getAtividades(c *gin.Context) {
	atividades := models.GetAtividades()

	if atividades == nil || len(atividades) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, atividades)
	}
}

// getAtividade pega uma atividade.
// getAtividade             godoc
// @Summary      Pega uma atividade
// @Description  Pega uma atividade
// @Tags         getAtividade
// @Produce      json
// @Param        id     path     string     true        "id"
// @Success      200  {array}  models.Atividade
// @Router       /atividade/{id} [get]
func getAtividade(c *gin.Context) {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)

	if err != nil {
		return
	}

	atividade := models.GetAtividade(convId)

	if atividade == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, atividade)
	}
}

// addAtividade adiciona uma atividade.
// addAtividade             godoc
// @Summary      Adiciona uma atividade
// @Description  Adiciona uma atividade
// @Tags         addAtividade
// @Produce      json
// @Param        body     body     models.Atividade     true        "Atividade"
// @Success      200  {array}  models.Atividade
// @Router       /addAtividade [post]
func addAtividade(c *gin.Context) {
	var atividade models.Atividade

	if err := c.BindJSON(&atividade); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddAtividade(atividade)
		c.IndentedJSON(http.StatusCreated, atividade)
	}
}

// getEntregas pega todas as entregas.
// getEntregas             godoc
// @Summary      Pega todas as entregas
// @Description  Pega todas as entregas
// @Tags         getEntregas
// @Produce      json
// @Success      200  {array}  models.Entrega
// @Router       /entregas [get]
func getEntregas(c *gin.Context) {
	entregas := models.GetEntregas()

	if entregas == nil || len(entregas) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, entregas)
	}
}

// getEntrega pega uma entrega.
// getEntrega             godoc
// @Summary      Pega uma entrega
// @Description  Pega uma entrega
// @Tags         getEntrega
// @Produce      json
// @Param        id     path     string     true        "id"
// @Success      200  {array}  models.Entrega
// @Router       /entrega/{id} [get]
func getEntrega(c *gin.Context) {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)

	if err != nil {
		return
	}

	entrega := models.GetEntrega(convId)

	if entrega == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, entrega)
	}
}

// addEntrega adiciona uma entrega.
// addEntrega             godoc
// @Summary      Adiciona uma entrega
// @Description  Adiciona uma entrega
// @Tags         addEntrega
// @Produce      json
// @Param        body     body     models.Entrega     true        "Entrega"
// @Success      200  {array}  models.Entrega
// @Router       /addEntrega [post]
func addEntrega(c *gin.Context) {
	var entrega models.Entrega

	if err := c.BindJSON(&entrega); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddEntrega(entrega)
		c.IndentedJSON(http.StatusCreated, entrega)
	}
}

// updateEntrega responds with the list of all books as JSON.
// updateEntrega             godoc
// @Summary      Atualiza a nota na entrega do aluno
// @Description  Atualiza a nota na entrega do aluno
// @Tags         updateEntrega
// @Produce      json
// @Param        body     body     models.EntregaRequest     true        "EntregaRequest"
// @Success      200  {array}  models.EntregaRequest
// @Router       /updateEntrega [put]
func updateEntrega(c *gin.Context) {
	var request models.EntregaRequest

	log.Println(string(request.Nota))

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.UpdateEntrega(request.Id, request.Nota)
		c.IndentedJSON(http.StatusCreated, request)
	}

	c.IndentedJSON(http.StatusOK, "Nota atualizada com sucesso!")
}

// getProfessores pega todos os professores.
// getProfessores             godoc
// @Summary      Pega todos os professores
// @Description  Pega todos os professores
// @Tags         getProfessores
// @Produce      json
// @Success      200  {array}  models.Professor
// @Router       /professores [get]
func getProfessores(c *gin.Context) {
	professores := models.GetProfessores()

	if professores == nil || len(professores) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, professores)
	}
}

// getProfessor pega um professor.
// getProfessor             godoc
// @Summary      Pega um professor
// @Description  Pega um professor
// @Tags         getProfessor
// @Produce      json
// @Param        id     path     string     true        "id"
// @Success      200  {array}  models.Professor
// @Router       /professor/{id} [get]
func getProfessor(c *gin.Context) {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)

	if err != nil {
		return
	}

	professor := models.GetProfessor(convId)

	if professor == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, professor)
	}
}

// addProfessor adiciona um professor.
// addProfessor             godoc
// @Summary      Adiciona um professor
// @Description  Adiciona um professor
// @Tags         addProfessor
// @Produce      json
// @Param        body     body     models.Professor     true        "Professor"
// @Success      200  {array}  models.Professor
// @Router       /addProfessor [post]
func addProfessor(c *gin.Context) {
	var professor models.Professor

	if err := c.BindJSON(&professor); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddProfessor(professor)
		c.IndentedJSON(http.StatusCreated, professor)
	}
}
