package models

import "database/sql"
import (
	"fmt"
	"time"
)

func RegisterSuperUser(db *sql.DB, username string, password string) {
	prepareStamtenet := "INSERT INTO user(username, password, level, role, description, created_by, created_date, last_update_date) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	statement, _ := db.Prepare(prepareStamtenet)
	result, _ := statement.Exec(username, password, 99, "Super User", "Created By CLI", -1, time.Now(), time.Now());
	lastID, _ := result.LastInsertId()
	affectedRows, _ := result.RowsAffected()
	fmt.Println("Users created, last_id : %s, affected_id: %s", lastID, affectedRows)
}