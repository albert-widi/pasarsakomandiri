package api

import (
	"net/http"
	"io/ioutil"
	"log"
    	"time"
    	"strings"
)

type IpCamera struct {
	Protocol string
	Host string
	Param string
	Username string
	Password string
    	PictureName string
	Picture []byte
}

func IpCamGetPicture(ipcam IpCamera) ([]byte, error) {
	var pictureByte []byte
	completeUrl := ipcam.Protocol+"://"+ipcam.Host+"/"+ipcam.Param

	client := http.Client{}
	req, err := http.NewRequest("GET", completeUrl, nil)
	req.SetBasicAuth(ipcam.Username, ipcam.Password)
	if err != nil {
		return pictureByte, err
	}

	res, err := client.Do(req)

	if err != nil {
		return pictureByte, err
	}

	defer res.Body.Close()

	pictureByte, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return pictureByte, err
	}

	return pictureByte, err
}

func ipCamGetPicture(ipcam *IpCamera) ([]byte, error) {
	var pictureByte []byte
	completeUrl := ipcam.Protocol+"://"+ipcam.Host+"/"+ipcam.Param

	client := http.Client{}
	req, err := http.NewRequest("GET", completeUrl, nil)
	req.SetBasicAuth(ipcam.Username, ipcam.Password)
	if err != nil {
		return pictureByte, err
	}

	res, err := client.Do(req)

	if err != nil {
		return pictureByte, err
	}

	defer res.Body.Close()

	pictureByte, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return pictureByte, err
	}

	return pictureByte, err
}

//call this function with goroutines
func (cam *IpCamera) GetPicture() []byte {
	picture, err := ipCamGetPicture(cam)

	if err != nil {
		log.Println(err)
	}
    
	//generate the name
	dateTimeName := time.Now().Format("2006-01-02 03:04:05 PM")
	pictureName := strings.Replace(dateTimeName, ":", "", 10)
    	cam.PictureName = pictureName
    	//send picture to channel
	cam.Picture = picture
}

//call this function with goroutines
func (cam *IpCamera) GetPictureWithChannel(c chan []byte) []byte {
	picture, err := ipCamGetPicture(cam)

	if err != nil {
		log.Println(err)
	}

	//generate the name
	dateTimeName := time.Now().Format("2006-01-02 03:04:05 PM")
	pictureName := strings.Replace(dateTimeName, ":", "", 10)
	cam.PictureName = pictureName
	cam.Picture = picture
	//send picture to channel
	c <- picture
}

