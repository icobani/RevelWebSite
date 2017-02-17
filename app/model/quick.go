package model

import (
	"encoding/json"
	"strings"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type QuickSettingUp struct {
	Id                int64 `json:"-"`
	UserId            int64 `json:"-"`

	Step              int `json:"step"`

	Company           string `json:"company_name" CaptionML:"enu=Company Name;trk=Firma Adı"`
	Name              string `json:"name" CaptionML:"enu=Name;trk=Surname"`
	Email             string `json:"email" CaptionML:"enu=E-Mail;trk=E-Mail"`
	Country           string  ` CaptionML:"enu=Country;trk=Ülke"`
	PhoneNumber       string `json:"phone_number" CaptionML:"enu=Phone Number;trk=Telefon Numarası"`
	ExpenseCategories string `json:"expense_categories" CaptionML:"enu=Expense Categories;trk=Masraf Kategorileriniz"`
	Branches          string `json:"branches" CaptionML:"enu=Branches;trk=Şubeler"`
	Projects          string `json:"projects" CaptionML:"enu=Projects;trk=Projeler"`
	Departments       string `json:"departments" CaptionML:"enu=Departments;trk=Departmanlar"`
	Invites           string `json:"stuff_emails" CaptionML:"enu=Invites;trk=Davetler"`
	Password          string  ` CaptionML:"enu=Password;trk=Şifre"`
}

type JQuickSettingUp QuickSettingUp

func (q QuickSettingUp) MarshalJSON() ([]byte, error) {
	if q.Id == 0 {
		return []byte("null"), nil
	}

	return json.Marshal(JQuickSettingUp(q))
}

func (q *QuickSettingUp) GetExpenseCategories() []string {
	if q.ExpenseCategories == "" {
		return []string{"Default"}
	}

	q.Invites = strings.Replace(q.ExpenseCategories," ,", ",", -1)
	q.Invites = strings.Replace(q.ExpenseCategories,", ", ",", -1)
	q.Invites = strings.Replace(q.ExpenseCategories," , ", ",", -1)
	return strings.Split(q.ExpenseCategories, ",")
}

func (q *QuickSettingUp) GetBranches() []string {
	if q.Branches == "" {
		return []string{"Default"}
	}

	q.Invites = strings.Replace(q.Branches," ,", ",", -1)
	q.Invites = strings.Replace(q.Branches,", ", ",", -1)
	q.Invites = strings.Replace(q.Branches," , ", ",", -1)
	return strings.Split(q.Branches, ",")
}

func (q *QuickSettingUp) GetProjects() []string {
	if q.Projects == "" {
		return []string{"Default"}
	}

	q.Invites = strings.Replace(q.Projects," ,", ",", -1)
	q.Invites = strings.Replace(q.Projects,", ", ",", -1)
	q.Invites = strings.Replace(q.Projects," , ", ",", -1)
	return strings.Split(q.Projects, ",")
}

func (q *QuickSettingUp) GetDepartments() []string {
	if q.Departments == "" {
		return []string{"Default"}
	}

	q.Invites = strings.Replace(q.Departments," ,", ",", -1)
	q.Invites = strings.Replace(q.Departments,", ", ",", -1)
	q.Invites = strings.Replace(q.Departments," , ", ",", -1)
	return strings.Split(q.Departments, ",")
}

func (q *QuickSettingUp) GetInvites() []string {
	if q.Invites == "" {
		return []string{}
	}
	q.Invites = strings.Replace(q.Invites," ,", ",", -1)
	q.Invites = strings.Replace(q.Invites,", ", ",", -1)
	q.Invites = strings.Replace(q.Invites," , ", ",", -1)
	return strings.Split(q.Invites, ",")
}

func (quick *QuickSettingUp) Save() {
	fmt.Println("Şirket bilgileri oluşturuluyor.")
	var company = Company{}
	app.DB.Where("name = ?", quick.Company).First(&company)
	if company.Id == 0 {
		company.Name = quick.Company
		app.DB.Create(&company)

		fmt.Println("Şubeler açılıyor")
		var branches []string = quick.GetBranches()

		for _, branchname := range branches {
			fmt.Println("Branch :", branchname)
			var branch = Branch{}
			branch.Name = branchname
			branch.CompanyId = company.Id
			app.DB.Create(&branch)

			fmt.Println("Her bir şube için belirtilen departmanlar açılıyor.")
			var departments []string = quick.GetDepartments()

			for _, departmentname := range departments {
				var department = Department{}
				department.CompanyId = company.Id
				department.BranchId = branch.Id
				department.Name = departmentname
				app.DB.Create(&department)

			} // departments loop
		} // branches loop


		fmt.Println("Masraf kategorilerini ekleyelim.")
		var expensecats []string = quick.GetExpenseCategories()

		for _, expensecatname := range expensecats {
			var expensecategory = ExpenseCategory{}
			expensecategory.CompanyId = company.Id
			expensecategory.Name = expensecatname
			app.DB.Create(&expensecategory)
		}

		fmt.Println("Projeler oluşturulur.")
		var projects []string = quick.GetProjects()

		for _, projectname := range projects {
			var project = Project{}
			project.CompanyId = company.Id
			project.Name = projectname
			app.DB.Create(&project)
		}

		fmt.Println("Admin kullanısı ekleniyor.")
		var user = User{}
		user.CompanyId = company.Id
		user.Name = quick.Name
		user.Email = quick.Email
		user.Country = quick.Country
		user.MobilePhone = quick.PhoneNumber
		user.Password = quick.Password
		user.Hash = app.GetMD5Hash(quick.Password)
		user.IsCompanyAdmin = true
		app.DB.Create(&user)

		var invites []string = quick.GetInvites()
		for _, invitemailaddress := range invites {
			var user = User{}
			user.CompanyId = company.Id
			user.Email = invitemailaddress
			user.Verified = false
			app.DB.Create(&user)
		}

	} // company control

	fmt.Println("--------------------")
}