package model

import (
	"encoding/json"
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Company struct {
	Id int64 `json:"-"`

	Name                     string     `json:"name" sql:"type:varchar(250);"`
	VatNumber                string     `json:"vat_number" sql:"type:varchar(50);"`
	VatArea                  string     `json:"vat_area" sql:"type:varchar(250);"`
	Address                  string     `json:"address" sql:"type:varchar(250);"`
	PostalCode               string     `json:"postal_code" sql:"type:varchar(50);"`
	CityId                   int64      `json:"city"`
	Country                  string     `json:"country" sql:"type:char(3)"`
	StartDateFiscalYear      *time.Time `json:"start_date_fiscal_year"`
	ContactPerson            string     `json:"contact_person" sql:"type:varchar(250);"`
	ContactFunction          string     `json:"contact_function" sql:"type:varchar(250);"`
	ContactEmail             string     `json:"contact_email" sql:"type:varchar(250);"`
	ContactPhone             string     `json:"contact_phone" sql:"type:varchar(50);"`
	StandardDecimalSeparator string     `json:"standard_decimal_separator" sql:"type:varchar(10);"`
	StandardDateFormat       string     `json:"standard_date_format" sql:"type:varchar(50);"`
	SubscriptionPlanId       int64      `json:"subscription_id"`
	NextPaymentDate          time.Time  `json:"-"`
	Color                    string     `json:"color" sql:"type:varchar(50);"`
	Active                   bool       `json:"active"`
	Logo                     string     `json:"logo" sql:"type:varchar(250);`
	LogoUrl                  string     `json:"logo_url" sql:"-" ss:"-"`

	CompanySettings *CompanySettings `json:"settings" ss:"-"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}


func (this Company) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Company Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Company Table Created")
}


type JCompany Company

func (c Company) MarshalJSON() ([]byte, error) {
	if c.Id == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(JCompany(c))
}

type CompanySettings struct {
	Id        int64 `json:"-"`
	CompanyId int64 `json:"-"`

	IsCategoryFieldRequired bool //`json:"-"`
	IsProjectFieldRequired  bool `json:"-"`
	IsPaidWithFieldRequired bool `json:"-"`
	ActiveApprovalModule    bool `json:"-"`
	ActiveControllerModule  bool `json:"-"`
	RemoveProjectField      bool `json:"-"`
	ActivateAllowances      bool `json:"-"`
}

func (this CompanySettings) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("CompanySettings Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("CompanySettings Table Created")
}