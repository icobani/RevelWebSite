package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

//note: members gibi alanlar bu strcutclar içnide tutulabilir, create edilirken for ile alınırlar yada update de
type Group struct {
	Id           int64    `json:"id" option:"id"`
	CompanyId    int64    `json:"-"`
	BranchId     int64    `json:"-"`
	DepartmentId int64    `json:"-"`
	Name         string   `json:"name" option:"value" sql:"type:varchar(250);"`
	Code         string   `json:"code" sql:"type:varchar(250);"`
	Email        string   `json:"email" sql:"type:varchar(250);"`
	Branch       *SBranch `json:"branch,omitempty" sql:"-" ss:"-"`
	Members      *[]SUser `json:"members,omitempty" sql:"-" ss:"-"`
}

func (this Group) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Group Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Group Table Created")
}
