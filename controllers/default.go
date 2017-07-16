package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gilbrit/blog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "linuxbrit.co.uk"
	c.Data["Email"] = "tom@linuxbrit.co.uk"
	c.Data["Title"] = "Tom Gilbert's blog"
	c.Data["Page"] = "index"
	c.Data["Images"] = models.GetImages(8)
	v, _ := c.GetBool("contact")
	if v {
		c.Data["Contact"] = true
	}
	c.TplName = "index.html"
}
