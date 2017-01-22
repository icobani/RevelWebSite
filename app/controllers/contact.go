/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 15/01/2017    
* Time      : 23:13
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"fmt"
)

type Contact struct {
	*revel.Controller
}

func (c Contact) Index() revel.Result {
	fmt.Println("Incoming Index")
	return c.Render()
}