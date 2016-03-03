package database
import (
	"strings"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"
	"github.com/pasarsakomandiri/shared/database/middleware"
)

var (
	Db *sqlx.DB
	err error
)

type Database struct {
	Type	string
	MySQL 	MySQLInfo
}

type MySQLInfo struct {
	Username 	string    `json:"Username"`
	Password 	string    `json:"Password"`
	Name 		string    `json:"Name"`
	Hostname	string    `json:"Hostname"`
	Port		string    `json:"Port"`
	DBParam	string    	  `json:"DBParam"`
}

func Connect(r *gin.Engine, database Database) {


	switch database.Type {
		case "MySQL":
			driver := strings.ToLower(database.Type)
			source := database.MySQL.Username+":"+database.MySQL.Password+"@tcp("+database.MySQL.Hostname+":"+database.MySQL.Port+")"+"/"+database.MySQL.Name//+db.MySQL.DBParam
			//fmt.Println(source)
			if Db, err = sqlx.Open(driver, source); err != nil {
				log.Println("Database driver error")
			}


			if err = Db.Ping(); err != nil {
				log.Println("Database error")
				log.Println("Database:", Db.Ping())
			} else {
				Db.SetMaxOpenConns(20)
				Db.SetMaxIdleConns(5)
				r.Use(middleware.AddDB(Db))
			}
	}
}

func DbInstance(c *gin.Context) *sqlx.DB {
	return c.MustGet("database").(*sqlx.DB)
}