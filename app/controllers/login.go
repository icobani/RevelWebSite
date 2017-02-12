/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 17/01/2017    
* Time      : 21:15
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"github.com/icobani/RevelWebSite/app/modelViews"
	"github.com/icobani/RevelWebSite/app"
	"github.com/icobani/RevelWebSite/app/model"
	"strconv"
)

type Login struct {
	*revel.Controller
}

func (c Login) Index() revel.Result {
	fmt.Println("Login")
	return c.Render()
}

func (c Login) LogOff() revel.Result {
	user := model.Me(c.Controller)
	delete(c.Session, "uid")

	fmt.Println(user.Name)
	return c.Redirect("/Login")
}

func (c Login) Post() revel.Result {
	var loginVM modelViews.LoginViewModel
	c.Params.Bind(&loginVM, "loginVM")

	user := model.User{}
	user.Hash = app.GetMD5Hash(loginVM.Password)

	app.DB.Where("Email = ? AND Hash = ?", loginVM.Username, user.Hash).First(&user)

	if user.Id != 0 {
		c.Session["uid"] = strconv.FormatInt(user.Id, 16)
		return c.Redirect("HomePage")
	} else {
		fmt.Println("Error")

		c.Flash.Error(c.Message("Login.InvalidUserOrPassword"))
		c.Validation.Keep()
		c.FlashParams()
		return c.RenderTemplate("Login/Index.html")
	}

	return c.Redirect("HomePage")
}