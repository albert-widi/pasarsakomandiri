package static
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Static struct {
	StaticAssets string		`json:"StaticAssets"`
	AssetsURI string		`json:"AssetsURI`
	StaticPage string  		`json:"StaticPage"`
	PageURI string			`json:"PageURI`
	CameraFile string 		`json:"CameraFile"`
	CameraURI string		`json:"CameraURI"`
}

func Configure(r *gin.Engine, s Static) {
	//set statics
	//r.Static("/testassets", "assets")
	r.StaticFS(s.StaticAssets, http.Dir(s.AssetsURI))
	r.StaticFS(s.StaticPage, http.Dir(s.PageURI))
	r.StaticFS(s.CameraFile, http.Dir(s.CameraURI))
}