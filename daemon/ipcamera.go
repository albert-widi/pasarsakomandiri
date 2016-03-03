package daemon

import (
	"time"
	"github.com/pasarsakomandiri/models"
	"os"
	"log"
)

//this ipcamera daemon creatred to delete ipcamera image automatically for every week/month
func InitPicturesDaemon() {
	for {
		log.Println("Pictures daemon started at " + time.Now().Format("2006-01-02 03:04:05"))
		//only check expired pictures between 24:00AM to 04:00AM
		if time.Now().Hour() >= 0 && time.Now().Hour() <= 4 {
			condition := "expired_date <= " + time.Now().Format("2006-01-02")
			pictures, err := models.PictureGetAll(condition)

			if err == nil {
				for _, elem := range pictures {
					//remove picture from directory
					err :=  os.Remove(elem.Filepath+string(os.PathSeparator)+elem.Filename+"."+elem.Format)

					if err != nil {
						log.Println(err)
						continue
					}
					//remove picture path from database
					err = models.PictureDelete(elem.Id)

					if err != nil {
						log.Println(err)
					}
				}
			} else {
				log.Println(err)
			}
		}
		//sleep the loop for every two hours
		time.Sleep(time.Hour * 2)
	}
}