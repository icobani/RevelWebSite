package model

import (
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Branch struct {
	Id        int64 `json:"id" option:"id"`
	CompanyId int64 `json:"-"`

	Code         string    `json:"code" sql:"type:varchar(250);"`
	Name         string    `json:"name" option:"value" sql:"type:varchar(250);"`
	Address      string    `json:"address" sql:"type:varchar(250);"`
	PostalCode   string    `json:"postal_code" sql:"type:varchar(50);"`
	CityId       int64     `json:"-"`
	City         *City     `json:"city" ss:"-"`
	CountryCode  string    `json:"-" sql:"type:char(3)"`
	Country      *Country  `json:"country" gorm:"foreignkey:CountryCode" ss:"-"`
	CurrencyCode string    `json:"-" sql:"type:varchar(50);"`
	Currency     *Currency `json:"currency" gorm:"foreignkey:CurrencyCode" ss:"-"`
	Active       bool      `json:"active"`

	AccountantId int64       `json:"-"`
	Accountant   *BranchUser `json:"accountant,omitempty" gorm:"foreignkey:AccountantId" ss:"-"`

	CreatedAt time.Time `json:"-"`
	// UpdatedAt time.Time `json:"-"`
}

//func CreateBranch(db *gorm.DB, cats interface{}) error {
//	return dbutils.MultipleInsert(db, "branches", Branch{}, cats)
//}

func (this Branch) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Branch Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Branch Table Created")
}





type BranchUser struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}


func (this BranchUser) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("BranchUser Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("BranchUser Table Created")
}