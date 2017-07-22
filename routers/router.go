package routers

import (
	"github.com/astaxie/beego"
	"github.com/gilbrit/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/gallery", &controllers.GalleryController{})
	beego.Router("/details", &controllers.DetailsController{})
	beego.Router("/contact", &controllers.ContactController{})
}
