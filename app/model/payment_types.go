package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type PaymentType struct {
	Id   int64  `option:"id" CaptionML:"enu=ID;trk=ID"`
	Name string `option:"value" sql:"type:varchar(100);" CaptionML:"enu=Payment Type;trk=Ödeme Tipi"`
}
func (this PaymentType) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("PaymentType Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("PaymentType Table Created")
	app.MakeCaptionML(this)
}