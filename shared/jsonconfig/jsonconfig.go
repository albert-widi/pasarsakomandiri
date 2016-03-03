package jsonconfig

import (
	//"io"
	//"os"
	"log"
	"io/ioutil"
	"github.com/pasarsakomandiri/shared/database"
	_"encoding/json"
	_"fmt"
	"github.com/pasarsakomandiri/shared/server"
	"github.com/pasarsakomandiri/shared/static"
	"github.com/pasarsakomandiri/shared/session"
	"github.com/pasarsakomandiri/shared/token"
)

type Parser interface {
	ParseJSON([]byte) error
}

type testconfig struct {
	Database database.Database		`json:"Database"`
	Server server.Server        	`json:"Server"`
	Static static.Static			`json:"Static`
	Session session.Session			`json:"Session`
	Token token.Token				`json:"Token`
}

func Load(fileName string, p Parser) {
	/*var err error
	var input = io.ReadCloser(os.Stdin)
	if input, err = os.Open(fileName); err != nil {
		log.Fatalln(err)
	}*/

	//read config file
	//jsonBytes, err := ioutil.ReadAll(input)
	//input.Close()
	/*if err != nil {
		log.Fatalln(err)
	}*/

	jsonBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	/*var test testconfig
	json.Unmarshal(jsonBytes, &test)*/
	//fmt.Printf("%+v", test)

	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", fileName, err)
	}
}