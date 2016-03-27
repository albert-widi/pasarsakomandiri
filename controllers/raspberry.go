package controllers

import (
    "github.com/pasarsakomandiri/api"
    "github.com/gin-gonic/gin"
)

func RaspberryPrintCheckOut(c *gin.Context) {
    raspberryPi := &api.RaspberryPi{}
    raspberryPi.Protocol="http"
    raspberryPi.Host = "192.168.0.177"
    raspberryPi.Port = "8888"
    raspberryPi.Token = "testing"
    raspberryPi.Param = ""
    raspberryPi.RaspberryPrintTicketOut()
}