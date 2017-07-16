package controllers

import (
	"github.com/astaxie/beego"
)

type DetailsController struct {
	beego.Controller
}

func (c *DetailsController) Get() {
	c.Data["Website"] = "linuxbrit.co.uk"
	c.Data["Email"] = "tom@linuxbrit.co.uk"
	c.Data["Title"] = "Details"
	c.Data["Page"] = "details"

	c.TplName = "details.html"
}
