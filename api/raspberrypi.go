package api

import (
	"net/http"
	"log"
    "io/ioutil"
)

type RaspberryPi struct {
    Protocol string
    Host string
    Port string
    Token string
    Param string
}

//RaspberryPrintTicketOut order raspberry to print ticket out
func (r *RaspberryPi) RaspberryPrintTicketOut() {
    path := r.fullPath()
    param := r.fullParam()
    completeURL := path + "/print/checkOut" + param
    
    client := http.Client{}
	req, err := http.NewRequest("POST", completeURL, nil)
    
    if err != nil {
        log.Println(err)
    }
    
    res, err := client.Do(req)
    
    if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
    
    responseBody, err := ioutil.ReadAll(res.Body)
    
    log.Println(string(responseBody))
}

func (r *RaspberryPi) fullPath() string {
    return r.Protocol + "://" + r.Host + ":" + r.Port
}

func (r *RaspberryPi) fullParam() string {
    return "?token=" + r.Token + r.Param
}