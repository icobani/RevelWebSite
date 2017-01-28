package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type PaymentType struct {
	Id   int64  `option:"id"`
	Name string `option:"value" sql:"type:varchar(100);"`
}
func (this PaymentType) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("PaymentType Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("PaymentType Table Created")
}