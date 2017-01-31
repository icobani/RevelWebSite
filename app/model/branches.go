package model

import (
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Branch struct {
	Id        int64 `json:"id" option:"id" CaptionML:"enu=ID;trk=ID"`
	CompanyId int64 `json:"-" CaptionML:"enu=Company ID;trk=Şirket ID"`

	Code         string    `json:"code" sql:"type:varchar(250);" CaptionML:"enu=Code;trk=Kod"`
	Name         string    `json:"name" option:"value" sql:"type:varchar(250);" CaptionML:"enu=Name;trk=Ad"`
	Address      string    `json:"address" sql:"type:varchar(250);" CaptionML:"enu=Address;trk=Adres"`
	PostalCode   string    `json:"postal_code" sql:"type:varchar(50);" CaptionML:"enu=Post Code;trk=Posta Kodu"`
	CityId       int64     `json:"-" CaptionML:"enu=City ID;trk=Şehir ID"`
	City         *City     `json:"city" ss:"-" CaptionML:"enu=City;trk=Şehir"`
	CountryCode  string    `json:"-" sql:"type:char(3)" CaptionML:"enu=Country Code;trk=Ülke Kodu"`
	Country      *Country  `json:"country" gorm:"foreignkey:CountryCode" ss:"-" CaptionML:"enu=Counrty;trk=Ülke"`
	CurrencyCode string    `json:"-" sql:"type:varchar(50);" CaptionML:"enu=Currency Code;trk=Döviz Kodu"`
	Currency     *Currency `json:"currency" gorm:"foreignkey:CurrencyCode" ss:"-" CaptionML:"enu=Currency;trk=Döviz"`
	Active       bool      `json:"active" CaptionML:"enu=Active;trk=Aktif"`

	AccountantId int64       `json:"-" CaptionML:"enu=Accountant ID;trk=Muhasebe ID"`
	Accountant   *BranchUser `json:"accountant,omitempty" gorm:"foreignkey:AccountantId" ss:"-" CaptionML:"enu=Accountant;trk=Muhasebeci"`

	CreatedAt time.Time `json:"-" CaptionML:"enu=Created At;trk=Oluşturma Tarihi"`
	UpdatedAt time.Time `json:"-" CaptionML:"enu=Updated Date;trk=Değiştirme Tarihi"`
}

//func CreateBranch(db *gorm.DB, cats interface{}) error {
//	return dbutils.MultipleInsert(db, "branches", Branch{}, cats)
//}

func (this Branch) CreateTable() {
	if app.DB.HasTable(&Branch{}) {
		app.DB.DropTable(this)
	}
	fmt.Println("Branch Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Branch Table Created")
	app.MakeCaptionML(this)
}





type BranchUser struct {
	Id       int64  `json:"id" CaptionML:"enu=ID;trk=ID"`
	Name     string `json:"name" CaptionML:"enu=Name;trk=Ad"`
	LastName string `json:"last_name" CaptionML:"enu=Last Name;trk=Soyadı"`
}


func (this BranchUser) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("BranchUser Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("BranchUser Table Created")
	app.MakeCaptionML(this)
}