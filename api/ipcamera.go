package api

import (
	"net/http"
	"io/ioutil"
	"log"
)

type IpCamera struct {
	Protocol string
	Host string
	Param string
	Username string
	Password string
	Picture chan []byte
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

func (cam *IpCamera) GetPicture() {
	picture, err := ipCamGetPicture(cam)

	if err != nil {
		log.Println(err)
	}

	cam.Picture <- picture
}
