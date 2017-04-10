/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 17/02/2017    
* Time      : 18:47
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"github.com/icobani/RevelWebSite/app/model"
	"github.com/icobani/RevelWebSite/app"
	"fmt"
	"github.com/icobani/RevelWebSite/app/modelViews"
	"strconv"
)

type Departments struct {
	*revel.Controller
}

// List Form
func (c Departments) Index() revel.Result {
	User := model.Me(c.Controller)

	var Departments []model.Department
	app.DB.Where("Company_Id = ?", User.CompanyId).Select("Id, Code, Name, Created_At, branch_id").Order("branch_id,code").Find(&Departments)

	for i, _ := range Departments {
		app.DB.Model(Departments[i]).Related(&Departments[i].Branch)
	}

	return c.Render(User, Departments)
}

//Show Edit Form
func (c Departments) Edit(id int64) revel.Result {
	User := model.Me(c.Controller)
	var Department model.Department

	app.DB.First(&Department, id)


	if Department.Id != 0 {
		master := modelViews.ModelReferance{LogicalName:"departments", Id:Department.Id, Name:Department.Name }

		// Şubeler
		Branch := model.Branch{Id:Department.BranchId}
		var BranchesComboItems []modelViews.ComboItem
		BranchesComboItems = Branch.GetComboValues(User, &master)

		flash := map[string]string{
			"Department.BranchId": strconv.FormatInt(Department.BranchId, 10),
		}


		return c.Render(User, Department, BranchesComboItems,flash)
	} else {
		return c.NotFound("Kayıt bulunamadı.")
	}
}


// Edit Post
func (c Departments) Post() revel.Result {

	var departmentForm model.Department
	c.Params.Bind(&departmentForm, "Department")

	var department model.Department
	app.DB.First(&department, departmentForm.Id)


	fmt.Println(department.BranchId)
	department.BranchId = departmentForm.BranchId
	department.Code = departmentForm.Code
	department.Name = departmentForm.Name

	app.DB.Save(&department)
	return c.Redirect(c.Request.Referer())
}


// Add Form
func (c Departments) Add(branchId int64) revel.Result {
	User := model.Me(c.Controller)
	var Department model.Department
	Department.BranchId = branchId
	fmt.Println(branchId)
	return c.Render(User, Department)
}




// Insert
func (c Departments) AddPost() revel.Result {
	User := model.Me(c.Controller)
	var departmentForm model.Department
	c.Params.Bind(&departmentForm, "Department")

	fmt.Println(departmentForm)
	var department model.Department

	department.CompanyId = User.CompanyId
	department.BranchId = departmentForm.BranchId
	department.Code = departmentForm.Code
	department.Name = departmentForm.Name

	if app.DB.NewRecord(department) {
		app.DB.Create(&department)
	}

	app.DB.NewRecord(department)

	return c.Redirect(fmt.Sprintf("/Branches/Edit/?id=%d", department.Id))
}

func (c Departments) Delete(id int64) revel.Result {
	var Department model.Department
	//var Department model.Department

	app.DB.First(&Department, id)

	//app.DB.Model(&Department).Related(&Department.Departments)

	//if len(Department.Departments) > 0 {
	//	c.Flash.Error("Bu şube ile ilişkili kayıtlar var silinemez")
	//	c.Validation.Keep()
	//	c.FlashParams()
	//	return c.Redirect(Branches.Index)
	//} else {
	app.DB.Delete(&Department)
	//}

	return c.Redirect(c.Request.Referer())

}

func (c Departments) Api_GetDepartments(branchid int64) revel.Result {
	User := model.Me(c.Controller)

	Depa := model.Department{}

	fmt.Println("BranchId",branchid)

	response := JsonResponse{}
	master := modelViews.ModelReferance{LogicalName:"branch",Id: branchid, Name:"" }
	response.Data = Depa.GetComboValues(User,&master)

	return c.RenderJson(response)

}