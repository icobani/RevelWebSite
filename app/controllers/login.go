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
)

type Login struct {
	*revel.Controller
}

func (c Login) Index() revel.Result {
	fmt.Println("Login")
	return c.Render()
}