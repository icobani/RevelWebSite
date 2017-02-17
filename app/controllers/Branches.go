/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 13/02/2017    
* Time      : 15:53
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"github.com/icobani/RevelWebSite/app/model"
	"github.com/icobani/RevelWebSite/app"
	"fmt"
)

type Branches struct {
	*revel.Controller
}


// List Form
func (c Branches) Index() revel.Result {
	User := model.Me(c.Controller)

	var Branches []model.Branch
	app.DB.Where("Company_Id = ?", User.CompanyId).Select("Id, Code, Name, Currency_Code, Created_At").Order("code").Find(&Branches)

	return c.Render(User, Branches)
}

// Edit Form
func (c Branches) Edit(id int64) revel.Result {
	User := model.Me(c.Controller)
	var Branch model.Branch

	app.DB.First(&Branch, id)

	var Departments []model.Department
	app.DB.Where("Company_Id = ? and Branch_Id = ?", User.CompanyId, id).Select("Id, Code, Name").Find(&Departments)

	if Branch.Id != 0 {
		return c.Render(User, Branch, Departments)
	} else {
		return c.Redirect("/Branches")
	}

}

// Add Form
func (c Branches) Add() revel.Result {
	return c.Render()
}


// Edit
func (c Branches) Post() revel.Result {
	var branchForm model.Branch
	c.Params.Bind(&branchForm, "Branch")

	var branch model.Branch
	app.DB.First(&branch, branchForm.Id)


	branch.Code = branchForm.Code
	branch.Name = branchForm.Name
	branch.CurrencyCode = branchForm.CurrencyCode
	branch.Address = branchForm.Address
	branch.City = branchForm.City
	branch.PostalCode = branchForm.PostalCode
	branch.Country = branchForm.Country

	app.DB.Save(&branch)
	return c.Redirect("/Branches")
}


// Insert
func (c Branches) AddPost() revel.Result {
	User := model.Me(c.Controller)
	var branchForm model.Branch
	c.Params.Bind(&branchForm, "Branch")

	var branch model.Branch

	branch.Code = branchForm.Code
	branch.Name = branchForm.Name
	branch.CurrencyCode = branchForm.CurrencyCode
	branch.CompanyId = User.CompanyId
	branch.Address = branchForm.Address
	branch.City = branchForm.City
	branch.PostalCode = branchForm.PostalCode
	branch.Country = branchForm.Country

	if app.DB.NewRecord(branch) {
		app.DB.Create(&branch)
	}

	app.DB.NewRecord(branch)

	return c.Redirect(fmt.Sprintf("/Branches/Edit/?id=%d", branch.Id))
}

func (c Branches) Delete(id int64) revel.Result {
	var Branch model.Branch
	//var Department model.Department

	app.DB.First(&Branch, id)

	app.DB.Model(&Branch).Related(&Branch.Departments)

	if len(Branch.Departments) > 0 {
		c.Flash.Error("Bu şube ile ilişkili kayıtlar var silinemez")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Branches.Index)
	} else {
		app.DB.Delete(&Branch)
	}

	return c.Redirect(Branches.Index)

}

func (c Branches) checkUser4Branches() revel.Result {
	user := model.Me(c.Controller)
	if user.Id == 0 {
		return c.Redirect(Login.Index)
	}
	return nil
}