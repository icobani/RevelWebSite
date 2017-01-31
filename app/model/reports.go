package model

import (
	"fmt"
	"time"
	"github.com/icobani/RevelWebSite/app"
)

type Report struct {
	Id        int64 `json:"id" CaptionML:"enu=ID;trk=ID"`
	UserId    int64 `json:"-" CaptionML:"enu=User;trk=Kullanıcı"`
	CompanyId int64 `json:"-" CaptionML:"enu=Company;trk=Şirket"`

	Date        time.Time `json:"date" CaptionML:"enu=Report Date;trk=Rapor Tarihi"`
	Name        string    `json:"name" CaptionML:"enu=Report Name;trk=Rapor Adı"`
	Description string    `json:"description" CaptionML:"enu=Report Description;trk=Rapor Açıklaması"`

	Total         float32 `json:"total" sql:"-" CaptionML:"enu=Total Amount;trk=Toplam Tutar"`
	Status        string  `json:"status" CaptionML:"enu=Report Status;trk=Rapor Durumu"`
	ApprovalState int     `json:"-" CaptionML:"enu=Approval State;trk=Onay Durumu"`

	ExpensesCount         int `json:"expenses_count" sql:"-" CaptionML:"enu=Expenses Count;trk=Masraf Adedi"`
	ApprovedExpensesCount int `json:"approved_expenses_count" sql:"-" CaptionML:"enu=Approved Expenses Count;trk=Onaylanan Masraf Adedi"`
	RefusedExpensesCount  int `json:"refused_expenses_count" sql:"-" CaptionML:"enu=Refused Expenses Count;trk=Reddedilen Masraf Adedi"`
	PendingExpensesCount  int `json:"pending_expenses_count" sql:"-" CaptionML:"enu=Pending Expenses Count;trk=Reddedilen Masraf Adedi"`

	SubmittedBy string `json:"submitted_by" CaptionML:"enu=Submited By;trk=Gönderen"`
	ApprovedBy  string `json:"approved_by" CaptionML:"enu=Approved By;trk=Onaylayan"`

	UserApprovalMode *bool `json:"user_approval_mode,omitempty" sql:"-" CaptionML:"enu=Approval Mode;trk=Onay Modu"`
	UserEditMode     *bool `json:"user_edit_mode,omitempty" sql:"-" CaptionML:"enu=Edit Mode;trk=Düzenleme Modu"`

	User struct {
		Id         int64             `json:"id" CaptionML:"enu=ID;trk=ID"`
		Name       string            `json:"name" CaptionML:"enu=Name;trk=Adı"`
		LastName   string            `json:"last_name" CaptionML:"enu=SurName;trk=Soyadı"`
		Department *ReportDepartment `json:"department,omitempty" CaptionML:"enu=Department;trk=Departman"`
	} `json:"user"`

	CreatedAt time.Time `json:"created_at" CaptionML:"enu=Created At;trk=Oluşturulma Tarihi"`
	UpdatedAt time.Time `json:"-" CaptionML:"enu=Updated At;trk=Güncelleme Tarihi"`
}

func (this Report) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Report Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Report Table Created")
	app.MakeCaptionML(this)
}

type ReportDepartment struct {
	Id   int64  `json:"id" CaptionML:"enu=ID;trk=ID"`
	Name string `json:"name" CaptionML:"enu=Name;trk=Ad"`
}

func (this ReportDepartment) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("ReportDepartment Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("ReportDepartment Table Created")
	app.MakeCaptionML(this)
}


type ReportHistory struct {
	Id       int64 `json:"id" CaptionML:"enu=ID;trk=ID"`
	ReportId int64 `json:"-" CaptionML:"enu=Report ID;trk=Rapor ID"`
	UserId   int64 `json:"-" CaptionML:"enu=User;trk=Kullanıcı"`

	Action  string `json:"action" sql:"type:text;" CaptionML:"enu=Action;trk=İşlem"`
	Comment string `json:"comment" sql:"type:text;" CaptionML:"enu=Comment;trk=Açıklamalar"`

	User ReportHistoryUser `json:"user" sql:"-" CaptionML:"enu=Report History User;trk=Rapor Kullanıcısı"`

	CreatedAt time.Time `json:"created_at" CaptionML:"enu=Created At;trk=Oluşturma Tarihi"`
}

func (this ReportHistory) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("ReportHistory Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("ReportHistory Table Created")
	app.MakeCaptionML(this)
}

type ReportHistoryUser struct {
	Id              int64  `json:"id" CaptionML:"enu=ID;trk=ID"`
	Name            string `json:"name" CaptionML:"enu=Name;trk=Ad"`
	LastName        string `json:"last_name" CaptionML:"enu=Last Name;trk=Soyad"`
	ProfileImage    string `json:"profile_image" CaptionML:"enu=Profile Image;trk=Profil Resmi"`
	ProfileImageUrl string `json:"profile_image_url" CaptionML:"enu=Profile Image URL;trk=Profile Resmi URL"`
	Branch          struct {
		Id   int64  `json:"id" CaptionML:"enu=ID;trk=ID"`
		Name string `json:"name" CaptionML:"enu=Name;trk=Adı"`
	} `json:"branch"`
	Department struct {
		Id   int64  `json:"id" CaptionML:"enu=ID;trk=ID"`
		Name string `json:"name" CaptionML:"enu=Name;trk=Ad"`
	} `json:"department"`
}

func (this ReportHistoryUser) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("ReportHistoryUser Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("ReportHistoryUser Table Created")
	app.MakeCaptionML(this)
}
