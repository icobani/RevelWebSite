package model

import (
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

//todo sqlstruct'da struct ve slice ve bunların ptr'ları otomatik ignore edilsin
type Project struct {
	Id           int64      `json:"id" option:"id" CaptionML:"enu=ID;trk=ID"`
	CompanyId    int64      `json:"-" CaptionML:"enu=Company;trk=Şirket"`
	BranchId     int64      `json:"-" CaptionML:"enu=Branch;trk=Şube"`
	DepartmentId int64      `json:"-" CaptionML:"enu=Department;trk=Departman"`
	Name         string     `json:"name" option:"value" sql:"type:varchar(250);" CaptionML:"enu=Project Name;trk=Proje Adı"`
	Code         string     `json:"code" sql:"type:varchar(250);" CaptionML:"enu=Project Code;trk=Proje Kodu"`
	StartDate    *time.Time `json:"start_date" CaptionML:"enu=Start Date;trk=Başlangıç Tarihi"`
	EndDate      *time.Time `json:"end_date" CaptionML:"enu=End Date;trk=Bitiş Tarihi"`
	IsActive     bool       `json:"is_active" CaptionML:"enu=Active;trk=Aktif"`
	Branch       *SBranch   `json:"branch" ss:"-" CaptionML:"enu=Branch;trk=Şube"`
}

func (this Project) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Project Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Project Table Created")
	app.MakeCaptionML(this)
}
