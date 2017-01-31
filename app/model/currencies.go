package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Currency struct {
	Name     string `json:"name" sql:"type:varchar(50);" CaptionML:"enu=Name;trk=Ad"`
	Code     string `json:"code" option:"id" gorm:"primary_key" sql:"type:varchar(3);" CaptionML:"enu=Code;trk=Code"`
	FullName string `json:"-" option:"value" CaptionML:"enu=Full Name;trk=Tam Adı"`
}

func (this Currency) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Currency Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Currency Table Created")
	app.MakeCaptionML(this)
}

