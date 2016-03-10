package models

import (
	"database/sql"
	"github.com/pasarsakomandiri/shared/database"
	"os"
)

type Picture struct {
	Id int64
	Filepath string
	Filename string
	Format string
	Expired_date string
	Created_by int64
	Created_date string
    	PictureFullPath string
}

func PictureGetAll(condition string) ([]Picture, error) {
	pictures := []Picture{}

	if condition == "" {
		condition = "1=1"
	}

	err := database.Db.Select(&pictures, "SELECT id, filepath, filename, expired_date, created_by, created_date FROM pictures WHERE "+condition)
	return pictures, err
}

func PictureGetById(pictureId int64) (Picture, error) {
	picture := Picture{}
	err :=  database.Db.Get(&picture, "SELECT id, filepath, filename, format, expired_date, created_by, created_date FROM pictures WHERE id=?", pictureId)
	return picture, err
}

func PictureUpdateName(pictureName string, pictureId int64) error {
	_, err := database.Db.Exec("UPDATE pictures SET filename=? WHERE id=?", pictureName, pictureId)
	return err
}

func PictureSave(pic Picture) (sql.Result, error) {
	result , err := database.Db.Exec("INSERT INTO pictures(filepath, filename, format, expired_date, created_by, created_date) VALUES(?, ?, ?, ?, ?, ?)", pic.Filepath, pic.Filename, pic.Format, pic.Expired_date, pic.Created_by, pic.Created_date)
	return result, err
}


func PictureDelete(picId int64) error {
	_, err := database.Db.Exec("DELETE FROM pictures WHERE id = ?", picId)
	return err
}

func (pic *Picture) GetFullPath() string {
	return pic.Filepath+string(os.PathSeparator)+pic.Filename+"."+pic.Format
}