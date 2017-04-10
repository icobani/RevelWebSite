/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 19/02/2017    
* Time      : 09:40
* Developer : ibrahimcobani
*
*******/
package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"github.com/icobani/RevelWebSite/app/model"
	"github.com/icobani/RevelWebSite/app"
	"github.com/icobani/RevelWebSite/app/modelViews"
	"strconv"
)

type ExpenseCategories struct {
	*revel.Controller
}


// List Form
func (c ExpenseCategories) Index() revel.Result {
	User := model.Me(c.Controller)
	fmt.Println("Hello")
	var ExpenseCategories []model.ExpenseCategory
	app.DB.Where("Company_Id = ?", User.CompanyId).Select("Id, Code, Name, department_id,branch_id, expense_types,main_category_id, accounting_number").Order("branch_id,code").Find(&ExpenseCategories)

	for i, _ := range ExpenseCategories {
		app.DB.Model(ExpenseCategories[i]).Related(&ExpenseCategories[i].Branch)
		app.DB.Model(ExpenseCategories[i]).Related(&ExpenseCategories[i].Department)
	}

	return c.Render(User, ExpenseCategories)
}

//Show Edit Form
func (c ExpenseCategories) Edit(id int64) revel.Result {

	User := model.Me(c.Controller)
	var ExpenseCategory model.ExpenseCategory

	//TODO: User Role mevzuları netleşince bunun gibi yerlerde çalışma yapmak lazım.
	app.DB.Where("Company_Id = ? and id = ?", User.CompanyId, id).First(&ExpenseCategory)

	if ExpenseCategory.Id != 0 {
		master := modelViews.ModelReferance{LogicalName:"expense_categories", Id:ExpenseCategory.Id, Name:ExpenseCategory.Name }


		// Şubeler
		Branch := model.Branch{Id:ExpenseCategory.BranchId}
		var BranchesComboItems []modelViews.ComboItem
		BranchesComboItems = Branch.GetComboValues(User, &master)

		// Departmanlar
		Department := model.Department{Id:ExpenseCategory.DepartmentId}
		var DepartmentsComboItems []modelViews.ComboItem
		DepartmentsComboItems = Department.GetComboValues(User, &master)

		// Kategoriler
		ECat := model.ExpenseCategory{Id:ExpenseCategory.MainCategoryId}
		var MainCategoriesComboItems []modelViews.ComboItem
		MainCategoriesComboItems = ECat.GetComboValues(User, &master)

		flash := map[string]string{
			"ExpenseCategory.BranchId": strconv.FormatInt(ExpenseCategory.BranchId, 10),
			"ExpenseCategory.DepartmentId":   strconv.FormatInt(ExpenseCategory.DepartmentId, 10),
			"ExpenseCategory.MainCategoryId":   strconv.FormatInt(ExpenseCategory.MainCategoryId, 10),
		}
		return c.Render(User, ExpenseCategory, BranchesComboItems, DepartmentsComboItems, MainCategoriesComboItems, flash)
	} else {
		return c.NotFound("Kayıt bulunamadı.")
	}
}


// Edit Post
func (c ExpenseCategories) Post() revel.Result {
	fmt.Println(c.Request.Form["ExpenseCategory.BranchId"])

	var expenseCategoryForm model.ExpenseCategory
	c.Params.Bind(&expenseCategoryForm, "ExpenseCategory")

	var expenseCategory model.ExpenseCategory
	app.DB.First(&expenseCategory, expenseCategoryForm.Id)

	expenseCategory.Code = expenseCategoryForm.Code
	expenseCategory.Name = expenseCategoryForm.Name
	expenseCategory.MainCategoryId = expenseCategoryForm.MainCategoryId

	expenseCategory.BranchId = expenseCategoryForm.BranchId
	expenseCategory.DepartmentId = expenseCategoryForm.DepartmentId
	expenseCategory.AccountingNumber = expenseCategoryForm.AccountingNumber

	expenseCategory.ExpenseTypes = expenseCategoryForm.ExpenseTypes
	expenseCategory.MaximumAmountPerReport = expenseCategoryForm.MaximumAmountPerReport
	expenseCategory.MaximumAmountPerMonth = expenseCategoryForm.MaximumAmountPerMonth

	expenseCategory.IsPublic = expenseCategoryForm.IsPublic
	expenseCategory.ReceiptsCheck = expenseCategoryForm.ReceiptsCheck
	expenseCategory.ProjectCheck = expenseCategoryForm.ProjectCheck
	expenseCategory.CommentsCheck = expenseCategoryForm.CommentsCheck

	app.DB.Save(&expenseCategory)
	return c.Redirect(c.Request.Referer())
}


// Add Form
func (c ExpenseCategories) Add(branchId int64) revel.Result {
	User := model.Me(c.Controller)
	var Department model.Department
	Department.BranchId = branchId
	fmt.Println(branchId)

	master := modelViews.ModelReferance{LogicalName:"expense_categories", Id:0, Name:"" }


	// Şubeler
	Branch := model.Branch{Id:branchId}
	var BranchesComboItems []modelViews.ComboItem
	BranchesComboItems = Branch.GetComboValues(User, &master)

	// Departmanlar

	var DepartmentsComboItems []modelViews.ComboItem
	DepartmentsComboItems = Department.GetComboValues(User, &master)

	// Ana Kategoriler
	ECat := model.ExpenseCategory{Id:0}
	var MainCategoriesComboItems []modelViews.ComboItem
	MainCategoriesComboItems = ECat.GetComboValues(User, &master)

	return c.Render(User, Department, BranchesComboItems, DepartmentsComboItems, MainCategoriesComboItems)
}




// Insert
func (c ExpenseCategories) AddPost() revel.Result {
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
	fmt.Println(c.Request.Referer())
	//return c.Redirect(fmt.Sprintf("/Branches/Edit/?id=%d", department.Id))
	return c.Redirect(fmt.Sprintf("/ExpenseCategories"))
	//return c.Redirect(c.Request.Referer())
}

func (c ExpenseCategories) Delete(id int64) revel.Result {
	var Department model.Department
	//var Department model.Department

	app.DB.First(&Department, id)

	//app.DB.Model(&Department).Related(&Department.ExpenseCategory)

	//if len(Department.ExpenseCategory) > 0 {
	//	c.Flash.Error("Bu şube ile ilişkili kayıtlar var silinemez")
	//	c.Validation.Keep()
	//	c.FlashParams()
	//	return c.Redirect(Branches.Index)
	//} else {
	app.DB.Delete(&Department)
	//}

	return c.Redirect(c.Request.Referer())

}