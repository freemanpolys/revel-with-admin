package controllers

import (
	"github.com/revel/revel"

	"github.com/revel/examples/booking/app/routes"
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


