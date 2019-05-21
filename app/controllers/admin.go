package controllers

import (
	"revel-with-admin/app/routes"
	"github.com/revel/revel"
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

func (c Admin) Table() revel.Result {

	return c.Render()
}

func (c Admin) Chart() revel.Result {
	return c.Render()
}

func (c Admin) General() revel.Result {
	return c.Render()
}

func (c Admin) Components() revel.Result {
	return c.Render()
}

func (c Admin) Buttons() revel.Result {
	return c.Render()
}

func (c Admin) Toastr() revel.Result {
	return c.Render()
}

func (c Admin) Fontawesome() revel.Result {
	return c.Render()
}

func (c Admin) IconIcons() revel.Result {
	return c.Render()
}

func (c Admin) Credits() revel.Result {
	return c.Render()
}
