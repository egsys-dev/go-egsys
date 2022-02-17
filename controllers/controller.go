package controllers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/fabiom2211/go-egsys/database"
	"github.com/fabiom2211/go-egsys/models"
	"github.com/gin-gonic/gin"
)

func Inicio(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Bem Vindo API"})

}

func ListFleets(c *gin.Context) {
	var fleet []models.Fleet
	//Falta retornar o model sem os relacionamentos
	database.DB.Find(&fleet)
	c.JSON(http.StatusOK, fleet)
}

func CreateFleets(c *gin.Context) {
	var fleet models.Fleet
	if err := c.ShouldBindJSON(&fleet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if fleet.Max_speed == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "A velocidade deve ser maior que 0."})
		return
	}

	database.DB.Create(&fleet)

	if fleet.Fleet_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fleets não cadastrado"})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id": fleet.Fleet_id})
		return
	}

}

func GetFleets(c *gin.Context) {
	var fleetAlert models.FleetAlert
	var fleetAlertList []models.FleetAlert

	fleetAlert.Fleet_id, _ = strconv.Atoi(c.Param("id"))

	if fleetAlert.Fleet_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	database.DB.Where(&models.FleetAlert{Fleet_id: fleetAlert.Fleet_id}).Find(&fleetAlertList)

	if len(fleetAlertList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, fleetAlertList)
}

func CreateFleetAlerts(c *gin.Context) {
	var fleet models.FleetAlert
	var fleetAlert models.FleetAlert

	if err := c.ShouldBindJSON(&fleetAlert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	} else if fleetAlert.Fleet_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	//valida URL
	u, err := url.Parse(fleetAlert.Webhook)
	if !(err == nil && u.Scheme != "" && u.Host != "") {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	fleet.Fleet_id, _ = strconv.Atoi(c.Param("id"))
	database.DB.First(fleet.Fleet_id)
	if fleet.Fleet_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	fleetAlert.Fleet_id = fleet.Fleet_id
	database.DB.Create(&fleetAlert)

	if fleetAlert.Fleet_alert_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id": fleetAlert.Fleet_alert_id})
		return
	}

}

func ListVehicles(c *gin.Context) {
	var vehicles []models.Vehicle
	database.DB.Find(&vehicles)
	//Precisa verificar se a velocidade do veiculo está preenchida, caso não esteja precisa pegar a velocidade da frota
	c.JSON(http.StatusOK, vehicles)
}

func CreateVehicles(c *gin.Context) {
	var vehicle models.Vehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&vehicle)

	if vehicle.Fleet_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Vehicle não cadastrado"})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id": vehicle.Vehicle_id})
		return
	}
}

//Parei o teste aqui
func GetVehiclesPosition(c *gin.Context) {
	var fleetAlert models.Vehicle
	var fleetAlertList []models.Vehicle

	fleetAlert.Fleet_id, _ = strconv.Atoi(c.Param("id"))

	if fleetAlert.Fleet_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	database.DB.Where(&models.FleetAlert{Fleet_id: fleetAlert.Fleet_id}).Find(&fleetAlertList)

	if len(fleetAlertList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, fleetAlertList)
}
