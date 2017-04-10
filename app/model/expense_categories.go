package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
	"github.com/icobani/RevelWebSite/app/modelViews"
)

type ExpenseCategory struct {
	Id                     int64             `json:"id" option:"id" CaptionML:"enu=ID;trk=ID"`
	Name                   string            `json:"name" option:"value" sql:"type:varchar(250);" CaptionML:"enu=Category Name;trk=Kategori Adı"`
	Code                   string            `json:"code" sql:"type:varchar(250);" CaptionML:"enu=Category Code;trk=Kategori Kod"`
	CompanyId              int64             `json:"-" CaptionML:"enu=Company;trk=Şirket"`
	BranchId               int64             `json:"-" CaptionML:"enu=Branch;trk=Şube"`
	DepartmentId           int64             `json:"-" CaptionML:"enu=Department;trk=Departman"`
	ExpenseTypes           string            `json:"expense_types" sql:"type:varchar(250);" CaptionML:"enu=Expense Type;trk=Masraf Tipi"`
	MainCategoryId         int64             `json:"-" CaptionML:"enu=Main Category;trk=Ana Kategori"`
	AccountingNumber       string            `json:"accounting_number" sql:"type:varchar(250);" CaptionML:"enu=Accounting Number;trk=Hesap Kodu"`
	MaximumAmountPerReport float32           `json:"maximum_amount_per_report" CaptionML:"enu=Maximum Amount Per Report;trk=Maksimum Tutar (Rapor Bazında)"`
	MaximumAmountPerMonth  float32           `json:"maximum_amount_per_month" CaptionML:"enu=Maximum Amount Per Month;trk=Maksimum Tutar (Aylık)"`
	Branch                 Branch           `json:"branch,omitempty" sql:"-" ss:"-" CaptionML:"enu=Branch;trk=Şube"`
	Department             Department       `json:"department,omitempty" sql:"-" ss:"-" CaptionML:"enu=Department;trk=Departman"`
	Parent                 *ExpenseCategory  `json:"parent,omitempty" sql:"-" ss:"-" CaptionML:"enu=Parent Category;trk=Ana Kategori"`
	IsPublic               bool              `json:"is_public" CaptionML:"enu=Public;trk=Genel"`
	ReceiptsCheck          bool              `CaptionML:"enu=Receipts Check;trk=Fiş Eklendi mi ?"`
	ProjectCheck           bool              `CaptionML:"enu=Project Check;trk=Proje Eklendi mi ?"`
	CommentsCheck          bool              `CaptionML:"enu=Comments Check;trk=Yorum Eklendi mi ?"`
}

func (this ExpenseCategory) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("ExpenseCategory Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("ExpenseCategory Table Created")
	app.MakeCaptionML(this)
}


func (this ExpenseCategory) GetComboValues(user User, master *modelViews.ModelReferance) []modelViews.ComboItem {
	var ExpenseCategories []ExpenseCategory
	var ComboItems []modelViews.ComboItem

	switch master.LogicalName {
	case "expense_categories":
		eCat := ExpenseCategory{Id:master.Id}
		if master.Id != 0 {
			app.DB.First(&eCat)
		}


		app.DB.
			Where(
				`
				Company_Id = ? and
				(Branch_Id = ? OR Branch_Id = 0) and
				(Department_Id= ? or Department_ID = 0) and
				Id <> ?

				`,
				user.CompanyId,
				eCat.BranchId,
				eCat.DepartmentId,
				eCat.Id).
			Select("Id, Name").
			Order("code").
			Find(&ExpenseCategories)

		break
	}



	ComboItems = append(ComboItems, modelViews.ComboItem{Id:0, Value:"Seçiniz", Selected:this.Id == 0})
	for _, item := range ExpenseCategories {
		ComboItems = append(ComboItems, modelViews.ComboItem{Id:item.Id, Value:item.Name, Selected:item.Id == this.Id})
	}
	return ComboItems
}