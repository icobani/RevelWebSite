package model

import (
	"fmt"
	"log"
	"strings"
	"time"
	"github.com/ilgooz/media"
	"github.com/icobani/RevelWebSite/app"
)

type User struct {
	Id int64 `json:"id" option:"id"`

	Email           string `json:"email" sql:"type:varchar(120)"`
	Name            string `json:"name" option:"value" sql:"type:varchar(120);"`
	LastName        string `json:"last_name" sql:"type:varchar(120);"`
	ProfileImage    string `json:"profile_image" sql:"type:varchar(250);`
	ProfileImageUrl string `json:"profile_image_url" ss:"profile_image_url"`
	Hash            string `json:"-" sql:"type:varchar(100);"`
	Password        string `json:"-" sql:"-" ss:"-"`
	IsActive        bool   `json:"is_active"`

	Verified             bool `json:"verified"`
	CompanyUserVerified  bool `json:"company_user_verified"`
	CompanyAdminVerified bool `json:"company_admin_verified"`

	BranchId                 int64  `json:"-"`
	DepartmentId             int64  `json:"-" sql:"index"`
	PositionId               int64  `json:"-"`
	RoleId                   int64  `json:"-"`
	IsCompanyAdmin           bool   `json:"is_company_admin"`
	DefaultProjectId         int64  `json:"-"`
	DefaultExpenseCategoryId int64  `json:"-"`
	JobTitle                 string `json:"job_title" sql:"type:varchar(250);"`
	AccountNumber            string `json:"-" sql:"type:varchar(250);"`
	Iban                     string `json:"iban" sql:"type:varchar(250);"`
	Language                 string `json:"language" sql:"type:varchar(250);"`
	MobilePhone              string `json:"mobile_phone" sql:"type:varchar(250);"`
	DistanceUnit             string `json:"distance_unit" sql:"type:varchar(50);"`
	DateFormat               string `json:"date_format" sql:"type:varchar(50);"`
	DecimalSeparator         string `json:"decimal_separator" sql:"type:varchar(50);"`
	ApprovalType             string `json:"-" sql:"type:varchar(250);"`
	ApprovalGroup            int64  `json:"-"`
	ApprovalUser1            int64  `json:"approval_user_1"`
	ApprovalUser2            int64  `json:"approval_user_2"`
	ApprovalUser3            int64  `json:"-"`
	ApprovalUser4            int64  `json:"-"`
	ApprovalUser5            int64  `json:"-"`
	ApprovalUser6            int64  `json:"-"`

	IsProfileCompleted bool `json:"is_profile_completed"`
	ShowQuickSettingUp bool `json:"show_quick_setting_up"`

	CompanyId int64    `json:"-" sql:"index"`
	Company   *Company `json:"company,omitempty" ss:"-"`

	//todo: web app de preferences'i user içinde sunabilirsin
	Preferences    *Preferences    `json:"preferences,omitempty" sql:"-" ss:"-"`
	QuickSettingUp *QuickSettingUp `json:"quick,omitempty" sql:"-" ss:"-"`

	Branch     *SBranch     `json:"branch,omitempty" ss:"-"`
	Projects   *[]SProject  `json:"projects,omitempty" gorm:"many2many:user_projects" ss:"-"`
	Groups     *[]SGroup    `json:"groups,omitempty" ss:"-" orm:"-"`
	Department *SDepartment `json:"department,omitempty" ss:"-" orm:"-"`

	//todo: remove omitempty and pointers
	Permissions *[]Permission `json:"permissions,omitempty" ss:"-" sql:"-"`
	Roles       *[]Role       `json:"roles,omitempty" ss:"-" sql:"-"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	//DeletedAt time.Time `json:"-"`
}

func (this User) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("User Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("User Table Created")
}


type UserRole struct {
	UserId int64
	RoleId int64
}

func (this UserRole) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("UserRole Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("UserRole Table Created")
}

type UserProject struct {
	UserId    int64 `sql:"index"`
	ProjectId int64 `sql:"index"`
}

func (this UserProject) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("UserProject Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("UserProject Table Created")
}


type UserGroup struct {
	UserId  int64 `sql:"index"`
	GroupId int64 `sql:"index"`
}

func (this UserGroup) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("UserGroup Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("UserGroup Table Created")
}

