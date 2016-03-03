package controllers

import (
	"strings"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/pasarsakomandiri/shared/response"
	"github.com/pasarsakomandiri/models"
	"github.com/pasarsakomandiri/shared/session"
)

func ParkingTransactionsPage(c *gin.Context) {
	session := session.Instance(c)
	c.HTML(http.StatusFound, "parking_transactions.tmpl", gin.H{"title":"Cahiser", "token":session.Get("token")})
}

func ParkingTransactionsGetAll(c *gin.Context) {
	tanggal := c.Query("tanggal")
	param := c.Query("param")
	queryParam := ""


	if tanggal != "" {
		tanggal = strings.Replace(tanggal, ",", "", 1)
		queryParam += "DATE_FORMAT(created_date, '%e %M %Y') = "+ "'" + tanggal + "'"


		switch param {
		case "All":
			queryParam += ""
		case  "Sudah keluar":
			queryParam += " AND parking_cost IS NOT NULL"
		case "Belum keluar":
			queryParam += " AND parking_cost IS NULL"
		}
	} else {
		log.Println("No date detected")
		c.JSON(http.StatusOK, response.NewSimpleResponse("Failed", "Date cannot be null"))
		return
	}

	transactionparking, err := models.ParkingTransactionGetAllAPI(queryParam)

	if err != nil{
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, transactionparking)
	return
}