package model

import (
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Branch struct {
	Id           int64 `json:"id" option:"id" CaptionML:"enu=ID;trk=ID"`
	CompanyId    int64 `json:"-" CaptionML:"enu=Company ID;trk=Şirket ID"`
	Code         string    `json:"code" sql:"type:varchar(250);" CaptionML:"enu=Code;trk=Kod"`
	Name         string    `json:"name" option:"value" sql:"type:varchar(250);" CaptionML:"enu=Name;trk=Ad"`
	Address      string    `json:"address" sql:"type:varchar(250);" CaptionML:"enu=Address;trk=Adres"`
	PostalCode   string    `json:"postal_code" sql:"type:varchar(50);" CaptionML:"enu=Post Code;trk=Posta Kodu"`
	City         string     `json:"city" ss:"-" CaptionML:"enu=City;trk=Şehir"`
	Country      string  `json:"country" gorm:"foreignkey:CountryCode" ss:"-" CaptionML:"enu=Counrty;trk=Ülke"`
	CurrencyCode string    `json:"-" sql:"type:varchar(50);" CaptionML:"enu=Currency Code;trk=Döviz Kodu"`
	Currency     *Currency `json:"currency" gorm:"foreignkey:CurrencyCode" ss:"-" CaptionML:"enu=Currency;trk=Döviz"`
	Active       bool      `json:"active" CaptionML:"enu=Active;trk=Aktif"`

	AccountantId int64       `json:"-" CaptionML:"enu=Accountant ID;trk=Muhasebe ID"`
	Accountant   *User `json:"accountant,omitempty" gorm:"foreignkey:AccountantId" ss:"-" CaptionML:"enu=Accountant;trk=Muhasebeci"`

	Departments []Department


	CreatedAt    time.Time `json:"-" CaptionML:"enu=Created At;trk=Oluşturma Tarihi"`
	UpdatedAt    time.Time `json:"-" CaptionML:"enu=Updated Date;trk=Değiştirme Tarihi"`
}


func (this Branch) CreateTable() {
	if app.DB.HasTable(&Branch{}) {
		app.DB.DropTable(this)
	}
	fmt.Println("Branch Table Dropped")
	app.DB.CreateTable(this)
	//app.DB.Model(&this).AddForeignKey("company_id","companies(id)","RESTRICT","NO ACTION")
	fmt.Println("Branch Table Created")
	app.MakeCaptionML(this)
}