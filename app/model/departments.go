package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Department struct {
	Id        int64    `json:"id" option:"id" CaptionML:"enu=ID;trk=ID"`
	CompanyId int64    `json:"-" CaptionML:"enu=Company ID;trk=Şirket ID"`
	BranchId  int64    `json:"-" CaptionML:"enu=Branch ID;trk=Şirket ID"`
	Name      string   `json:"name" option:"value" sql:"type:varchar(250);" CaptionML:"enu=Name;trk=Adı"`
	Code      string   `json:"code" sql:"type:varchar(50);" CaptionML:"enu=Code;trk=Code"`
	Branch    *SBranch `json:"branch,omitempty" sql:"-" ss:"-" CaptionML:"enu=Branch;trk=Şube"`
	Members   *[]SUser `json:"members,omitempty" sql:"-" ss:"-" CaptionML:"enu=Members;trk=Personeller"`
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