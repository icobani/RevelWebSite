/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/02/2017    
* Time      : 12:16
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"fmt"
)

func init() {
	fmt.Println("Hello from controller init method")
	revel.InterceptMethod(HomePage.checkUser, revel.BEFORE)
}
