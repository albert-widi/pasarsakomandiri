package api

import (
	"net/http"
	"log"
    "io/ioutil"
    "net/url"
    "strings"
)

type RaspberryPi struct {
    Protocol string
    Host string
    Port string
    Token string
    Param string
    Data map[string]string
}

//RaspberryPrintTicketOut order raspberry to print ticket out
func (r *RaspberryPi) RaspberryPrintTicketOut() {
    path := r.fullPath()
    //param := r.fullParam()
    completeURL := path + "/print/checkOut"
    
    data := url.Values{}

    for key, value := range r.Data {
        data.Add(key, value)
    }
    
    client := http.Client{}
	req, err := http.NewRequest("POST", completeURL, strings.NewReader(data.Encode()))
    
    if err != nil {
        log.Println(err) 
        return
    }
    
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    
    
    res, err := client.Do(req)
    
    if err != nil {
		log.Println(err)
        return
	}
    
    _, err = ioutil.ReadAll(res.Body)
    res.Body.Close()
    //log.Println(string(responseBody))
}

func (r *RaspberryPi) fullPath() string {
    return r.Protocol + "://" + r.Host + ":" + r.Port
}

func (r *RaspberryPi) fullParam() string {
    return "?token=" + r.Token + r.Param
}
