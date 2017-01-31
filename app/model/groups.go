package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

//note: members gibi alanlar bu strcutclar içnide tutulabilir, create edilirken for ile alınırlar yada update de
type Group struct {
	Id           int64    `json:"id" option:"id" CaptionML:"enu=ID;trk=ID"`
	CompanyId    int64    `json:"-" CaptionML:"enu=Company;trk=Şirket"`
	BranchId     int64    `json:"-" CaptionML:"enu=Branch;trk=Şube"`
	DepartmentId int64    `json:"-" CaptionML:"enu=Department;trk=Departman"`
	Name         string   `json:"name" option:"value" sql:"type:varchar(250);" CaptionML:"enu=Group Name;trk=Grup Adı"`
	Code         string   `json:"code" sql:"type:varchar(250);" CaptionML:"enu=Group Code;trk=Grup Kodu"`
	Email        string   `json:"email" sql:"type:varchar(250);" CaptionML:"enu=E-Mail;trk=E-Mail"`
	Branch       *SBranch `json:"branch,omitempty" sql:"-" ss:"-" CaptionML:"enu=Branch;trk=Şube"`
	Members      *[]SUser `json:"members,omitempty" sql:"-" ss:"-" CaptionML:"enu=Members;trk=Üyeler"`
}

func (this Group) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Group Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Group Table Created")
	app.MakeCaptionML(this)
}
