/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 19/01/2017    
* Time      : 19:50
* Developer : ibrahimcobani
*
*******/
package controllers

import "github.com/revel/revel"

type Accounts struct {
	*revel.Controller
}

func (c Accounts)Create() revel.Result  {
	return c.Render()
}
