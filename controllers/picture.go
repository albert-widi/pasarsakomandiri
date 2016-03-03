package controllers

type IpCamPicture struct {
	Id int64
	Camera_id int
	Filepath string
	Filename string
	Expired_date string
	Created_by int64
	Created_date string
	PictureBytes []byte
}

func PictureSaveToDB() {

}

func PictureSaveToFS() {

}

func PictureDelete() {

}
