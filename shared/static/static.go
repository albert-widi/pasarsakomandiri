package static
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"log"
)

type Static struct {
	StaticAssets string		`json:"StaticAssets"`
	AssetsURI string		`json:"AssetsURI`
	StaticPage string  		`json:"StaticPage"`
	PageURI string			`json:"PageURI`
	CameraFile string 		`json:"CameraFile"`
	CameraURI string		`json:"CameraURI"`
}

var stat Static

func Configure(r *gin.Engine, s Static) {
	//set statics
	stat = s
	//r.Static("/testassets", "assets")
	r.StaticFS(s.StaticAssets, http.Dir(s.AssetsURI))
	r.StaticFS(s.StaticPage, http.Dir(s.PageURI))
	r.StaticFS(s.CameraFile, http.Dir(s.CameraURI))
}

func SaveFileToStaticFS(file []byte, fileFullPath string) error {
	var err error
	err = ioutil.WriteFile(fileFullPath, file, 0666)

	if err != nil {
		log.Println(err)
	}
	return err
}