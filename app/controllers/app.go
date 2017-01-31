package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"net/http"
	"github.com/icobani/RevelWebSite/app/model"
	"github.com/icobani/RevelWebSite/app"
	"crypto/md5"
	"encoding/hex"
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

	//Şirket bilgileri oluşturuluyor.
	var company = model.Company{}
	app.DB.Where("name = ?", quick.Company).First(&company)
	if company.Id == 0 {
		company.Name = quick.Company
		app.DB.Create(&company)


		// Şubeler açılıyor
		var branches []string = quick.GetBranches()

		for _, branchname := range branches {
			var branch = model.Branch{}
			branch.Name = branchname
			branch.CompanyId = company.Id
			app.DB.Create(&branch)


			//Her bir şube için belirtilen departmanlar açılıyor.
			var departments []string = quick.GetDepartments()

			for _, departmentname := range departments {
				var department = model.Department{}
				department.CompanyId = company.Id
				department.BranchId = branch.Id
				department.Name = departmentname
				app.DB.Create(&department)

				// Her bir departman bazında Masraf kategorilerini ekleyelim.
				var expensecats []string = quick.GetExpenseCategories()

				for _, expensecatname := range expensecats {
					var expensecategory = model.ExpenseCategory{}
					expensecategory.CompanyId = company.Id
					expensecategory.BranchId = branch.Id
					expensecategory.DepartmentId = department.Id
					expensecategory.Name = expensecatname
					app.DB.Create(&expensecategory)
				}

				// Her bir departman bazında Projeler oluşturulur.
				var projects []string = quick.GetProjects()

				for _, projectname := range projects {
					var project = model.Project{}
					project.CompanyId = company.Id
					project.BranchId = branch.Id
					project.DepartmentId = department.Id
					project.Name = projectname
					app.DB.Create(&project)
				}

			} // departments loop
		} // branches loop

		//Admin kullanısı ekleniyor.
		var user = model.User{}
		user.CompanyId = company.Id
		user.Name = quick.Name
		user.Email = quick.Email
		user.Country = quick.Country
		user.MobilePhone = quick.PhoneNumber
		user.Password = quick.Password
		user.Hash = GetMD5Hash(quick.Password)
		user.IsCompanyAdmin = true
		app.DB.Create(&user)

		var invites []string = quick.GetInvites()
		for _, invitemailaddress := range invites {
			var user = model.User{}
			user.CompanyId = company.Id
			user.Email = invitemailaddress
			user.Verified = false
			app.DB.Create(&user)
		}

	} // company control

	fmt.Println("--------------------")

	return c.RenderTemplate("Login/Index.html")
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (c App) AutoMigrate() revel.Result {

	model.Company{}.CreateTable()







	/*
	model.Branch{}.CreateTable()
	model.CompanySettings{}.CreateTable()
	model.Currency{}.CreateTable()
	model.Department{}.CreateTable()
	model.UserDepartment{}.CreateTable()
	model.ExpenseCategory{}.CreateTable()
	model.Expense{}.CreateTable()
	model.Receipt{}.CreateTable()
	model.Import{}.CreateTable()
	model.ExpenseUser{}.CreateTable()
	model.Export{}.CreateTable()
	model.Country{}.CreateTable()
	model.City{}.CreateTable()
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
	*/
	return c.Render()
}





// Set for Turkish Lang
func (c App)SetLang_TR() revel.Result {
	new_cookie := &http.Cookie{Name: "REVEL_LANG", Value: "tr"}
	c.SetCookie(new_cookie)
	// Return the refered page
	return c.Redirect(c.Request.Referer())
}

//Set for English Lang
func (c App)SetLang_EN() revel.Result {
	new_cookie := &http.Cookie{Name: "REVEL_LANG", Value: "en"}
	c.SetCookie(new_cookie)
	// Return the refered page
	return c.Redirect(c.Request.Referer())
}