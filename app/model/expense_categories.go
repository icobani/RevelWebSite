package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type ExpenseCategory struct {
	Id                     int64             `json:"id" option:"id"`
	Name                   string            `json:"name" option:"value" sql:"type:varchar(250);"`
	Code                   string            `json:"code" sql:"type:varchar(250);"`
	CompanyId              int64             `json:"-"`
	BranchId               int64             `json:"-"`
	DepartmentId           int64             `json:"-"`
	ExpenseTypes           string            `json:"expense_types" sql:"type:varchar(250);"`
	ParentCategoryId       int64             `json:"-"`
	AccountingNumber       string            `json:"accounting_number" sql:"type:varchar(250);"`
	MaximumAmountPerReport float32           `json:"maximum_amount_per_report"`
	Branch                 *SBranch          `json:"branch,omitempty" sql:"-" ss:"-"`
	Parent                 *SExpenseCategory `json:"parent,omitempty" sql:"-" ss:"-"`
	IsPublic               bool              `json:"is_public"`
}

func (this ExpenseCategory) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("ExpenseCategory Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("ExpenseCategory Table Created")
}
