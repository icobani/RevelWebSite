package model

import (
	"encoding/json"
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

// string, array(comma spared string), bool, time.Time(string)
type Preferences struct {
	Id     int64 `json:"-"`
	UserId int64 `json:"-"`

	AddExpenseType           string `json:"add_expence_type" sql:"type:varchar(25);default:'standard'"`
	MainMenuPinStatus        string `json:"main_menu_pin_status" sql:"type:varchar(25);default:'passive'"`
	DashboardChart           string `json:"dashboard_chart" sql:"type:varchar(100);default:'expenses-by-approval'"`
	ExpenseFiltersVisibility *bool  `json:"expense_filters_visibility" sql:"default:'1'"`
	ReportsFiltersVisibility *bool  `json:"reports_filters_visibility" sql:"default:'1'"`

	ExpenseCategory    string `json:"-" sql:"type:text"`
	ExpenseProject     string `json:"-" sql:"type:text"`
	ExpenseDepartment  string `json:"-" sql:"type:text"`
	ExpenseTag         string `json:"-" sql:"type:text"`
	ExpenseCurrency    string `json:"-" sql:"type:text"`
	ExpensePaymentType string `json:"-" sql:"type:text"`
	Expense            struct {
		Category    string `json:"category"`
		Project     string `json:"project"`
		Department  string `json:"department"`
		Tag         string `json:"tag"`
		Currency    string `json:"currency"`
		PaymentType string `json:"payment_type"`
	} `json:"expense"`

	ExpenseFilterMerchant    string     `json:"-" sql:"type:text`
	ExpenseFilterScope       string     `json:"-" sql:"type:varchar(50)`
	ExpenseFilterCategories  string     `json:"-" sql:"type:text`
	ExpenseFilterProjects    string     `json:"-" sql:"type:text`
	ExpenseFilterPaymentType string     `json:"-" sql:"type:text`
	ExpenseFilterStartDate   *time.Time `json:"-"`
	ExpenseFilterEndDate     *time.Time `json:"-"`
	ExpenseFilters           struct {
		Merchant    string     `json:"merchant"`
		Scope       string     `json:"scope"`
		Categories  string     `json:"categories"`
		Projects    string     `json:"projects"`
		PaymentType string     `json:"payment_type"`
		StartDate   *time.Time `json:"start_date"`
		EndDate     *time.Time `json:"end_date"`
	} `json:"expense_filters"`

	ReportFilterStatus string `json:"-" sql:"type:varchar(50)"`
	ReportFilters      struct {
		Status string `json:"status"`
	} `json:"report_filters"`

	UserFilterName       string `json:"-" sql:"type:text`
	UserFilterGroup      string `json:"-" sql:"type:text`
	UserFilterDepartment string `json:"-" sql:"type:text`
	UserFilterBranch     string `json:"-" sql:"type:text`
	UserFilterStatus     string `json:"-" sql:"type:text`
	UserFilterOrder      string `json:"-" sql:"type:text`
	UserFilters          struct {
		Name       string `json:"name"`
		Group      string `json:"group"`
		Department string `json:"department"`
		Branch     string `json:"branch"`
		Status     string `json:"status"`
		Order      string `json:"order"`
	} `json:"user_filters"`

	ExpenseCategoryOrder   string `json:"-" sql:"type:text`
	ExpenseCategoryFilters struct {
		Order string `json:"order"`
	} `json:"expense_category_filters"`
}

func (this Preferences) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Preferences Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Preferences Table Created")
}

type JPreferences Preferences

func (p Preferences) MarshalJSON() ([]byte, error) {
	if p.Id == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(JPreferences(p))
}
