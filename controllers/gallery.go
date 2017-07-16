package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gilbrit/blog/models"
)

type GalleryController struct {
	beego.Controller
}

func (c *GalleryController) Get() {
	c.Data["Website"] = "linuxbrit.co.uk"
	c.Data["Email"] = "tom@linuxbrit.co.uk"
	c.Data["Title"] = "Gallery"
	c.Data["Page"] = "gallery"
	c.Data["Images"] = models.GetImages(0)

	c.TplName = "gallery.html"
}
