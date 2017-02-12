package model

import (
	"fmt"
	"time"
	"github.com/icobani/RevelWebSite/app"
	"github.com/revel/revel"
	"strconv"
)

type User struct {
	Id                       int64 `json:"id" option:"id" CaptionML:"enu=ID;trk=ID"`

	Email                    string `json:"email" sql:"type:varchar(120)" CaptionML:"enu=Email;trk=Email"`
	Name                     string `json:"name" option:"value" sql:"type:varchar(120);" CaptionML:"enu=Name;trk=Ad"`
	LastName                 string `json:"last_name" sql:"type:varchar(120);" CaptionML:"enu=Last Name;trk=Soyadı"`
	ProfileImage             string `json:"profile_image" sql:"type:varchar(250);" CaptionML:"enu=Profile Image;trk=Profil Resmi"`
	ProfileImageUrl          string `json:"profile_image_url" ss:"profile_image_url" CaptionML:"enu=Profile Image URL;trk=Profil Resmi URL"`
	Hash                     string `json:"-" sql:"type:varchar(100);" CaptionML:"enu=Hash;trk=Şifre"`
	Password                 string `json:"-" sql:"-" ss:"-" CaptionML:"enu=Password;trk=Şifre"`
	IsActive                 bool   `json:"is_active" CaptionML:"enu=Active;trk=Aktif"`

	Verified                 bool `json:"verified" CaptionML:"enu=Verified;trk=Geçerli"`
	CompanyUserVerified      bool `json:"company_user_verified" CaptionML:"enu=Company User Verified;trk=Şirket Kullanıcısı Geçerli"`
	CompanyAdminVerified     bool `json:"company_admin_verified" CaptionML:"enu=Company Admin Verified;trk=Şirket Yetkilisi Geçerli"`

	BranchId                 int64  `json:"-" CaptionML:"enu=Branch;trk=Şube"`
	DepartmentId             int64  `json:"-" sql:"index" CaptionML:"enu=Department;trk=Departman"`
	PositionId               int64  `json:"-" CaptionML:"enu=Position;trk=Position"`
	RoleId                   int64  `json:"-" CaptionML:"enu=Role;trk=Role"`
	IsCompanyAdmin           bool   `json:"is_company_admin" CaptionML:"enu=Company Admin;trk=Şirket Yetkilisi"`
	DefaultProjectId         int64  `json:"-" CaptionML:"enu=Default Project;trk=Varsayılan Proje"`
	DefaultExpenseCategoryId int64  `json:"-" CaptionML:"enu=Default Expense Category ID;trk=Varsayılan Masraf Kategorisi"`
	JobTitle                 string `json:"job_title" sql:"type:varchar(250);" CaptionML:"enu=Job Title;trk=İş Ünvanı"`
	BankAccountNo            string `json:"-" sql:"type:varchar(250);" CaptionML:"enu=Bank Account No.;trk=Banka Hesap No."`
	Iban                     string `json:"iban" sql:"type:varchar(250);" CaptionML:"enu=IBAN;trk=IBAN"`
	Language                 string `json:"language" sql:"type:varchar(250);" CaptionML:"enu=Language;trk=Dil"`
	MobilePhone              string `json:"mobile_phone" sql:"type:varchar(250);" CaptionML:"enu=Mobile Phone;trk=Cep Telefonu"`
	Country                  string  ` CaptionML:"enu=Country;trk=Ülke"`
	DistanceUnit             string `json:"distance_unit" sql:"type:varchar(50);" CaptionML:"enu=Distance Unit;trk=Uzaklık Birimi"`
	DateFormat               string `json:"date_format" sql:"type:varchar(50);" CaptionML:"enu=Date Format;trk=Tarih Formatı"`
	DecimalSeparator         string `json:"decimal_separator" sql:"type:varchar(50);" CaptionML:"enu=Decimal Separator;trk=Ondalık Ayıracı"`
	ApprovalType             string `json:"-" sql:"type:varchar(250);" CaptionML:"enu=Approval Type;trk=Onay Tipi"`
	ApprovalGroup            int64  `json:"-" CaptionML:"enu=Approval Group;trk=Onay Grubu"`
	ApprovalUser1            int64  `json:"approval_user_1" CaptionML:"enu=Approval User 1;trk=Onaylayan 1"`
	ApprovalUser2            int64  `json:"approval_user_2" CaptionML:"enu=Approval User 2;trk=Onaylayan 2"`
	ApprovalUser3            int64  `json:"-" CaptionML:"enu=Approval User 3;trk=Onaylayan 3"`
	ApprovalUser4            int64  `json:"-" CaptionML:"enu=Approval User 4;trk=Onaylayan 4"`
	ApprovalUser5            int64  `json:"-" CaptionML:"enu=Approval User 5;trk=Onaylayan 5"`
	ApprovalUser6            int64  `json:"-" CaptionML:"enu=Approval User 6;trk=Onaylayan 6"`

	IsProfileCompleted       bool `json:"is_profile_completed" CaptionML:"enu=Profile Completed;trk=Profil Tamamlandı"`
	ShowQuickSettingUp       bool `json:"show_quick_setting_up" CaptionML:"enu=Show Quick Setting Up;trk=Hızlı Tanımlamayı Göster"`

	CompanyId                int64    `json:"-" sql:"index" CaptionML:"enu=Comapany;trk=Şirket"`
	Company                  *Company `json:"company,omitempty" ss:"-" CaptionML:"enu=Company;trk=Şirket"`

	//todo: web app de preferences'i user içinde sunabilirsin
	Preferences              *Preferences    `json:"preferences,omitempty" sql:"-" ss:"-" CaptionML:"enu=Preferences;trk=Tercihler"`
	QuickSettingUp           *QuickSettingUp `json:"quick,omitempty" sql:"-" ss:"-" CaptionML:"enu=Quick Setting Up;trk=Pratik Tanımlama"`

	Branch                   *SBranch     `json:"branch,omitempty" ss:"-" CaptionML:"enu=Branch;trk=Şube"`
	Projects                 *[]SProject  `json:"projects,omitempty" gorm:"many2many:user_projects" ss:"-" CaptionML:"enu=Project;trk=Proje"`
	Groups                   *[]SGroup    `json:"groups,omitempty" ss:"-" orm:"-" CaptionML:"enu=Groups;trk=Gruplar"`
	Department               *SDepartment `json:"department,omitempty" ss:"-" orm:"-" CaptionML:"enu=Department;trk=Departman"`

	//todo: remove omitempty and pointers
	Permissions              *[]Permission `json:"permissions,omitempty" ss:"-" sql:"-" CaptionML:"enu=Permissions;trk=Yetkiler"`
	Roles                    *[]Role       `json:"roles,omitempty" ss:"-" sql:"-" CaptionML:"enu=Roles;trk=Roller"`

	CreatedAt                time.Time `json:"-" CaptionML:"enu=Created At;trk=Oluşturma Tarihi"`
	UpdatedAt                time.Time `json:"-" CaptionML:"enu=Updated At;trk=Değiştirme Tarihi"`
	//DeletedAt time.Time `json:"-"`
}

func Me(c *revel.Controller) User {



	user := User{}

	uid := c.Session["uid"]

	if uid == ""{
		return User{}
	}


	i, _ := strconv.ParseInt(uid, 10, 64)
	user.Id = i
	app.DB.Find(&user)

	if user.Id == 0{
		return User{}
	}

	return user
}

func (this User) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("User Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("User Table Created")
	app.MakeCaptionML(this)
}

type UserRole struct {
	UserId int64 ` CaptionML:"enu=User;trk=Kullanıcı"`
	RoleId int64 ` CaptionML:"enu=Role;trk=Rol"`
}

func (this UserRole) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("UserRole Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("UserRole Table Created")
	app.MakeCaptionML(this)
}

type UserProject struct {
	UserId    int64 `sql:"index" CaptionML:"enu=User;trk=Kullanıcı"`
	ProjectId int64 `sql:"index" CaptionML:"enu=Project;trk=Proje"`
}

func (this UserProject) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("UserProject Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("UserProject Table Created")
	app.MakeCaptionML(this)
}

type UserGroup struct {
	UserId  int64 `sql:"index" CaptionML:"enu=User;trk=Kullanıcı"`
	GroupId int64 `sql:"index" CaptionML:"enu=Group;trk="`
}

func (this UserGroup) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("UserGroup Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("UserGroup Table Created")
	app.MakeCaptionML(this)
}

