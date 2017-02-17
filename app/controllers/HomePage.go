/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 31/01/2017    
* Time      : 12:44
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"github.com/icobani/RevelWebSite/app/model"
)

type HomePage struct {
	*revel.Controller
}

func (c HomePage) Index() revel.Result {
	User := model.Me(c.Controller)
	return c.Render(User)
}

func (c HomePage) checkUser() revel.Result {
	fmt.Println("Hello from intercept method")
	user := model.Me(c.Controller)
	if user.Id == 0 {
		return c.Redirect(Login.Index)
	}
	return nil
}