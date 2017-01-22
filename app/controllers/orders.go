/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 19/01/2017    
* Time      : 19:33
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"github.com/icobani/RevelWebSite/app/models"
	"encoding/json"
)

type Orders struct {
	*revel.Controller
}

func (c Orders)Create() revel.Result {

	//new_cookie := &http.Cookie{Name: "REVEL_LANG", Value: "en",}
	//c.SetCookie(new_cookie)

	//new_cookie := &http.Cookie{Name: "i18n.cookie", Value: "en",}
	//c.SetCookie(new_cookie)

	fmt.Println(c.Request.Header.Get("Accept-Language"))
	// Localization information
	//c.RenderArgs["acceptLanguageHeader"] = c.Request.Header.Get("Accept-Language")
	//c.RenderArgs["acceptLanguageHeaderParsed"] = c.Request.AcceptLanguages.String()
	//c.RenderArgs["acceptLanguageHeaderMostQualified"] = c.Request.AcceptLanguages[1]
	//c.RenderArgs["controllerCurrentLocale"] = c.Request.Locale







	return c.Render()
}

func (c Orders)Create_POST() revel.Result {
	fmt.Println("Coming Create_POST")
	var order models.Order
	//Bind özelliği çok iyi
	//gelen parametreleri harika bir şekilde bizim modelimize uyguluyor.
	c.Params.Bind(&order, "order")

	c.Validation.Required(order.LastName).Message("Lütfen soyadınızı giriniz.")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
	}
	fmt.Printf("Has error : %v\n", c.Validation.HasErrors())

	fmt.Printf("Order Info : %v\n", order)
	return c.RenderTemplate("orders/create.html")

}

func (c Orders) Create_Api_POST() revel.Result {
	fmt.Println("api create order")
	var order models.Order

	dec := json.NewDecoder(c.Request.Body)
	dec.Decode(&order)
	c.Validation.Required(order.LastName)
	c.Validation.Length(order.PostCode, 4)

	fmt.Printf("Has error : %v\n", c.Validation.HasErrors())

	fmt.Printf("Order Info : %v\n", order)
	return c.RenderText(fmt.Sprintf("Order Info : %v\n", order))

}

func (c Orders) GetPayments(orderID int) revel.Result {
	fmt.Println("The order ID : ", orderID)

	return c.RenderTemplate("orders/create.html")
}