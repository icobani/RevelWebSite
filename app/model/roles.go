package model

import (
	"encoding/json"

	"github.com/guregu/null"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Role struct {
	Id   null.Int    `json:"id" option:"id" CaptionML:"enu=ID;trk=ID"`
	Name null.String `json:"name" option:"value" sql:"type:varchar(150)" CaptionML:"enu=Name;trk=Ad"`
}

func (this Role) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Role Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Role Table Created")
	app.MakeCaptionML(this)
}

type Permission struct {
	Id   null.Int      `CaptionML:"enu=ID;trk=ID"`
	Name null.String  `CaptionML:"enu=Name;trk=Ad"`
}

func (this Permission) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Permission Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Permission Table Created")
	app.MakeCaptionML(this)
}

func (p Permission) MarshalJSON() ([]byte, error) {
	if p.Id.Int64 == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(p.Name)
}

type RolePermission struct {
	RoleId       int64 ` CaptionML:"enu=ID;trk=ID"`
	PermissionId int64 ` CaptionML:"enu=Permission ID;trk=Yetki"`
}

func (this RolePermission) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("RolePermission Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("RolePermission Table Created")
	app.MakeCaptionML(this)
}