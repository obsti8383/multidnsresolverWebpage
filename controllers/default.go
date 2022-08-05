package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "none"
	c.Data["Email"] = "none"
	inputValues, err := c.Input()
	domainName := ""
	if err == nil && inputValues["domain"] != nil {
		domainName = inputValues["domain"][0]
		if domainName != "" {
			c.Data["Output"] = QueryMultipleDNS(domainName)
		}
	}
	c.Data["Servername"] = domainName
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	domainName := c.GetString("Servername")
	c.Ctx.Redirect(302, "?domain="+domainName)
}
