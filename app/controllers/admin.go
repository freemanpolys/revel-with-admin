package controllers

import (
	"github.com/revel/revel"

	"akhosto/app/routes"
)

type Admin struct {
	Application
}

func (c Admin) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.Application.Index())
	}
	return nil
}

func (c Admin) Index() revel.Result {
	c.Log.Info("Fetching index")

	return c.Render()
}


func (c Admin) ChartsChartjs() revel.Result {
	return c.Render()
}


func (c Admin) ElementsButtonsStandard() revel.Result {
	return c.Render()
}

func (c Admin) ElementsIcons() revel.Result {
	return c.Render()
}


