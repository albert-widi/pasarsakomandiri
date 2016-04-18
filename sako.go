package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pasarsakomandiri/daemon"
	"github.com/pasarsakomandiri/router"
	"github.com/pasarsakomandiri/shared/database"
	"github.com/pasarsakomandiri/shared/jsonconfig"
	"github.com/pasarsakomandiri/shared/server"
	"github.com/pasarsakomandiri/shared/session"
	"github.com/pasarsakomandiri/shared/static"
	"github.com/pasarsakomandiri/shared/token"
	"github.com/pasarsakomandiri/shared/view"
)

func init() {
	//activate all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	//will need this random seed later
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	r := gin.Default()
	jsonconfig.Load("./configfile"+string(os.PathSeparator)+"testconfig.json", config)
	database.Connect(r, config.Database)
	session.Configure(r, config.Session)
	token.Configure(config.Token)
	static.Configure(r, config.Static)
	static.SetDefaultTemplate(r)
	view.Configure(config.View)
	router.Initialize(r)
	//init daemon
	go daemon.InitPicturesDaemon()

	//set default template

	server.Run(r, config.Server)
}

//================ CONFIGURATION ================

var config = &configuration{}

type configuration struct {
	Database database.Database `json:"Database"`
	Server   server.Server     `json:"Server"`
	Static   static.Static     `json:"Static"`
	Session  session.Session   `json:"Session"`
	Token    token.Token       `json:"Token"`
	View     view.View         `json:"View"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
