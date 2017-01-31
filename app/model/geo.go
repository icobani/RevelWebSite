package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Country struct {
	Code         string `json:"code" option:"id" gorm:"primary_key" sql:"type:char(2);" CaptionML:"enu=Code;trk=Kod"`
	Description         string `json:"description" option:"value" sql:"type:varchar(52);" CaptionML:"enu=Name;trk=Açıklama"`
	CurrencyCode string `json:"currency_code" option:"currency_code" sql:"type:varchar(3);" CaptionML:"enu=Currency Code;trk=Döviz Kodu"`
}
func (this Country) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Country Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Country Table Created")
	app.MakeCaptionML(this)
}


type City struct {
	Id          int64  `json:"id" option:"id" CaptionML:"enu=ID;trk=ID"`
	Name        string `json:"name" option:"value" sql:"type:varchar(35);" CaptionML:"enu=City Name;trk=Şehir Adı"`
	CountryCode string `json:"country_code" sql:"type:char(2);" CaptionML:"enu=Country Code;trk=Ülke Kodu"`
}

func (this City) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("City Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("City Table Created")
	app.MakeCaptionML(this)
}