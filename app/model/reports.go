package model

import (
	"fmt"
	"time"
	"github.com/icobani/RevelWebSite/app"
)

type Report struct {
	Id        int64 `json:"id"`
	UserId    int64 `json:"-"`
	CompanyId int64 `json:"-"`

	Date        time.Time `json:"date"`
	Name        string    `json:"name"`
	Description string    `json:"description"`

	Total         float32 `json:"total" sql:"-"`
	Status        string  `json:"status"`
	ApprovalState int     `json:"-"`

	ExpensesCount         int `json:"expenses_count" sql:"-"`
	ApprovedExpensesCount int `json:"approved_expenses_count" sql:"-"`
	RefusedExpensesCount  int `json:"refused_expenses_count" sql:"-"`
	PendingExpensesCount  int `json:"pending_expenses_count" sql:"-"`

	SubmittedBy string `json:"submitted_by"`
	ApprovedBy  string `json:"approved_by"`

	UserApprovalMode *bool `json:"user_approval_mode,omitempty" sql:"-"`
	UserEditMode     *bool `json:"user_edit_mode,omitempty" sql:"-"`

	User struct {
		Id         int64             `json:"id"`
		Name       string            `json:"name"`
		LastName   string            `json:"last_name"`
		Department *ReportDepartment `json:"department,omitempty"`
	} `json:"user"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

func (this Report) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Report Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Report Table Created")
}

type ReportDepartment struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (this ReportDepartment) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("ReportDepartment Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("ReportDepartment Table Created")
}


type ReportHistory struct {
	Id       int64 `json:"id"`
	ReportId int64 `json:"-"`
	UserId   int64 `json:"-"`

	Action  string `json:"action" sql:"type:text;"`
	Comment string `json:"comment" sql:"type:text;"`

	User ReportHistoryUser `json:"user" sql"-"`

	CreatedAt time.Time `json:"created_at"`
}

func (this ReportHistory) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("ReportHistory Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("ReportHistory Table Created")
}

type ReportHistoryUser struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	LastName        string `json:"last_name"`
	ProfileImage    string `json:"profile_image"`
	ProfileImageUrl string `json:"profile_image_url"`
	Branch          struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"branch"`
	Department struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"department"`
}

func (this ReportHistoryUser) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("ReportHistoryUser Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("ReportHistoryUser Table Created")
}
