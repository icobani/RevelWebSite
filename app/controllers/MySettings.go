/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 11/02/2017    
* Time      : 21:21
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"github.com/icobani/RevelWebSite/app/model"
	"fmt"
	"github.com/icobani/RevelWebSite/app/BusinessLibs"
	"github.com/icobani/RevelWebSite/app"
)

type MySettings struct {
	*revel.Controller
}

func (c MySettings)Index() revel.Result {
	User := model.Me(c.Controller)
	fmt.Println(User.Name)

	if User.LastName == "" {
		var hparser = BusinessLibs.HumanParser{}
		hparser.Set(User.Name)

		fmt.Println(hparser.Name)
		fmt.Println(hparser.LastName)
		User.Name = hparser.Name
		User.LastName = hparser.LastName
		app.DB.Save(&User)
	}

	return c.Render(User)
}

func (c MySettings) Post() revel.Result {
	var UserForm model.User
	c.Params.Bind(&UserForm, "User")
	//TODO: Bu şekilde kaydettiğimiz zaman formda gelen bilgiler db içindeki değerleri eziyor. bununun için belki form viewleri ayrı tasarlamak lazım.
	app.DB.Save(&UserForm)
	return c.Redirect("/MySettings")
}

func (c MySettings) checkUser4MySettings() revel.Result {
	fmt.Println("Hello from intercept method")
	user := model.Me(c.Controller)
	if user.Id == 0 {
		return c.Redirect(Login.Index)
	}
	return nil
}