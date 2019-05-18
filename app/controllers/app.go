package controllers

import (
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"revel-with-admin/app/models"
	"revel-with-admin/app/routes"
)

type Application struct {
	gormc.TxnController
}

func (c Application) setCurrentUser () revel.Result {
	c.ViewArgs["currentUrl"] = c.Action
	return nil
}

func (c Application) AutoMigrate() revel.Result {

	c.Txn.AutoMigrate(&models.User{})
	bcryptPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("demo"), bcrypt.DefaultCost)
	adminUser := &models.User{UserId:0,Name:"Admin user",Username: "admin", Password:"demo", HashedPassword:bcryptPassword}

	if c.Txn.Where("username = ?", adminUser.Username).First(adminUser).RecordNotFound()  {
		c.Txn.Save(adminUser)
	}

	return nil
}
func (c Application) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.ViewArgs["user"] = user
	}
	return nil
}

func (c Application) connected() *models.User {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username.(string))
	}
	return nil
}

func (c Application) getUser(username string) (user *models.User) {
	user = &models.User{}
	_,  err := c.Session.GetInto("fulluser", user, false)
	if user.Username == username {
		return user
	}

	if c.Txn.Where("username = ?", username).First(user).RecordNotFound() {
		c.Log.Error("Failed to find user", "user", username, "error",err)

	}

	c.Session["fulluser"] = user
	return
}

func (c Application) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.Admin.Index())
	}
	c.Flash.Error("Please log in first")
	return c.Render()
}

func (c Application) Register() revel.Result {
	return c.Render()
}

func (c Application) SaveUser(user models.User, verifyPassword string) revel.Result {
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).
		MessageKey("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Application.Register())
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	err := c.Txn.Save(&user)
	if err != nil {
		panic(err)
	}

	c.Session["user"] = user.Username
	c.Flash.Success("Welcome, " + user.Name)
	return c.Redirect(routes.Admin.Index())
}

func (c Application) Login(username, password string, remember bool) revel.Result {
	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.Admin.Index())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.Application.Index())
}

func (c Application) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Application.Index())
}
func (c Application) About() revel.Result {
    c.ViewArgs["Msg"]="Revel Speaks"
	return c.Render()
}
