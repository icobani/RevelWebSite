package model

import (
	"encoding/json"
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

// string, array(comma spared string), bool, time.Time(string)
type Preferences struct {
	Id     int64 `json:"-" CaptionML:"enu=;trk="`
	UserId int64 `json:"-" CaptionML:"enu=;trk="`

	AddExpenseType           string `json:"add_expence_type" sql:"type:varchar(25);default:'standard'" CaptionML:"enu=;trk="`
	MainMenuPinStatus        string `json:"main_menu_pin_status" sql:"type:varchar(25);default:'passive'" CaptionML:"enu=;trk="`
	DashboardChart           string `json:"dashboard_chart" sql:"type:varchar(100);default:'expenses-by-approval'" CaptionML:"enu=;trk="`
	ExpenseFiltersVisibility *bool  `json:"expense_filters_visibility" sql:"default:'1'" CaptionML:"enu=;trk="`
	ReportsFiltersVisibility *bool  `json:"reports_filters_visibility" sql:"default:'1'" CaptionML:"enu=;trk="`

	ExpenseCategory    string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	ExpenseProject     string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	ExpenseDepartment  string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	ExpenseTag         string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	ExpenseCurrency    string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	ExpensePaymentType string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	Expense            struct {
		Category    string `json:"category" CaptionML:"enu=;trk="`
		Project     string `json:"project" CaptionML:"enu=;trk="`
		Department  string `json:"department" CaptionML:"enu=;trk="`
		Tag         string `json:"tag" CaptionML:"enu=;trk="`
		Currency    string `json:"currency" CaptionML:"enu=;trk="`
		PaymentType string `json:"payment_type" CaptionML:"enu=;trk="`
	} `json:"expense"`

	ExpenseFilterMerchant    string     `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	ExpenseFilterScope       string     `json:"-" sql:"type:varchar(50)" CaptionML:"enu=;trk="`
	ExpenseFilterCategories  string     `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	ExpenseFilterProjects    string     `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	ExpenseFilterPaymentType string     `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	ExpenseFilterStartDate   *time.Time `json:"-" CaptionML:"enu=;trk="`
	ExpenseFilterEndDate     *time.Time `json:"-" CaptionML:"enu=;trk="`
	ExpenseFilters           struct {
		Merchant    string     `json:"merchant" CaptionML:"enu=;trk="`
		Scope       string     `json:"scope" CaptionML:"enu=;trk="`
		Categories  string     `json:"categories" CaptionML:"enu=;trk="`
		Projects    string     `json:"projects" CaptionML:"enu=;trk="`
		PaymentType string     `json:"payment_type" CaptionML:"enu=;trk="`
		StartDate   *time.Time `json:"start_date" CaptionML:"enu=;trk="`
		EndDate     *time.Time `json:"end_date" CaptionML:"enu=;trk="`
	} `json:"expense_filters"`

	ReportFilterStatus string `json:"-" sql:"type:varchar(50)" CaptionML:"enu=;trk="`
	ReportFilters      struct {
		Status string `json:"status" CaptionML:"enu=;trk="`
	} `json:"report_filters" CaptionML:"enu=;trk="`

	UserFilterName       string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	UserFilterGroup      string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	UserFilterDepartment string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	UserFilterBranch     string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	UserFilterStatus     string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	UserFilterOrder      string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	UserFilters          struct {
		Name       string `json:"name" CaptionML:"enu=;trk="`
		Group      string `json:"group" CaptionML:"enu=;trk="`
		Department string `json:"department" CaptionML:"enu=;trk="`
		Branch     string `json:"branch" CaptionML:"enu=;trk="`
		Status     string `json:"status" CaptionML:"enu=;trk="`
		Order      string `json:"order" CaptionML:"enu=;trk="`
	} `json:"user_filters" CaptionML:"enu=;trk="`

	ExpenseCategoryOrder   string `json:"-" sql:"type:text" CaptionML:"enu=;trk="`
	ExpenseCategoryFilters struct {
		Order string `json:"order"`
	} `json:"expense_category_filters" CaptionML:"enu=;trk="`
}

func (this Preferences) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Preferences Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Preferences Table Created")
	app.MakeCaptionML(this)
}

type JPreferences Preferences

func (p Preferences) MarshalJSON() ([]byte, error) {
	if p.Id == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(JPreferences(p))
}
