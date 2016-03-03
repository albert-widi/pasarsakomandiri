package api
import (
	"net/http"
	"net/url"
	"fmt"
	"bytes"
	"strconv"
	"io/ioutil"
)

func CameraTakePicture(cameraIP string) {

}

func ContactClientCheckOut() {
	data := url.Values{}

	client := &http.Client{}
	r, _ := http.NewRequest("GET", "http://192.168.0.177:8080", bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Body)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}

