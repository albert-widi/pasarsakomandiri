package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pasarsakomandiri/shared/database"
)

type User struct {
	Id           int64
	Username     string
	Password     string
	Level        int
	Role         string
	Description  string
	Created_date string
	Created_by   int64
	Status       int
}

func UserGetAll() ([]User, error) {
	var user = []User{}
	err := database.Db.Select(&user, "SELECT id, username, password, level, role, description, created_date FROM user")
	return user, err
}

//UserGetAllLimitByLevel
func UserGetAllLimitByLevel(level int) ([]User, error) {
	var user = []User{}
	//log.Println(level)
	err := database.Db.Select(&user, "SELECT id, username, password, level, role, description, created_date FROM user WHERE level <= ?", level)
	//fmt.Printf("%+v", user)
	return user, err
}

//UpdateUserStatus is only temporary way to block user doing multiple login
func UpdateUserStatus(id int64, status int) error {
	_, err := database.Db.Exec("UPDATE user SET status=? WHERE id=?", status, id)
	return err
}

func UserGetByUsername(username string) (User, error) {
	var user = User{}
	err := database.Db.Get(&user, "SELECT id, username, password, level, role, description, created_date, status FROM user WHERE username=?", username)
	return user, err
}

func UserGetByID(id int64) (User, error) {
	var user = User{}
	err := database.Db.Get(&user, "SELECT id, username, password, level, role ,description, created_date FROM user WHERE id=?", id)
	return user, err
}

func UserCreateNew(user User) (sql.Result, error) {
	result, err := database.Db.Exec("INSERT INTO user(username, password, level, role, description, created_by, created_date, status) VALUES(?, ?, ?, ?, ?, ?, ?, 0)", user.Username, user.Password, user.Level, user.Role, user.Description, user.Created_by, user.Created_date)
	return result, err
}

func UserLevelByRole(role string) (int, error) {
	var level int
	err := database.Db.Get(&level, "SELECT role_level FROM roles WHERE role_name=?", role)

	return level, err
}

func UserUpdate(user User) error {
	_, err := database.Db.Exec("UPDATE user SET username=?, password=?, role=?, description=? WHERE id=?", user.Username, user.Password, user.Role, user.Description, user.Id)
	return err
}

func UserResetLoginStatus() error {
	_, err := database.Db.Exec("UPDATE user SET status = ?", 0)
	return err
}
