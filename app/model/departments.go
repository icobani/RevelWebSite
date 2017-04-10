package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
	"time"
	"github.com/icobani/RevelWebSite/app/modelViews"
)
//TODO: bu tabloda created_at modified_at alanları eksik.
//TODO: Entity Referance diye bir struct oluşturup bunun içinde id ve Name alanını referans tablolarına ekleyelim.

type Department struct {
	Id        int64    `json:"id" option:"id" CaptionML:"enu=ID;trk=ID"`
	CompanyId int64    `json:"-" CaptionML:"enu=Company ID;trk=Şirket ID"`
	BranchId  int64    `json:"-" CaptionML:"enu=Branch;trk=Şube"`
	Name      string   `json:"name" option:"value" sql:"type:varchar(250);" CaptionML:"enu=Name;trk=Adı"`
	Code      string   `json:"code" sql:"type:varchar(50);" CaptionML:"enu=Code;trk=Code"`
	Branch    Branch `json:"branch,omitempty" gorm:"foreignkey:BranchId"  sql:"-" ss:"-" CaptionML:"enu=Branch;trk=Şube"`
	Members   *[]SUser `json:"members,omitempty" sql:"-" ss:"-" CaptionML:"enu=Members;trk=Personeller"`

	CreatedAt time.Time `json:"-" CaptionML:"enu=Created At;trk= Oluşturulma Tarihi"`
	UpdatedAt time.Time `json:"-" CaptionML:"enu=Updated At;trk=Güncelleme Tarihi"`
}


// Combolara çıkacak olan değerler burada hazırlanıyor.
func (this Department) GetComboValues(user User, master *modelViews.ModelReferance) []modelViews.ComboItem {
	var Departments []Department
	var ComboItems []modelViews.ComboItem

	switch master.LogicalName {
	case "expense_categories":
		eCat := ExpenseCategory{Id:master.Id}
		app.DB.First(&eCat)
		app.DB.Where("Company_Id = ? and Branch_Id = ?", user.CompanyId, eCat.BranchId).Select("Id, Name").Order("code").Find(&Departments)
		break
	case "branch":
		app.DB.Where("Company_Id = ? and Branch_Id = ?", user.CompanyId, master.Id).Select("Id, Name").Order("code").Find(&Departments)
		break
	}

	ComboItems = append(ComboItems, modelViews.ComboItem{Id:0, Value:"Seçiniz", Selected:this.Id == 0})
	for _, item := range Departments {
		ComboItems = append(ComboItems, modelViews.ComboItem{Id:item.Id, Value:item.Name, Selected:item.Id == this.Id})
	}
	return ComboItems
}

func (this Department) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Department Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Department Table Created")
	app.MakeCaptionML(this)
}

type UserDepartment struct {
	UserId       int64 `sql:"index" CaptionML:"enu=UserId;trk=Kullanıcı ID"`
	DepartmentId int64 `sql:"index" CaptionML:"enu=DepartmentID;trk=Departman ID"`
}

func (this UserDepartment) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("UserDepartment Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("UserDepartment Table Created")
	app.MakeCaptionML(this)
}