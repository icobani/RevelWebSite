package model

import (
	"encoding/json"
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Company struct {
	Id                       int64 `json:"-"`

	Name                     string     `json:"name" sql:"type:varchar(250);" CaptionML:"trk=Firma Adı;enu=Company Name"`
	VatNumber                string     `json:"vat_number" sql:"type:varchar(50);" CaptionML:"trk=Vergi No.;enu=VAT Number"`
	VatArea                  string     `json:"vat_area" sql:"type:varchar(250);" CaptionML:"trk=Vergi Dairesi;enu=VAT Area"`
	Address                  string     `json:"address" sql:"type:varchar(250);" CaptionML:"trk=Adres;enu=Address"`
	PostalCode               string     `json:"postal_code" sql:"type:varchar(50);" CaptionML:"trk=Posta Kodu;enu=Postal Code"`
	CityId                   string      `json:"city" CaptionML:"trk=Şehir;enu=City"`
	Country                  string     `json:"country" sql:"type:char(50)" CaptionML:"trk=Ülke;enu=Country"`
	StartDateFiscalYear      time.Time `json:"start_date_fiscal_year" CaptionML:"trk=Mali Yıl Başlangıç Tarihi;enu=Start Date Fiscal Year"`
	ContactPerson            string     `json:"contact_person" sql:"type:varchar(250);" CaptionML:"enu=Contact Person;trk=İlgili Kişi"`
	ContactFunction          string     `json:"contact_function" sql:"type:varchar(250);" CaptionML:"enu=Function;trk=Görevi"`
	ContactEmail             string     `json:"contact_email" sql:"type:varchar(250);" CaptionML:"enu=E-Mail;trk=E-Mail"`
	ContactPhone             string     `json:"contact_phone" sql:"type:varchar(50);" CaptionML:"enu=Phone;trk=Telefon"`
	StandardDecimalSeparator string     `json:"standard_decimal_separator" sql:"type:varchar(10);" CaptionML:"enu=Standart Decimal Seperator;trk=Standart Ondalık Ayıracı"`
	StandardDateFormat       string     `json:"standard_date_format" sql:"type:varchar(50);default:'DayMonth';" CaptionML:"enu=Standart Date Format;trk=Standart Tarih Formatı"`
	SubscriptionPlanId       int64      `json:"subscription_id" CaptionML:"enu=Subscription Plan ID;trk=Abonelik Plan ID"`
	NextPaymentDate          time.Time  `json:"-" CaptionML:"enu=Next Payment Date;trk=Bir Sonraki Ödeme Tarihi"`
	Color                    string     `json:"color" sql:"type:varchar(50);" CaptionML:"enu=Color;trk=Renk"`
	Active                   bool       `json:"active" CaptionML:"enu=Active;trk=Aktif"`
	Logo                     string     `json:"logo" sql:"type:varchar(250);" CaptionML:"enu=Logo;trk=Logo"`
	LogoUrl                  string     `json:"logo_url" sql:"-" ss:"-" CaptionML:"enu=Logo URL;trk=Logo URL"`

	CompanySettings          *CompanySettings `json:"settings" ss:"-"`

	CreatedAt                time.Time `json:"-" CaptionML:"enu=Created At;trk= Oluşturulma Tarihi"`
	UpdatedAt                time.Time `json:"-" CaptionML:"enu=Updated At;trk=Güncelleme Tarihi"`
}

func (this Company) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Company Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Company Table Created")
	app.MakeCaptionML(this)

}



type JCompany Company

func (c Company) MarshalJSON() ([]byte, error) {
	if c.Id == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(JCompany(c))
}

type CompanySettings struct {
	Id                      int64 `json:"-" CaptionML:"enu=ID;trk=ID"`
	CompanyId               int64 `json:"-" CaptionML:"enu=Company ID;trk=Şirket ID"`

	IsCategoryFieldRequired bool `CaptionML:"enu=Is Category Field Required;trk=Kategory Alanı Zorunlu"`
	IsProjectFieldRequired  bool `json:"-" CaptionML:"enu=Make the project field required;trk=Proje Kodu Zorunlu"`
	IsPaidWithFieldRequired bool `json:"-" CaptionML:"enu=Is Paid With Field Required;trk="`
	ActiveApprovalModule    bool `json:"-" CaptionML:"enu=;trk="`
	ActiveControllerModule  bool `json:"-" CaptionML:"enu=;trk="`
	RemoveProjectField      bool `json:"-" CaptionML:"enu=;trk="`
	ActivateAllowances      bool `json:"-" CaptionML:"enu=;trk="`
}

func (this CompanySettings) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("CompanySettings Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("CompanySettings Table Created")
	app.MakeCaptionML(this)
}