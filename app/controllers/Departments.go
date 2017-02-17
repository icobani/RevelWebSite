/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 17/02/2017    
* Time      : 18:47
* Developer : ibrahimcobani
*
*******/
package Departments

import (
	"github.com/revel/revel"
	"github.com/icobani/RevelWebSite/app/model"
)

type Departments struct {
	*revel.Controller
}

func (c Departments) Index() revel.Result  {
	User := model.User{}

	return c.Render(User)
}