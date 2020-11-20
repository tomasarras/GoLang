package agencyController

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tomasarras/GoLang/controller"
	"github.com/tomasarras/GoLang/entity"
	"github.com/tomasarras/GoLang/service/agencyService"
)

var serviceAgency agencyService.ServiceAgency

func Start(db *sql.DB) {
	serviceAgency, _ = agencyService.New(db)
}

func SaveAgencyHandler(c *gin.Context) {
	var reqBody entity.Agency
	err := c.ShouldBindJSON(&reqBody)

	if err != nil || !reqBody.IsValid() {
		controller.BadRequest(c)
		return
	}

	response, _ := serviceAgency.Save(reqBody)
	c.JSON(http.StatusOK, response.ToJson())
}

func FindByIdAgencyHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		controller.BadRequest(c)
		return
	}

	a, _ := serviceAgency.FindByID(id)
	if a.ID == 0 {
		controller.NotFound(c, id)
		return
	}
	c.JSON(http.StatusOK, a.ToJson())
}

func FindAllAgencyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, serviceAgency.FindAll())
}

func UpdateAgencyHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		controller.BadRequest(c)
		return
	}

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		controller.BadRequest(c)
		return
	}

	var requestAgency entity.Agency
	json.Unmarshal(jsonData, &requestAgency)
	if !requestAgency.IsValid() {
		controller.BadRequest(c)
		return
	}

	requestAgency.ID = int64(id)
	serviceAgency.Update(requestAgency)
	c.JSON(http.StatusOK, requestAgency.ToJson())
}

func RemoveAgencyHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		controller.BadRequest(c)
		return
	}

	serviceAgency.Remove(id)
	c.JSON(http.StatusOK, "")
}
