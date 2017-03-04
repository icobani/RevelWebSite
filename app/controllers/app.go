package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"net/http"
	"github.com/icobani/RevelWebSite/app/model"
	"github.com/icobani/RevelWebSite/app"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	fmt.Println("Incoming Index")
	return c.Render()
}

func (c App) Pricing() revel.Result {
	fmt.Println("Incoming Pricing")
	return c.Render()
}

func (c App) Tour() revel.Result {
	return c.Render()
}

func (c App) Us() revel.Result {
	return c.Render()
}

func (c App) Why() revel.Result {
	fmt.Println("Why....")
	return c.Render()
}

func (c App) Pet() revel.Result {
	fmt.Println("Pet....")
	return c.Render()
}

func (c App) QuickRegister() revel.Result {
	fmt.Println("QuickRegister....")
	return c.Render()
}

func (c App) QuickRegister_POST() revel.Result {
	fmt.Println("QuickRegister_POST")
	var quick model.QuickSettingUp
	c.Params.Bind(&quick, "quick")

	fmt.Println("--------------------")
	fmt.Println(quick)

	quick.Save()

	return c.RenderTemplate("Login/Index.html")
}

func (c App) AutoMigrate() revel.Result {

	model.Company{}.CreateTable()
	model.Branch{}.CreateTable()
	model.Currency{}.CreateTable()
	model.Department{}.CreateTable()
	model.UserDepartment{}.CreateTable()
	model.ExpenseCategory{}.CreateTable()
	model.Expense{}.CreateTable()
	model.Receipt{}.CreateTable()
	model.Import{}.CreateTable()
	model.ExpenseUser{}.CreateTable()
	model.Export{}.CreateTable()
	model.Group{}.CreateTable()
	model.MileageCategory{}.CreateTable()
	model.PaymentType{}.CreateTable()
	model.Preferences{}.CreateTable()
	model.Project{}.CreateTable()
	model.Report{}.CreateTable()
	model.ReportDepartment{}.CreateTable()
	model.ReportHistory{}.CreateTable()
	model.ReportHistoryUser{}.CreateTable()
	model.Role{}.CreateTable()
	model.Permission{}.CreateTable()
	model.RolePermission{}.CreateTable()
	model.TokenStorage{}.CreateTable()
	model.User{}.CreateTable()
	model.UserRole{}.CreateTable()
	model.UserProject{}.CreateTable()
	model.UserGroup{}.CreateTable()

	var quick model.QuickSettingUp
	// First Demo Referance Records
	quick.Company = "B1 Yönetim Sistemleri"
	quick.Name = "Ibrahim ÇOBANİ"
	quick.Email = "ibrahim@imaconsult.com"
	quick.Country = "90"
	quick.PhoneNumber = "532 540 1194"
	quick.ExpenseCategories = "Yemek,Taksi,Otopark,Otel,Diğer"
	quick.Branches = "Esentepe,Maslak"
	quick.Projects = "B1,DO&CO,Saint Gobain Rigips,Lifewath,Manfred Geldbach,LMG,Mardav,Parge,Saygın,Tora Petrol,Sodexo"
	quick.Departments = "Yazılım,Danışmanlık Servisleri,Dijital"
	quick.Invites = "gokhan@imaconsult.com,bahadir@b1.com.tr,basak@imaconsult.com,aycan@b1.com.tr"
	quick.Password = "azura777"

	quick.Save()
	return c.Redirect("/Login")
}





// Set for Turkish Lang
func (c App)SetLang_TR() revel.Result {
	new_cookie := &http.Cookie{Name: "REVEL_LANG", Value: "tr"}
	c.SetCookie(new_cookie)

	User := model.Me(c.Controller)
	if User.Id != 0 {
		User.Language = "tr";
		app.DB.Save(&User)
	}

	// Return the refered page
	return c.Redirect(c.Request.Referer())
}

//Set for English Lang
func (c App)SetLang_EN() revel.Result {

	new_cookie := &http.Cookie{Name: "REVEL_LANG", Value: "en"}
	c.SetCookie(new_cookie)

	User := model.Me(c.Controller)
	if User.Id != 0 {
		User.Language = "tr";
		app.DB.Save(&User)
	}

	// Return the refered page
	return c.Redirect(c.Request.Referer())
}