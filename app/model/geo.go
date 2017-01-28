package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Country struct {
	Code         string `json:"code" option:"id" gorm:"primary_key" sql:"type:char(2);"`
	Name         string `json:"name" option:"value" sql:"type:varchar(52);"`
	CurrencyCode string `json:"currency_code" option:"currency_code" sql:"type:varchar(3);"`
}
func (this Country) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Country Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Country Table Created")
}


type City struct {
	Id          int64  `json:"id" option:"id"`
	Name        string `json:"name" option:"value" sql:"type:varchar(35);"`
	CountryCode string `json:"country_code" sql:"type:char(2);"`
}

func (this City) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("City Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("City Table Created")
}