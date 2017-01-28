package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Department struct {
	Id        int64    `json:"id" option:"id"`
	CompanyId int64    `json:"-"`
	BranchId  int64    `json:"-"`
	Name      string   `json:"name" option:"value" sql:"type:varchar(250);"`
	Code      string   `json:"code" sql:"type:varchar(50);"`
	Branch    *SBranch `json:"branch,omitempty" sql:"-" ss:"-"`
	Members   *[]SUser `json:"members,omitempty" sql:"-" ss:"-"`
}

func (this Department) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Department Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Department Table Created")
}


type UserDepartment struct {
	UserId       int64 `sql:"index"`
	DepartmentId int64 `sql:"index"`
}

func (this UserDepartment) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("UserDepartment Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("UserDepartment Table Created")
}