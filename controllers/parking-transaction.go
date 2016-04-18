package controllers

import (
	_ "fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/models"
	"github.com/pasarsakomandiri/shared/response"
	"github.com/pasarsakomandiri/shared/view"
)

type ParkingTransactionsPerUser struct {
	Verified_by  int64
	Username     string
	NOC          int
	NOM          int
	Parking_cost int
}

func ParkingTransactionsPage(c *gin.Context) {
	v := view.New(c)
	v.Name = "parking_transactions"
	v.Render()
	//session := session.Instance(c)
	//c.HTML(http.StatusFound, "parking_transactions.tmpl", gin.H{"title":"Cahiser", "token":session.Get("token")})
}

func ParkingTransactionsGetAll(c *gin.Context) {
	tanggal := c.Query("tanggal")
	param := c.Query("param")
	queryParam := ""

	if tanggal != "" {
		tanggal = strings.Replace(tanggal, ",", "", 1)
		queryParam += "DATE_FORMAT(created_date, '%e %M %Y') = " + "'" + tanggal + "'"

		switch param {
		case "All":
			queryParam += ""
		case "Sudah keluar":
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

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, transactionparking)
	return
}

func ParkingTransactionCashier(c *gin.Context) {
	date := c.Query("tanggal")
	hour := c.Query("jam")

	date = strings.Replace(date, ",", "", 1)

	if date == "" || hour == "" {
		c.JSON(http.StatusOK, response.NewSimpleResponse("failed", "Date or hour cannot be null"))
		return
	}

	fixDate := date + " " + hour

	data, err := models.UserParkingTransactions(fixDate)

	//fmt.Printf("%+v", data)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.NewSimpleResponse("failed", err.Error()))
		return
	}

	dataSet := []ParkingTransactionsPerUser{}

	for _, value := range data {
		count, err := models.PTGetVehicleCountByDate(value.Verified_by, fixDate, value.Vehicle_id)
		found := false

		if err == nil {
			for key, valueSet := range dataSet {
				if valueSet.Username == value.Username {
					if err == nil {
						dataSet[key].Parking_cost += value.Parking_cost

						if value.Vehicle_id == 1 {
							dataSet[key].NOM = count

						}

						if value.Vehicle_id == 2 {
							dataSet[key].NOC = count
						}
					} else {
						log.Println(err)
					}
					found = true
					break
				}
			}

			if found {
				continue
			}

			//insert new data
			userData := ParkingTransactionsPerUser{}
			userData.Username = value.Username
			userData.Verified_by = value.Verified_by
			userData.Parking_cost = value.Parking_cost

			if value.Vehicle_id == 1 {
				userData.NOM = count
			} else {
				userData.NOC = count
			}

			dataSet = append(dataSet, userData)
		} else {
			log.Println(err)
		}
	}

	c.JSON(http.StatusOK, dataSet)
}
