package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Currency struct {
	Name     string `json:"name" sql:"type:varchar(50);"`
	Code     string `json:"code" option:"id" gorm:"primary_key" sql:"type:varchar(3);"`
	FullName string `json:"-" option:"value"`
}

func (this Currency) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Currency Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Currency Table Created")
}

