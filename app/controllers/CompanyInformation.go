/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 31/01/2017    
* Time      : 20:48
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"github.com/icobani/RevelWebSite/app/model"
	"github.com/icobani/RevelWebSite/app"
	"github.com/icobani/RevelWebSite/app/modelViews"
)

type CompanyInformation struct {
	*revel.Controller
}

func (c CompanyInformation) Index() revel.Result {
	user := model.Me(c.Controller)

	Company := model.Company{}
	Company.Id = user.CompanyId
	app.DB.Find(&Company)
	if Company.StandardDateFormat == "" {
		Company.StandardDateFormat = "DayMonth"
	}

	// Standart Tarih Formatı Default Combo Değerleri..
	var StandartDateFormats = []modelViews.ComboItem{
		modelViews.ComboItem{
			Code:"DayMonth",
			Value:c.Message("CompanyInfo.DayMonthYear"),
		}, modelViews.ComboItem{
			Code:"MonthDay",
			Value:c.Message("CompanyInfo.MonthDayYear"),
		},
	}

	return c.Render(user, Company, StandartDateFormats)
}

func (c CompanyInformation) Post() revel.Result {
	var cInfo model.Company
	c.Params.Bind(&cInfo, "Company")
	app.DB.Save(&cInfo)
	return c.Redirect("/CompanyInformation")
}