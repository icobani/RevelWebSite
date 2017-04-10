/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/03/2017    
* Time      : 00:18
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"github.com/icobani/RevelWebSite/app/model"
	"fmt"
)

type Currencies struct {
	*revel.Controller
}

func (c Currencies) Init() revel.Result {
	var cur = model.Currency{}
	fmt.Println("Oluşturuluyor.")
	cur.InitTable()
	fmt.Println("Döviz kurları oluşturuldu")
	return c.RenderText("Döviz kurları oluşturuldu..")
}

