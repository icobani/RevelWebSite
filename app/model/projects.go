package model

import (
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

//todo sqlstruct'da struct ve slice ve bunların ptr'ları otomatik ignore edilsin
type Project struct {
	Id           int64      `json:"id" option:"id"`
	CompanyId    int64      `json:"-"`
	BranchId     int64      `json:"-"`
	DepartmentId int64      `json:"-"`
	Name         string     `json:"name" option:"value" sql:"type:varchar(250);"`
	Code         string     `json:"code" sql:"type:varchar(250);"`
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	IsActive     bool       `json:"is_active"`
	Branch       *SBranch   `json:"branch" ss:"-"`
}

func (this Project) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Project Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Project Table Created")
}
