package controllers

import "github.com/revel/revel"

func init() {
	revel.InterceptMethod(Application.AddUser, revel.BEFORE)
	revel.InterceptMethod(Application.AutoMigrate, revel.BEFORE)
	revel.InterceptMethod(Admin.checkUser, revel.BEFORE)
}
