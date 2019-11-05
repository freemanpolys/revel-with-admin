package controllers

import (
	"sms-engine/app/models"
	"sms-engine/app/routes"
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
	var label  models.GroupLabel
	label.Name = "Tesrrrrr"

	c.Txn.Debug().Save(&label)
	c.Log.Info("Fetching index")

	return c.Render()
}


