package main

import (
	"github.com/astaxie/beego"
	_ "github.com/gilbrit/blog/routers"
)

func main() {
	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/gallery", &controllers.GalleryController{})
	// beego.Router("/details", &controllers.DetailsController{})
	// beego.Router("/contact", &controllers.ContactController{})
	beego.Run()
}
